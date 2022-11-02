package user

import (
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	// router做cache
	// 本機很快
	store := persistence.NewInMemoryStore(time.Second)
	// 外部redis會慢一點
	// store := persistence.NewRedisCache("192.168.50.92:6379", "", time.Second)

	controller := Controller{}

	routerGroup := router.Group("/user")
	{
		// withQuery:page&size
		routerGroup.GET("/", cache.CachePage(store, time.Minute, controller.List))

		routerGroup.POST("/", controller.Create)

		// withOutQuery
		routerGroup.GET("/:id", GetPolicy(), cache.CachePageWithoutQuery(store, time.Minute*10, controller.Get))

		routerGroup.PUT("/:id", UpdatePolicy(), controller.Update, ClearCache(store))

		routerGroup.DELETE("/:id", DeletePolicy(), controller.Delete, ClearCache(store))

		routerGroup.POST("/:id/avatar", controller.UploadAvatar, ClearCache(store))
	}
}

// 清除快取
func ClearCache(store *persistence.InMemoryStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Next() 上面的代碼先進先出執行
		c.Next()
		// Next() 下面的代碼先進後出執行

		// 清除 GET:/user/:id 產生的舊快取
		store.Delete(cache.CreateKey("/user/" + c.Param("id")))
	}
}
