package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"online.eric/ginWebServer/models"
	"online.eric/ginWebServer/stores"
)

func GetUser(c *gin.Context) {
	user := c.Params.ByName("name")
	value, ok := stores.UsersDb[user]
	if ok {
		c.JSON(http.StatusOK, models.User{Name: user, Value: value, Status: "ok"})
	} else {
		c.JSON(http.StatusOK, models.User{Name: user, Status: "no value"})
	}
}
