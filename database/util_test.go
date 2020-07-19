package database

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestDBValues(t *testing.T) {
	testCases := []struct {
		name   string
		config Config
		want   map[string]string
	}{
		{
			name: "empty config",
			want: make(map[string]string),
		},
		{
			name: "some config",
			config: Config{
				Name:     "testingDb",
				Password: "password123",
				Port:     "3242",
			},
			want: map[string]string{
				"dbname":   "testingDb",
				"password": "password123",
				"port":     "3242",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := dbValues(&tc.config)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("mismatch (-want, +got): \n%s", diff)
			}
		})
	}
}
