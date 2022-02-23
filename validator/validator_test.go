package validator

import (
	"reflect"
	"testing"
)

func TestValidator(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"empty":            {input: "", want: false},
		"invalid_start":    {input: "}[]{}", want: false},
		"invalid_end":      {input: "{[]}(", want: false},
		"invalid_sequence": {input: "{{[}]}", want: false},
		"valid":            {input: "{()[]}{()}", want: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Validate(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
