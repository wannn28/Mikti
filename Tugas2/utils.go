package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// Fungsi untuk membersihkan layar terminal
func bersihkanLayar() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Fungsi untuk mencetak garis pemisah
func cetakGaris() {
	fmt.Println(strings.Repeat("=", 40))
}

// Fungsi untuk mencetak header dengan garis pemisah
func cetakHeader(judul string) {
	cetakGaris()
	fmt.Printf("%s\n", judul)
	cetakGaris()
}

// Fungsi validasi input untuk angka
func validasiInputNumerik(input string) (float64, error) {
	input = strings.TrimSpace(input)          // Menghapus spasi
	re := regexp.MustCompile(`^\d+(\.\d+)?$`) // Mengizinkan angka dan angka desimal
	if !re.MatchString(input) {
		return 0, fmt.Errorf("input tidak valid, harus berupa angka")
	}
	parsedValue, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("input tidak dapat dikonversi menjadi angka")
	}
	return parsedValue, nil
}

// Fungsi untuk encoding informasi pesanan ke dalam base64
func encodeDetailPesanan(detail string) string {
	return base64.StdEncoding.EncodeToString([]byte(detail))
}

// Fungsi untuk menangani error menggunakan recover
func safeInput(handler func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error:", r)
		}
	}()
	handler()
}
