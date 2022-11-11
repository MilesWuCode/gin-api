package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"

	"gin-api/auth"
	"gin-api/model"
	"gin-api/plugin"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Controller struct{}

func (ctrl *Controller) List(c *gin.Context) {
	// path
	fmt.Println(c.Request.URL.RequestURI())

	var service Service

	// 預設值
	p := plugin.Pagination{Page: 1, Size: 3, Sort: 0}

	// Bind query string or post data
	c.ShouldBind(&p)

	if list, err := service.List(p); err != nil {
		// Abort(), AbortWithStatusJSON() 不執行後面的middleware不執行
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": list})
	}
}

func (ctrl *Controller) Create(c *gin.Context) {
	// form
	type CreateData struct {
		Name     string `form:"name" json:"name" validate:"required,max=20" label:"名稱"`
		Email    string `form:"email" json:"email" validate:"required,email" label:"帳號"`
		Password string `form:"password" json:"password" validate:"required" label:"密碼"`
	}

	var data CreateData

	// Bind form-data request
	// c.Bind(&user)

	// Bind query string or post data
	c.ShouldBind(&data)

	// 回傳簡單錯誤,gin預設
	// if err := c.ShouldBind(&user); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	// 	return
	// }

	// 回傳複雜錯誤,validator套件
	if err := plugin.Validate(data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})

		return
	}

	var service Service

	var user model.User

	// 表單值,key-value,複雜需要做客制代碼
	user.Name = data.Name

	user.Email = data.Email

	if password, err := auth.SetPassword(data.Password); err != nil {
		// Abort(), AbortWithStatusJSON() 不執行後面的middleware不執行
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})

		return
	} else {
		user.Password = password
	}

	if err := service.Create(&user); err != nil {
		// Abort(), AbortWithStatusJSON() 不執行後面的middleware不執行
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

func (ctrl *Controller) Get(c *gin.Context) {
	// path
	fmt.Println(c.Request.URL.Path)

	// Parameters in path
	id := c.Param("id")

	var service Service

	var user model.User

	if err := service.Get(id, &user); err != nil {
		// Abort(), AbortWithStatusJSON() 不執行後面的middleware不執行
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

func (ctrl *Controller) Update(c *gin.Context) {
	// path
	fmt.Println(c.Request.URL.Path)

	// form
	type UpdateData struct {
		Name string `json:"name" form:"name" validate:"required,max=20" label:"名稱"`
	}

	// Parameters in path
	id := c.Param("id")

	var data UpdateData

	c.ShouldBind(&data)

	if err := plugin.Validate(data); err != nil {
		// Abort(), AbortWithStatusJSON() 不執行後面的middleware不執行
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})

		return
	}

	// 表單值,key-value,複雜需要做客制代碼
	dataMap := plugin.StructToMapString(data)

	var service Service

	var user model.User

	if err := service.Update(id, dataMap, &user); err != nil {
		// Abort(), AbortWithStatusJSON() 不執行後面的middleware不執行
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

func (ctrl *Controller) Delete(c *gin.Context) {
	// Parameters in path
	id := c.Param("id")

	var service Service

	if err := service.Delete(id); err != nil {
		// Abort(), AbortWithStatusJSON() 不執行後面的middleware不執行
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		// Status() 只回應 http
		c.Status(http.StatusNoContent)
	}
}

func (ctrl *Controller) UploadAvatar(c *gin.Context) {
	// Parameters in path
	id := c.Param("id")

	var user model.User

	var service Service

	if err := service.Get(id, &user); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})

		return
	}

	file, _ := c.FormFile("file")

	fileName := plugin.StringRand(32) + filepath.Ext(file.Filename)
	filePath := "/tmp/" + fileName
	objectName := "/" + id + "/" + fileName
	contentType := file.Header.Get("Content-Type")
	bucketName := "gin-api"

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})

		return
	}

	minioClient := plugin.InitMinio()

	ctx := context.Background()

	// Upload the zip file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})

	if err != nil {
		logger := plugin.InitLog()

		defer logger.Sync()

		logger.Error("service.Delete(id)", zap.String("err", err.Error()))

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "file upload error"})

		return
	}

	type UploadFile struct {
		Name        string `json:"name"`
		Size        int64  `json:"size"`
		Path        string `json:"path"`
		Url         string `json:"url"`
		ContentType string `json:"contentType"`
	}

	uploadFile := UploadFile{
		Name:        file.Filename,
		Size:        info.Size,
		Path:        fmt.Sprintf("/%s/%s/%s", bucketName, id, fileName),
		Url:         fmt.Sprintf("%s/%s/%s/%s", "https://storage", bucketName, id, fileName),
		ContentType: contentType,
	}

	c.JSON(http.StatusOK, gin.H{"data": uploadFile})
}
