package main

import (
	"fmt"
)

func Faktöriyel(X int) int {
	sonuc := 1

	if X >= 0 {

		for i := 1; i <= X; i++ {
			sonuc *= i

		}

	} else {
		println("Sıfırdan küçük değer girildi!")

	}

	return sonuc

}
func main() {
	var a int
	fmt.Println("Faktöriyelini almak istediğiniz sayıyı girin:")
	fmt.Scan(&a)
	fmt.Println("Sonuç:", Faktöriyel(a))

}
