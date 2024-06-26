// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package worker

import (
	"context"
	"fmt"
	"github.com/Duke1616/ecmdb/internal/worker/internal/event"
	"github.com/Duke1616/ecmdb/internal/worker/internal/event/watch"
	"github.com/Duke1616/ecmdb/internal/worker/internal/repository"
	"github.com/Duke1616/ecmdb/internal/worker/internal/repository/dao"
	"github.com/Duke1616/ecmdb/internal/worker/internal/service"
	"github.com/Duke1616/ecmdb/internal/worker/internal/web"
	"github.com/Duke1616/ecmdb/pkg/mongox"
	"github.com/ecodeclub/mq-api"
	"github.com/google/wire"
	"go.etcd.io/etcd/client/v3"
	"sync"
)

// Injectors from wire.go:

func InitModule(q mq.MQ, db *mongox.Mongo, etcdClient *clientv3.Client) (*Module, error) {
	taskWorkerEventProducer, err := event.NewTaskRunnerEventProducer(q)
	if err != nil {
		return nil, err
	}
	workerDAO := InitWorkerDAO(db, taskWorkerEventProducer)
	workerRepository := repository.NewWorkerRepository(workerDAO)
	serviceService := service.NewService(q, workerRepository, taskWorkerEventProducer)
	taskWorkerWatch := initWatch(etcdClient, serviceService)
	handler := web.NewHandler(serviceService)
	module := &Module{
		Svc: serviceService,
		w:   taskWorkerWatch,
		Hdl: handler,
	}
	return module, nil
}

// wire.go:

var ProviderSet = wire.NewSet(web.NewHandler, service.NewService, repository.NewWorkerRepository)

func initWatch(etcdClient *clientv3.Client, svc service.Service) *watch.TaskWorkerWatch {
	task, err := watch.NewTaskWorkerWatch(etcdClient, svc)
	if err != nil {
		panic(err)
	}

	go task.Watch()
	return task
}

var (
	daoOnce = sync.Once{}
	d       dao.WorkerDAO
)

func InitProducer(producer event.TaskWorkerEventProducer) {
	wt, err := d.ListWorkerTopic(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(wt)
}

func InitWorkerDAO(db *mongox.Mongo, producer event.TaskWorkerEventProducer) dao.WorkerDAO {
	daoOnce.Do(func() {
		d = dao.NewWorkerDAO(db)
		InitProducer(producer)
	})

	return d
}
