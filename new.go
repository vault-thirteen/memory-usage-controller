package muc

import (
	"errors"
)

func NewMemoryUsageController(
	memUsageLimitMb uint,
	memoryUsedToLimitRatioThreshold float64,
	memoryUsageCriterion byte,
	verboseMode bool,
) (result *MemoryUsageController, err error) {

	if memUsageLimitMb == 0 {
		err = errors.New(ErrUsageLimitError)
		return
	}
	if memoryUsedToLimitRatioThreshold <= 0 {
		err = errors.New(ErrUsageRatioThresholdError)
		return
	}

	result = new(MemoryUsageController)

	// Set the Memory Usage Criteria.
	result.memoryUsageCriterion = NewMemoryUsageCriterion(memoryUsageCriterion)
	if !result.memoryUsageCriterion.IsValid() {
		err = errors.New(ErrMemoryUsageCriterionInvalid)
		result = nil
		return
	}

	// Set other Fields.
	result.memUsageLimitMb = memUsageLimitMb
	result.memoryUsedToLimitRatioThreshold = memoryUsedToLimitRatioThreshold
	result.verboseMode = verboseMode

	return
}
