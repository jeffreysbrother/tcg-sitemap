package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/snabb/sitemap"
)

var baseURL = "https://www.truthfinder.com/people-search"

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

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func main() {
	// delete sitemaps dir if it already exists
	directoryExists, err := exists("sitemaps")
	if err != nil {
		log.Println(err)
	}

	if directoryExists {
		os.RemoveAll("sitemaps")
	}

	// retrieve file name from Args
	// and handle error if user fails pass an arg or passes too many
	var csvFile string
	if len(os.Args) == 2 {
		csvFile = os.Args[1]
	} else if len(os.Args) < 2 {
		log.Fatalf("Please include a file name as an argument")
	} else {
		log.Fatalf("Too many arguments!")
	}

	// read CSV file
	sliceOfCSVLines, err := readLines(csvFile)
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

	// create sitemap dir
	err = os.Mkdir("sitemaps", 0755)
	if err != nil {
		fmt.Println(err)
	}

	// create and write to sitemap index file
	f, err := os.Create("sitemaps/sitemap.xml")
	if err != nil {
		fmt.Println(err)
		return
	}

	sitemapIndex.WriteTo(f)

	// create a sitemap for every index of the map
	for ind, el := range mapOfSitemapContents {
		sm := sitemap.New()
		for _, innerEl := range el {
			sm.Add(&sitemap.URL{
				Loc: innerEl,
			})
		}
		// create and write to individual sitemap files (with the above naving convention)
		f, err := os.Create(fmt.Sprintf("sitemaps/icm-ppl%v-sitemap.xml", ind))
		if err != nil {
			fmt.Println(err)
			return
		}

		sm.WriteTo(f)
	}

	fmt.Println("Sitemaps created!")
}
