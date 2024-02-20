package model

import "testing"

func TestError(t *testing.T) {
	table := []struct {
		name       string
		createFunc func(error) PocError
		isFunc     func(error) bool
	}{
		{
			name:       "NotFoundError",
			createFunc: NewNotFoundError,
			isFunc:     IsNotFoundError,
		},
		{
			name:       "InvalidInputError",
			createFunc: NewInvalidInputError,
			isFunc:     IsInvalidInputError,
		},
	}

	for _, tt := range table {
		err := tt.createFunc(nil)
		if !tt.isFunc(err) {
			t.Errorf("Expected error to be of the correct type: %s", tt.name)
		}
	}
}
