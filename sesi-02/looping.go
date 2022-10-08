package main

import "fmt"

func main() {
	//First Way of Looping
	for i := 0; i < 3; i++ {
		fmt.Println("Angka", i)
	}
	fmt.Println("======================================================")

	//Second Way of Looping
	var u = 0

	for u < 3 {
		fmt.Println("Angka", u)
		u++
	}
	fmt.Println("======================================================")

	//Third Way of Looping
	var e = 0

	for {
		fmt.Println("Angka", e)

		e++
		if e == 3 {
			break

		}
	}
	fmt.Println("======================================================")

}
