package gojapanese

import (
	"testing"
)


func TestNoun_Negative(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Happy Path", fields{"石"}, "石じゃない",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			N := &Noun{
				Name: tt.fields.Name,
			}
			if got := N.Negative(); got != tt.want {
				t.Errorf("Negative() = %v, want %v", got, tt.want)
			}
		})
	}
}