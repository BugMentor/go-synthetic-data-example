package models

import "time"

// BudgetRequest defines the structured schema for a synthetic government budgeting record.
// All fields use JSON tags for easy serialization/injection.
type BudgetRequest struct {
	RequestID       string    `json:"request_id"`
	TenantID        string    `json:"tenant_id"`        // Critical for multi-tenancy testing
	Department      string    `json:"department"`       // Categorical data
	RequestedAmount float64 `json:"requested_amount"` // Statistical data
	IsApproved      bool      `json:"is_approved"`      // Boolean/Status field
	CreatedAt       time.Time `json:"created_at"`
	Narrative       string    `json:"narrative"`        // Placeholder text
}

// ValidationResult holds the status of a single record validation.
type ValidationResult struct {
	RequestID string
	IsValid   bool
	Errors    []string // List of validation errors encountered
}
