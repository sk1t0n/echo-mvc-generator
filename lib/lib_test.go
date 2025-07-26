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

func TestIsLower(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"lower case", "test", true},
		{"pascal case", "Test", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lib.IsLower(tt.s)
			if got != tt.want {
				t.Errorf("IsLower(%s) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestRemoveLastSlash(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"ends with a slash", "test/", "test"},
		{"ends with a backslash", "test\\", "test"},
		{"does not end with a slash", "test", "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lib.RemoveLastSlash(tt.s)
			if got != tt.want {
				t.Errorf("RemoveLastSlash(%s) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
