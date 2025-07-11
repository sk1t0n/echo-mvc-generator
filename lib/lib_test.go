package lib_test

import (
	"testing"

	"github.com/sk1t0n/echo-mvc-generator/lib"
)

func TestGetEntityName(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{"file: snake case", "home_controller", "HomeController"},
		{"file: pascal case", "HomeController", "HomeController"},
		{"file with dirs: snake case", "controllers/home_controller", "HomeController"},
		{"file with dirs: pascal case", "./controllers/HomeController.go", "HomeController"},
		{"file: lower case", "user", "User"},
		{"file: pascal case", "User", "User"},
		{"file with dirs: lower case", "models/user", "User"},
		{"file with dirs: pascal case", "./models/User.go", "User"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lib.GetEntityName(tt.path)
			if got != tt.want {
				t.Errorf("GetEntityName(%s) = %s, want = %s", tt.path, got, tt.want)
			}
		})
	}
}
