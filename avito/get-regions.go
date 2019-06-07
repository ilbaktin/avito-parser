package avito

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"univer/avito-parser/model"
)

//func GetAllRegions() ([]*model.Region, error) {
//	reqUrl := prepareUrlWithPath("rossiya")
//
//	resp, err := http.Get(reqUrl.String())
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//
//	doc, err := goquery.NewDocumentFromReader(resp.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	script := doc.Find("div.catalog-counts__section > script").Text()
//	if len(script) == 0 {
//		return nil, fmt.Errorf("can't find script on page")
//	}
//
//	jsonStartIdx := strings.Index(script, "[{")
//	jsonEndIdx := strings.Index(script, "}]")
//
//	if jsonStartIdx == -1 || jsonEndIdx == -1 {
//		return nil, fmt.Errorf("can't find json with counts on page")
//	}
//
//	jsonStr := script[jsonStartIdx:jsonEndIdx+2]
//	jsonStr = strings.Replace(jsonStr, `\`, "", -1)
//
//	var regions []*model.Region
//
//	err = json.Unmarshal([]byte(jsonStr), &regions)
//	if err != nil {
//		return nil, err
//	}
//
//	for _, reg := range regions {
//		qPos := strings.Index(reg.Url, "?")
//		if qPos != -1 {
//			reg.Url = reg.Url[:qPos]
//		}
//	}
//
//	return regions, nil
//}

func GetAllRegions() ([]*model.Region, error) {
	reqUrl := prepareUrlWithPath("web/1/slocations?limit=100000")

	resp, err := http.Get(reqUrl.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, err := ioutil.ReadAll(resp.Body)

	var regions *model.RegionResponse

	err = json.Unmarshal(jsonBytes, &regions)
	if err != nil {
		return nil, err
	}
	return regions.ToRegionSlice(), nil
}
