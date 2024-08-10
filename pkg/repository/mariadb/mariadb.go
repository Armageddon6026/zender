package mariadb

import (
	"database/sql"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
)

type SqlCommand struct {
	sqlCommand      string
	queryClauseArgs []any
}

type pool struct {
	Db   *sql.DB
	once sync.Once
}

var ourConnectionStr string
var ourPool = new(pool)

func New(config *mysql.Config) {
	ourConnectionStr = config.FormatDSN()
}

func getDB() (*sql.DB, error) {
	var err error
	ourPool.once.Do(func() {
		ourPool.Db, err = sql.Open("mysql", ourConnectionStr)
		ourPool.Db.SetMaxOpenConns(25)
		ourPool.Db.SetMaxIdleConns(25)
		ourPool.Db.SetConnMaxLifetime(5 * time.Minute)
		err = ourPool.Db.Ping()
	})
	return ourPool.Db, err
}

func Select(tableName string) *SqlCommand {
	m := new(SqlCommand)
	m.sqlCommand = "SELECT * FROM " + tableName
	return m
}

func Delete(tableName string) *SqlCommand {
	m := new(SqlCommand)
	m.sqlCommand = "DELETE FROM " + tableName
	return m
}

func Insert(tableName string, data any) *SqlCommand {
	valueOfData := reflect.ValueOf(data)
	//supprt for Ptr data
	if valueOfData.Kind() == reflect.Ptr {
		valueOfData = valueOfData.Elem()
	}
	dataColumnLen := valueOfData.NumField()
	if dataColumnLen <= 0 {
		return nil
	}
	conn := new(SqlCommand)
	conn.queryClauseArgs = make([]any, 0, dataColumnLen)
	columnList := make([]string, 0, dataColumnLen)
	valueList := make([]string, 0, dataColumnLen)
	for i := range dataColumnLen {
		if valueOfData.Field(i).IsZero() {
			continue
		}
		columnList = append(columnList, strings.ToLower(valueOfData.Type().Field(i).Name))
		valueList = append(valueList, "?")
		conn.queryClauseArgs = append(conn.queryClauseArgs, valueOfData.Field(i).Interface())
	}
	conn.sqlCommand = "INSERT INTO " + tableName + "(" + strings.Join(columnList, ",") + ") VALUES(" + strings.Join(valueList, ",") + ")"
	return conn
}

func Update(tableName string, data any) *SqlCommand {
	valueOfData := reflect.ValueOf(data)
	//supprt for Ptr data
	if valueOfData.Kind() == reflect.Ptr {
		valueOfData = valueOfData.Elem()
	}
	dataColumnLen := valueOfData.NumField()
	if dataColumnLen <= 0 {
		return nil
	}
	conn := new(SqlCommand)
	conn.queryClauseArgs = make([]any, 0, dataColumnLen)
	columnList := make([]string, 0, dataColumnLen)
	for i := range dataColumnLen {
		if valueOfData.Field(i).IsZero() {
			continue
		}
		columnList = append(columnList, strings.ToLower(valueOfData.Type().Field(i).Name)+"=?")
		conn.queryClauseArgs = append(conn.queryClauseArgs, valueOfData.Field(i).Interface())
	}
	conn.sqlCommand = "UPDATE " + tableName + " SET " + strings.Join(columnList, ",")
	return conn
}

func (Conn *SqlCommand) Where(queryClause any) *SqlCommand {
	valueOfClause := reflect.ValueOf(queryClause)
	//supprt for Ptr data
	if valueOfClause.Kind() == reflect.Ptr {
		valueOfClause = valueOfClause.Elem()
	}
	clauseColumnLen := valueOfClause.NumField()
	if clauseColumnLen <= 0 {
		return Conn
	}
	clauseList := make([]string, 0, clauseColumnLen)
	if Conn.queryClauseArgs == nil {
		Conn.queryClauseArgs = make([]any, 0, clauseColumnLen)
	}
	for i := range clauseColumnLen {
		if valueOfClause.Field(i).IsZero() {
			continue
		}
		clauseList = append(clauseList, valueOfClause.Type().Field(i).Name+"=?")
		Conn.queryClauseArgs = append(Conn.queryClauseArgs, valueOfClause.Field(i).Interface())
	}
	Conn.sqlCommand += " WHERE " + strings.Join(clauseList, " AND ")
	return Conn
}

// Exec executes a SQL command and returns the number of rows affected by an update, insert, or delete.
func Exec(Conn *SqlCommand) (int64, error) {
	db, err := getDB()
	if err != nil {
		return 0, err
	}
	stmt, err := db.Prepare(Conn.sqlCommand)
	defer stmt.Close()
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(Conn.queryClauseArgs...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// Execute SQL command and returns the query results as a <T>slice
func Query[T any](Conn *SqlCommand) ([]T, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	stmt, err := db.Prepare(Conn.sqlCommand)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	dataRows, err := stmt.Query(Conn.queryClauseArgs...)
	if err != nil {
		return nil, err
	}

	//Make return object list
	ret := make([]T, 0)
	retType := reflect.TypeOf(ret).Elem()
	retRow := reflect.New(retType).Elem()
	//Make T fieldName -> value cache
	retRowFieldCache := make(map[string]reflect.Value, retType.NumField())
	for i := 0; i < retType.NumField(); i++ {
		retRowFieldCache[retType.Field(i).Name] = retRow.Field(i)
	}
	//Create data Scanner
	dataColumns, _ := dataRows.Columns()
	row := make([]any, len(dataColumns))
	for i := range row {
		row[i] = new(any)
	}
	//Add each Row Data to return list
	for dataRows.Next() {
		dataRows.Scan(row...)
		for i, v := range row {
			retValue := retRowFieldCache[stringToCamel(dataColumns[i])]
			rawData := *v.(*any)
			if !retValue.IsValid() || rawData == nil {
				continue
			}
			switch retValue.Kind() {
			case reflect.String:
				retValue.SetString(string(rawData.([]byte)))
			case reflect.Int:
				retValue.SetInt(rawData.(int64))
			default:
				retValue.Set(reflect.ValueOf(rawData))
			}
		}
		ret = append(ret, retRow.Interface().(T))
	}
	return ret, nil
}

func stringToCamel(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}
