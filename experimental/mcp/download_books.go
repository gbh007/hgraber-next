package mcp

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/parsing"
	"github.com/gbh007/hgraber-next/pkg"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func (c *Controller) downloadBooksTool() server.ServerTool {
	return server.ServerTool{
		Tool: mcp.NewTool(
			"download books",
			mcp.WithDescription("download books to system by urls, by direct links or attribute links"),
			mcp.WithArray(
				"urls",
				mcp.Required(),
				mcp.WithStringItems(),
			),
			mcp.WithString(
				"links type",
				mcp.Required(),
				mcp.Enum(
					"book",
					"attribute",
				),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			urls, err := pkg.MapWithError(request.GetStringSlice("urls", nil), func(s string) (url.URL, error) {
				u, err := url.Parse(s)
				if err != nil {
					return url.URL{}, fmt.Errorf("parse url: %w", err)
				}

				return *u, nil
			})
			if err != nil {
				return nil, fmt.Errorf("parse urls: %w", err)
			}

			switch request.GetString("links type", "") {
			case "book":
				result, err := c.bookParserUseCases.NewBooks(ctx, urls, parsing.ParseFlags{})
				if err != nil {
					return nil, fmt.Errorf("download books: %w", err)
				}

				return mcp.NewToolResultStructuredOnly(map[string]any{
					"downloaded": result.LoadedCount,
					"duplicates": result.DuplicateCount,
					"failed":     result.ErrorCount,
				}), nil
			case "attribute":
				result, err := c.bookParserUseCases.NewBooksMulti(ctx, urls, parsing.ParseFlags{})
				if err != nil {
					return nil, fmt.Errorf("download books: %w", err)
				}

				return mcp.NewToolResultStructuredOnly(map[string]any{
					"downloaded": result.Details.LoadedCount,
					"duplicates": result.Details.DuplicateCount,
					"failed":     result.Details.ErrorCount,
				}), nil
			}

			return nil, fmt.Errorf("unknown link type: %s", request.GetString("links type", ""))
		},
	}
}
