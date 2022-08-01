package main

import (
	"fmt"
)

func main() {

	fmt.Println("Çarpanlarına ayırmak istediğiniz sayıyı girin!")
	var a int
	fmt.Scan(&a)

	for i := 1; i <= a; i += 1 {

		if a%i == 0 {
			fmt.Println(i, "*", a/i, "=", a)
		}

	}

}
