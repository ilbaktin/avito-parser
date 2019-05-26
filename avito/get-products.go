package avito

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
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

	doc.Find("div.js-catalog_serp > div.item").Each(func (i int, s *goquery.Selection) {
		title := s.Find("h3.title > a > span").Text()
		priceCurrency := s.Find("div.about > span[itemprop=priceCurrency]").AttrOr("content", "unknown")
		priceStr := s.Find("div.about > span[itemprop=price]").AttrOr("content", "0")

		dataPTag := s.Find("div.data > p")
		category := strings.TrimSpace(firstImmediateText(dataPTag.First()))
		//category := strings.TrimSpace(dataPTag.First().Text())
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

