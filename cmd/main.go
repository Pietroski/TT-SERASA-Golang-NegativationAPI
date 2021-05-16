package main

import (
	"fmt"
	"github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/test/mock"
)

func init() {
	//
}

func main() {
	d := mock.RandomData.GenerateRandomDocument(14)
	fmt.Println("Random Document ->", d, len(d))
}
