package controller

import (
	"context"
	"fmt"
	"gin-test/model"
	"gin-test/plugin"
	"gin-test/service"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
	"log"
	"net/http"
)

var logger *zap.Logger

func init() {
	logger = plugin.InitLog()

	defer logger.Sync()
}

type UserController struct{}

func (ctrl *UserController) List(c *gin.Context) {
	var userService service.UserService

	// 預設值
	p := plugin.Pagination{Page: 1, Size: 3}

	// Bind query string or post data
	c.ShouldBind(&p)

	if list, err := userService.List(p); err != nil {
		logger.Error("userService.List(p)", zap.String("err", err.Error()))

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": list})
	}
}

func (ctrl *UserController) Create(c *gin.Context) {
	var user model.User

	// Bind form-data request
	// c.Bind(&user)

	// Bind query string or post data
	c.ShouldBind(&user)

	// 回傳簡單錯誤,gin預設
	// if err := c.ShouldBind(&user); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	// 	return
	// }

	// 回傳複雜錯誤,validator套件
	if err := plugin.Validate(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})

		return
	}

	var userService service.UserService

	if err := userService.Create(&user); err != nil {
		logger.Error("userService.Create(&user)", zap.String("err", err.Error()))

		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

func (ctrl *UserController) Get(c *gin.Context) {
	// Parameters in path
	id := c.Param("id")

	var userService service.UserService

	var user model.User

	if err := userService.Get(id, &user); err != nil {
		logger.Error("userService.Get(id)", zap.String("err", err.Error()))

		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

func (ctrl *UserController) Update(c *gin.Context) {
	type UpdateData struct {
		Name string `json:"name" form:"name" validate:"required,max=20" label:"名稱"`
	}

	// Parameters in path
	id := c.Param("id")

	var userService service.UserService

	var data UpdateData

	c.ShouldBind(&data)

	if err := plugin.Validate(data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})

		return
	}

	dataMap := plugin.StructToMapString(data)

	var user model.User

	if err := userService.Update(id, dataMap, &user); err != nil {
		logger.Error("userService.Update(id, &updateUser)", zap.String("err", err.Error()))

		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

func (ctrl *UserController) Delete(c *gin.Context) {
	// Parameters in path
	id := c.Param("id")

	var userService service.UserService

	if err := userService.Delete(id); err != nil {
		logger.Error("userService.Delete(id)", zap.String("err", err.Error()))

		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	} else {
		c.AbortWithStatus(http.StatusNoContent)
	}
}

func (ctrl *UserController) UploadAvatar(c *gin.Context) {
	// Parameters in path
	id := c.Param("id")

	file, _ := c.FormFile("file")

	objectName := file.Filename
	filePath := "/tmp/" + file.Filename
	contentType := "application/json"
	size := file.Size
	bucketName := "gin-api"

	c.SaveUploadedFile(file, filePath)

	fmt.Println(id, objectName, contentType, size, file.Header)

	minioClient := plugin.InitMinio()
	ctx := context.Background()
	// Upload the zip file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})

	if err != nil {
		fmt.Printf("%v", err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

	c.JSON(http.StatusOK, gin.H{"data": file})
}
