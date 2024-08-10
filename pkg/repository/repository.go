package repository

import (
	"github.com/Armageddon6026/zender/pkg/repository/mariadb"

	"github.com/go-sql-driver/mysql"
)

type repository struct {
	user         UserRepository
	group        GroupRepository
	service      ServiceRepository
	loginService LoginServiceRepository
}

func New(dbConfig *mysql.Config) Repository {
	mariadb.New(dbConfig)
	return &repository{
		user:         getUserRpository(),
		group:        getGroupRepository(),
		service:      getServiceRepository(),
		loginService: getLoginServiceRepository(),
	}
}

func (r *repository) User() UserRepository {
	return r.user
}

func (r *repository) Group() GroupRepository {
	return r.group
}

func (r *repository) Service() ServiceRepository {
	return r.service
}

func (r *repository) LoginService() LoginServiceRepository {
	return r.loginService
}
