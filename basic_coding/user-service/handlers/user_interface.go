package handler

import "github.com/gin-gonic/gin"

type UserRestHandler interface {
	RegisterUser(c *gin.Context)
}
