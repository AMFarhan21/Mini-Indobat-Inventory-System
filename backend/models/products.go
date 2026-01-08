package models

type Products struct {
	Id       int     `json:"id" gorm:"primaryKey;autoIncrement"`
	NamaObat string  `json:"nama_obat"`
	Stok     *int    `json:"stok"`
	Harga    float64 `json:"harga"`
}
