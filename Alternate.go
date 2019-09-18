package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	// read CSV file
	sliceOfCSVLines, err := readLines("seo_gte5.csv")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// analyze CSV file to see how many blocks of 49,000 entires there are
	countNames := len(sliceOfCSVLines)
	completeBlocks := countNames / 49999
	partialBlock := countNames % 49999
	var totalSitemaps int

	if partialBlock != 0 {
		totalSitemaps = completeBlocks + 1
	} else {
		totalSitemaps = completeBlocks
	}

	// create map of slices corresponding to each sitemap's content
	var sitemapBlocks = make(map[int][]string, totalSitemaps)
	var keys []string

	// populate map with sitemap content
	var sitemapCounter int
	var keysCounter int

	for i := 0; i < countNames; i++ {
		if sitemapCounter < sitemapCounter*49999+49999 {
			keys[keysCounter] = sliceOfCSVLines[i]
			keysCounter++
		}
		sitemapBlocks[sitemapCounter] = keys
		sitemapCounter++
		keysCounter = 0
	}

	fmt.Println(len(sitemapBlocks))

	// fmt.Println(countNames, completeBlocks, partialBlock, totalSitemaps)
	// for k := range sitemapBlocks {
	// 	fmt.Println(sitemapBlocks[k])
	// }

	// create sitemap
	// sm := sitemap.New()

	// // loop over slice of CSV file lines
	// for _, line := range sliceOfCSVLines {
	// 	sliceOfStrings := strings.Split(line, ",")
	// 	sliceNames := sliceOfStrings[:len(sliceOfStrings)-1]
	// 	formattedName := strings.Join(sliceNames, "-")

	// 	url := fmt.Sprintf("https://www.truthfinder.com/people-search/%v/", strings.ToLower(formattedName))

	// 	// add URL to sitemap
	// 	sm.Add(&sitemap.URL{
	// 		Loc: url,
	// 	})
	// }

	// sm.WriteTo(os.Stdout)
}
