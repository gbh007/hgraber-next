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
				"filter.name",
				mcp.Description("name for filtering by ILIKE"),
			),
			mcp.WithArray(
				"filter.authors",
				mcp.WithStringItems(),
			),
			mcp.WithArray(
				"filter.categories",
				mcp.WithStringItems(),
			),
			mcp.WithArray(
				"filter.characters",
				mcp.WithStringItems(),
			),
			mcp.WithArray(
				"filter.groups",
				mcp.WithStringItems(),
			),
			mcp.WithArray(
				"filter.languages",
				mcp.WithStringItems(),
			),
			mcp.WithArray(
				"filter.parodies",
				mcp.WithStringItems(),
			),
			mcp.WithArray(
				"filter.tags",
				mcp.WithStringItems(),
			),
			mcp.WithString(
				"sort.by",
				mcp.Enum(
					"page count",
					"name",
					"creation date",
				),
			),
			mcp.WithString(
				"sort.order",
				mcp.Enum(
					"asc",
					"desc",
				),
				mcp.DefaultString("asc"),
			),
			mcp.WithNumber("limit"),
			mcp.WithNumber("offset"),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			filter := core.BookFilter{
				ShowDeleted:      core.BookFilterShowTypeExcept,
				ShowWithoutPages: core.BookFilterShowTypeExcept,
				Fields: core.BookFilterFields{
					Name: request.GetString("filter.name", ""),
				},
				Desc:   request.GetString("sort order", "") == "desc",
				Limit:  request.GetInt("limit", 0),
				Offset: request.GetInt("offset", 0),
			}

			if v := request.GetStringSlice("filter.authors", nil); len(v) > 0 {
				filter.Fields.Attributes = append(filter.Fields.Attributes, core.BookFilterAttribute{
					Code:   core.AttributeCodeAuthor,
					Type:   core.BookFilterAttributeTypeLike,
					Values: v,
				})
			}

			if v := request.GetStringSlice("filter.categories", nil); len(v) > 0 {
				filter.Fields.Attributes = append(filter.Fields.Attributes, core.BookFilterAttribute{
					Code:   core.AttributeCodeCategory,
					Type:   core.BookFilterAttributeTypeLike,
					Values: v,
				})
			}

			if v := request.GetStringSlice("filter.characters", nil); len(v) > 0 {
				filter.Fields.Attributes = append(filter.Fields.Attributes, core.BookFilterAttribute{
					Code:   core.AttributeCodeCharacter,
					Type:   core.BookFilterAttributeTypeLike,
					Values: v,
				})
			}

			if v := request.GetStringSlice("filter.groups", nil); len(v) > 0 {
				filter.Fields.Attributes = append(filter.Fields.Attributes, core.BookFilterAttribute{
					Code:   core.AttributeCodeGroup,
					Type:   core.BookFilterAttributeTypeLike,
					Values: v,
				})
			}

			if v := request.GetStringSlice("filter.languages", nil); len(v) > 0 {
				filter.Fields.Attributes = append(filter.Fields.Attributes, core.BookFilterAttribute{
					Code:   core.AttributeCodeLanguage,
					Type:   core.BookFilterAttributeTypeLike,
					Values: v,
				})
			}

			if v := request.GetStringSlice("filter.parodies", nil); len(v) > 0 {
				filter.Fields.Attributes = append(filter.Fields.Attributes, core.BookFilterAttribute{
					Code:   core.AttributeCodeParody,
					Type:   core.BookFilterAttributeTypeLike,
					Values: v,
				})
			}

			if v := request.GetStringSlice("filter.tags", nil); len(v) > 0 {
				filter.Fields.Attributes = append(filter.Fields.Attributes, core.BookFilterAttribute{
					Code:   core.AttributeCodeTag,
					Type:   core.BookFilterAttributeTypeLike,
					Values: v,
				})
			}

			switch request.GetString("sort by", "") {
			case "page count":
				filter.OrderBy = core.BookFilterOrderByPageCount
			case "name":
				filter.OrderBy = core.BookFilterOrderByName
			case "creation date":
				filter.OrderBy = core.BookFilterOrderByCreated
			}

			list, err := c.bffUseCases.BookList(ctx, filter)
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
