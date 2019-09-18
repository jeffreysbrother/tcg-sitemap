// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"

// 	"github.com/snabb/sitemap"
// )

// func readLines(path string) ([]string, error) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	var lines []string
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}
// 	return lines, scanner.Err()
// }

// func main() {
// 	sm := sitemap.New()

// 	lines, err := readLines("seo_gte5.csv")
// 	if err != nil {
// 		log.Fatalf("readLines: %s", err)
// 	}

// 	for _, line := range lines {
// 		sliceOfStrings := strings.Split(line, ",")
// 		sliceNames := sliceOfStrings[:len(sliceOfStrings)-1]
// 		formattedName := strings.Join(sliceNames, "-")

// 		url := fmt.Sprintf("https://www.truthfinder.com/people-search/%v/", strings.ToLower(formattedName))

// 		// add URL to sitemap
// 		sm.Add(&sitemap.URL{
// 			Loc: url,
// 		})
// 	}

// 	sm.WriteTo(os.Stdout)
// }
