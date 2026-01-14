package mcp

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func (c *Controller) hProxyBookTool() server.ServerTool {
	return server.ServerTool{
		Tool: mcp.NewTool(
			"origin book details",
			mcp.WithDescription("get origin book details by url, including attributes (tags etc)"),
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

			pageLimit := 0

			book, err := c.hProxyUseCases.Book(ctx, *url, &pageLimit)
			if err != nil {
				return nil, fmt.Errorf("get book: %w", err)
			}

			result := hProxyBookData{
				Name:      book.Name,
				PageCount: book.PageCount,
				SystemIDs: book.ExistsIDs,
				OriginURL: book.ExURL.String(),
			}

			for _, attr := range book.Attributes {
				switch attr.Code {
				case core.AttributeCodeTag:
					for _, v := range attr.Values {
						var u string

						if v.ExtURL != nil {
							u = v.ExtURL.String()
						}

						result.Tags = append(result.Tags, hProxyAttributeValue{
							Value:     v.Name,
							OriginURL: u,
						})
					}

				case core.AttributeCodeAuthor:
					for _, v := range attr.Values {
						var u string

						if v.ExtURL != nil {
							u = v.ExtURL.String()
						}

						result.Authors = append(result.Authors, hProxyAttributeValue{
							Value:     v.Name,
							OriginURL: u,
						})
					}

				case core.AttributeCodeCategory:
					for _, v := range attr.Values {
						var u string

						if v.ExtURL != nil {
							u = v.ExtURL.String()
						}

						result.Categories = append(result.Categories, hProxyAttributeValue{
							Value:     v.Name,
							OriginURL: u,
						})
					}

				case core.AttributeCodeCharacter:
					for _, v := range attr.Values {
						var u string

						if v.ExtURL != nil {
							u = v.ExtURL.String()
						}

						result.Characters = append(result.Characters, hProxyAttributeValue{
							Value:     v.Name,
							OriginURL: u,
						})
					}

				case core.AttributeCodeGroup:
					for _, v := range attr.Values {
						var u string

						if v.ExtURL != nil {
							u = v.ExtURL.String()
						}

						result.Groups = append(result.Groups, hProxyAttributeValue{
							Value:     v.Name,
							OriginURL: u,
						})
					}

				case core.AttributeCodeLanguage:
					for _, v := range attr.Values {
						var u string

						if v.ExtURL != nil {
							u = v.ExtURL.String()
						}

						result.Languages = append(result.Languages, hProxyAttributeValue{
							Value:     v.Name,
							OriginURL: u,
						})
					}

				case core.AttributeCodeParody:
					for _, v := range attr.Values {
						var u string

						if v.ExtURL != nil {
							u = v.ExtURL.String()
						}

						result.Parodies = append(result.Parodies, hProxyAttributeValue{
							Value:     v.Name,
							OriginURL: u,
						})
					}
				}
			}

			return mcp.NewToolResultStructuredOnly(result), nil
		},
	}
}
