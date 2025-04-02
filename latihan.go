package main
import "fmt"

func formatRupiah(n int) {
	var hasil string
	count := 0
	for n > 0 {
		digit := n % 10
		n /= 10
		hasil = fmt.Sprintf("%d%s", digit, hasil)

		count++
		if count%3 == 0 && n > 0 {
			hasil = fmt.Sprintf(".%s", hasil) 
    	}
	}

	fmt.Printf("Rp %s\n", hasil)
}

func main() {
	var num int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&num)
	formatRupiah(num)
}