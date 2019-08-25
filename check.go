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
