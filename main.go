package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"univer/avito-parser/avito"
)

func main() {
	regions, err := avito.GetAllRegions()
	if err != nil {
		fmt.Printf("can't get regions, err='%v'", err)
		return
	}
	regionsJsonBytes, err := json.MarshalIndent(&regions, "", "\t")
	if err != nil {
		fmt.Printf("can't marshal json (regions), err='%v'", err)
		return
	}
	err = ioutil.WriteFile("locations.json", regionsJsonBytes, os.ModePerm)
	if err != nil {
		fmt.Printf("can't save file (regions), err='%v'", err)
		return
	}


	categories, err := avito.GetCategoriesTree()
	if err != nil {
		fmt.Printf("can't get categories, err='%v'", err)
		return
	}
	categoriesJsonBytes, err := json.MarshalIndent(&categories, "", "\t")
	if err != nil {
		fmt.Printf("can't marshal json (categories), err='%v'", err)
		return
	}
	err = ioutil.WriteFile("categories.json", categoriesJsonBytes, os.ModePerm)
	if err != nil {
		fmt.Printf("can't save file (categories), err='%v'", err)
		return
	}


	categoriesMoscow, err := avito.GetCategoriesWithCountsForRegion("moskva")
	if err != nil {
		fmt.Printf("can't get categories counts, err='%v'", err)
		return
	}
	catsMoscowJsonBytes, err := json.MarshalIndent(&categoriesMoscow, "", "\t")
	if err != nil {
		fmt.Printf("can't marshal json (categories counts Moscow), err='%v'", err)
		return
	}
	err = ioutil.WriteFile("categoriesMoscow.json", catsMoscowJsonBytes, os.ModePerm)
	if err != nil {
		fmt.Printf("can't save file (categories counts Moscow), err='%v'", err)
		return
	}


	products, err := avito.GetProducts("moskva", "деревья", 5)
	if err != nil {
		fmt.Printf("can't get products, err='%v'", err)
		return
	}
	productsJsonBytes, err := json.MarshalIndent(&products, "", "\t")
	if err != nil {
		fmt.Printf("can't marshal json (products), err='%v'", err)
		return
	}
	err = ioutil.WriteFile("products.json", productsJsonBytes, os.ModePerm)
	if err != nil {
		fmt.Printf("can't save file (products), err='%v'", err)
		return
	}
}
