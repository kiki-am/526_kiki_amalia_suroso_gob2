package main

import "fmt"

func main() {
	//Temporary variable (if else)
	var currentYear = 2022

	if age := currentYear - 1998; age < 18 {
		fmt.Println("kamu belum boleh membuat kartu sim")
		/*fmt.Printf("kamu belum boleh membuat kartu sim %d \n", age)
		fmt.Printf("kamu belum boleh membuat kartu sim %T \n", age)*/
	} else {
		fmt.Println("kamu sudah bisa membuat kartu sim")
		/*fmt.Printf("kamu sudah boleh membuat kartu sim %d \n", age)
		fmt.Printf("kamu sudah boleh membuat kartu sim %T \n", age)*/
	}
	fmt.Println("======================================================")

	//Switch
	var score = 8

	switch score {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("not bad")
	}
	fmt.Println("======================================================")

	//Switch with Relational Operator
	var nilai = 6

	switch {
	case nilai == 8:
		fmt.Println("Perfect")
	case (nilai < 8) && (nilai > 3):
		fmt.Println("not bad")
	default:
		{
			fmt.Println("Study harder")
			fmt.Println("you need to learn more")
		}
	}
	fmt.Println("======================================================")

	//Falltrough Keyword
	var value = 6

	switch {
	case value == 8:
		fmt.Println("Perfect")
	case (value < 8) && (value > 3):
		fmt.Println("not bad")
		fallthrough
	case value < 0:
		fmt.Println("it's oke, but please study harder")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you don't have a good score yet")
		}
	}
	fmt.Println("======================================================")

	//Nested condition
	var hasil = 10

	if hasil > 7 {
		switch hasil {
		case 10:
			fmt.Println("Perfect")
		default:
			fmt.Println("nice")
		}
	} else {
		if hasil == 5 {
			fmt.Println("not bad")
		} else if hasil == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if hasil == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
