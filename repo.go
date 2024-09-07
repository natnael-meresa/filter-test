package main

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

var (
	ErrNoTableAFound = ErrNoRecordFound.New("table a not found")
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	// auto migrate
	err := db.AutoMigrate(&TableA{})
	if err != nil {
		panic(err)
	}
	return &Repo{db: db}
}

func (r *Repo) List(ctx context.Context, listParam FilterParam) ([]TableA, error) {
	var tableA []TableA
	db := r.db.WithContext(ctx)

	db = r.ApplyFilters(db, listParam.Filter)
	db = r.ApplySort(db, listParam.Sort)
	db = r.ApplyPagination(db, listParam.PageSize, listParam.PageNum)
	err := db.Find(&tableA).Error

	if err != nil {
		return nil, err
	}

	return tableA, nil
}

func (r *Repo) Get(ctx context.Context, id int) (*TableA, error) {
	var tableA TableA
	db := r.db.WithContext(ctx)
	err := db.First(&tableA, id).Error
	if err != nil {
		fmt.Println("err", err)
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNoTableAFound
		}
		return nil, err
	}

	if tableA == (TableA{}) {
		return nil, ErrNoTableAFound
	}

	fmt.Println("tableA", tableA)
	return &tableA, nil
}

func (r *Repo) ApplyFilters(db *gorm.DB, filter []Filter) *gorm.DB {
	for _, f := range filter {
		db = db.Where(f.Key+" "+f.Operator+" ?", f.Value)
	}
	return db
}

func (r *Repo) ApplySort(db *gorm.DB, sort []Sort) *gorm.DB {
	for _, s := range sort {
		db = db.Order(s.Key + " " + s.Order)
	}
	return db
}

func (r *Repo) ApplyPagination(db *gorm.DB, pageSize, pageNum int) *gorm.DB {
	if pageSize > 0 {
		db = db.Limit(pageSize)
	}
	if pageNum > 0 {
		db = db.Offset((pageNum - 1) * pageSize)
	}
	return db
}
