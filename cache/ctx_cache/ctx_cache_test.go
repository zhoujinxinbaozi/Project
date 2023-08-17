package ctx_cache

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"testing"

	"zjx/project/rpc"
)

const (
	ctxCacheKey = "ctx_cache"
)

// 初始化 ctx，存储使用 syncMap
func InitCtxCache(ctx context.Context) context.Context {
	cacheMap := new(sync.Map)
	return context.WithValue(ctx, ctxCacheKey, cacheMap)
}

// 根据key从缓存中获取内容
func Get(ctx context.Context, key string) (interface{}, bool) {
	cacheMap, ok := ctx.Value(ctxCacheKey).(*sync.Map)
	if !ok {
		return nil, false
	}

	return cacheMap.Load(key)
}

// 存储内容
func Store(ctx context.Context, key string, obj interface{}) {
	cacheMap, ok := ctx.Value(ctxCacheKey).(*sync.Map)
	if !ok {
		// common.Logger(ctx).Warnf("cache map not found, will not cache the key %s", key)
		return
	}

	cacheMap.Store(key, obj)
}

// 删除内容
func Delete(ctx context.Context, key string) {
	cacheMap, ok := ctx.Value(ctxCacheKey).(*sync.Map)
	if !ok {
		// common.Logger(ctx).Warnf("cache map not found, will not cache the key %s", key)
		return
	}

	cacheMap.Delete(key)
}

// 单测
func TestCtxCache(t *testing.T) {
	ctx := context.Background()
	ctx = InitCtxCache(ctx)
	for i := 0; i < 10; i++ {
		if i < 5 {
			GetResultWithCache(ctx, 0)
			continue
		}
		GetResultWithCache(ctx, 1)
	}
}

// 获取内容 有缓存走缓存，没有缓存rpc获取
func GetResultWithCache(ctx context.Context, request int) int {
	v, ok := Get(ctx, makeCtxKey(request))
	if ok {
		fmt.Printf("get from ctx cache, value : %v\n", v.(int))
		return v.(int)
	}
	rpcResult := rpc.GetRpcResult(request)
	Store(ctx, makeCtxKey(request), rpcResult)
	fmt.Printf("get from rpc result : %v \n", rpcResult)
	return rpcResult
}

// 构造缓存key
func makeCtxKey(key int) string {
	return fmt.Sprintf("ctx_key_%v", strconv.FormatInt(int64(key), 10))
}
