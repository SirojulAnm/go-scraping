package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type Article2 struct {
	Title  string
	Author string
}

func main() {
	res, err := http.Get("http://quotes.toscrape.com/")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	rows := make([]Article2, 0)
	doc.Find(".row .quote").Each(func(i int, sel *goquery.Selection) {
		row := new(Article2)
		row.Title = sel.Find(".text").Text()
		row.Author = sel.Find(".author").Text()
		rows = append(rows, *row)
	})

	csvFile, err := os.Create("./result.csv")

	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	for _, usance := range rows {
		var row []string
		row = append(row, usance.Title)
		row = append(row, usance.Author)
		writer.Write(row)
	}
	writer.Flush()
}
