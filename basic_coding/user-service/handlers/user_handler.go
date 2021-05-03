package handler

import (
	user_usecase "user-service/usecases"
	models "user-service/models"
	"strconv"
	"github.com/gin-gonic/gin"
)

type restHandler struct {
	su   *user_usecase.UserUsecase
}

func UserNewHandler(su *user_usecase.UserUsecase) UserRestHandler {
	return &restHandler{su : su}
}

func (r *restHandler) RegisterUser(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)

	if err != nil {
		c.JSON(401, gin.H{
			"status": "failed",
			"Error":  "There is error on query string",
		})
		return
	}

	status_user := r.su.RegisterUser(user)

	if status_user == nil {
		c.JSON(403, gin.H{
			"message": "Insert failed",
		})
	} else {
		c.JSON(200, gin.H{
			"message" : "Can insert user with id : " + strconv.Itoa(status_user.UserId),
		})
	}

}
