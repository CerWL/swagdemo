package handler 

import (
	"net/http"
	"github.com/CerWL/swagdemo/user/model"
	"github.com/CerWL/swagdemo/db"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

// ListUser is the handler of list user endpoint 
// @Summary List Users
// @Description list all the users bsased on filter given
// @Tags user
// @Produce json
// @Param user body model.CreateUser true "create user"
// @Success 200 {object} model.UserDetail
// @Failure 400 {object} httputil.HTTPError
// @Router /user/ [post]

func ListUser() func(c *gin.Context){
	user := model.UserDetail{db.User{UserID:1,Name:"Test",Email:"test@gmail.com"}}
	users := model.UserList{[]model.UserDetails{user}}
	return func(c *gin.Context) {
		c.JSON(200, users)
	}
}

func CreateUser() func(c *gin.Context) {
	user := mode.UserDetail{}
	return func(c *gin.Context) {
		var req model.CreateUser
		if err := c.ShouldBindJSon(&req); err != nil {
			httputil.NewError(c, http.StatusBadRequest, err)
			return
		}
		user.UserID = 1
		user.Name = req.Name
		user.Email = req.Email
		c.JSON(200, user)
	}
}
