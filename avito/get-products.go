package avito

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"univer/avito-parser/model"
)

func GetProducts(regionPath, query string, page int) ([]*model.Product, error) {
	if page <= 0 {
		page = 1
	}

	reqUrl := prepareUrlWithPath(regionPath)
	q := reqUrl.Query()
	q.Add("q", query)
	q.Add("p", strconv.Itoa(page))
	reqUrl.RawQuery = q.Encode()

	resp, err := http.Get(reqUrl.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	products := make([]*model.Product, 0, 100)

	doc.Find("div.js-catalog_serp > div.item").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h3.title > a > span").Text()
		priceCurrency := s.Find("div.about > span[itemprop=priceCurrency]").AttrOr("content", "unknown")
		priceStr := s.Find("div.about > span[itemprop=price]").AttrOr("content", "0")

		dataPTag := s.Find("div.data > p")
		category := strings.TrimSpace(firstImmediateText(dataPTag.First()))
		location := strings.TrimSpace(dataPTag.First().Next().Text())
		date := strings.TrimSpace(s.Find(".js-item-date").AttrOr("data-absolute-date", ""))

		product := &model.Product{}
		product.Name = title
		product.Currency = priceCurrency
		price, _ := strconv.Atoi(priceStr)
		product.Price = price
		product.Category = category
		product.Location = location
		product.Date = date

		products = append(products, product)
	})

	return products, nil
}

func GetProductsExtended(query string, categoryId, locationId, page int) ([]*model.Product, error) {
	urlToReq, err := url.Parse("https://www.avito.ru/search")
	if err != nil {
		return nil, err
	}

	data := url.Values{}
	if categoryId == 0 {
		data.Set("category_id", "")
	} else {
		data.Set("category_id", strconv.Itoa(categoryId))
	}
	if locationId == 0 {
		data.Set("location_id", "621540")
	} else {
		data.Set("location_id", strconv.Itoa(locationId))
	}
	data.Set("map", "")
	data.Set("s", "101")
	data.Set("sgtd", "")
	data.Set("name", query)
	fmt.Println(data.Encode())

	req, err := http.NewRequest("POST", urlToReq.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	defer func(r *http.Response) {
		r.Body.Close()
	}(resp)

	if resp.StatusCode != 301 {
		return nil, fmt.Errorf("wrong resp code for search query: got %d, want 301", resp.StatusCode)
	}
	newLocation := resp.Header.Get("Location")
	fmt.Println(newLocation)
	urlToReq, err = url.Parse("https://avito.ru" + newLocation)
	if err != nil {
		return nil, err
	}
	q := urlToReq.Query()
	q.Add("p", strconv.Itoa(page))
	urlToReq.RawQuery = q.Encode()
	fmt.Println(urlToReq.String())

	resp, err = http.Get(urlToReq.String())
	if err != nil {
		return nil, err
	}
	defer func(r *http.Response) {
		r.Body.Close()
	}(resp)

	products := make([]*model.Product, 0, 100)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("div.js-catalog_serp > div.item").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h3.title > a > span").Text()
		priceCurrency := s.Find("div.about > span[itemprop=priceCurrency]").AttrOr("content", "unknown")
		priceStr := s.Find("div.about > span[itemprop=price]").AttrOr("content", "0")
		link := s.Find("a.item-description-title-link").AttrOr("href", "unknown")
		if link != "unknown" {
			link = "https://avito.ru" + link
		}

		dataPTag := s.Find("div.data > p")
		category := strings.TrimSpace(firstImmediateText(dataPTag.First()))
		location := strings.TrimSpace(dataPTag.First().Next().Text())
		date := strings.TrimSpace(s.Find(".js-item-date").AttrOr("data-absolute-date", ""))

		product := &model.Product{}
		product.Name = title
		product.Currency = priceCurrency
		price, _ := strconv.Atoi(priceStr)
		product.Price = price
		product.Category = category
		product.Location = location
		product.Date = date
		product.Link = link

		products = append(products, product)
	})

	wg := sync.WaitGroup{}
	wg.Add(len(products))
	for _, product := range products {
		go func(p *model.Product) {
			p.FullText = GetFulltextForProduct(p.Link)
			wg.Done()
		}(product)
	}
	wg.Wait()

	return products, nil
}

func GetFulltextForProduct(rawUrl string) string {
	resp, err := http.Get(rawUrl)
	if err != nil {
		return ""
	}
	defer func(r *http.Response) {
		r.Body.Close()
	}(resp)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return ""
	}

	return strings.Trim(doc.Find("div.item-description").Text(), "\n\t ")
}
