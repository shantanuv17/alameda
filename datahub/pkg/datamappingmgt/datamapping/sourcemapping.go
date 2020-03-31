package datamapping

type SourceMapping struct {
	Source  Source
	Mapping string
}

func NewSourceMapping(source Source, mapping string) *SourceMapping {
	sourceMapping := SourceMapping{}
	sourceMapping.Source = source
	sourceMapping.Mapping = mapping
	return &sourceMapping
}
