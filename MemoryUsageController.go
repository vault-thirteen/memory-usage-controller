// MemoryUsageController.go.

////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019..2020 by Vault Thirteen.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
////////////////////////////////////////////////////////////////////////////////
//
// Web Site Address:	https://github.com/vault-thirteen.
//
////////////////////////////////////////////////////////////////////////////////

package muc

// Memory Usage Controller.

import (
	"errors"
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
