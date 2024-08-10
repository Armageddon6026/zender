package repository

import (
	"errors"

	"github.com/Armageddon6026/zender/pkg/model"
	"github.com/Armageddon6026/zender/pkg/repository/mariadb"
)

type userRpository struct {
	tableName string
}

func getUserRpository() UserRepository {
	return &userRpository{
		tableName: "users",
	}
}

func (u *userRpository) SelectUsers() (model.UserList, error) {
	sqlCmd := mariadb.Select(u.tableName)
	res, err := mariadb.Query[model.UserInfo](sqlCmd)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userRpository) SelectUserByAuth(account, password string) (*model.UserInfo, error) {
	queryClause := &model.UserTable{Account: account, Password: password, Auth: true}
	sqlCmd := mariadb.Select(u.tableName).Where(queryClause)
	res, err := mariadb.Query[model.UserInfo](sqlCmd)
	if err != nil {
		return nil, err
	}
	if len(res) != 1 {
		return nil, errors.New("Wrong account or password")
	}
	return &res[0], nil
}

func (u *userRpository) SelectUserByAccount(account string) (*model.UserInfo, error) {
	queryClause := &model.UserTable{Account: account}
	sqlCmd := mariadb.Select(u.tableName).Where(queryClause)
	res, err := mariadb.Query[model.UserInfo](sqlCmd)
	if err != nil {
		return nil, err
	}
	if len(res) != 1 {
		return nil, errors.New("Wrong user")
	}
	return &res[0], nil
}

func (u *userRpository) UpdateUserByAccount(account string, data *model.UserUpdate) error {
	queryClause := &model.UserTable{Account: account}
	sqlCmd := mariadb.Update(u.tableName, data).Where(queryClause)
	affectedNum, err := mariadb.Exec(sqlCmd)
	if err != nil {
		return err
	}
	if affectedNum <= 0 {
		return errors.New("update " + account + " fail")
	}
	return nil
}

func (u *userRpository) UpdateUserAuthByAccount(account string, data *model.UserAuthUpdate) error {
	queryClause := &model.UserTable{Account: account}
	sqlCmd := mariadb.Update(u.tableName, data).Where(queryClause)
	affectedNum, err := mariadb.Exec(sqlCmd)
	if err != nil {
		return err
	}
	if affectedNum <= 0 {
		return errors.New("update " + account + " fail")
	}
	return nil
}
