package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"online.eric/ginWebServer/models"
	"online.eric/ginWebServer/stores"
)

func PostSetUser(c *gin.Context) {
	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/

	user := c.MustGet(gin.AuthUserKey).(string)
	// Parse JSON
	var userValue models.UserValue

	if err := c.Bind(&userValue); err == nil {
		stores.UsersDb[user] = userValue.Value
		c.JSON(http.StatusOK, models.Status{Status: "ok", Value: userValue.Value})
	} else {
		c.JSON(http.StatusBadRequest, models.Status{Status: "error", Error: err.Error()})
	}
}
