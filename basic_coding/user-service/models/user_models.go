
package models


import (
	"time"

)


type User struct {
	UserId        int          `json:"user_id"`
	UserAddress      string       `json:"user_address"`
	Created *time.Time   `json:"created_at,omitempty"`
}
