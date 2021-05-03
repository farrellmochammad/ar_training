package handler

import (
	supplier_usecase "supplier-service/usecases"
	models "supplier-service/models"
	"strconv"
	"github.com/gin-gonic/gin"
)

type restHandler struct {
	su   *supplier_usecase.SupplierUsecase
}

func SupplierNewHandler(su *supplier_usecase.SupplierUsecase) SupplierRestHandler {
	return &restHandler{su : su}
}

func (r *restHandler) RegisterSeller(c *gin.Context) {
	var supplier models.Supplier
	err := c.Bind(&supplier)

	if err != nil {
		c.JSON(401, gin.H{
			"status": "failed",
			"Error":  "There is error on query string",
		})
		return
	}

	status_supplier := r.su.RegisterSeller(supplier)

	if status_supplier == nil {
		c.JSON(403, gin.H{
			"message": "Insert failed",
		})
	} else {
		c.JSON(200, gin.H{
			"message" : "Can insert supplier with id : " + strconv.Itoa(status_supplier.SupplierId),
		})
	}

}

func (r *restHandler) RegisterStore(c *gin.Context) {
	if c.GetBool("isnonvalid") {
		c.JSON(500, gin.H{
			"status": "failed",
			"Error":  "Invalid Id",
		})
		return
	}

	if !c.GetBool("isvalid") {
		c.JSON(500, gin.H{
			"status": "failed",
			"Error":  "Supplier Id is not valid",
		})
		return
	}

	var supplierstore models.SupplierStore
	err := c.Bind(&supplierstore)

	if err != nil {
		c.JSON(401, gin.H{
			"status": "failed",
			"Error":  "There is error on query string",
		})
		return
	}

	paramid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"status": "failed",
			"Error":  "Invalid Id",
		})
		return
	}

	status_supplierstore := r.su.RegisterStore(supplierstore, paramid)

	if status_supplierstore == nil {
		c.JSON(403, gin.H{
			"message": "Insert failed",
		})
	} else {
		c.JSON(200, gin.H{
			"message" : "Can insert supplier with id : " + strconv.Itoa(status_supplierstore.StoreId),
		})
	}

}
