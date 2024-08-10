package repository

import (
	"errors"

	"github.com/Armageddon6026/zender/pkg/model"
	"github.com/Armageddon6026/zender/pkg/repository/mariadb"
)

type serviceRepository struct {
	tableName string
}

func getServiceRepository() ServiceRepository {
	return &serviceRepository{
		tableName: "service_infomation",
	}
}

func (s *serviceRepository) SelectServiceByAuth(uuid, privateKey string) (*model.ServiceInfo, error) {
	queryClause := &model.ServiceTable{Uuid: uuid, PrivateKey: privateKey}
	sqlCmd := mariadb.Select(s.tableName).Where(queryClause)
	res, err := mariadb.Query[model.ServiceInfo](sqlCmd)
	if err != nil {
		return nil, err
	}
	if len(res) != 1 {
		return nil, errors.New("Wrong uuid & privateKey")
	}

	return &res[0], nil
}

func (s *serviceRepository) SelectServices() (model.ServciceList, error) {
	sqlCmd := mariadb.Select(s.tableName)
	res, err := mariadb.Query[model.ServiceInfo](sqlCmd)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *serviceRepository) SelectServicesByGroup(group int) (model.ServciceList, error) {
	queryClause := &model.ServiceTable{GroupId: group}
	sqlCmd := mariadb.Select(s.tableName).Where(queryClause)
	res, err := mariadb.Query[model.ServiceInfo](sqlCmd)
	if err != nil {
		return nil, err
	}
	if len(res) <= 0 {
		return nil, errors.New("Wrong group")
	}

	return res, nil
}

func (s *serviceRepository) InsertService(data *model.ServiceInsert) error {
	sqlCmd := mariadb.Insert(s.tableName, data)
	affectedNum, err := mariadb.Exec(sqlCmd)
	if err != nil {
		return err
	}
	if affectedNum == 0 {
		return errors.New("No Service Insert")
	}

	return nil
}

func (s *serviceRepository) UpdaateServiceByUuid(uuid string, data *model.ServiceUpdate) error {
	queryClause := &model.ServiceTable{Uuid: uuid}
	sqlCmd := mariadb.Update(s.tableName, data).Where(queryClause)
	affectedNum, err := mariadb.Exec(sqlCmd)
	if err != nil {
		return err
	}
	if affectedNum == 0 {
		return errors.New("No Service Update")
	}

	return nil
}

func (s *serviceRepository) DeleteServiceByUuid(uuid string) error {
	queryClause := &model.ServiceTable{Uuid: uuid}
	sqlCmd := mariadb.Delete(s.tableName).Where(queryClause)
	affectedNum, err := mariadb.Exec(sqlCmd)
	if err != nil {
		return err
	}
	if affectedNum == 0 {
		return errors.New("No Service Delete")
	}

	return nil
}
