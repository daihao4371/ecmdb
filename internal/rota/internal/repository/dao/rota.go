package dao

import (
	"context"
	"fmt"
	"github.com/Duke1616/ecmdb/pkg/mongox"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	RotaCollection = "c_rota"
)

type RotaDao interface {
	Create(ctx context.Context, req Rota) (int64, error)
	List(ctx context.Context, offset, limit int64) ([]Rota, error)
	Count(ctx context.Context) (int64, error)
	Update(ctx context.Context, req Rota) (int64, error)
	Detail(ctx context.Context, id int64) (Rota, error)
	Delete(ctx context.Context, id int64) (int64, error)

	FindOrAddSchedulingRule(ctx context.Context, id int64, rr RotaRule) (int64, error)
	UpdateSchedulingRule(ctx context.Context, id int64, rotaRules []RotaRule) (int64, error)

	FindOrAddAdjustmentRule(ctx context.Context, id int64, rr RotaAdjustmentRule) (int64, error)
	UpdateAdjustmentRule(ctx context.Context, id int64, rotaRules []RotaAdjustmentRule) (int64, error)
}

func NewRotaDao(db *mongox.Mongo) RotaDao {
	return &rotaDao{
		db: db,
	}
}

type rotaDao struct {
	db *mongox.Mongo
}

func (dao *rotaDao) Delete(ctx context.Context, id int64) (int64, error) {
	col := dao.db.Collection(RotaCollection)
	filter := bson.M{"id": id}

	result, err := col.DeleteOne(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("删除文档错误: %w", err)
	}

	return result.DeletedCount, nil
}

func (dao *rotaDao) Update(ctx context.Context, req Rota) (int64, error) {
	col := dao.db.Collection(RotaCollection)
	updateDoc := bson.M{
		"$set": bson.M{
			"name":    req.Name,
			"desc":    req.Desc,
			"owner":   req.Owner,
			"enabled": req.Enabled,
			"utime":   time.Now().UnixMilli(),
		},
	}
	filter := bson.M{"id": req.Id}
	count, err := col.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return 0, fmt.Errorf("修改文档操作: %w", err)
	}

	return count.ModifiedCount, nil
}

func (dao *rotaDao) UpdateAdjustmentRule(ctx context.Context, id int64, rotaRules []RotaAdjustmentRule) (int64, error) {
	col := dao.db.Collection(RotaCollection)
	updateDoc := bson.M{
		"$set": bson.M{
			"adjustment_rules": rotaRules,
			"utime":            time.Now().UnixMilli(),
		},
	}
	filter := bson.M{"id": id}
	count, err := col.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return 0, fmt.Errorf("修改文档操作: %w", err)
	}

	return count.ModifiedCount, nil
}

func (dao *rotaDao) UpdateSchedulingRule(ctx context.Context, id int64, rotaRules []RotaRule) (int64, error) {
	col := dao.db.Collection(RotaCollection)
	updateDoc := bson.M{
		"$set": bson.M{
			"rules": rotaRules,
			"utime": time.Now().UnixMilli(),
		},
	}
	filter := bson.M{"id": id}
	count, err := col.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return 0, fmt.Errorf("修改文档操作: %w", err)
	}

	return count.ModifiedCount, nil
}

func (dao *rotaDao) Detail(ctx context.Context, id int64) (Rota, error) {
	col := dao.db.Collection(RotaCollection)
	var rota Rota
	filter := bson.M{"id": id}

	if err := col.FindOne(ctx, filter).Decode(&rota); err != nil {
		return Rota{}, fmt.Errorf("解码错误，%w", err)
	}

	return rota, nil
}

func (dao *rotaDao) List(ctx context.Context, offset, limit int64) ([]Rota, error) {
	col := dao.db.Collection(RotaCollection)
	filter := bson.M{}
	opts := &options.FindOptions{
		Sort:  bson.D{{Key: "ctime", Value: -1}},
		Limit: &limit,
		Skip:  &offset,
	}

	cursor, err := col.Find(ctx, filter, opts)
	defer cursor.Close(ctx)
	if err != nil {
		return nil, fmt.Errorf("查询错误, %w", err)
	}

	var result []Rota
	if err = cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("解码错误: %w", err)
	}
	if err = cursor.Err(); err != nil {
		return nil, fmt.Errorf("游标遍历错误: %w", err)
	}
	return result, nil
}

func (dao *rotaDao) Count(ctx context.Context) (int64, error) {
	col := dao.db.Collection(RotaCollection)
	filter := bson.M{}

	count, err := col.CountDocuments(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("文档计数错误: %w", err)
	}

	return count, nil
}

func (dao *rotaDao) FindOrAddSchedulingRule(ctx context.Context, id int64, rr RotaRule) (int64, error) {
	col := dao.db.Collection(RotaCollection)
	filter := bson.M{"id": id}

	// 然后执行实际的更新操作
	update := bson.M{
		"$set": bson.M{
			"utime": time.Now().UnixMilli(),
		},
		"$push": bson.M{
			"rules": rr,
		},
	}

	result := col.FindOneAndUpdate(ctx, filter, update,
		options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(true),
	)

	if result.Err() != nil {
		return 0, result.Err()
	}

	var updatedRota Rota
	if err := result.Decode(&updatedRota); err != nil {
		return 0, err
	}

	return updatedRota.Id, nil
}

func (dao *rotaDao) FindOrAddAdjustmentRule(ctx context.Context, id int64, rr RotaAdjustmentRule) (int64, error) {
	col := dao.db.Collection(RotaCollection)
	filter := bson.M{"id": id}
	update := bson.M{
		"$push": bson.M{
			"adjustment_rules": rr,
		},
		"$set": bson.M{
			"utime": time.Now().UnixMilli(),
		},
	}

	result := col.FindOneAndUpdate(ctx, filter, update,
		options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		return 0, result.Err()
	}

	var updatedRota Rota
	if err := result.Decode(&updatedRota); err != nil {
		return 0, err
	}

	return updatedRota.Id, nil
}

func (dao *rotaDao) Create(ctx context.Context, r Rota) (int64, error) {
	now := time.Now()
	r.Ctime, r.Utime = now.UnixMilli(), now.UnixMilli()
	r.Id = dao.db.GetIdGenerator(RotaCollection)
	col := dao.db.Collection(RotaCollection)

	_, err := col.InsertOne(ctx, r)
	if err != nil {
		return 0, fmt.Errorf("插入数据错误: %w", err)
	}

	return r.Id, nil
}

type Rota struct {
	Id              int64                `bson:"id"`
	Name            string               `bson:"name"`
	Desc            string               `bson:"desc"`
	Enabled         bool                 `bson:"enabled"`
	Owner           int64                `bson:"owner"`
	Rules           []RotaRule           `bson:"rules"`
	AdjustmentRules []RotaAdjustmentRule `bson:"adjustment_rules"`
	Ctime           int64                `bson:"ctime"`
	Utime           int64                `bson:"utime"`
}

// RotaRule 值班规则
type RotaRule struct {
	RotaGroups []RotaGroup `bson:"rota_groups"`
	Rotate     Rotate      `bson:"rotate"`
	StartTime  int64       `bson:"start_time"`
	EndTime    int64       `bson:"end_time"`
}

// RotaAdjustmentRule 临时值班规则
type RotaAdjustmentRule struct {
	StartTime int64     `bson:"start_time"`
	EndTime   int64     `bson:"end_time"`
	RotaGroup RotaGroup `bson:"rota_group"`
}

// Rotate 轮换相关参数
type Rotate struct {
	TimeUnit     uint8 `bson:"time_unit"`
	TimeDuration uint8 `bson:"time_duration"`
}

// RotaGroup 值班组
type RotaGroup struct {
	Id      int64   `bson:"id"`
	Name    string  `bson:"name"`
	Members []int64 `bson:"members"`
}
