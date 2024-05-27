package main

import "fmt"

// ================ FUNGSI BENTUK UMUM PADA GOLANG ================
func main() {

	HasilKonversi := hitungMataUangUSD(3)
	hasilKonversi, _ := hitungMataUang(3, "JPY")

	fmt.Println(HasilKonversi)
	fmt.Println(hasilKonversi)
}

func hitungMataUangUSD(uang int) float32 {
	
	hasil := uang * 15000

	return float32(hasil)
}

// ================ FUNGSI LEBIH DARI SATU INPUT DAN RETURN PADA GOLANG ================
func hitungMataUang(uang int, mataUang string) (float32, string) {
	
	var hasil float32

	switch  mataUang {
	case "USD":
		hasil = float32(uang) * 15000
	case "JPY":
		hasil = float32(uang) * 103.96
	default:
		hasil = 0
		mataUang = "Tidak Ditemukan"
	}

	return hasil, mataUang
}