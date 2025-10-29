package validator

import (
	"go-synthetic-data-tool/internal/models"
	"strings"
	"sync"
)

// ValidateBudgetRequest runs all constraint checks on a single generated record.
// This ensures that the generated data is compliant with business rules.
func ValidateBudgetRequest(r models.BudgetRequest, allowedTenants []string) models.ValidationResult {
	result := models.ValidationResult{
		RequestID: r.RequestID,
		IsValid:   true,
		Errors:    []string{},
	}

	// Rule 1: RequestedAmount must be non-negative and greater than a minimum threshold.
	if r.RequestedAmount < 500.00 {
		result.Errors = append(result.Errors, "RequestedAmount must be at least 500.00.")
	}

	// Rule 2: TenantID must be one of the configured, known tenants (Multi-Tenancy Isolation Check)
	if !isTenantAllowed(r.TenantID, allowedTenants) {
		result.Errors = append(result.Errors, "TenantID is invalid or not in the approved list for testing.")
	}

	// Rule 3: RequestID must have the correct prefix and length (Structural check)
	if !strings.HasPrefix(r.RequestID, "BREQ-") || len(r.RequestID) < 10 {
		result.Errors = append(result.Errors, "RequestID must start with 'BREQ-' and be at least 10 characters long.")
	}

	// Rule 4: Narrative must not be empty (Content presence check)
	if len(strings.TrimSpace(r.Narrative)) == 0 {
		result.Errors = append(result.Errors, "Narrative field cannot be empty.")
	}

	if len(result.Errors) > 0 {
		result.IsValid = false
	}

	return result
}

// isTenantAllowed checks if the TenantID is in the allowed list (demonstrates lookup logic).
func isTenantAllowed(tenantID string, allowedTenants []string) bool {
	for _, allowed := range allowedTenants {
		if tenantID == allowed {
			return true
		}
	}

	return false
}

// ValidateAll processes all generated records and returns a summary.
func ValidateAll(data []models.BudgetRequest, allowedTenants []string) []models.ValidationResult {
	results := make([]models.ValidationResult, len(data))
	// Validation can also be parallelized for high-volume checks
	var wg sync.WaitGroup // Use sync.WaitGroup from the imported sync package

	for i, record := range data {
		wg.Add(1)
		go func(index int, rec models.BudgetRequest) {
			defer wg.Done()
			results[index] = ValidateBudgetRequest(rec, allowedTenants)
		}(i, record)
	}

	wg.Wait()

	return results
}
