package dao

import (
	"context"
	"fmt"
	"github.com/Duke1616/ecmdb/pkg/mongox"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type ModelDAO interface {
	CreateModel(ctx context.Context, m Model) (int64, error)
	GetModelById(ctx context.Context, id int64) (Model, error)
	ListModels(ctx context.Context, offset, limit int64) ([]Model, error)
	Count(ctx context.Context) (int64, error)
}

func NewModelDAO(db *mongox.Mongo) ModelDAO {
	return &modelDAO{
		db: db,
	}
}

type modelDAO struct {
	db *mongox.Mongo
}

func (dao *modelDAO) CreateModel(ctx context.Context, m Model) (int64, error) {
	session, err := dao.db.DBClient.StartSession()
	if err != nil {
		return 0, fmt.Errorf("无法创建会话: %w", err)
	}
	defer session.EndSession(ctx)

	// 开始事务
	err = session.StartTransaction()
	if err != nil {
		return 0, fmt.Errorf("无法开始事务: %w", err)
	}

	now := time.Now()
	m.Ctime, m.Utime = now.UnixMilli(), now.UnixMilli()
	m.Id = dao.db.GetIdGenerator(ModelCollection)
	col := dao.db.Collection(ModelCollection)

	_, err = col.InsertOne(ctx, m)
	if err != nil {
		return 0, fmt.Errorf("插入数据错误: %w", err)
	}

	// 提交事务
	err = session.CommitTransaction(ctx)
	if err != nil {
		return 0, fmt.Errorf("提交事务错误: %w", err)
	}

	return m.Id, nil
}

func (dao *modelDAO) GetModelById(ctx context.Context, id int64) (Model, error) {
	col := dao.db.Collection(ModelCollection)
	filter := bson.M{"id": id}

	var m Model
	if err := col.FindOne(ctx, filter).Decode(&m); err != nil {
		return Model{}, fmt.Errorf("解码错误: %w", err)
	}

	return m, nil
}

func (dao *modelDAO) ListModels(ctx context.Context, offset, limit int64) ([]Model, error) {
	col := dao.db.Collection(ModelCollection)
	filer := bson.M{}
	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "ctime", Value: -1}},
		Limit: &limit,
		Skip:  &offset,
	}
	cursor, err := col.Find(ctx, filer, opt)
	defer cursor.Close(ctx)
	if err != nil {
		return nil, fmt.Errorf("查询错误, %w", err)
	}

	var result []Model
	if err = cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("解码错误: %w", err)
	}
	if err = cursor.Err(); err != nil {
		return nil, fmt.Errorf("游标遍历错误: %w", err)
	}
	return result, nil
}

func (dao *modelDAO) Count(ctx context.Context) (int64, error) {
	col := dao.db.Collection(ModelCollection)
	filter := bson.M{}

	count, err := col.CountDocuments(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("文档计数错误: %w", err)
	}

	return count, nil
}
