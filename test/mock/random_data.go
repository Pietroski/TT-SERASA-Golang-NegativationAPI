package mock

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type sUTF8Range struct {
	Min int
	Max int
}

type sUTF8RangesDetail struct {
	Description string
	UTF8Range sUTF8Range
}

var (
	RandomData iRandData = &sRandData{}

	UTF8MapRanges = map[int]sUTF8RangesDetail{
		0: {
			Description: "Numbers",
			UTF8Range: sUTF8Range{48, 57},
		},
		1: {
			Description: "Upper case letters",
			UTF8Range: sUTF8Range{65, 90},
		},
		2: {
			Description: "Lower case letters",
			UTF8Range: sUTF8Range{97, 122},
		},
	}
)

type iRandData interface {
	GenerateRandomDocument(length int) string
	GenerateCompanyName() string
	GenerateRandomContract() string
	GenerateRandomDebt(min int, max int) float64
	GenerateRandomDate() time.Time
	GenerateRandomID(limit int64) int64

	genRandFloatNum(min int, max int) float64
	genRandNum(min int, max int) int
	genRandNumDiffFrom1() int
	genRandIntSlice(length int, min int, max int) []int
	stringifyIntSlice(intSlice []int) string

	generateRandomString(length int, min int, max int) string
}

type sRandData struct {}

func (rd *sRandData) GenerateRandomDocument(length int) string {
	min, max := 0, 9
	ii := rd.genRandIntSlice(length, min, max)
	str := rd.stringifyIntSlice(ii)

	return str
}

func (rd *sRandData) GenerateCompanyName() string {
	var strBuilder strings.Builder
	min, max := 0, 1

	rLen := rd.genRandNum(3, 4)
	rStr := rd.generateRandomString(rLen, min, max)

	strBuilder.WriteString(rStr)

	// TODO: chose randomly the company type such as S.A., L.T.D.A and so on...
	strBuilder.WriteString(" S.A.")

	companyName := strBuilder.String()
	return companyName
}

// GenerateRandomContract generates a new random contract hash
// TODO: Decouple and improve this method respecting SOLID principles
func (rd *sRandData) GenerateRandomContract() string {
	var strBuilder strings.Builder

	// temporary method
	// TODO: Implement a method to generate a random number between 2 different number ranges
	randNum := rd.genRandNumDiffFrom1()

	// Generate the first document part
	for i := 0; i < 8; i++ {
		randNum = rd.genRandNumDiffFrom1()
		chars := UTF8MapRanges[randNum]
		rn := rd.genRandNum(chars.UTF8Range.Min, chars.UTF8Range.Max)
		char := rune(rn)
		strBuilder.WriteString(string(char))
	}

	strBuilder.WriteString("-")
	// Generate the second document part
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			randNum = rd.genRandNumDiffFrom1()
			chars := UTF8MapRanges[randNum]
			rn := rd.genRandNum(chars.UTF8Range.Min, chars.UTF8Range.Max)
			char := rune(rn)
			strBuilder.WriteString(string(char))
		}
		strBuilder.WriteString("-")
	}

	// Generate the third document part
	for i := 0; i < 11; i++ {
		randNum = rd.genRandNumDiffFrom1()
		chars := UTF8MapRanges[randNum]
		rn := rd.genRandNum(chars.UTF8Range.Min, chars.UTF8Range.Max)
		char := rune(rn)
		strBuilder.WriteString(string(char))
	}

	contract := strBuilder.String()
	return contract
}

func (rd *sRandData) GenerateRandomDebt(min int, max int) float64 {
	rfn := rd.genRandFloatNum(min, max)
	return rfn
}

func (rd *sRandData) GenerateRandomID(limit int64) int64 {
	min := int64(1)
	return rand.Int63n(limit - min) + min
}

func (rd *sRandData) generateRandomString(length int, min int, max int) string {
	var strBuilder strings.Builder

	for i := 0; i < length; i++ {
		randType := rd.genRandNum(min, max)
		chars := UTF8MapRanges[randType]
		rn := rd.genRandNum(chars.UTF8Range.Min, chars.UTF8Range.Max)
		char := rune(rn)
		strBuilder.WriteString(string(char))
	}

	str := strBuilder.String()
	return str
}

func (rd *sRandData) stringifyIntSlice(intSlice []int) string {
	var strBuilder strings.Builder

	for _, n := range intSlice {
		strBuilder.WriteString(strconv.Itoa(n))
	}

	stringifiedIntSlice := strBuilder.String()
	return stringifiedIntSlice
}

func (rd *sRandData) genRandIntSlice(length int, min int, max int) []int {
	intSlice := make([]int, length)
	for i := 0; i < length; i++ {
		intSlice[i] = rd.genRandNum(min, max)
	}
	return intSlice
}

func (rd *sRandData) genRandNumDiffFrom1() int {
	min, max := 0, 2
	randNum := rd.genRandNum(min, max)

	for randNum == 1 {
		randNum = rd.genRandNum(min, max)
	}

	return randNum
}

func (rd *sRandData) genRandNum(min int, max int) int {
	random := rand.Intn(max-min+1) + min
	return random
}

func (rd *sRandData) genRandFloatNum(min int, max int) float64 {
	minF, maxF := float64(min), float64(max)

	r := minF + rand.Float64() * (maxF - minF)
	fmtR := fmt.Sprintf("%.2f", r)

	convR, err := strconv.ParseFloat(fmtR, 64)
	if err != nil {
		panic(err)
	}

	return convR
}

func (rd *sRandData) GenerateRandomDate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix() // time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0) // .Format("2006-01-02T15:04:05Z")
}