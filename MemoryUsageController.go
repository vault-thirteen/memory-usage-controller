package muc

// Memory Usage Controller.

import (
	"sync"
)

const (
	MB = 1 * 1000 * 1000 // 1 M.
)

// Errors.
const (
	ErrUsageLimitError             = "Usage Limit Error"
	ErrUsageRatioThresholdError    = "Usage Ratio Threshold Error"
	ErrUsageIsBeyondLimits         = "Memory Usage is beyond the Limits"
	ErrMemoryUsageCriterionInvalid = "Memory Usage Criterion is not valid"
	ErrMemoryUsageCriterionUnknown = "Memory Usage Criterion is unknown"
)

const (
	fReportVerbose = "RAM Usage has been minimized: %d -> %d MB.\r\n"
)

type MemoryUsageController struct {
	memoryUsageCriterion            MemoryUsageCriterion
	memUsageLimitMb                 uint
	memoryUsedToLimitRatioThreshold float64
	verboseMode                     bool

	// Locks.
	usageLock sync.Mutex
	freeLock  sync.Mutex
}
