package defaults

import (
	"prophetstor.com/alameda/datahub/pkg/entities"
)

func DefaultSchemaFedemeter() []interface{} {
	schemas := make([]interface{}, 0)

	// Calculation
	schemas = append(schemas, &entities.FedemeterCalculationInstance{})
	schemas = append(schemas, &entities.FedemeterCalculationStorage{})

	// Recommendation
	schemas = append(schemas, &entities.FedemeterRecommendationJERI{})

	// Resource history
	schemas = append(schemas, &entities.FedemeterResourceHistoryCostApp{})
	schemas = append(schemas, &entities.FedemeterResourceHistoryCostNamespace{})

	// Resource prediction
	schemas = append(schemas, &entities.FedemeterResourcePredictionCostApp{})
	schemas = append(schemas, &entities.FedemeterResourcePredictionCostNamespace{})

	return schemas
}
