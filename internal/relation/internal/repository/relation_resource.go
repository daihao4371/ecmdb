package repository

import (
	"context"
	"github.com/Duke1616/ecmdb/internal/relation/internal/domain"
	"github.com/Duke1616/ecmdb/internal/relation/internal/repository/dao"
	"github.com/ecodeclub/ekit/slice"
)

type RelationResourceRepository interface {
	CreateResourceRelation(ctx context.Context, req domain.ResourceRelation) (int64, error)

	// ListSrcResources 查询资源列表
	ListSrcResources(ctx context.Context, modelUid string, id int64) ([]domain.ResourceRelation, error)
	ListDstResources(ctx context.Context, modelUid string, id int64) ([]domain.ResourceRelation, error)

	// ListSrcAggregated 聚合查询关联列表
	ListSrcAggregated(ctx context.Context, modelUid string, id int64) ([]domain.ResourceAggregatedAssets, error)
	ListDstAggregated(ctx context.Context, modelUid string, id int64) ([]domain.ResourceAggregatedAssets, error)

	// ListSrcRelated 查询当前已经关联的数据，新增资源关联使用
	ListSrcRelated(ctx context.Context, modelUid, relationName string, id int64) ([]int64, error)
	ListDstRelated(ctx context.Context, modelUid, relationName string, id int64) ([]int64, error)
}

func NewRelationResourceRepository(dao dao.RelationResourceDAO) RelationResourceRepository {
	return &resourceRepository{
		dao: dao,
	}
}

type resourceRepository struct {
	dao dao.RelationResourceDAO
}

func (r *resourceRepository) CreateResourceRelation(ctx context.Context, req domain.ResourceRelation) (int64, error) {
	return r.dao.CreateResourceRelation(ctx, r.toEntity(req))
}

func (r *resourceRepository) ListSrcAggregated(ctx context.Context, modelUid string, id int64) ([]domain.ResourceAggregatedAssets, error) {
	rrs, err := r.dao.ListSrcAggregated(ctx, modelUid, id)
	return slice.Map(rrs, func(idx int, src dao.ResourceAggregatedAsset) domain.ResourceAggregatedAssets {
		return r.toAggregatedAssetsDomain(src)
	}), err
}

func (r *resourceRepository) ListDstAggregated(ctx context.Context, modelUid string, id int64) ([]domain.ResourceAggregatedAssets, error) {
	rrs, err := r.dao.ListDstAggregated(ctx, modelUid, id)
	return slice.Map(rrs, func(idx int, src dao.ResourceAggregatedAsset) domain.ResourceAggregatedAssets {
		return r.toAggregatedAssetsDomain(src)
	}), err
}

func (r *resourceRepository) ListSrcResources(ctx context.Context, modelUid string, id int64) ([]domain.ResourceRelation, error) {
	rrs, err := r.dao.ListSrcResources(ctx, modelUid, id)
	return slice.Map(rrs, func(idx int, src dao.ResourceRelation) domain.ResourceRelation {
		return r.toResourceDomain(src)
	}), err
}

func (r *resourceRepository) ListDstResources(ctx context.Context, modelUid string, id int64) ([]domain.ResourceRelation, error) {
	rrs, err := r.dao.ListDstResources(ctx, modelUid, id)
	return slice.Map(rrs, func(idx int, src dao.ResourceRelation) domain.ResourceRelation {
		return r.toResourceDomain(src)
	}), err
}

func (r *resourceRepository) ListSrcRelated(ctx context.Context, modelUid, relationName string, id int64) ([]int64, error) {
	return r.dao.ListSrcRelated(ctx, modelUid, relationName, id)
}

func (r *resourceRepository) ListDstRelated(ctx context.Context, modelUid, relationName string, id int64) ([]int64, error) {
	return r.dao.ListDstRelated(ctx, modelUid, relationName, id)
}

func (r *resourceRepository) toEntity(req domain.ResourceRelation) dao.ResourceRelation {
	return dao.ResourceRelation{
		RelationName:     req.RelationName,
		SourceResourceID: req.SourceResourceID,
		TargetResourceID: req.TargetResourceID,
	}
}

func (r *resourceRepository) toResourceDomain(resourceDao dao.ResourceRelation) domain.ResourceRelation {
	return domain.ResourceRelation{
		ID:               resourceDao.Id,
		SourceModelUID:   resourceDao.SourceModelUID,
		TargetModelUID:   resourceDao.TargetModelUID,
		SourceResourceID: resourceDao.SourceResourceID,
		TargetResourceID: resourceDao.TargetResourceID,
		RelationTypeUID:  resourceDao.RelationTypeUID,
		RelationName:     resourceDao.RelationName,
	}
}

func (r *resourceRepository) toAggregatedAssetsDomain(src dao.ResourceAggregatedAsset) domain.ResourceAggregatedAssets {
	return domain.ResourceAggregatedAssets{
		RelationName: src.RelationName,
		ModelUid:     src.ModelUid,
		Count:        src.Count,
		ResourceIds:  src.ResourceIds,
	}
}
