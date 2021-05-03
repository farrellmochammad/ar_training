package repositories

import (
	"supplier-service/models"
	"supplier-service/drivers/sqllite"
)

type SupplierRepository struct {
	db *sqllite.SqlLiteDatasource
}

func NewSupplierRepository(db *sqllite.SqlLiteDatasource) *SupplierRepository {
	return &SupplierRepository{db}
}

func (r *SupplierRepository) RegisterSeller(supplier models.Supplier) *models.Supplier {
	supplier_data := r.db.InsertSupplier(supplier)

	if supplier_data == nil {
        return nil
    }

	return supplier_data
}

func (r *SupplierRepository) RegisterStore(supplierstore models.SupplierStore,supplierid int) *models.SupplierStore {
	supplierstore_data := r.db.InsertStore(supplierstore,supplierid)

	if supplierstore_data == nil {
        return nil
    }

	return supplierstore_data
}

func (r *SupplierRepository) ValidateSupplierId(id int) bool  {
	status := r.db.ValidateSupplierId(id)

	return status
}