package main

import "fmt"

func main() {
	// make для оптимизации памяти
tr := make([]string,0,2)
fmt.Println(len(tr),cap(tr))
tr = append(tr, "1")
fmt.Println(len(tr),cap(tr))
tr = append(tr, "2")
fmt.Println(len(tr),cap(tr))
tr = append(tr, "3")
fmt.Println(len(tr),cap(tr))
tr = append(tr, "4")
fmt.Println(len(tr),cap(tr))
tr = append(tr, "5")
fmt.Println(len(tr),cap(tr))
fmt.Println(tr)

	transactions := []float64{}
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
}
balance := calculateBalance(transactions)
fmt.Printf("Ваш баланс: %.2f", balance)
}

func scanTransaction() float64 {
	var transaction float64
		fmt.Print("Введите транзакцию (или n для выхода)")
		fmt.Scan(&transaction)
		return transaction
}

func calculateBalance(transactions []float64) float64 {
	balance := 0.0
  for _,value := range transactions {
    balance += value
	}
	return balance
}