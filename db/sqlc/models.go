// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	UserName     string    `json:"user_name"`
	FullName     string    `json:"full_name"`
	HashPassword string    `json:"hash_password"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"created_at"`
}

type Customer struct {
	ID      int64          `json:"id"`
	Phone   string         `json:"phone"`
	Name    string         `json:"name"`
	Address sql.NullString `json:"address"`
}

type Invoice struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	CustomersID int64     `json:"customers_id"`
	TotalMoney  int64     `json:"total_money"`
	HadPaid     int64     `json:"had_paid"`
	IsDone      bool      `json:"is_done"`
	IsDeleted   bool      `json:"is_deleted"`
}

type InvoiceDetail struct {
	ID          int64   `json:"id"`
	InvoiceID   int64   `json:"invoice_id"`
	ProductID   int64   `json:"product_id"`
	PriceAtSell int64   `json:"price_at_sell"`
	Amount      float64 `json:"amount"`
	TotalPrice  int64   `json:"total_price"`
	Discount    float64 `json:"discount"`
	LastPrice   int64   `json:"last_price"`
}

type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Unit        string    `json:"unit"`
	PriceImport int64     `json:"price_import"`
	Amount      int64     `json:"amount"`
	Price       int64     `json:"price"`
	Warehouse   string    `json:"warehouse"`
	CreatedAt   time.Time `json:"created_at"`
	IDSupplier  int64     `json:"id_supplier"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	UserName     string    `json:"user_name"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	CreatedAt    time.Time `json:"created_at"`
	ExpiredAt    time.Time `json:"expired_at"`
}

type Supplier struct {
	ID      int64          `json:"id"`
	Name    string         `json:"name"`
	Phone   string         `json:"phone"`
	Address sql.NullString `json:"address"`
	Notes   sql.NullString `json:"notes"`
}
