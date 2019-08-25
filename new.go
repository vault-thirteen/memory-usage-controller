////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019 by Vault Thirteen.
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
