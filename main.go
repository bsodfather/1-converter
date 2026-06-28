package main

import "fmt"

func main() {
	const USDtoEURConversionRate float64 = 0.123
	
	const USDtoRUBConversionRate float64 = 9.234

	var amount float64 

	fmt.Println("Введите сумму EUR для конвертации в RUB: ")

	fmt.Scan(&amount)

	conversionResult := amount / USDtoEURConversionRate * USDtoRUBConversionRate

	fmt.Printf("%v EUR = %0.2f RUB\n", amount, conversionResult)
}

