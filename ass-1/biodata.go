package main

import (
	"fmt"
	"os"
)

type Biodata struct {
	No, Name, Address, Profession, Reason string
}

func (bio Biodata) printBiodata() {
	fmt.Println("No. Absen	:", bio.No)
	fmt.Println("Nama		:", bio.Name)
	fmt.Println("Alamat		:", bio.Address)
	fmt.Println("Pekerjaan	:", bio.Profession)
	fmt.Println("Alasan memilih kelas Golang	:", bio.Reason)
}

func main() {
	var Absen1 Biodata
	Absen1.No = "1"
	Absen1.Name = "Kiki"
	Absen1.Address = "Bekasi"
	Absen1.Profession = "Belum bekerja"
	Absen1.Reason = "Mempelajari skill baru"

	var Absen2 Biodata
	Absen2.No = "2"
	Absen2.Name = "Amalia"
	Absen2.Address = "Gombong"
	Absen2.Profession = "Mahasiswa"
	Absen2.Reason = "Memperdalam ilmu pemrogramman"

	var Absen3 Biodata
	Absen3.No = "3"
	Absen3.Name = "Suroso"
	Absen3.Address = "Rowokele"
	Absen3.Profession = "Penyuluh Pertanian Lapangan"
	Absen3.Reason = "Membuat e-commerce produk pertanian kelompok tani yang saya suluh"

	var Absen4 Biodata
	Absen4.No = "4"
	Absen4.Name = "Andre"
	Absen4.Address = "Sukomulyo"
	Absen4.Profession = "Pelajar"
	Absen4.Reason = "Belajar ngoding"

	var Absen5 Biodata
	Absen5.No = "5"
	Absen5.Name = "Puji"
	Absen5.Address = "Makassar"
	Absen5.Profession = "PNS"
	Absen5.Reason = "Mempelajari ilmu/hal baru"

	var argsRaw = os.Args
	var argsProgram = argsRaw[1]

	if argsRaw == nil {
		fmt.Println("ketikan angka absen setelah spasi dibelakang biodata.go")
	} else if argsProgram == "1" {
		Absen1.printBiodata()
	} else if argsProgram == "2" {
		Absen2.printBiodata()
	} else if argsProgram == "3" {
		Absen3.printBiodata()
	} else if argsProgram == "4" {
		Absen4.printBiodata()
	} else if argsProgram == "5" {
		Absen5.printBiodata()
	} else {
		fmt.Println("tolong ketikan angka n yang memnuhi=(1<= n <=5) setelah spasi, tepat dibelakang perintah go run biodata.go")
	}
}
