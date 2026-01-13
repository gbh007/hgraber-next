package mcp

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) attributesCountTool() server.ServerTool {
	return server.ServerTool{
		Tool: mcp.NewTool(
			"attributes count",
			mcp.WithDescription("get attributes fount by filter, has desc order"),
			mcp.WithString("filter.type"),
			mcp.WithNumber("limit"),
			mcp.WithNumber("offset"),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			variants, err := c.attrUseCases.AttributesCount(ctx)
			if err != nil {
				return nil, fmt.Errorf("get attributes count: %w", err)
			}

			// FIXME: перенести фильтрацию в репозиторий
			if code := request.GetString("filter.type", ""); code != "" {
				variants = pkg.SliceFilter(variants, func(av core.AttributeVariant) bool {
					return av.Code == code
				})
			}

			if offset := request.GetInt("offset", 0); offset > len(variants) {
				variants = nil
			} else if offset > 0 {
				variants = variants[offset:]
			}

			if limit := request.GetInt("limit", 0); limit < len(variants) && limit > 0 {
				variants = variants[:limit]
			}

			return mcp.NewToolResultStructuredOnly(map[string]any{
				"attributes": pkg.Map(variants, func(av core.AttributeVariant) attributeData {
					return attributeData{
						Type:  av.Code,
						Value: av.Value,
						Count: av.Count,
					}
				}),
			}), nil
		},
	}
}
