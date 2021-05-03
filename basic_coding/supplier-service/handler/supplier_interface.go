package handler

import "github.com/gin-gonic/gin"

type SupplierRestHandler interface {
	RegisterSeller(c *gin.Context)
	RegisterStore(c *gin.Context)
}
