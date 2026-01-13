package mcp

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/gbh007/hgraber-next/domain/bff"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) bookListTool() server.ServerTool {
	return server.ServerTool{
		Tool: mcp.NewTool(
			"book list",
			mcp.WithDescription("get book list by filter"),
			mcp.WithString(
				"author",
				mcp.Required(),
				mcp.Description("book`s author"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			list, err := c.bffUseCases.BookList(ctx, core.BookFilter{
				ShowDeleted:      core.BookFilterShowTypeExcept,
				ShowWithoutPages: core.BookFilterShowTypeExcept,
				Fields: core.BookFilterFields{
					Attributes: []core.BookFilterAttribute{
						{
							Code:   core.AttributeCodeAuthor,
							Type:   core.BookFilterAttributeTypeLike,
							Values: []string{request.GetString("author", "")},
						},
					},
				},
			})
			if err != nil {
				return nil, fmt.Errorf("get book list: %w", err)
			}

			return mcp.NewToolResultStructuredOnly(map[string]any{
				"books": pkg.Map(list.Books, func(b bff.BookShort) bookData {
					data := bookData{
						ID:        b.Book.ID,
						Name:      b.Book.Name,
						PageCount: b.Book.PageCount,
					}

					if values, ok := b.AttributesRaw[core.AttributeCodeTag]; ok {
						data.Tags = values
					}

					if values, ok := b.AttributesRaw[core.AttributeCodeAuthor]; ok {
						data.Authors = values
					}

					if values, ok := b.AttributesRaw[core.AttributeCodeCategory]; ok {
						data.Categories = values
					}

					if values, ok := b.AttributesRaw[core.AttributeCodeCharacter]; ok {
						data.Characters = values
					}

					if values, ok := b.AttributesRaw[core.AttributeCodeGroup]; ok {
						data.Groups = values
					}

					if values, ok := b.AttributesRaw[core.AttributeCodeLanguage]; ok {
						data.Languages = values
					}

					if values, ok := b.AttributesRaw[core.AttributeCodeParody]; ok {
						data.Parodies = values
					}

					return data
				}),
			}), nil
		},
	}
}
