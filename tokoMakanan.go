package main

import (
	"fmt"
	"strings"
)

func main() {
	var pilih string
	var makanan []string
	var yt string
	var x = true
	for x {
		fmt.Println("TOKO MAKANAN INDONESIA")
		fmt.Println("===========================")
		fmt.Println("Tahu")
		fmt.Println("Tempe")
		fmt.Println("Sambal")
		fmt.Println("Nasi")
		fmt.Println("Lele")
		fmt.Println("Ayam")
		fmt.Println("============================")
		fmt.Print("Masukkan Pilihan Menu Anda Dalam Huruf (eg: tahu,) : ")
		fmt.Scan(&pilih)

		pilih = strings.ToLower(pilih)
		if pilih == "tahu" || pilih == "tempe" || pilih == "sambal" || pilih == "nasi" || pilih == "lele" || pilih == "ayam" {
			makanan = append(makanan, pilih)
		} else {
			fmt.Println("\nPesanan yang diinputkan tidak terdaftar dalam menu.")
		}

		fmt.Print("\nLanjut memesan? (Y/T): ")
		fmt.Scan(&yt)

		yt = strings.ToUpper(yt)
		if yt == "T" {
			x = false
		}
	}
	if len(makanan) <= 0 {
		fmt.Printf("\nTidak ada pesanan yang diinputkaan")
	} else {
		for i := 0; i < len(makanan); i++ {
			fmt.Printf("Pesanan anda \t: %s\n", makanan[i])
		}
		fmt.Println("\nTerimakasih atas Pesanannya.")
	}

}
