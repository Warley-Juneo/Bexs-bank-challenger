package dto

import (
	"net/http"
	"strconv"
	"strings"
)

type PaginationRequestParms struct {
	Search				string `json:"search"`
	Descending		[]string `json:"descending"`
	Page					int `json:"page"`
	ItemsPerPage	int `json:"itemsPerPage"`
	Sort 					[]string `json:"sortBy"`
}

func FromValuePaginationRequestParms(request *http.Request) (*PaginationRequestParms, error) {
	page, _ := strconv.Atoi(request.FormValue("page"))
	itemsPerPage, _ := strconv.Atoi(request.FormValue("itemsPerPage"))

	paginationRequestParms := PaginationRequestParms {
		Search: request.FormValue("search"),
		Descending: strings.Split(request.FormValue("descending"), ","),
		Page: page,
		ItemsPerPage: itemsPerPage,
		Sort: strings.Split(request.FormValue("sort"), ","),
	}

	return &paginationRequestParms, nil
}