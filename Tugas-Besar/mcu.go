package main
//IF-11-01 kelompok 11
// Nisrina Amalia Iffatunnisa (2311102156)
// Muhammad Dani Ayubi	(2311102003)
// Naufal Luthfi Assary	(2311102125)

import (
	"fmt"
	"strings"
	"time"
)

type PaketMCU struct {
	NamaPaket string
	Harga     int
	Layanan   string
}

type Pasien struct {
	Nama         string
	Usia         int
	PaketDipilih PaketMCU
	TanggalMCU   time.Time
	HasilMCU     string
}

var daftarPaket = []PaketMCU{
	{"Silver", 200000, "Basic Check-Up"},
	{"Gold", 500000, "Basic + Heart Check-Up"},
	{"Platinum", 750000, "Full Body Check-Up"},
	{"Diamond", 1000000, "VIP Full Check-Up"},
}

var daftarPasien []Pasien

func tampilkanMenu() {
	fmt.Println("\n=== Aplikasi Medical Check-Up ===")
	fmt.Println("1. Kelola Data Pasien")
	fmt.Println("2. Kelola Data Paket MCU")
	fmt.Println("3. Kelola Hasil Rekap MCU")
	fmt.Println("4. Pencarian dan Laporan")
	fmt.Println("5. Keluar")
	fmt.Print("Pilih menu: ")
}

func kelolaPasien() {
	fmt.Println("\n=== Kelola Data Pasien ===")
	fmt.Println("1. Tambah Pasien")
	fmt.Println("2. Edit Pasien")
	fmt.Println("3. Hapus Pasien")
	fmt.Println("4. Tampil Pasien")
	fmt.Println("5. Kembali")
	fmt.Print("Pilih menu: ")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		tambahPasien()
	case 2:
		editPasien()
	case 3:
		hapusPasien()
	case 4:
		tampilkanDaftarPasien()
	case 5:
		return
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}

func tambahPasien() {
	var nama string
	var usia, pilihanPaket int
	var tanggal string

	fmt.Print("Masukkan nama pasien: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan usia pasien: ")
	fmt.Scanln(&usia)

	tampilkanPaket()
	fmt.Print("Pilih paket (1-", len(daftarPaket), "): ")
	fmt.Scanln(&pilihanPaket)

	if pilihanPaket < 1 || pilihanPaket > len(daftarPaket) {
		fmt.Println("Pilihan paket tidak valid!")
		return
	}

	fmt.Print("Masukkan tanggal MCU (YYYY-MM-DD): ")
	fmt.Scanln(&tanggal)
	tanggalParsed, err := time.Parse("2006-01-02", tanggal)

	if err != nil {
		fmt.Println("Format tanggal tidak valid!")
		return
	}

	pasien := Pasien{nama, usia, daftarPaket[pilihanPaket-1], tanggalParsed, "Belum ada hasil"}
	daftarPasien = append(daftarPasien, pasien)
	fmt.Println("Data pasien berhasil ditambahkan!")
}

func tampilkanDaftarPasien() {
	if len(daftarPasien) == 0 {
		fmt.Println("Belum ada data pasien.")
		return
	}

	fmt.Printf("%-5s %-20s %-5s %-20s %-15s %-10s\n", "No", "Nama", "Usia", "Paket", "Tanggal MCU", "Hasil")
	for i := 0; i < 80; i++ {
		fmt.Print("-")
	}
	fmt.Println()

	for i := 0; i < len(daftarPasien); i++ {
		pasien := daftarPasien[i]
		fmt.Printf("%-5d %-20s %-5d %-20s %-15s %-10s\n", i+1, pasien.Nama, pasien.Usia, pasien.PaketDipilih.NamaPaket, pasien.TanggalMCU.Format("2006-01-02"), pasien.HasilMCU)
	}
}

func editPasien() {
	tampilkanDaftarPasien()
	fmt.Print("Pilih nomor pasien yang akan diedit: ")
	var index int
	fmt.Scanln(&index)

	if index < 1 || index > len(daftarPasien) {
		fmt.Println("Pilihan tidak valid!")
		return
	}

	pasien := &daftarPasien[index-1]
	fmt.Printf("Nama baru (%s): ", pasien.Nama)
	fmt.Scanln(&pasien.Nama)
	fmt.Printf("Usia baru (%d): ", pasien.Usia)
	fmt.Scanln(&pasien.Usia)

	tampilkanPaket()
	fmt.Print("Pilih paket baru (1-", len(daftarPaket), "): ")
	var pilihanPaket int
	fmt.Scanln(&pilihanPaket)

	if pilihanPaket >= 1 && pilihanPaket <= len(daftarPaket) {
		pasien.PaketDipilih = daftarPaket[pilihanPaket-1]
	}

	fmt.Print("Masukkan tanggal MCU baru (YYYY-MM-DD): ")
	var tanggal string
	fmt.Scanln(&tanggal)
	tanggalParsed, err := time.Parse("2006-01-02", tanggal)

	if err == nil {
		pasien.TanggalMCU = tanggalParsed
	}

	fmt.Println("Data pasien berhasil diperbarui!")
}

func hapusPasien() {
	tampilkanDaftarPasien()
	fmt.Print("Pilih nomor pasien yang akan dihapus: ")
	var index int
	fmt.Scanln(&index)

	if index < 1 || index > len(daftarPasien) {
		fmt.Println("Pilihan tidak valid!")
		return
	}

	daftarPasien = append(daftarPasien[:index-1], daftarPasien[index:]...)
	fmt.Println("Data pasien berhasil dihapus!")
}

// ====== Kelola Data Paket ======
func kelolaPaket() {
	fmt.Println("\n=== Kelola Data Paket ===")
	fmt.Println("1. Tambah Paket")
	fmt.Println("2. Edit Paket")
	fmt.Println("3. Hapus Paket")
	fmt.Println("4. Tampil Paket")
	fmt.Println("5. Kembali")
	fmt.Print("Pilih menu: ")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		tambahPaket()
	case 2:
		editPaket()
	case 3:
		hapusPaket()
	case 4:
		tampilkanPaket()
	case 5:
		return
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}

func tambahPaket() {
	var nama string
	var harga int
	var layanan string

	fmt.Print("Masukkan nama paket: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan harga paket: ")
	fmt.Scanln(&harga)
	fmt.Print("Masukkan layanan paket: ")
	fmt.Scanln(&layanan)

	paket := PaketMCU{nama, harga, layanan}
	daftarPaket = append(daftarPaket, paket)
	fmt.Println("Paket berhasil ditambahkan!")
}

func editPaket() {
	tampilkanPaket()
	fmt.Print("Pilih nomor paket yang akan diedit: ")
	var index int
	fmt.Scanln(&index)

	if index < 1 || index > len(daftarPaket) {
		fmt.Println("Pilihan tidak valid!")
		return
	}

	paket := &daftarPaket[index-1]
	fmt.Printf("Nama baru (%s): ", paket.NamaPaket)
	fmt.Scanln(&paket.NamaPaket)
	fmt.Printf("Harga baru (%d): ", paket.Harga)
	fmt.Scanln(&paket.Harga)
	fmt.Printf("Layanan baru (%s): ", paket.Layanan)
	fmt.Scanln(&paket.Layanan)

	fmt.Println("Paket berhasil diperbarui!")
}

func hapusPaket() {
	tampilkanPaket()
	fmt.Print("Pilih nomor paket yang akan dihapus: ")
	var index int
	fmt.Scanln(&index)

	if index < 1 || index > len(daftarPaket) {
		fmt.Println("Pilihan tidak valid!")
		return
	}

	daftarPaket = append(daftarPaket[:index-1], daftarPaket[index:]...)
	fmt.Println("Paket berhasil dihapus!")
}

func tampilkanPaket() {
	fmt.Println("\nDaftar Paket MCU:")

	fmt.Printf("%-5s %-20s %-10s %s\n", "No", "Nama Paket", "Harga(Rp)", "Layanan")
	fmt.Println(strings.Repeat("-", 50))

	for i, paket := range daftarPaket {
		fmt.Printf("%-5d %-20s %-10d %s\n", i+1, paket.NamaPaket, paket.Harga, paket.Layanan)
	}
}

func rekapHasilMCU() {
	fmt.Println("\n=== Rekap Hasil MCU ===")
	tampilkanDaftarPasien()
	fmt.Print("Pilih nomor pasien untuk memasukkan hasil MCU: ")
	var index int
	fmt.Scanln(&index)

	if index < 1 || index > len(daftarPasien) {
		fmt.Println("Pilihan tidak valid!")
		return
	}

	pasien := &daftarPasien[index-1]
	fmt.Printf("Masukkan hasil MCU untuk pasien %s: ", pasien.Nama)
	fmt.Scanln(&pasien.HasilMCU)
	fmt.Println("Hasil MCU berhasil diperbarui!")
}

func pencarianLaporan() {
	fmt.Println("\n=== Pencarian dan Laporan ===")
	fmt.Println("1. Cari pasien berdasarkan nama")
	fmt.Println("2. Cari pasien berdasarkan paket")
	fmt.Println("3. Cari pasien berdasarkan periode tanggal")
	fmt.Println("4. Tampilkan pasien terurut berdasarkan tanggal MCU")
	fmt.Println("5. Tampilkan pasien terurut berdasarkan paket MCU")
	fmt.Println("6. Laporan Pemasukkan")
	fmt.Println("7. Kembali")
	fmt.Print("Pilih menu: ")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		cariPasienBerdasarkanNama()
	case 2:
		cariPasienBerdasarkanPaket()
	case 3:
		cariPasienBerdasarkanPeriode()
	case 4:
		urutkanPasienBerdasarkanTanggal()
	case 5:
		urutkanPasienBerdasarkanPaket()
	case 6:
		laporanPemasukanPeriode()
	case 7:
		return
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}

func cariPasienBerdasarkanNama() {
	fmt.Print("Masukkan nama pasien: ")
	var nama string
	fmt.Scanln(&nama)

	fmt.Println("\nHasil pencarian:")
	found := false
	for i := 0; i < len(daftarPasien); i++ {
		if daftarPasien[i].Nama == nama {
			fmt.Printf("%-5d %-20s %-5d %-20s %-15s %-10s\n", i+1, daftarPasien[i].Nama, daftarPasien[i].Usia, daftarPasien[i].PaketDipilih.NamaPaket, daftarPasien[i].TanggalMCU.Format("2006-01-02"), daftarPasien[i].HasilMCU)
			found = true
		}
	}
	if !found {
		fmt.Println("Pasien tidak ditemukan.")
	}
}

func cariPasienBerdasarkanPaket() {
	urutkanPasienBerdasarkanPaket()

	tampilkanPaket()
	fmt.Print("Pilih nomor paket: ")
	var pilihan int
	fmt.Scanln(&pilihan)

	if pilihan < 1 || pilihan > len(daftarPaket) {
		fmt.Println("Pilihan tidak valid!")
		return
	}

	namaPaketDicari := daftarPaket[pilihan-1].NamaPaket
	low, high := 0, len(daftarPasien)-1
	found := false

	for low <= high {
		mid := (low + high) / 2
		if daftarPasien[mid].PaketDipilih.NamaPaket == namaPaketDicari {
			for i := mid; i >= 0 && daftarPasien[i].PaketDipilih.NamaPaket == namaPaketDicari; i-- {
				cetakPasien(i)
				found = true
			}
			for i := mid + 1; i < len(daftarPasien) && daftarPasien[i].PaketDipilih.NamaPaket == namaPaketDicari; i++ {
				cetakPasien(i)
				found = true
			}
			break
		} else if daftarPasien[mid].PaketDipilih.NamaPaket < namaPaketDicari {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Println("Tidak ada pasien yang memilih paket ini.")
	}
}

func cetakPasien(i int) {
	fmt.Printf("%-5d %-20s %-5d %-20s %-15s %-10s\n", i+1, daftarPasien[i].Nama, daftarPasien[i].Usia, daftarPasien[i].PaketDipilih.NamaPaket, daftarPasien[i].TanggalMCU.Format("2006-01-02"), daftarPasien[i].HasilMCU)
}

func cariPasienBerdasarkanPeriode() {
	fmt.Print("Masukkan tanggal awal (YYYY-MM-DD): ")
	var awal string
	fmt.Scanln(&awal)

	fmt.Print("Masukkan tanggal akhir (YYYY-MM-DD): ")
	var akhir string
	fmt.Scanln(&akhir)

	var awalParsed time.Time
	var akhirParsed time.Time
	var err error

	awalParsed, err = time.Parse("2006-01-02", awal)
	if err != nil {
		fmt.Println("Format tanggal awal tidak valid!")
		return
	}

	akhirParsed, err = time.Parse("2006-01-02", akhir)
	if err != nil {
		fmt.Println("Format tanggal akhir tidak valid!")
		return
	}

	fmt.Println("\nHasil pencarian:")
	fmt.Printf("%-5s %-20s %-5s %-20s %-15s %-10s\n", "No", "Nama", "Usia", "Paket", "Tanggal MCU", "Hasil")
	for i := 0; i < 80; i++ {
		fmt.Print("-")
	}
	fmt.Println()

	found := false
	for i := 0; i < len(daftarPasien); i++ {
		tanggal := daftarPasien[i].TanggalMCU
		if !tanggal.Before(awalParsed) && !tanggal.After(akhirParsed) {
			fmt.Printf("%-5d %-20s %-5d %-20s %-15s %-10s\n", i+1, daftarPasien[i].Nama, daftarPasien[i].Usia, daftarPasien[i].PaketDipilih.NamaPaket, tanggal.Format("2006-01-02"), daftarPasien[i].HasilMCU)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada pasien dalam periode ini.")
	}
}

func urutkanPasienBerdasarkanTanggal() {
	var pilihan int
	fmt.Println("\nPilih urutan:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Print("Pilih (1/2): ")
	fmt.Scanln(&pilihan)

	if pilihan != 1 && pilihan != 2 {
		fmt.Println("Pilihan tidak valid! Harap pilih 1 atau 2.")
		return
	}

	isAscending := pilihan == 1

	if len(daftarPasien) == 0 {
		fmt.Println("Tidak ada data pasien untuk diurutkan.")
		return
	}

	if isAscending {
		fmt.Println("\nDaftar pasien terurut berdasarkan tanggal MCU (Ascending):")
	} else {
		fmt.Println("\nDaftar pasien terurut berdasarkan tanggal MCU (Descending):")
	}

	for i := 0; i < len(daftarPasien)-1; i++ {
		for j := i + 1; j < len(daftarPasien); j++ {
			if (isAscending && daftarPasien[i].TanggalMCU.After(daftarPasien[j].TanggalMCU)) ||
				(!isAscending && daftarPasien[i].TanggalMCU.Before(daftarPasien[j].TanggalMCU)) {
				// Swap elemen
				temp := daftarPasien[i]
				daftarPasien[i] = daftarPasien[j]
				daftarPasien[j] = temp
			}
		}
	}

	tampilkanDaftarPasien()
}

func urutkanPasienBerdasarkanPaket() {
	var order string
	fmt.Println("\nPilih urutan:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Print("Pilih (1/2): ")
	fmt.Scanln(&order)

	if order != "1" && order != "2" {
		fmt.Println("Pilihan tidak valid!")
		return
	}

	ascending := order == "1"

	for i := 1; i < len(daftarPasien); i++ {
		key := daftarPasien[i]
		j := i - 1

		for j >= 0 && ((ascending && daftarPasien[j].PaketDipilih.NamaPaket > key.PaketDipilih.NamaPaket) ||
			(!ascending && daftarPasien[j].PaketDipilih.NamaPaket < key.PaketDipilih.NamaPaket)) {
			daftarPasien[j+1] = daftarPasien[j]
			j = j - 1
		}
		daftarPasien[j+1] = key
	}
	fmt.Println("\nDaftar pasien terurut berdasarkan paket MCU:")
	tampilkanDaftarPasien()
}

func laporanPemasukanPeriode() {
	fmt.Print("Masukkan tanggal awal (YYYY-MM-DD): ")
	var awal string
	fmt.Scanln(&awal)
	fmt.Print("Masukkan tanggal akhir (YYYY-MM-DD): ")
	var akhir string
	fmt.Scanln(&akhir)

	awalParsed, err1 := time.Parse("2006-01-02", awal)
	akhirParsed, err2 := time.Parse("2006-01-02", akhir)

	if err1 != nil || err2 != nil {
		fmt.Println("Format tanggal tidak valid!")
		return
	}

	totalPemasukan := 0
	for i := 0; i < len(daftarPasien); i++ {
		tanggal := daftarPasien[i].TanggalMCU
		if !tanggal.Before(awalParsed) && !tanggal.After(akhirParsed) {
			totalPemasukan += daftarPasien[i].PaketDipilih.Harga
		}
	}

	fmt.Printf("\nTotal pemasukan dari %s hingga %s adalah: Rp%d\n", awalParsed.Format("2006-01-02"), akhirParsed.Format("2006-01-02"), totalPemasukan)
}

func main() {
	for {
		tampilkanMenu()
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			kelolaPasien()
		case 2:
			kelolaPaket()
		case 3:
			rekapHasilMCU()
		case 4:
			pencarianLaporan()
		case 5:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}