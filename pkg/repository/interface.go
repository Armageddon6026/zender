package repository

import "github.com/Armageddon6026/zender/pkg/model"

type Repository interface {
	User() UserRepository
	Group() GroupRepository
	Service() ServiceRepository
	LoginService() LoginServiceRepository
}

type UserRepository interface {
	SelectUsers() (model.UserList, error)
	SelectUserByAuth(string, string) (*model.UserInfo, error)
	SelectUserByAccount(string) (*model.UserInfo, error)
	UpdateUserByAccount(string, *model.UserUpdate) error
	UpdateUserAuthByAccount(string, *model.UserAuthUpdate) error
}

type GroupRepository interface {
	SelectGroups() (model.GroupList, error)
	SelectGroupById(int) (*model.GroupInfo, error)
	SelectGroupByName(string) (*model.GroupInfo, error)
	InsertGroup(*model.GroupInfo) error
	UpdateGroupById(int, *model.GroupUpdate) error
	DeleteGroupById(int) error
}

type ServiceRepository interface {
	SelectServices() (model.ServciceList, error)
	SelectServicesByGroup(int) (model.ServciceList, error)
	SelectServiceByAuth(string, string) (*model.ServiceInfo, error)
	InsertService(*model.ServiceInsert) error
	UpdaateServiceByUuid(string, *model.ServiceUpdate) error
	DeleteServiceByUuid(string) error
}

type LoginServiceRepository interface {
	SelectLoginServices() model.LoginServiceList
	SelectLoginServiceByUuid(string) (model.LoginService, error)
	InsertLoginService(*model.LoginService) error
	UpdateLoginService(*model.LoginService) error
	DeleteLoginServiceByUuid(string)
}
