package local_cache

import (
	"fmt"
	"testing"
	"time"

	"github.com/bluele/gcache"
)

var (
	localCache = gcache.New(3).Build() // 到达指定数量随机淘汰
	// localCacheWithLRU = gcache.New(3).LRU().Build()  // 到达指定数量LRU淘汰
	// localCacheWithLRUAndExpir = gcache.New(3).Expiration(time.Duration(5)*time.Second).LRU().Build() // 到达指定数量LRU淘汰并且有过期时间
)

func TestLocalCache(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := localCache.Set(i, i)
		if err != nil {
			fmt.Printf("set find err : %v", err)
			return
		}
		localCache.Get(0)
		if i >= 5 {
			time.Sleep(time.Duration(1) * time.Second)
		}
		fmt.Println("size : ", localCache.Keys(true))
	}
}
