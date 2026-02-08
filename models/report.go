package models

type SalesReport struct {
	TotalRevenue   int          `json:"total_revenue"`
	TotalTransaksi int          `json:"total_transaksi"`
	ProdukTerlaris *BestSelling `json:"produk_terlaris"`
}

type BestSelling struct {
	Nama       string `json:"nama"`
	QtyTerjual int    `json:"qty_terjual"`
}
