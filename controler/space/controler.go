package space

import (
	"net/http"

	"github.com/Neniel/mock-api/service"
	"github.com/gin-gonic/gin"
)

func GetPeopleInSpaceData(c *gin.Context) {
	r := service.GetPeopleInSpaceDataMock()
	c.JSON(http.StatusOK, r)
}
