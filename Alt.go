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
	// read argument (file name)
	fileNameCSV := os.Args[1]

	// TODO: handle error if user doesn't pass an arg or passes too many

	// read CSV file
	sliceOfCSVLines, err := readLines(fileNameCSV)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// set max number of entries per sitemap
	maxPerSitemap := 49999

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
	// NOT SURE IF THIS IS WORKING
	for _, el := range mapOfSitemapContents {
		sm := sitemap.New()
		for _, innerEl := range el {
			sm.Add(&sitemap.URL{
				Loc: innerEl,
			})
		}
		sm.WriteTo(os.Stdout)
	}

	// figure out a way to write the sitemap and sitemap index to files

	// name the sitemap files appropriately

	// fmt.Println(mapOfSitemapContents[0][2])
	// sitemapIndex.WriteTo(os.Stdout)
}
