package muc

import (
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_NewMemoryUsageController(t *testing.T) {

	var err error
	var mc *MemoryUsageController
	var mcExpected *MemoryUsageController
	var tst *tester.Test

	tst = tester.New(t)
	tst.MustBeNoError(err)

	// Test #1. Positive.
	mc, err = NewMemoryUsageController(
		128,
		float64(0.5),
		MemoryUsageCriterionWorkingSet,
		true,
	)
	mcExpected = &MemoryUsageController{
		memoryUsageCriterion:            MemoryUsageCriterion(MemoryUsageCriterionWorkingSet),
		memUsageLimitMb:                 128,
		memoryUsedToLimitRatioThreshold: float64(0.5),
		verboseMode:                     true,
	}
	tst.MustBeNoError(err)
	tst.MustBeEqual(mc, mcExpected)

	// Test #2. Negative: Usage Limit is Zero.
	mc, err = NewMemoryUsageController(
		0,
		float64(0.5),
		MemoryUsageCriterionWorkingSet,
		true,
	)
	mcExpected = nil
	tst.MustBeAnError(err)
	tst.MustBeEqual(err.Error(), ErrUsageLimitError)
	tst.MustBeEqual(mc, mcExpected)

	// Test #3. Negative: Ratio is Zero.
	mc, err = NewMemoryUsageController(
		128,
		float64(0.0),
		MemoryUsageCriterionWorkingSet,
		true,
	)
	mcExpected = nil
	tst.MustBeAnError(err)
	tst.MustBeEqual(err.Error(), ErrUsageRatioThresholdError)
	tst.MustBeEqual(mc, mcExpected)

	// Test #4. Negative: Usage Criteria is not valid.
	mc, err = NewMemoryUsageController(
		128,
		float64(0.5),
		0,
		true,
	)
	mcExpected = nil
	tst.MustBeAnError(err)
	tst.MustBeEqual(err.Error(), ErrMemoryUsageCriterionInvalid)
	tst.MustBeEqual(mc, mcExpected)
}
