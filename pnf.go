package main

type FilterParam struct {
	Filter       []Filter
	Sort         []Sort
	PageSize     int
	PageNum      int
	Search       string
	SearchFields []string
	LinkOperator string
}

type Filter struct {
	Key      string
	Value    string
	Operator string
}

type Sort struct {
	Key   string
	Order string
}

func NewFilterParam() FilterParam {
	return FilterParam{
		Filter:       []Filter{},
		Sort:         []Sort{},
		PageSize:     10,
		PageNum:      1,
		Search:       "",
		SearchFields: []string{},
		LinkOperator: "and",
	}
}
