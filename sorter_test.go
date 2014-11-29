package sorter_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/disintegration/sorter"
)

func ExampleCreateIndex() {
	albums := []struct {
		Title string
		Year  int
	}{
		{"Space Oddity", 1969},
		{"Heroes", 1977},
		{"The Man Who Sold the World", 1970},
		{"Never Let Me Down", 1987},
		{"The Next Day", 2013},
		{"Black Tie White Noise", 1993},
		{"Young Americans", 1975},
		{"Aladdin Sane", 1973},
	}

	// Sort the albums by title.
	fmt.Println("Sorted by title:")
	byTitle := sorter.CreateIndex(
		len(albums),
		func(i, j int) bool {
			return albums[i].Title < albums[j].Title
		},
	)
	for _, i := range byTitle {
		fmt.Println(albums[i].Year, albums[i].Title)
	}

	fmt.Println()

	// Sort the albums by year.
	fmt.Println("Sorted by year:")
	byYear := sorter.CreateIndex(
		len(albums),
		func(i, j int) bool {
			return albums[i].Year < albums[j].Year
		},
	)
	for _, i := range byYear {
		fmt.Println(albums[i].Year, albums[i].Title)
	}

	// Output:
	//
	// Sorted by title:
	// 1973 Aladdin Sane
	// 1993 Black Tie White Noise
	// 1977 Heroes
	// 1987 Never Let Me Down
	// 1969 Space Oddity
	// 1970 The Man Who Sold the World
	// 2013 The Next Day
	// 1975 Young Americans
	//
	// Sorted by year:
	// 1969 Space Oddity
	// 1970 The Man Who Sold the World
	// 1973 Aladdin Sane
	// 1975 Young Americans
	// 1977 Heroes
	// 1987 Never Let Me Down
	// 1993 Black Tie White Noise
	// 2013 The Next Day
}

type TestStruct struct {
	A int
	B float64
	C byte
	D string
	E bool
}

const dataSize = 10000

var testData []TestStruct

type TestSliceByA []TestStruct

func (s TestSliceByA) Len() int           { return len(s) }
func (s TestSliceByA) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s TestSliceByA) Less(i, j int) bool { return s[i].A < s[j].A }

func BenchmarkStructStandard(b *testing.B) {
	testData = make([]TestStruct, dataSize)
	for i := 0; i < b.N; i++ {
		for j := range testData {
			testData[j] = TestStruct{A: rand.Intn(dataSize)}
		}
		sort.Sort(TestSliceByA(testData))
	}
}

func BenchmarkStructSorter(b *testing.B) {
	testData = make([]TestStruct, dataSize)
	for i := 0; i < b.N; i++ {
		for j := range testData {
			testData[j] = TestStruct{A: rand.Intn(dataSize)}
		}
		_ = sorter.CreateIndex(len(testData), func(i, j int) bool {
			return testData[i].A < testData[j].A
		})
	}
}
