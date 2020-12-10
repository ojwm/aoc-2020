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
	fmt.Println(expensereport.Find(expenses, 2020))
}
