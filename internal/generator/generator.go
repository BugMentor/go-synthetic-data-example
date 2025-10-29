package generator

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"go-synthetic-data-tool/internal/models"
)

// Config holds parameters for the generator, making it highly configurable.
type Config struct {
	NumRecords  int
	NumWorkers  int
	Departments []string
	TenantIDs   []string
}

// Generator manages the creation of synthetic data.
type Generator struct {
	Config Config
}

// NewGenerator initializes a new Generator with the provided configuration.
func NewGenerator(cfg Config) *Generator {
	rand.Seed(time.Now().UnixNano())
	return &Generator{Config: cfg}
}

// generateSingleRecord creates one BudgetRequest instance based on rules.
func (g *Generator) generateSingleRecord() models.BudgetRequest {
	cfg := g.Config
	
	// 1. Categorical Data
	dept := cfg.Departments[rand.Intn(len(cfg.Departments))]
	tenant := cfg.TenantIDs[rand.Intn(len(cfg.TenantIDs))]

	// 2. Statistical Data: Simulating a Normal Distribution for request amounts
	// Base amount is 10,000.00, with a standard deviation of 3,000.00
	baseAmount := 10000.00 + rand.NormFloat64()*3000.00 
	if baseAmount < 500.00 {
		baseAmount = 500.00 // Minimum value constraint
	}
    
	// 3. Status/Boolean Data (e.g., 70% chance of being approved)
	approved := rand.Intn(100) < 70 

	// 4. Time Data (within the last 90 days)
	daysAgo := time.Duration(rand.Intn(90)) * 24 * time.Hour
	createdAt := time.Now().Add(-daysAgo)

	return models.BudgetRequest{
		RequestID:    fmt.Sprintf("BREQ-%d-%d", time.Now().UnixNano()/int64(time.Millisecond), rand.Intn(10000)),
		TenantID:     tenant,
		Department:   dept,
		RequestedAmount: float64(int(baseAmount*100)) / 100, // Round to two decimal places
		IsApproved:   approved,
		CreatedAt:    createdAt,
		Narrative:    fmt.Sprintf("Request for annual %s budget for Q%d operations.", dept, rand.Intn(4)+1),
	}
}

// GenerateConcurrent executes data generation across multiple goroutines (worker pool).
func (g *Generator) GenerateConcurrent() []models.BudgetRequest {
	count := g.Config.NumRecords
	data := make([]models.BudgetRequest, count)
	jobs := make(chan int, count)
	var wg sync.WaitGroup

	// Start worker goroutines
	for w := 0; w < g.Config.NumWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobs {
				data[i] = g.generateSingleRecord()
			}
		}()
	}

	// Queue all jobs
	for i := 0; i < count; i++ {
		jobs <- i
	}
	close(jobs)

	// Wait for all workers to complete
	wg.Wait()
	return data
}
