package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("http://samskipnaden.no/meny")
	if err != nil {
		fmt.Println(err)
		return
	}

	doc.Find("li.views-row").Each(func(i int, s *goquery.Selection) {

		kantine := s.Find(".views-field-field-kantine .field-content").Text()
		kantine = strings.TrimSpace(kantine)
		fmt.Println(kantine)
		s.Find(".field-collection-view").Each(func(i int, s *goquery.Selection) {
			rett := s.Find(".field-name-field-dish").Text()
			pris := s.Find(".field-name-field-price").Text()

			rett = strings.TrimSpace(rett)
			pris = strings.TrimSpace(pris)

			if rett != "" && pris != "" {
				fmt.Println(rett, "\t", pris)
			}
		})

		fmt.Println()

	})
}
