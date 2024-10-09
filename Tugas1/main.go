package main

import (
	"GolangAplication/function"
	"fmt"
)

func main() {
	var menu int

	for {
		fmt.Println("Aplikasi Sederhana berbasis CLI")
		fmt.Println("Menu Utama :")
		fmt.Println("1. Tampilkan Hello World")
		fmt.Println("2. Operasi Matematika Sederhana")
		fmt.Println("3. Simpan dan tampilkan data pengguna")
		fmt.Println("4. Hitung Faktorial (Rekursif Function)")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&menu)

		switch menu {
		case 1:
			function.ShowHelloWorld()
		case 2:
			function.SimpleMathOperation()
		case 3:
			function.SaveAndShowUserData()
		case 4:
			function.CalculateFactorial()
		case 5:
			fmt.Println("Terima kasih! Program selesai.")
			return
		default:
			fmt.Println("Menu tidak valid, coba lagi.")
		}
		fmt.Println()
	}
}
