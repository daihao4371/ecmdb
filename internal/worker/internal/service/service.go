package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/Duke1616/ecmdb/internal/worker/internal/domain"
	"github.com/Duke1616/ecmdb/internal/worker/internal/event"
	"github.com/Duke1616/ecmdb/internal/worker/internal/repository"
	"github.com/ecodeclub/mq-api"
	"github.com/gotomicro/ego/core/elog"
	"golang.org/x/sync/errgroup"
)

type Service interface {
	Register(ctx context.Context, req domain.Worker) (int64, error)
	FindOrRegisterByName(ctx context.Context, req domain.Worker) (domain.Worker, error)
	FindOrRegisterByKey(ctx context.Context, req domain.Worker) (domain.Worker, error)
	ListWorker(ctx context.Context, offset, limit int64) ([]domain.Worker, int64, error)
	UpdateStatus(ctx context.Context, id int64, status uint8) (int64, error)
	ValidationByName(ctx context.Context, name string) (bool, error)
	FindByName(ctx context.Context, name string) (domain.Worker, error)

	// Execute 推送消息到Kafka
	Execute(ctx context.Context, req domain.Execute) error
}

type service struct {
	repo     repository.WorkerRepository
	logger   *elog.Component
	producer event.TaskWorkerEventProducer
	mq       mq.MQ
}

func (s *service) FindByName(ctx context.Context, name string) (domain.Worker, error) {
	return s.repo.FindByName(ctx, name)
}

func (s *service) Execute(ctx context.Context, req domain.Execute) error {
	evt := event.EworkRunnerExecuteEvent{
		Language:  req.Language,
		Code:      req.Code,
		TaskId:    req.TaskId,
		Args:      req.Args,
		Variables: req.Variables,
	}

	err := s.producer.Produce(ctx, req.Topic, evt)
	if err != nil {
		s.logger.Debug("工作节点发送指令失败",
			elog.FieldErr(err),
			elog.Any("event", evt),
		)
	}

	return err
}

func NewService(mq mq.MQ, repo repository.WorkerRepository, producer event.TaskWorkerEventProducer) Service {
	return &service{
		mq:       mq,
		repo:     repo,
		logger:   elog.DefaultLogger,
		producer: producer,
	}
}

func (s *service) Register(ctx context.Context, req domain.Worker) (int64, error) {
	return s.repo.CreateWorker(ctx, req)
}

func (s *service) FindOrRegisterByName(ctx context.Context, req domain.Worker) (domain.Worker, error) {
	worker, err := s.repo.FindByName(ctx, req.Name)
	if !errors.Is(err, repository.ErrUserNotFound) {
		if req.Status != worker.Status {
			_, err = s.repo.UpdateStatus(ctx, worker.Id, domain.Status.ToUint8(req.Status))
			if err != nil {
				return worker, fmt.Errorf("修改状态失败: %x", err)
			}
			worker.Status = req.Status
		}

		return worker, err
	}

	// 新增工作节点
	id, err := s.repo.CreateWorker(ctx, req)
	if err != nil {
		return domain.Worker{}, fmt.Errorf("创建节点失败: %x", err)
	}
	worker.Id = id

	// 新增 Topic
	if err = s.mq.CreateTopic(ctx, req.Topic, 1); err != nil {
		return domain.Worker{}, fmt.Errorf("创建Topic失败: %x", err)
	}

	// 新增 producer 监听
	if err = s.producer.AddProducer(req.Topic); err != nil {
		return domain.Worker{}, fmt.Errorf("创建Topic失败: %x", err)
	}
	return worker, nil
}

func (s *service) FindOrRegisterByKey(ctx context.Context, req domain.Worker) (domain.Worker, error) {
	worker, err := s.repo.FindByKey(ctx, req.Key)
	if !errors.Is(err, repository.ErrUserNotFound) {
		if req.Status != worker.Status {
			_, err = s.repo.UpdateStatus(ctx, worker.Id, domain.Status.ToUint8(req.Status))
			if err != nil {
				return worker, fmt.Errorf("修改状态失败: %x", err)
			}
			worker.Status = req.Status
		}

		return worker, err
	}

	// 新增工作节点
	id, err := s.repo.CreateWorker(ctx, req)
	if err != nil {
		return domain.Worker{}, fmt.Errorf("创建节点失败: %x", err)
	}
	worker.Id = id

	// 新增 Topic
	if err = s.mq.CreateTopic(ctx, req.Topic, 1); err != nil {
		return domain.Worker{}, fmt.Errorf("创建Topic失败: %x", err)
	}

	// 新增 producer 监听
	if err = s.producer.AddProducer(req.Topic); err != nil {
		return domain.Worker{}, fmt.Errorf("创建Topic失败: %x", err)
	}
	return worker, nil
}

func (s *service) UpdateStatus(ctx context.Context, id int64, status uint8) (int64, error) {
	return s.repo.UpdateStatus(ctx, id, status)
}

func (s *service) ValidationByName(ctx context.Context, name string) (bool, error) {
	_, err := s.repo.FindByName(ctx, name)
	if !errors.Is(err, repository.ErrUserNotFound) {
		return true, err
	}

	return false, err
}

func (s *service) ListWorker(ctx context.Context, offset, limit int64) ([]domain.Worker, int64, error) {
	var (
		eg    errgroup.Group
		ts    []domain.Worker
		total int64
	)
	eg.Go(func() error {
		var err error
		ts, err = s.repo.ListWorker(ctx, offset, limit)
		return err
	})

	eg.Go(func() error {
		var err error
		total, err = s.repo.Total(ctx)
		return err
	})
	if err := eg.Wait(); err != nil {
		return ts, total, err
	}
	return ts, total, nil
}
