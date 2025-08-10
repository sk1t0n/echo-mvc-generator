package lib_test

import (
	"testing"

	"github.com/sk1t0n/fiber-mvc-generator/lib"
)

func TestGetEntityName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		path   string
		format int
		want   string
	}{
		{
			"file:snake_case_to_lower_case",
			"blog_category_controller",
			lib.FormatEntityNameLowerCase,
			"blogcategory",
		},
		{
			"file:snake_case_to_pascal_case",
			"blog_category_controller",
			lib.FormatEntityNamePascalCase,
			"BlogCategory",
		},
		{
			"file_with_dirs:snake_case_to_lower_case",
			"controllers/home_controller",
			lib.FormatEntityNameLowerCase,
			"home",
		},
		{
			"file_with_dirs:snake_case_to_pascal_case",
			"./controllers/home_controller.go",
			lib.FormatEntityNamePascalCase,
			"Home",
		},
		{
			"file_with_dirs:lower_case_to_pascal_case",
			"internal/entity/user",
			lib.FormatEntityNamePascalCase,
			"User",
		},
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
