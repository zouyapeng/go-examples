package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CORS() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", context.Request.Header.Get("Origin"))
		context.Writer.Header().Set("Access-Control-Max-Age", "86400")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Apitoken")
		context.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(http.StatusOK)
		} else {
			context.Next()
		}
	}
}

func AuthSessionMiddle(context *gin.Context) {
	token := context.Request.Header.Get("token")
	if token == "test-token" {
		context.Set("auth", true)
		context.Next()
	} else {
		context.Set("auth", false)
		context.AbortWithStatus(http.StatusUnauthorized)
	}
}

// url上的参数
type APITestQueryInputs struct {
	Limit int `json:"limit" form:"limit"`
	Page  int `json:"page" form:"page"`
}

func GetTest(context *gin.Context) {
	// 默认值
	inputs := APITestQueryInputs{
		Limit: 500,
		Page:  1,
	}
	// 这样也行，但是都是string
	// limit := context.DefaultQuery("limit", "1")
	// Page := context.DefaultQuery("page", "500")
	// fmt.Println(limit, Page)

	if err := context.Bind(&inputs); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "error", "msg": err.Error()})
		return
	}
	// do something
	fmt.Printf("do something\n")
	context.JSON(http.StatusOK, gin.H{"status": "ok"})
}

type APITestReqData struct {
	Name       string     `json:"name" binding:"required"`
	Tests      []string   `json:"tests" binding:"required"`
	Count      int        `json:"count" binding:"required"`
	CreateTime *time.Time `json:"create_time"`
	UpdateTime *time.Time `json:"update_time"`
}

func (a APITestReqData) CheckFormat() (err error) {
	switch {
	case !strings.HasPrefix(a.Name, "test"):
		err = errors.New("name do not start with 'test'")
	case a.Count > 10:
		err = errors.New("'count' cannot greater than 10")
	}
	return
}

func PostTest(context *gin.Context) {
	var inputs APITestReqData
	if err := context.Bind(&inputs); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "error", "msg": err.Error()})
		return
	}

	// 检查入参
	err := inputs.CheckFormat()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "error", "msg": err.Error()})
		return
	}

	// do something
	fmt.Printf("do something\n")

	context.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func PutTest(context *gin.Context) {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "error", "msg": err.Error()})
	}
	if idInt <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{"status": "error", "msg": "illegal parameter"})
	}
	// do something
	fmt.Printf("do something\n")
	context.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func DeleteTest(context *gin.Context) {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "error", "msg": err.Error()})
	}
	if idInt <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{"status": "error", "msg": "illegal parameter"})
	}
	// do something
	fmt.Printf("do something\n")
	context.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func GenerateRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	api.Use(AuthSessionMiddle)

	// curl -H Content-Type:application/json -H token:test-token 'http://127.0.0.1:8080/api/v1/test?limit=1&page=5'
	api.GET("/test", GetTest)

	// curl -X POST -H Content-Type:application/json -H token:test-token -d '{"name":"test-1111","tests":["111","222"],"count":10,"create_time":"2019-12-25T10:10:10Z","update_time":"2019-12-25T10:10:11Z"}' http://127.0.0.1:8080/api/v1/test
	api.POST("/test", PostTest)

	// curl -X PUT -H Content-Type:application/json -H token:test-token http://127.0.0.1:8080/api/v1/test/1
	api.PUT("/test/:id", PutTest)

	// curl -X DELETE -H Content-Type:application/json -H token:test-token http://127.0.0.1:8080/api/v1/test/1
	api.DELETE("/test/:id", DeleteTest)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORS())

	// string
	// curl http://127.0.0.1:8080
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, I'm Gin\n")
	})

	// json
	// curl http://127.0.0.1:8080/json
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, I'm Gin"})
	})

	GenerateRoutes(r)

	return r
}

func main() {
	r := setupRouter()
	_ = r.Run(":8080") //默认端口是8080
}
