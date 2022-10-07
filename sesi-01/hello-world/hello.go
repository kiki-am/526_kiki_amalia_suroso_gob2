package main

import "fmt"

func main() {
	fmt.Println("Sugeng Rawuh")
	fmt.Println("Sugeng Rawuh", "wonten ing", "Go")
	fmt.Println("============================")
	fmt.Println("Kiki Amalia")
	fmt.Println("Kiki", "Amalia")
	fmt.Print("Kiki Amalia\n")
	fmt.Print("Kiki", " Amalia\n")
	fmt.Print("Kiki", " ", "Amalia\n")
	fmt.Println("============================")
	//Komentar kode
	//fmt.Println("Baris ini tidak akan ditampilkan")
	/*
		Untuk komentar yang panjang
		gunakan bentuk komen seperti ini
	*/

	//VARIABLE
	//with Data Type
	var name string = "Kiki Amalia"
	var age int = 26
	var lulus bool = true
	fmt.Println("Nama saya:", name)
	fmt.Println("Usia saya:", age)
	fmt.Println("Saya sudah lulus kuliah:", lulus)
	fmt.Println("============================")

	//reassign nilai dengan menggunakan tipe data yang sama
	var sekolah string = "UT"
	sekolah = "Universitas Terbuka"
	fmt.Println("saya lulusan dari:", sekolah)
	fmt.Println("============================")

	//without Data Type
	var suka = "makan"
	var sehari = 5
	var gendut = true
	fmt.Printf("%T, %T, %T \n", suka, sehari, gendut)
	fmt.Println("============================")

	//Short Declaration
	pobia := "langsing"
	ngemilSnack := 7
	kurus := false
	fmt.Printf("%T, %T, %T \n", pobia, ngemilSnack, kurus)
	fmt.Println("============================")

	//Multiple Declaration
	var candu, buah, rasa string = "makan", "salak", "sepet"
	beratBadan, tinggiBadan, nomorSepatu := 79, 249, 23
	var tinggi, hitam, langsing = true, false, true
	fmt.Println(candu, buah, rasa)
	fmt.Println(beratBadan, tinggiBadan, nomorSepatu)
	fmt.Println(tinggi, hitam, langsing)
	fmt.Println("============================")

	//Underscore Variable
	var sehat bool
	var baik, hati, ramah = "sangat", 1, false
	_, _, _, _ = sehat, baik, hati, ramah
	fmt.Println("============================")

	//fmt.Printf Usage
	rambut, mata, sipit := "keriting", 2, true
	fmt.Printf("Saya memiliki rambut yang ==> %s, Mata saya berjumlah ==> %d, Bisa jadi mata saya tidak belo ==> %t\n", rambut, mata, sipit)
	fmt.Println("============================")
	telinga, jariTangan, pintar := 2, "sepuluh", true
	fmt.Printf("tipe data pada telinga adalah %T \n", telinga)
	fmt.Printf("tipe data pada jari tangan adalah %T \n", jariTangan)
	fmt.Printf("tipe data pada kepintaran saya adalah %T \n", pintar)
	fmt.Println("============================")

	//DATA TYPE
	//Basic Type : Number
	//integer dan floating number
	var first = 234
	var second = -11
	fmt.Printf("tipe data number dari first adalah: %T\n", first)
	fmt.Printf("tipe data number dari second adalah: %T\n", second)
	fmt.Println("============================")
	var decimalNumber = 2.11
	fmt.Printf("decimal number: %f\n", decimalNumber)
	fmt.Printf("decimal number: %.3f\n", decimalNumber)
	fmt.Println("============================")

	//Basic Type : Boolean
	condition := true
	fmt.Printf("is it permitted? %t\n", condition)
	fmt.Println("============================")

	//Basic Type : String
	pesanSaya := "Jaga kesehatanmu yah! tidak masalah makan banyak asal jangan lupa olahraga"
	pesanSaya = `Kamu tau kan!.
	Aku dan Kamu jadi Kita :)
	"Selamat berjuang!"`
	fmt.Println(pesanSaya)
	fmt.Println("============================")

	//CONSTANT dan OPERATOR
	//operator matematika
	var operator = 2 + 3*3
	var berhitung = (2 + 3) * 3
	var a = 25
	var b = 16
	var hasil = a / b
	fmt.Println(operator, berhitung, hasil)
	fmt.Println("============================")

	//Relational operator
	var bool1 = operator == berhitung
	var bool2 = a > b
	var bool3 = a+b != hasil
	fmt.Printf("hasil dari bool1 adalah: %t \nhasil dari bool2 adalah: %t\nhasil dari bool3 adalah: %t \n", bool1, bool2, bool3)
	fmt.Println("============================")

	//Logical operator
	var logicDan = bool2 && bool3
	logicOr := bool1 || bool2
	logicLawan := !bool3
	fmt.Printf("hasil dari logic dan adalah (%t) \nhasil dari locic or adalah (%t) \nhasil dari logic kebalikan adalah (%t) \n", logicDan, logicOr, logicLawan)
}
