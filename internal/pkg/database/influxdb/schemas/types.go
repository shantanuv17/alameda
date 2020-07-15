package schemas

type SchemaMeta struct {
	Scope    Scope
	Category string
	Type     string
}

func NewSchemaMeta(scope Scope, category, schemaType string) *SchemaMeta {
	schemaMeta := SchemaMeta{}
	schemaMeta.Scope = scope
	schemaMeta.Category = category
	schemaMeta.Type = schemaType
	return &schemaMeta
}

func CompareSchemaMeta(old *SchemaMeta, new *SchemaMeta) bool {
	if old.Scope == new.Scope && old.Category == new.Category && old.Type == new.Type {
		return true
	}
	return false
}

func CompareMeasurement(old *Measurement, new *Measurement) bool {
	if old.Name == new.Name && old.MetricType == new.MetricType && old.Boundary == new.Boundary && old.Quota == new.Quota {
		return true
	}
	return false
}
