package main

import (
	"encoding/json"
	"net/url"
)

type FilterParamDTO struct {
	Filter       string `json:"filter" form:"filter"`
	Sort         string `json:"sort" form:"sort"`
	Page         int    `json:"page" form:"page"`
	Size         int    `json:"size" form:"size"`
	Search       string `json:"search" form:"search"`
	LinkOperator string `json:"link_operator" form:"link_operator"`
}

func (f *FilterParamDTO) ToFilterParam() (FilterParam, error) {
	var filter []Filter
	if f.Filter != "" {
		err := json.Unmarshal([]byte(f.Filter), &filter)
		if err != nil {
			return FilterParam{}, err
		}
	}

	var sort []Sort
	if f.Sort != "" {
		sortString, err := url.QueryUnescape(f.Sort)
		if err != nil {
			return FilterParam{}, err
		}
		err = json.Unmarshal([]byte(sortString), &sort)
		if err != nil {
			return FilterParam{}, err
		}
	}
	return FilterParam{
		Filter:       filter,
		Sort:         sort,
		PageSize:     f.Size,
		PageNum:      f.Page,
		Search:       f.Search,
		LinkOperator: f.LinkOperator,
	}, nil
}
