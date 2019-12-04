package domain

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func (app *App) parse() ([]House, error) {
	defer timeTrack(time.Now(), "Parsing")
	var all []House
	doc, err := getPage(0)

	if err != nil {
		return nil, err
	}

	// Get number of pages for that search
	// should we include all the results? or just 4 beds?
	pages := getNumPages(doc)

	// Iterate for each page
	next := 0
	for i := 0; i < pages; i++ {
		p, err := getPage(next)
		if err != nil {
			return nil, err
		}

		houses, err := getAdverts(p)
		if err != nil {
			return nil, err
		}
		all = append(all, houses...)
		next += 20
	}

	for _, h := range all {
		getHouseDetails(&h)
		err := h.save(app.db)
		if err != nil {
			return nil, err
		}
		//printHouse(h)
	}

	return all, nil
}

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
	brandLink := house.BrandLink
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
		house.Meters = meters
	} else {
		house.Meters = "N/A"
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
			house.BrandLink = brandlink
		} else {
			//find another way to get link
			link, ok := s.Find(".PropertyImage__link").Attr("href")
			if ok {
				house.BrandLink = link
			}
		}
		price := s.Find(".PropertyInformationCommonStyles__costAmountCopy").Text()
		house.Price = price
		date := s.Find(".PropertyInformationCommonStyles__startDate").Text()
		house.Date = date
		if s.Find(".PropertyInformationCommonStyles__newDevelopmentLabel").Size() > 0 {
			house.NewDevelopment = true
		} else {
			house.NewDevelopment = false
		}
		houses = append(houses, house)
	})

	return houses, nil
}
