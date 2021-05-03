package middlewares

import (
	"github.com/gin-gonic/gin"
	supplier_usecase "supplier-service/usecases"
	"strconv"
	"fmt"
)

type middleWareHandler struct {
	su   *supplier_usecase.SupplierUsecase
}

func SupplierMiddlewareNewHandler(su *supplier_usecase.SupplierUsecase) *middleWareHandler {
	return &middleWareHandler{su : su}
}

func (r *middleWareHandler) AuthHandler(su *supplier_usecase.SupplierUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		supplierid := c.Param("id")

		result,err := strconv.Atoi(supplierid)

		status := r.su.ValidateSupplierId(result)

		c.Set("isvalid", status)

		fmt.Println("Status : ",status)

		if err != nil {
			c.Set("isnonvalid", true)
		}

		c.Next()

	}
}