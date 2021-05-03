package usecase

import (
	"supplier-service/repositories"
	"supplier-service/models"
)

type SupplierUsecase struct {
	rp *repositories.SupplierRepository
}

func NewSupplierUsecase(rp *repositories.SupplierRepository) *SupplierUsecase {
	return &SupplierUsecase{rp}
}


func (r *SupplierUsecase) RegisterSeller(supplier models.Supplier) *models.Supplier {

	status_supplier := r.rp.RegisterSeller(supplier)

	return status_supplier
}

func (r *SupplierUsecase) RegisterStore(supplierstore models.SupplierStore,supplierid int) *models.SupplierStore {

	status_supplierstore := r.rp.RegisterStore(supplierstore,supplierid)

	return status_supplierstore
}

func (r *SupplierUsecase) ValidateSupplierId(id int) bool {
	status := r.rp.ValidateSupplierId(id)
	return status
}