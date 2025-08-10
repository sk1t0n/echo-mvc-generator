package cmd

import (
	"os"
	"strings"
	"testing"

	"github.com/sk1t0n/fiber-mvc-generator/lib"
)

func Test_updateRoutes(t *testing.T) {
	tests := []struct {
		name      string
		f         string
		modelName string
		wantErr   bool
	}{
		{"case1", "routes.go", "blog_post", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := "func (r *Router) RegisterRoutes() {}"
			err := os.WriteFile(tt.f, []byte(data), 0666)
			if err != nil {
				t.Fatalf("updateRoutes(%s, %s) failed: %v", tt.f, tt.modelName, err)
			}

			gotErr := updateRoutes(tt.f, tt.modelName)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("updateRoutes(%s, %s) failed: %v", tt.f, tt.modelName, gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatalf("updateRoutes(%s, %s) succeeded unexpectedly", tt.f, tt.modelName)
			}

			content, err := os.ReadFile(tt.f)
			if err != nil {
				t.Fatalf("updateRoutes(%s, %s) failed: %v", tt.f, tt.modelName, err)
			}

			if !strings.Contains(
				string(content),
				lib.GetEntityName(tt.modelName, lib.FormatEntityNamePascalCase)+"Controller())\n}",
			) {
				t.Errorf("updateRoutes(%s, %s), content is invalid", tt.f, tt.modelName)
			}
		})
	}

	t.Cleanup(func() {
		_ = os.Remove("routes.go")
	})
}
