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
