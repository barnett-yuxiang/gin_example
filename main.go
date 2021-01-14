package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Body struct {
	Tag string `json:"tag"`
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		log.Println("ok ping")
	})
	router.POST("/monitor", func(c *gin.Context) {
		//log.Println(c.Query("tag"))
		//buf := make([]byte, 1024)
		//n, _ := c.Request.Body.Read(buf)
		//log.Println("yuli: ", string(buf[0:n]))

		// ---> 绑定数据
		var data Body
		if err := c.ShouldBindJSON(&data); err != nil {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}

		// --> 返回
		c.JSON(http.StatusOK, gin.H{"msg": "ok"})

		log.Printf("%+v", data)

		log.Println("ok monitor")
	})
	router.Run(":8080")
}
