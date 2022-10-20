package user

import (
	"log"
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
		routerGroup.GET("/", cache.CachePage(store, time.Minute*10, controller.List))

		routerGroup.POST("/", controller.Create)

		// withOutQuery
		routerGroup.GET("/:id", cache.CachePageWithoutQuery(store, time.Minute*10, controller.Get))

		routerGroup.PUT("/:id", controller.Update, func(c *gin.Context) {
			if err := store.Delete(cache.CreateKey("/user/1")); err != nil {
				log.Println("remove cache:", err)
			}

			c.Next()
		})

		routerGroup.DELETE("/:id", controller.Delete)

		routerGroup.POST("/:id/avatar", controller.UploadAvatar)
	}
}
