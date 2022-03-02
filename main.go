package main

import (
	"github.com/Neniel/mock-api/controler/space"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.GET("/space/people-in-space", space.GetPeopleInSpaceData)

	e.Run()
}
