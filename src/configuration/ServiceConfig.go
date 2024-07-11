package configuration

import (
	"UseGoft/src/daos"
	"UseGoft/src/service"
)

type ServiceConfig struct {
}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}

func (this *ServiceConfig) UserDao() *daos.UserDAO {
	return daos.NewUserDAO()
}

func (this *ServiceConfig) UserService() *service.UserService {
	return service.NewUserService()
}
