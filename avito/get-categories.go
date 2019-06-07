package avito

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
	"univer/avito-parser/model"
)

func GetCategoriesTree() ([]*model.Category, error) {
	reqUrl := prepareUrlWithPath("rossiya")

	resp, err := http.Get(reqUrl.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	categories := make([]*model.Category, 0, 30)
	var lastRoot *model.Category

	doc.Find("div.form-select-v2 > select[name=category_id] > option").Each(func(i int, s *goquery.Selection) {
		//val, err := s.Html()
		//fmt.Printf("%#v, err='%v'\n", val, err)
		//fmt.Println(goquery.NodeName(s))
		//fmt.Println(goquery.OuterHtml(s))
		if s.HasClass("opt-group") || lastRoot == nil {
			lastRoot = model.NewCategory()
			lastRoot.Name = s.Text()
			lastRoot.Id = s.AttrOr("value", "")
			categories = append(categories, lastRoot)
		} else {
			curCategory := model.NewCategory()
			curCategory.Name = s.Text()
			curCategory.Id = s.AttrOr("value", "")
			lastRoot.Subcategories = append(lastRoot.Subcategories, curCategory)
		}
	})
	categories = append(categories, lastRoot)

	return categories, nil
}

func GetCategoriesWithCountsForRegion(regionPath string) ([]*model.Category, error) {
	reqUrl := prepareUrlWithPath(regionPath)

	resp, err := http.Get(reqUrl.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	categories := make([]*model.Category, 0, 30)

	doc.Find("div.catalog-counts__row > ul > li").Each(func(i int, s *goquery.Selection) {
		if s.Contents().Length() < 2 {
			return
		}
		aTag := s.Find("a")
		spanTag := s.Find("span")
		category := model.NewCategory()
		category.Name = strings.Trim(aTag.Text(), "\n\t ")
		count, _ := strconv.Atoi(strings.Replace(spanTag.Text(), " ", "", -1))
		category.Count = &count
		categories = append(categories, category)
	})

	return categories, nil
}
