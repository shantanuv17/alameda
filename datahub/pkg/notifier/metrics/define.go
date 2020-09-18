package metrics

type CriteriaType int

const (
	CriteriaTypeUndefined  CriteriaType = 0
	CriteriaTypeGauge      CriteriaType = 1
	CriteriaTypeContinuous CriteriaType = 2
)
