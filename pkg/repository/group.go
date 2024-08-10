package repository

import (
	"errors"

	"github.com/Armageddon6026/zender/pkg/model"
	"github.com/Armageddon6026/zender/pkg/repository/mariadb"
)

type groupRepository struct {
	tableName string
}

func getGroupRepository() GroupRepository {
	return &groupRepository{
		tableName: "service_group",
	}
}

func (g *groupRepository) SelectGroups() (model.GroupList, error) {
	sqlCmd := mariadb.Select(g.tableName)
	res, err := mariadb.Query[model.GroupInfo](sqlCmd)
	if err != nil {
		return nil, err
	}
	if len(res) <= 0 {
		return nil, errors.New("Wrong query")
	}
	return res, nil
}

func (g *groupRepository) SelectGroupById(id int) (*model.GroupInfo, error) {
	queryClause := &model.GroupTable{Id: id}
	sqlCmd := mariadb.Select(g.tableName).Where(queryClause)
	res, err := mariadb.Query[model.GroupInfo](sqlCmd)
	if err != nil {
		return nil, err
	}
	if len(res) != 1 {
		return nil, errors.New("Wrong id")
	}
	return &res[0], nil
}

func (g *groupRepository) SelectGroupByName(name string) (*model.GroupInfo, error) {
	queryClause := &model.GroupTable{Name: name}
	sqlCmd := mariadb.Select(g.tableName).Where(queryClause)
	res, err := mariadb.Query[model.GroupInfo](sqlCmd)
	if err != nil {
		return nil, err
	}
	if len(res) != 1 {
		return nil, errors.New("Wrong name")
	}
	return &res[0], nil
}

func (g *groupRepository) InsertGroup(data *model.GroupInfo) error {
	sqlCmd := mariadb.Insert(g.tableName, data)
	affectedNum, err := mariadb.Exec(sqlCmd)
	if err != nil {
		return err
	}
	if affectedNum == 0 {
		return errors.New("No group Insert")
	}

	return nil
}

func (g *groupRepository) UpdateGroupById(id int, data *model.GroupUpdate) error {
	queryClause := &model.GroupTable{Id: id}
	sqlCmd := mariadb.Update(g.tableName, data).Where(queryClause)
	affectedNum, err := mariadb.Exec(sqlCmd)
	if err != nil {
		return err
	}
	if affectedNum == 0 {
		return errors.New("No group Update")
	}

	return nil
}

func (g *groupRepository) DeleteGroupById(id int) error {
	queryClause := &model.GroupTable{Id: id}
	sqlCmd := mariadb.Delete(g.tableName).Where(queryClause)
	affectedNum, err := mariadb.Exec(sqlCmd)
	if err != nil {
		return err
	}
	if affectedNum == 0 {
		return errors.New("No group Delete")
	}

	return nil
}
