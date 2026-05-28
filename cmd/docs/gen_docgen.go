package docs

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/toaweme/docgen"
	"github.com/toaweme/docgen/core/extractor"
	"github.com/toaweme/docgen/sdk"
)

// GenerateWithDocgen runs sintax documentation generation using the docgen SDK.
// This replaces the inline DocsCommand.Run logic with composable SDK calls.
func GenerateWithDocgen(ctx context.Context, dir string) error {
	client := sdk.New()

	content, err := client.Read(ctx, docgen.Source{Dir: dir})
	if err != nil {
		return fmt.Errorf("failed to read sintax source: %w", err)
	}

	ext := extractor.NewConstFunc(extractor.ConstFuncConfig{
		ConstPrefix: "ModifierName",
		Groups:      docgenGroups(),
	})
	model, err := client.Extract(content, ext)
	if err != nil {
		return fmt.Errorf("failed to extract modifiers: %w", err)
	}

	model = client.ApplyGroups(model, docgen.Groups{
		Items: docgenGroups(),
		Order: groupOrder,
	})

	mdx, err := client.RenderMDX(model, docgen.MDXRendering{
		Grouping: "category",
		Indexes:  true,
	})
	if err != nil {
		return fmt.Errorf("failed to render modifier MDX: %w", err)
	}

	return client.WriteFiles(mdx, filepath.Join(dir, "_data/docs/sintax"))
}

func docgenGroups() map[string]docgen.Group {
	out := make(map[string]docgen.Group, len(groups))
	for k, v := range groups {
		out[k] = docgen.Group{
			Title:       v.Title,
			Description: v.Description,
		}
	}
	return out
}
