package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/snabb/sitemap"
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

	// create map of slices corresponding to each sitemap's content
	mapOfSitemapContents := make(map[int][]string)
	mapIndex := 0

	for ind, el := range sliceOfCSVLines {
		// formatting
		sliceOfStrings := strings.Split(el, ",")
		sliceNames := sliceOfStrings[:len(sliceOfStrings)-1]
		url := fmt.Sprintf("%v/%v-%v/", baseURL, strings.ToLower(sliceNames[1]), strings.ToLower(sliceNames[0]))

		if ind < (maxPerSitemap*mapIndex + maxPerSitemap) {
			mapOfSitemapContents[mapIndex] = append(mapOfSitemapContents[mapIndex], url)
		} else {
			mapIndex++
		}
	}

	// create a sitemap index
	// naming convention: icm-ppl10-sitemap.xml
	sitemapIndex := sitemap.NewSitemapIndex()
	for ind := range mapOfSitemapContents {
		sitemapIndex.Add(&sitemap.URL{
			Loc: fmt.Sprintf("icm-ppl%v-sitemap.xml", ind),
		})
	}

	// create a sitemap for every index of the map
	// for _, el := range mapOfSitemapContents {

	// }

	// figure out a way to write the sitemap and sitemap index to files

	// name the sitemap files appropriately

	fmt.Println(countNames, completeBlocks, partialBlock, totalSitemaps, mapOfSitemapContents[0][2])
	sitemapIndex.WriteTo(os.Stdout)
}
