package main

import (
	"github.com/gin-gonic/gin"
	"kgm-backend/config"
	"kgm-backend/routes"
)

func main() {
	config.InitDB()

	route := gin.Default()
	route.MaxMultipartMemory = 10 << 20
	v1 := route.Group("api/v1")
	{
		books := v1.Group("/books")
		{
			books.POST("/books", routes.CreateBook)
			books.GET("/all-books", routes.GetBooks)
			books.PUT("/book/:id", routes.EditBook)
			books.DELETE("/book/:id", routes.DeleteBook)
		}
	}
	route.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
