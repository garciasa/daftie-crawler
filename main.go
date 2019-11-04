package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getPage(page int) (io.Reader, error) {
	nbed := "4"
	url := ""
	if page == 0 {
		url = fmt.Sprintf("https://www.daft.ie/wexford/property-for-sale/wexford-town/?s%%5Bmnb%%5D=%s&s%%5Badvanced%%5D=1&s%%5Bsort_by%%5D=date&s%%5Bsort_type%%5D=d&searchSource=sale", nbed)
	} else {
		page := page * 20
		url = fmt.Sprintf("https://www.daft.ie/wexford/property-for-sale/wexford-town/?s%%5Bmnb%%5D=%s&s%%5Badvanced%%5D=1&s%%5Bsort_by%%5D=date&s%%5Bsort_type%%5D=d&searchSource=sale&offset=%s",
			nbed, strconv.FormatInt(int64(page), 10))
	}

	fmt.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		defer resp.Body.Close()
		return nil, err
	}

	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	return resp.Body, nil

}

func main() {
	data, err := getPage(1)

	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".PropertyCardContainer__container").Each(func(i int, s *goquery.Selection) {
		brandlink, ok := s.Find(".brandLink").Attr("href")
		if ok {
			fmt.Println(brandlink)
		}
		price := s.Find(".PropertyInformationCommonStyles__costAmountCopy").Text()
		date := s.Find(".PropertyInformationCommonStyles__startDate").Text()
		if s.Find(".PropertyInformationCommonStyles__newDevelopmentLabel").Size() > 0 {
			fmt.Println("New Development")
		} else {
			fmt.Printf("Price: %s - Date: %s\n", price, date)
		}
	})
	var pages []int
	doc.Find(".paging.clear").Each(func(i int, s *goquery.Selection) {
		s.Find("li").Each(func(i int, p *goquery.Selection) {
			if !strings.Contains(p.Text(), "Next") && !strings.Contains(p.Text(), "Previous") {
				p, err := strconv.Atoi(p.Text())
				if err != nil {
					log.Fatal("Error converting...")
				}
				pages = append(pages, p)
			}
		})
	})

	// Getting rest of the pages
	fmt.Println(pages)

}
