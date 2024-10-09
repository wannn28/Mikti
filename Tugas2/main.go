package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	defer fmt.Println("Aplikasi selesai")
	pesananRestoran := &PesananRestoran{}

	// Menambahkan beberapa item ke menu
	pesananRestoran.DaftarMenu = []MenuItem{
		{Nama: "Nasi Goreng", Harga: 25000.00},
		{Nama: "Mie Goreng", Harga: 22000.00},
		{Nama: "Ayam Bakar", Harga: 30000.00},
	}

	// Menampilkan menu yang tersedia
	bersihkanLayar()
	cetakHeader("    SELAMAT DATANG DI RESTO KAMI.")
	fmt.Println("Menu yang tersedia:")
	for _, item := range pesananRestoran.DaftarMenu {
		fmt.Printf("- %s: Rp.%.2f\n", item.Nama, item.Harga)
	}
	cetakGaris()

	// Menerima pesanan dari pengguna
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Masukkan perintah (tambah/ubah/hapus/selesai): ")
		scanner.Scan()
		perintah := strings.TrimSpace(strings.ToLower(scanner.Text()))

		if perintah == "selesai" {
			break
		}

		switch perintah {
		case "tambah":
			fmt.Print("Masukkan nama item: ")
			scanner.Scan()
			namaItem := strings.TrimSpace(scanner.Text())

			fmt.Print("Masukkan jumlah: ")
			scanner.Scan()
			jumlahPesanan := scanner.Text()

			safeInput(func() {
				jumlah, err := strconv.Atoi(jumlahPesanan)
				if err != nil {
					panic("Jumlah tidak valid.")
				}

				var itemDitemukan *MenuItem
				for i := range pesananRestoran.DaftarMenu {
					if strings.EqualFold(strings.ReplaceAll(pesananRestoran.DaftarMenu[i].Nama, " ", ""), strings.ReplaceAll(namaItem, " ", "")) {
						itemDitemukan = &pesananRestoran.DaftarMenu[i]
						itemDitemukan.Jumlah += jumlah
						fmt.Printf("Jumlah item %s berhasil ditambah menjadi %d\n", itemDitemukan.Nama, itemDitemukan.Jumlah)
						break
					}
				}

				if itemDitemukan == nil {
					fmt.Println("Item tidak ditemukan. Pastikan nama yang dimasukkan sesuai dengan yang ada di menu.")
				}
			})

		case "ubah":
			fmt.Print("Masukkan nama item yang ingin diubah jumlahnya: ")
			scanner.Scan()
			namaItem := strings.TrimSpace(scanner.Text())

			fmt.Print("Masukkan jumlah baru: ")
			scanner.Scan()
			inputJumlahBaru := scanner.Text()

			safeInput(func() {
				jumlahBaru, err := strconv.Atoi(inputJumlahBaru)
				if err != nil {
					panic("Jumlah tidak valid.")
				}

				var itemDitemukan *MenuItem
				for i := range pesananRestoran.DaftarMenu {
					if strings.EqualFold(strings.ReplaceAll(pesananRestoran.DaftarMenu[i].Nama, " ", ""), strings.ReplaceAll(namaItem, " ", "")) {
						itemDitemukan = &pesananRestoran.DaftarMenu[i]
						if jumlahBaru >= 0 {
							itemDitemukan.Jumlah = jumlahBaru
							fmt.Printf("Jumlah item %s berhasil diubah menjadi %d\n", itemDitemukan.Nama, itemDitemukan.Jumlah)
						} else {
							fmt.Println("Jumlah harus bernilai positif atau nol.")
						}
						break
					}
				}

				if itemDitemukan == nil {
					fmt.Println("Item tidak ditemukan. Pastikan nama yang dimasukkan sesuai dengan yang ada di pesanan.")
				}
			})

		case "hapus":
			fmt.Print("Masukkan nama item yang ingin dihapus: ")
			scanner.Scan()
			namaItem := strings.TrimSpace(scanner.Text())

			pesananRestoran.HapusItem(namaItem)

		default:
			fmt.Println("Perintah tidak dikenal. Silakan masukkan 'tambah', 'ubah', 'hapus', atau 'selesai'.")
		}

		// Menampilkan pesanan sementara
		fmt.Println("\nPesanan Sementara Anda:")
		for _, item := range pesananRestoran.DaftarMenu {
			if item.Jumlah > 0 {
				fmt.Printf("- %s (x%d)\n", item.Nama, item.Jumlah)
			}
		}
		cetakGaris()
	}

	// Menampilkan pesanan terakhir
	bersihkanLayar()
	cetakHeader("          DETAIL PESANAN ANDA")
	for _, item := range pesananRestoran.DaftarMenu {
		if item.Jumlah > 0 {
			fmt.Printf("- %s (x%d)\n", item.Nama, item.Jumlah)
		}
	}
	cetakGaris()
	totalHarga := 0.0
	for _, item := range pesananRestoran.DaftarMenu {
		if item.Jumlah > 0 {
			totalHarga += float64(item.Jumlah) * item.Harga
		}
	}
	fmt.Printf("Total Harga: Rp%.2f\n", totalHarga)
	cetakGaris()

	// Menggunakan goroutine untuk memproses pesanan
	ch := make(chan string, len(pesananRestoran.DaftarMenu))
	var wg sync.WaitGroup

	for _, item := range pesananRestoran.DaftarMenu {
		if item.Jumlah > 0 {
			wg.Add(1)
			go prosesPesanan(item, ch, &wg)
		}
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	// Menerima hasil dari goroutine
	for msg := range ch {
		fmt.Println(msg)
	}
	cetakGaris()

	// Pembayaran
	fmt.Print("Masukkan jumlah yang dibayar: ")
	scanner.Scan()
	inputPembayaran := scanner.Text()

	safeInput(func() {
		pembayaran, err := validasiInputNumerik(inputPembayaran)
		if err != nil {
			panic("Input pembayaran tidak valid")
		}

		if pembayaran < totalHarga {
			fmt.Println("Jumlah yang dibayar kurang.")
		} else {
			kembalian := pembayaran - totalHarga
			fmt.Printf("Jumlah yang dibayar valid. Kembalian: Rp%.2f\n", kembalian)
		}
	})
	cetakGaris()

	fmt.Println("Memproses pesanan di goroutine lain...")
}
