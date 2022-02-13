package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getUser(c *gin.Context) {
	name := c.Param("name")

	log.Printf("in getUser name = %q", name)

	user := &user{
		Name: name,
		Age:  42,
	}

	c.JSON(http.StatusOK, user)
}

func main() {
	router := gin.Default()

	router.GET("/user/:name", getUser)

	router.Run()
}
