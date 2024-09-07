package main

import (
	"time"

	"gorm.io/gorm"
)

// SeedData seeds the database with initial data.
func SeedData(db *gorm.DB) error {
	// Check if data already exists to avoid duplicating it
	var count int64
	db.Model(TableA{}).Count(&count)
	if count > 0 {
		return nil // Data already seeded
	}

	// Define seed data for TableA
	tableAData := []TableA{
		{
			Column1:   "Value1",
			Column2:   10,
			Column3:   "Alpha",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Column1:   "Value2",
			Column2:   20,
			Column3:   "Beta",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Column1:   "Value3",
			Column2:   30,
			Column3:   "Gamma",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Column1:   "Value4",
			Column2:   40,
			Column3:   "Delta",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	// Insert seed data into the database
	if err := db.Create(&tableAData).Error; err != nil {
		return err
	}

	return nil
}
