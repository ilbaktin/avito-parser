package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"univer/avito-parser/avito"
	"univer/avito-parser/model"
)

const INVALID_PARAMS_MSG = "invalid parse GET params"
const INTERNAL_ERROR = "Sorry, internal error"

func GetAllRegions(w http.ResponseWriter, r *http.Request) {
	regions, err := avito.GetAllRegions()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s: %v", INTERNAL_ERROR, err)))
		w.WriteHeader(500)
		return
	}

	writeObjectToResp(w, regions)
}

func GetAllRegionsExtended(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := ioutil.ReadFile("samples/regions.json")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s: %v", INTERNAL_ERROR, err)))
		w.WriteHeader(500)
		return
	}

	var categories []*model.RegionExt

	err = json.Unmarshal(jsonBytes, &categories)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s: %v", INTERNAL_ERROR, err)))
		w.WriteHeader(500)
		return
	}

	writeObjectToResp(w, categories)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.Write([]byte(INVALID_PARAMS_MSG))
		w.WriteHeader(400)
		return
	}

	region := params.Get("region")
	query := params.Get("q")
	page := params.Get("p")

	if region == "" {
		region = "rossiya"
	}
	if query == "" {
		w.Write([]byte(fmt.Sprintf("%s, 'q' param is required, 'region' and 'p' is optional", INVALID_PARAMS_MSG)))
		w.WriteHeader(400)
		return
	}
	if page == "" {
		page = "1"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt <= 0 {
		w.Write([]byte(fmt.Sprintf("%s, 'p' param should be positive integer", INVALID_PARAMS_MSG)))
		w.WriteHeader(400)
		return
	}

	products, err := avito.GetProducts(region, query, pageInt)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s: %v", INTERNAL_ERROR, err)))
		w.WriteHeader(500)
		return
	}

	writeObjectToResp(w, products)
}

func GetProductsExtended(w http.ResponseWriter, r *http.Request) {
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.Write([]byte(INVALID_PARAMS_MSG))
		w.WriteHeader(400)
		return
	}

	categoryId := params.Get("cid")
	locationId := params.Get("lid")
	query := params.Get("q")
	page := params.Get("p")

	if categoryId == "" {
		categoryId = "0"
	}
	if locationId == "" {
		locationId = "0"
	}
	if page == "" {
		page = "1"
	}
	if query == "" {
		w.Write([]byte(fmt.Sprintf("%s, 'q' param is required, 'cid', 'lid' and 'p' is optional", INVALID_PARAMS_MSG)))
		w.WriteHeader(400)
		return
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt <= 0 {
		w.Write([]byte(fmt.Sprintf("%s, 'p' param should be positive integer", INVALID_PARAMS_MSG)))
		w.WriteHeader(400)
		return
	}
	categoryIdInt, err := strconv.Atoi(categoryId)
	if err != nil || pageInt <= 0 {
		w.Write([]byte(fmt.Sprintf("%s, 'cid' param should be positive integer", INVALID_PARAMS_MSG)))
		w.WriteHeader(400)
		return
	}
	locationIdInt, err := strconv.Atoi(locationId)
	if err != nil || pageInt <= 0 {
		w.Write([]byte(fmt.Sprintf("%s, 'lid' param should be positive integer", INVALID_PARAMS_MSG)))
		w.WriteHeader(400)
		return
	}

	products, err := avito.GetProductsExtended(query, categoryIdInt, locationIdInt, pageInt)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s: %v", INTERNAL_ERROR, err)))
		w.WriteHeader(500)
		return
	}

	writeObjectToResp(w, products)
}

func GetCategoriesTree(w http.ResponseWriter, r *http.Request) {
	categories, err := avito.GetCategoriesTree()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s: %v", INTERNAL_ERROR, err)))
		w.WriteHeader(500)
		return
	}

	writeObjectToResp(w, categories)
}

func GetCategoriesWithCountsForRegion(w http.ResponseWriter, r *http.Request) {
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.Write([]byte(INVALID_PARAMS_MSG))
		w.WriteHeader(400)
		return
	}

	region := params.Get("region")
	if region == "" {
		region = "rossiya"
	}

	categories, err := avito.GetCategoriesWithCountsForRegion(region)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s: %v", INTERNAL_ERROR, err)))
		w.WriteHeader(500)
		return
	}

	writeObjectToResp(w, categories)
}
