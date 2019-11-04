package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// House Info about House advert
type House struct {
	brandLink      string
	price          string
	date           string
	newDevelopment bool
	meters         string
	eircode        string
}

// DOMAIN main domain to use
const DOMAIN = "https://www.daft.ie"

func getPage(page int) (*goquery.Document, error) {
	nbed := "4"
	url := ""
	if page == 0 {
		url = fmt.Sprintf("https://www.daft.ie/wexford/property-for-sale/wexford-town/?s%%5Bmnb%%5D=%s&s%%5Badvanced%%5D=1&s%%5Bsort_by%%5D=date&s%%5Bsort_type%%5D=d&searchSource=sale", nbed)
	} else {
		url = fmt.Sprintf("https://www.daft.ie/wexford/property-for-sale/wexford-town/?s%%5Bmnb%%5D=%s&s%%5Badvanced%%5D=1&s%%5Bsort_by%%5D=date&s%%5Bsort_type%%5D=d&searchSource=sale&offset=%s",
			nbed, strconv.FormatInt(int64(page), 10))
	}

	resp, err := http.Get(url)

	if err != nil {
		defer resp.Body.Close()
		return nil, err
	}

	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil

}

func getHouseDetails(house *House) {
	brandLink := house.brandLink
	url := fmt.Sprintf("%s%s", DOMAIN, brandLink)
	resp, err := http.Get(url)
	if err != nil {
		defer resp.Body.Close()
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	details := doc.Find(".PropertyOverview__propertyOverviewDetails").Text()
	validMeters := regexp.MustCompile(`[0-9].+ m2`)
	meters := validMeters.FindString(details)
	if validMeters.MatchString(details) {
		house.meters = meters
	} else {
		house.meters = "N/A"
	}
}

func getNumPages(doc *goquery.Document) int {
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

	return len(pages)

}

func getAdverts(doc *goquery.Document) ([]House, error) {
	var houses []House
	doc.Find(".PropertyCardContainer__container").Each(func(i int, s *goquery.Selection) {
		var house House
		brandlink, ok := s.Find(".brandLink").Attr("href")
		if ok {
			house.brandLink = brandlink
		} else {
			//find another way to get link
			link, ok := s.Find(".PropertyImage__link").Attr("href")
			if ok {
				house.brandLink = link
			}
		}
		price := s.Find(".PropertyInformationCommonStyles__costAmountCopy").Text()
		house.price = price
		date := s.Find(".PropertyInformationCommonStyles__startDate").Text()
		house.date = date
		if s.Find(".PropertyInformationCommonStyles__newDevelopmentLabel").Size() > 0 {
			house.newDevelopment = true
		} else {
			house.newDevelopment = false
		}
		houses = append(houses, house)
	})

	return houses, nil
}

func printHouse(house House) {
	if house.newDevelopment {
		fmt.Printf("%s - New Development\n", house.brandLink)
	} else {
		fmt.Printf("%s - %s - %s - %s\n", house.brandLink, house.date, house.price, house.meters)
	}
}

func main() {
	var allHouses []House
	doc, err := getPage(0)
	if err != nil {
		log.Fatal(err)
	}

	// Get number of pages for that search
	// should we include all the results? or just 4 bed?
	pages := getNumPages(doc)

	// Iterate for each page
	next := 0
	for i := 0; i < pages; i++ {
		p, err := getPage(next)
		if err != nil {
			log.Fatal(err)
		}

		houses, err := getAdverts(p)
		if err != nil {
			log.Fatal(err)
		}
		allHouses = append(allHouses, houses...)
		next += 20
	}

	for _, h := range allHouses {
		// Check if that result is already in
		// If so, check new price and date

		// If it's not in, add it
		getHouseDetails(&h)
		printHouse(h)
	}
	fmt.Println(len(allHouses))
	// repeat every hour

}
