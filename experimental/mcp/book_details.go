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

			if book.Book.OriginURL != nil {
				result.OriginURL = book.Book.OriginURL.String()
			}

			for _, attr := range book.Attributes {
				switch attr.Code {
				case core.AttributeCodeTag:
					for _, v := range attr.Values {
						result.Tags = append(result.Tags, v.Name)
					}

				case core.AttributeCodeAuthor:
					for _, v := range attr.Values {
						result.Authors = append(result.Authors, v.Name)
					}

				case core.AttributeCodeCategory:
					for _, v := range attr.Values {
						result.Categories = append(result.Categories, v.Name)
					}

				case core.AttributeCodeCharacter:
					for _, v := range attr.Values {
						result.Characters = append(result.Characters, v.Name)
					}

				case core.AttributeCodeGroup:
					for _, v := range attr.Values {
						result.Groups = append(result.Groups, v.Name)
					}

				case core.AttributeCodeLanguage:
					for _, v := range attr.Values {
						result.Languages = append(result.Languages, v.Name)
					}

				case core.AttributeCodeParody:
					for _, v := range attr.Values {
						result.Parodies = append(result.Parodies, v.Name)
					}
				}
			}

			return mcp.NewToolResultStructuredOnly(result), nil
		},
	}
}
