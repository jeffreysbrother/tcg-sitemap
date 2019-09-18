package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var baseURL string = "https://www.truthfinder.com/people-search"

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
	maxPerSitemap := 49999
	completeBlocks := countNames / maxPerSitemap
	partialBlock := countNames % maxPerSitemap
	var totalSitemaps int

	if partialBlock != 0 {
		totalSitemaps = completeBlocks + 1
	} else {
		totalSitemaps = completeBlocks
	}

	// sitemapCounter := 0
	// arrayIndexCounter := 0

	// create map of slices corresponding to each sitemap's content
	mapOfSitemapContents := make(map[int][]string)
	mapIndex := 0

	for ind, el := range sliceOfCSVLines {
		sliceOfStrings := strings.Split(el, ",")
		sliceNames := sliceOfStrings[:len(sliceOfStrings)-1]

		url := fmt.Sprintf("%v/%v-%v/", baseURL, strings.ToLower(sliceNames[1]), strings.ToLower(sliceNames[0]))
		if ind < (maxPerSitemap*mapIndex + maxPerSitemap) {
			mapOfSitemapContents[mapIndex] = append(mapOfSitemapContents[mapIndex], url)
		} else {
			mapIndex++
		}
	}

	fmt.Println(countNames, completeBlocks, partialBlock, totalSitemaps, mapOfSitemapContents)

	// var sitemapBlocks = make(map[int][]string, totalSitemaps)
	// var keys []string

	// populate map with sitemap content

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
