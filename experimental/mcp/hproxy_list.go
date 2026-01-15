package mcp

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func (c *Controller) hProxyListTool() server.ServerTool {
	return server.ServerTool{
		Tool: mcp.NewTool(
			"hproxy list books",
			mcp.WithDescription("get origin books by url, without attributes (tags etc)"),
			mcp.WithString(
				"url",
				mcp.Required(),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			url, err := url.Parse(request.GetString("url", ""))
			if err != nil {
				return nil, fmt.Errorf("parse book id: %w", err)
			}

			list, err := c.hProxyUseCases.List(ctx, *url)
			if err != nil {
				return nil, fmt.Errorf("get book: %w", err)
			}

			result := hProxyListBookData{}

			if list.NextPage != nil {
				result.NextPage = list.NextPage.String()
			}

			for _, page := range list.Pagination {
				result.Pages = append(result.Pages, hProxyValue{
					Value:     page.Name,
					OriginURL: page.ExtURL.String(),
				})
			}

			for _, book := range list.Books {
				result.Books = append(result.Books, hProxyBookData{
					Name:       book.Name,
					SystemIDs:  book.ExistsIDs,
					Downloaded: len(book.ExistsIDs) > 0,
					OriginURL:  book.ExtURL.String(),
				})
			}

			return mcp.NewToolResultStructuredOnly(result), nil
		},
	}
}
