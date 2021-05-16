package mock

import (
	"fmt"
	"testing"
)

func TestGenerateCompanyName(t *testing.T) {
	fmt.Println("TestGenerateCompanyName")

	cn := RandomData.GenerateCompanyName()
	fmt.Println(cn)
}

func TestGenerateRandomContract(t *testing.T) {
	fmt.Println("TestGenerateRandomContract")

	c := RandomData.GenerateRandomContract()
	fmt.Println(c)
}

func TestGenerateRandomDebt(t *testing.T) {
	fmt.Println("TestGenerateRandomDebt")

	min, max := 1, 5_000
	rfn := RandomData.GenerateRandomDebt(min, max)
	fmt.Println(rfn)
}

func TestGenerateRandomDate(t *testing.T) {
	fmt.Println("TestGenerateRandomDate")

	randDate := RandomData.GenerateRandomDate()
	fmt.Println(randDate)
}

func TestGenerateRandomString(t *testing.T) {
	fmt.Println("TestGenerateRandomString")

	min, max := 0, 1
	rLen := RandomData.genRandNum(3, 4)
	rstr := RandomData.generateRandomString(rLen, min, max)
	fmt.Println("Random string", rstr)
}

func TestGenRandNum(t *testing.T) {
	fmt.Println("TestGenRandNum")
	min, max := 0, 10

	rn1 := RandomData.genRandNum(min, max)
	rn2 := RandomData.genRandNum(min, max)

	fmt.Println(rn1, rn2)
}
