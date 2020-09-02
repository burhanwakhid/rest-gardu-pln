package main

import (
	"net/http"

	"burhanwakhid.space/controller"
	"burhanwakhid.space/model"
	"github.com/gin-gonic/gin"
)

//CORSMiddleware ...
//CORS (Cross-Origin Resource Sharing)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			// fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {

	r := gin.Default()

	//models
	db := model.SetUpModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.Use(CORSMiddleware())

	v1 := r.Group("/v1/gardu", gin.BasicAuth(gin.Accounts{
		"foo": "foo",
	}))
	{
		/*** START Gardu ***/
		gardu := new(controller.GarduController)

		v1.GET("/insert-gardu", gardu.AddListGardu)
		v1.POST("/list-gardu", gardu.ListGardu)

	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "OK"})
	})

	r.Run()
}
