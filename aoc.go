package main

import (
	"expensereport"
	"fmt"
	"log"
	"os"
	"readinput"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	expenses := readinput.GetInts(path + "/input/day1/expenses.in")
	sum := 2020
	fmt.Println("1.1: ", expensereport.Find(expenses, sum, 2))
	fmt.Println("1.2: ", expensereport.Find(expenses, sum, 3))
}
