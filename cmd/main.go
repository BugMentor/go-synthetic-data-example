package main

import (
	"fmt"
	"time"
	"go-synthetic-data-tool/internal/generator"
	"go-synthetic-data-tool/internal/validator"
)

func main() {
	// --- 1. Configuration ---
	cfg := generator.Config{
		NumRecords:  50000, // Generate 50,000 records
		NumWorkers:  8,     // Use 8 concurrent goroutines
		Departments: []string{"Police", "Fire", "Public Works", "Parks & Rec", "IT", "Finance", "Legal"},
		TenantIDs:   []string{"NYC-101", "LA-202", "CHI-303", "BOS-404", "DAL-505"},
	}

	// --- 2. Generation ---
	fmt.Printf("--- Starting Synthetic Data Generation (%d Records, %d Workers) ---\n", cfg.NumRecords, cfg.NumWorkers)
	gen := generator.NewGenerator(cfg)
	startTime := time.Now()
	syntheticData := gen.GenerateConcurrent()
	duration := time.Since(startTime)

	fmt.Printf("Generation Complete in %s. Throughput: %.2f records/second\n\n", 
		duration, float64(cfg.NumRecords)/duration.Seconds())

	// --- 3. Validation ---
	fmt.Println("--- Starting Concurrent Data Validation ---")
	validationResults := validator.ValidateAll(syntheticData, cfg.TenantIDs)
	
	// --- 4. Reporting ---
	
	var validCount, invalidCount int
	for _, res := range validationResults {
		if res.IsValid {
			validCount++
		} else {
			invalidCount++
		}
	}

	fmt.Printf("Validation Complete. Total Records: %d\n", len(syntheticData))
	fmt.Printf("   ✅ Valid Records: %d\n", validCount)
	fmt.Printf("   ❌ Invalid Records: %d\n", invalidCount)
	
	if invalidCount > 0 {
		fmt.Println("\n--- Sample of Failures ---")

		for _, res := range validationResults {
			if !res.IsValid {
				fmt.Printf("[ID: %s] Errors: %v\n", res.RequestID, res.Errors)
			}
		}
	} else {
		fmt.Println("\nAll generated records passed structural and business validation.")
	}
}
