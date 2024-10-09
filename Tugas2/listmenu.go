package main

import (
	"fmt"
	"strings"
)

// Struct untuk representasi item menu
type MenuItem struct {
	Nama   string
	Harga  float64
	Jumlah int
}

// Struct untuk daftar menu
type PesananRestoran struct {
	DaftarMenu []MenuItem
}

// Interface untuk mengelola pesanan
type KelolaPesanan interface {
	TambahItem(item MenuItem)
	UbahItem(nama string, field string, newValue interface{})
	HapusItem(nama string)
}

func (pr *PesananRestoran) TambahItem(item MenuItem) {
	pr.DaftarMenu = append(pr.DaftarMenu, item)
	fmt.Println("Item berhasil ditambahkan:", item.Nama)
}

func (pr *PesananRestoran) UbahItem(nama string, field string, newValue interface{}) {
	for i, item := range pr.DaftarMenu {
		if strings.EqualFold(strings.ReplaceAll(item.Nama, " ", ""), strings.ReplaceAll(nama, " ", "")) {
			switch field {
			case "harga":
				if value, ok := newValue.(float64); ok {
					pr.DaftarMenu[i].Harga = value
					fmt.Printf("Harga item %s berhasil diupdate menjadi Rp%.2f\n", nama, value)
				} else {
					fmt.Println("Kesalahan: Harga harus berupa angka desimal.")
				}
			case "jumlah":
				if value, ok := newValue.(int); ok {
					pr.DaftarMenu[i].Jumlah = value
					fmt.Printf("Jumlah item %s berhasil diupdate menjadi %d\n", nama, value)
				} else {
					fmt.Println("Kesalahan: Jumlah harus berupa angka bulat.")
				}
			default:
				fmt.Println("Field yang diminta tidak ditemukan.")
			}
			return
		}
	}
	fmt.Println("Item tidak ditemukan:", nama)
}

func (pr *PesananRestoran) HapusItem(nama string) {
	for i, item := range pr.DaftarMenu {
		if strings.EqualFold(strings.ReplaceAll(item.Nama, " ", ""), strings.ReplaceAll(nama, " ", "")) {
			pr.DaftarMenu = append(pr.DaftarMenu[:i], pr.DaftarMenu[i+1:]...)
			fmt.Println("Item berhasil dihapus:", nama)
			return
		}
	}
	fmt.Println("Item tidak ditemukan:", nama)
}
