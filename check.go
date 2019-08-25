package muc

import (
	"errors"
	"log"
)

func (muc *MemoryUsageController) CheckFreeMemory() (err error) {

	var memUsageAMb uint
	var memUsageBMb uint

	// Get the Memory Usage.
	memUsageAMb, err = muc.getMemoryUsage()
	if err != nil {
		return
	}

	// Perform a Check.
	if float64(muc.memUsageLimitMb)/float64(memUsageAMb) < muc.memoryUsedToLimitRatioThreshold {
		// Memory Usage is below the Threshold.
		return
	}

	// Try to free some Memory.
	muc.freeMemory()

	// Get the Memory Usage.
	memUsageBMb, err = muc.getMemoryUsage()
	if err != nil {
		return
	}

	// Verbose Memory Usage Change Report.
	if muc.verboseMode {
		log.Printf(
			fReportVerbose,
			memUsageAMb,
			memUsageBMb,
		)
	}

	// Have we succeeded in freeing the Memory?
	if memUsageBMb < muc.memUsageLimitMb {
		return
	}

	// Memory Usage breaks the Limit.
	err = errors.New(ErrUsageIsBeyondLimits)
	return
}
