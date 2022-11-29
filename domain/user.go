package domain

import (
	"time"
)

type User struct {
	IdCustomer       int       `json:"id_customer,omitempty"`
	NamaDepan        string    `json:"nama_depan,omitempty"`
	NamaTengah       string    `json:"nama_tengah,omitempty"`
	NamaBelakang     string    `json:"nama_belakang,omitempty"`
	AlamatCustomer   string    `json:"alamat_customer,omitempty"`
	NikCustomer      string    `json:"nik_customer,omitempty"`
	TlpCustomer      string    `json:"tlp_customer,omitempty"`
	EmailCustomer    string    `json:"email_customer,omitempty"`
	Username         string    `json:"username,omitempty"`
	Password         string    `json:"password,omitempty"`
	LastLogin        time.Time `json:"last_login,omitempty"`
	PasswordUpdateAt time.Time `json:"password_update_at,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	Cuuid            string    `json:"cuuid,omitempty"`
}
