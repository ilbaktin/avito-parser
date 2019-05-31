package rest

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"univer/avito-parser/avito"
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
