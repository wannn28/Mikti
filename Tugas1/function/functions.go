package function

import (
	"fmt"
)

func ShowHelloWorld() {
	fmt.Println("Hello World!")
}

func SimpleMathOperation() {
	var num1, num2 int
	operations := []string{"+", "-", "*", "/", "%"}
	var results []float64

	fmt.Println("Operasi Matematika Sederhana")
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scan(&num1)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scan(&num2)

	results = append(results, float64(num1+num2))
	results = append(results, float64(num1-num2))
	results = append(results, float64(num1*num2))

	if num2 != 0 {
		results = append(results, float64(num1)/float64(num2))
		results = append(results, float64(num1%num2))
	} else {
		results = append(results, 0)
		results = append(results, 0)
	}

	for i, operation := range operations {
		if operation == "/" || operation == "%" {
			if num2 == 0 {
				fmt.Printf("Operasi %d %s %d tidak valid (pembagi nol)\n", num1, operation, num2)
				continue
			}
		}
		fmt.Printf("Hasil dari %d %s %d = %.2f\n", num1, operation, num2, results[i])
	}
}

func SaveAndShowUserData() {
	var name, address string
	var age int
	var userData [3]string

	fmt.Println("Simpan dan Tampilkan Data Pengguna")
	fmt.Print("Masukkan nama: ")
	fmt.Scan(&name)
	userData[0] = name

	fmt.Print("Masukkan umur: ")
	fmt.Scan(&age)
	userData[1] = fmt.Sprintf("%d", age)

	fmt.Print("Masukkan alamat: ")
	fmt.Scan(&address)
	userData[2] = address

	fmt.Println("Data Pengguna:")
	fmt.Printf("Nama: %s\nUmur: %s\nAlamat: %s\n", userData[0], userData[1], userData[2])
}


func Factorial(numbers ...int) []int {
	var results []int

	var fact func(n int) int
	fact = func(n int) int {
		if n == 0 {
			return 1
		}
		return n * fact(n-1)
	}

	for _, num := range numbers {
		results = append(results, fact(num))
	}
	return results
}

func CalculateFactorial() {
	var count, number int

	fmt.Print("Berapa banyak angka yang ingin dihitung faktorialnya? ")
	fmt.Scan(&count)

	var numbers []int
	for i := 0; i < count; i++ {
		fmt.Printf("Masukkan angka ke-%d: ", i+1)
		fmt.Scan(&number)
		numbers = append(numbers, number)
	}


	factorials := Factorial(numbers...)


	for i, fact := range factorials {
		fmt.Printf("Faktorial dari %d adalah %d\n", numbers[i], fact)
	}
}
