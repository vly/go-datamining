package utilities

// imports and writes arff files for Weka
type ArffReader struct {
	Data *Dataset
}

func (r *ArffReader) LoadArff(file string) bool {
	return true
}

func (r *ArffReader) ExportArff(file string) bool {
	return true
}
