package lib_test

import (
	"testing"

	"github.com/sk1t0n/echo-mvc-generator/lib"
)

func TestGetEntityName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		path   string
		format int
		want   string
	}{
		{"file:pascal_case", "BlogPostController", lib.FormatEntityNamePascalCase, "BlogPost"},
		{"file:pascal_case", "BlogPostController", lib.FormatEntityNameSnakeCase, "blog_post"},
		{"file:snake_case", "blog_post_controller", lib.FormatEntityNameSnakeCase, "blog_post"},
		{"file:snake_case", "blog_post_controller", lib.FormatEntityNamePascalCase, "BlogPost"},
		{
			"file_with_dirs:snake_case",
			"controllers/home_controller",
			lib.FormatEntityNamePascalCase,
			"Home",
		},
		{
			"file_with_dirs:pascal_case",
			"./controllers/HomeController.go",
			lib.FormatEntityNamePascalCase,
			"Home",
		},
		{"file:lower_case", "user", lib.FormatEntityNamePascalCase, "User"},
		{"file:pascal_case", "User", lib.FormatEntityNamePascalCase, "User"},
		{"file_with_dirs:lower_case", "models/user", lib.FormatEntityNamePascalCase, "User"},
		{"file_with_dirs:pascal_case", "./models/User.go", lib.FormatEntityNamePascalCase, "User"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lib.GetEntityName(tt.path, tt.format)
			if got != tt.want {
				t.Errorf("GetEntityName(%s, %d) = %s, want = %s", tt.path, tt.format, got, tt.want)
			}
		})
	}
}

func TestIsLower(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"lower_case", "test", true},
		{"pascal_case", "Test", false},
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
	t.Parallel()

	tests := []struct {
		name string
		s    string
		want string
	}{
		{"ends_with_a_slash", "test/", "test"},
		{"ends_with_a_backslash", "test\\", "test"},
		{"does_not_end_with_a_slash", "test", "test"},
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
