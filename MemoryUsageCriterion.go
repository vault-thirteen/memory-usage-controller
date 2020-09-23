// MemoryUsageCriterion.go.

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

// Memory Usage Criterion.
const (
	MemoryUsageCriterionWorkingSet     = 1
	MemoryUsageCriterionResidentMemory = 2
)

type MemoryUsageCriterion byte

func NewMemoryUsageCriterion(
	c byte,
) MemoryUsageCriterion {
	return MemoryUsageCriterion(c)
}

func (c MemoryUsageCriterion) IsValid() bool {
	if (c == MemoryUsageCriterionWorkingSet) ||
		(c == MemoryUsageCriterionResidentMemory) {
		return true
	}
	return false
}
