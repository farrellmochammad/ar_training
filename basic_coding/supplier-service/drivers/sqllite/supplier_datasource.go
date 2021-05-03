package sqllite

import (
	"database/sql"
	"supplier-service/models"
	"time"
)

type SqlLiteDatasource struct {
	db        *sql.DB
}

func NewSqlLiteDatasource(db *sql.DB) *SqlLiteDatasource {
    return &SqlLiteDatasource{db}
}

func (r *SqlLiteDatasource) InsertSupplier(supplier models.Supplier) *models.Supplier {
	statement, _ := r.db.Prepare("INSERT INTO supplier_seller (name, created) VALUES (?, ?)")
    statement.Exec(supplier.Name, time.Now())
    rows, _ := r.db.Query("SELECT supplier_id, name FROM supplier_seller")
    var id int
    var name string
    for rows.Next() {
        rows.Scan(&id, &name)
    }

	defer statement.Close()
	defer rows.Close()

	return &models.Supplier{SupplierId : id, Name : name}
}


func (r *SqlLiteDatasource) InsertStore(supplierstore models.SupplierStore, supplierid int) *models.SupplierStore {
	statement, _ := r.db.Prepare("INSERT INTO supplier_store ( supplier_id,store_name, store_address,created) VALUES (?, ?, ?, ?)")
    statement.Exec(supplierid, supplierstore.StoreName, supplierstore.StoreAddress, time.Now())
    rows, _ := r.db.Query("SELECT store_id FROM supplier_store")
    var id int
    for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
		  // handle this error
		  panic(err)
		}
    }

	defer statement.Close()
	defer rows.Close()

	return &models.SupplierStore{StoreId : id}
}

func (r *SqlLiteDatasource) ValidateSupplierId(id int) bool {

    rows, _ := r.db.Query("SELECT supplier_id FROM supplier_seller")
	
    var supplierid int
    for rows.Next() {
        err := rows.Scan(&supplierid)
		if err != nil {
		  // handle this error
		  panic(err)
		}
		if (supplierid == id) {
			return true
		}
		return false
    }

	defer rows.Close()

	return false
}