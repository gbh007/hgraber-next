package mcp

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/google/uuid"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func (c *Controller) bookDetailsTool() server.ServerTool {
	return server.ServerTool{
		Tool: mcp.NewTool(
			"book details",
			mcp.WithDescription("get book details by id, including attributes (tags etc)"),
			mcp.WithString(
				"id",
				mcp.Required(),
				mcp.Description("book`s id"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			id, err := uuid.Parse(request.GetString("id", ""))
			if err != nil {
				return nil, fmt.Errorf("parse book id: %w", err)
			}

			book, err := c.bffUseCases.BookDetails(ctx, id)
			if err != nil {
				return nil, fmt.Errorf("get book: %w", err)
			}

			result := bookData{
				ID:        book.Book.ID,
				Name:      book.Book.Name,
				PageCount: book.Book.PageCount,
			}

			for _, attr := range book.Attributes {
				if attr.Code != core.AttributeCodeTag {
					continue
				}

				for _, v := range attr.Values {
					result.Tags = append(result.Tags, v.Name)
				}
			}

			return mcp.NewToolResultStructuredOnly(result), nil
		},
	}
}
