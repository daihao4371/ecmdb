// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package user

import (
	"github.com/Duke1616/ecmdb/internal/department"
	"github.com/Duke1616/ecmdb/internal/policy"
	"github.com/Duke1616/ecmdb/internal/user/internal/repository"
	"github.com/Duke1616/ecmdb/internal/user/internal/repository/cache"
	"github.com/Duke1616/ecmdb/internal/user/internal/repository/dao"
	"github.com/Duke1616/ecmdb/internal/user/internal/service"
	"github.com/Duke1616/ecmdb/internal/user/internal/web"
	"github.com/Duke1616/ecmdb/internal/user/ldapx"
	"github.com/Duke1616/ecmdb/pkg/mongox"
	"github.com/RediSearch/redisearch-go/v2/redisearch"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitModule(db *mongox.Mongo, redisClient *redisearch.Client, ldapConfig ldapx.Config, policyModule *policy.Module, departmentModule *department.Module) (*Module, error) {
	userDAO := dao.NewUserDao(db)
	userRepository := repository.NewResourceRepository(userDAO)
	serviceService := service.NewService(userRepository)
	redisearchLdapUserCache := InitLdapUserCache(redisClient)
	ldapService := service.NewLdapService(ldapConfig, redisearchLdapUserCache)
	service2 := policyModule.Svc
	service3 := departmentModule.Svc
	handler := web.NewHandler(serviceService, ldapService, service2, service3)
	module := &Module{
		Hdl: handler,
		Svc: serviceService,
	}
	return module, nil
}

// wire.go:

var ProviderSet = wire.NewSet(service.NewLdapService, service.NewService, repository.NewResourceRepository, dao.NewUserDao, web.NewHandler)

func InitLdapUserCache(conn *redisearch.Client) cache.RedisearchLdapUserCache {
	return cache.NewRedisearchLdapUserCache(conn)
}
