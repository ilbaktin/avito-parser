package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func StartServer(port uint16) error {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(setJsonHeadersHandler)

	router.HandleFunc("/regions", GetAllRegions).Methods("GET")
	router.HandleFunc("/regions_ext", GetAllRegionsExtended).Methods("GET")
	router.HandleFunc("/categories_tree", GetCategoriesTree).Methods("GET")
	router.HandleFunc("/categories_counts", GetCategoriesWithCountsForRegion).Methods("GET")
	router.HandleFunc("/products", GetProducts).Methods("GET")
	router.HandleFunc("/products_ext", GetProductsExtended).Methods("GET")

	server := http.Server{Handler: router, Addr: fmt.Sprintf("0.0.0.0:%d", port)}

	fmt.Println(fmt.Sprintf("Starting http server on :%d...", port))
	return server.ListenAndServe()
}
