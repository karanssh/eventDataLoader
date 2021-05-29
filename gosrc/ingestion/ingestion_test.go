package ingestion

import "testing"

func Test_readCSVFile(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "default test case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// readCSVFile() disabled
		})
	}
}
