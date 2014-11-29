# Sorter

Package sorter provides a function to create a sorted index of data elements.
It is useful for iteration over slices of arbitrary data structures in sorted order.

## Example

```go
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
```
