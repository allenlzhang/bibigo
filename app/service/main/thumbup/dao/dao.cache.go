// Code generated by $GOPATH/src/go-common/app/tool/cache/gen. DO NOT EDIT.

/*
  Package dao is a generated cache proxy package.
  It is generated from:
  type _cache interface {
		// 用户点赞列表
		// cache: -singleflight=true -ignores=||start,end
		userLikeList(c context.Context, mid int64, businessID int64, state int8, start, end int) (res []*model.ItemLikeRecord, err error)
	}
*/

package dao

import (
	"context"

	"go-common/app/service/main/thumbup/model"
	"go-common/library/stat/prom"

	"golang.org/x/sync/singleflight"
)

var _ _cache
var cacheSingleFlights = [1]*singleflight.Group{{}}

// userLikeList 用户点赞列表
func (d *Dao) userLikeList(c context.Context, id int64, businessID int64, state int8, start, end int) (res []*model.ItemLikeRecord, err error) {
	addCache := true
	res, err = d.CacheUserLikeList(c, id, businessID, state, start, end)
	if err != nil {
		addCache = false
		err = nil
	}
	if len(res) != 0 {
		prom.CacheHit.Incr("userLikeList")
		return
	}
	var rr interface{}
	sf := d.cacheSFuserLikeList(id, businessID, state, start, end)
	rr, err, _ = cacheSingleFlights[0].Do(sf, func() (r interface{}, e error) {
		prom.CacheMiss.Incr("userLikeList")
		r, e = d.RawUserLikeList(c, id, businessID, state, start, end)
		return
	})
	res = rr.([]*model.ItemLikeRecord)
	if err != nil {
		return
	}
	miss := res
	if !addCache {
		return
	}
	d.cache.Do(c, func(c context.Context) {
		d.AddCacheUserLikeList(c, id, miss, businessID, state)
	})
	return
}
