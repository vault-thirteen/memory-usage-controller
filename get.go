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
	"runtime"

	"github.com/prometheus/procfs"
)

// Get the Memory Usage Amount. Result is set in MB.
func (muc *MemoryUsageController) GetMemoryUsage() (usageMb uint, err error) {
	return muc.getMemoryUsage()
}

// Get the Memory Usage Amount. Result is set in MB.
func (muc *MemoryUsageController) getMemoryUsage() (usageMb uint, err error) {

	muc.usageLock.Lock()
	defer muc.usageLock.Unlock()

	switch muc.memoryUsageCriterion {

	case MemoryUsageCriterionWorkingSet:
		return muc.getMemoryUsageWorkingSet()

	case MemoryUsageCriterionResidentMemory:
		return muc.getMemoryUsageResident()

	default:
		err = errors.New(ErrMemoryUsageCriterionUnknown)
		return
	}
}

// Returns the Memory Usage as the "Working Set".
// This Parameter is often used in Windows OS.
func (muc *MemoryUsageController) getMemoryUsageWorkingSet() (usageMb uint, err error) {

	var memoryUsageStatistics *runtime.MemStats

	// Update the Memory Usage Statistics using the built-in Go Mechanism.
	memoryUsageStatistics = new(runtime.MemStats)
	runtime.ReadMemStats(memoryUsageStatistics)

	// Calculate the Working Set.
	return uint(memoryUsageStatistics.HeapInuse+
		memoryUsageStatistics.StackInuse+
		memoryUsageStatistics.MSpanInuse+
		memoryUsageStatistics.MCacheInuse+
		memoryUsageStatistics.BuckHashSys+
		memoryUsageStatistics.GCSys+
		memoryUsageStatistics.OtherSys) / MB, nil
}

// Returns the Memory Usage as the "Resident Memory".
// This Parameter is often used in Linux OS.
func (muc *MemoryUsageController) getMemoryUsageResident() (usageMb uint, err error) {

	var osProcess procfs.Proc
	var processStatistics procfs.ProcStat

	// Get the Memory Usage Statistics using the external Library which makes
	// a special Call to the Operating System. Such a Call is not supported in
	// Operating Systems of the 'Windows' Family.
	osProcess, err = procfs.Self()
	if err != nil {
		return
	}
	processStatistics, err = osProcess.Stat()
	if err != nil {
		return
	}

	// Calculate the resident Memory.
	return uint(processStatistics.ResidentMemory()) / MB, nil
}
