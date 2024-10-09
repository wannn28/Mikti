package main

import (
    "fmt"
    "sync"
    "time"
)

// Fungsi untuk memproses pesanan
func prosesPesanan(pesanan MenuItem, ch chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    select {
    case <-time.After(5 * time.Second): // Timeout untuk pemrosesan pesanan
        ch <- fmt.Sprintf("Proses pesanan timeout: %s", pesanan.Nama)
    default:
        time.Sleep(2 * time.Second) // Simulasi pemrosesan pesanan
        detailEncoded := encodeDetailPesanan(fmt.Sprintf("Pesanan: %s, Jumlah: %d", pesanan.Nama, pesanan.Jumlah))
        ch <- fmt.Sprintf("Pesanan diproses: %s (encoded: %s)", pesanan.Nama, detailEncoded)
    }
}
