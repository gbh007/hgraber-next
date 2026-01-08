package mcp

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/gbh007/hgraber-next/domain/bff"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/pkg"
)

type bffUseCases interface {
	BookDetails(ctx context.Context, bookID uuid.UUID) (bff.BookDetails, error)
	BookList(ctx context.Context, filter core.BookFilter) (bff.BookList, error)
}

type Controller struct {
	logger      *slog.Logger
	bffUseCases bffUseCases
	addr        string
}

func New(
	logger *slog.Logger,
	bffUseCases bffUseCases,
	addr string,
) *Controller {
	return &Controller{
		bffUseCases: bffUseCases,
		logger:      logger,
		addr:        addr,
	}
}

func (c *Controller) Start(ctx context.Context) (chan struct{}, error) {
	done := make(chan struct{})

	s := server.NewMCPServer(
		"hgraber-next",
		"0.0.1",
	)

	s.AddTool(
		mcp.NewTool(
			"book list",
			mcp.WithDescription("get book list by filter"),
			mcp.WithString(
				"author",
				mcp.Required(),
				mcp.Description("book`s author"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			list, err := c.bffUseCases.BookList(ctx, core.BookFilter{
				ShowDeleted:      core.BookFilterShowTypeExcept,
				ShowWithoutPages: core.BookFilterShowTypeExcept,
				Fields: core.BookFilterFields{
					Attributes: []core.BookFilterAttribute{
						{
							Code:   "author",
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
					return bookData{
						ID:        b.Book.ID,
						Name:      b.Book.Name,
						PageCount: b.Book.PageCount,
					}
				}),
			}), nil
		},
	)

	s.AddTool(
		mcp.NewTool(
			"book details",
			mcp.WithDescription("get book details by id"),
			mcp.WithString(
				"id",
				mcp.Required(),
				mcp.Description("book`s id"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
				if attr.Code != "tag" {
					continue
				}

				for _, v := range attr.Values {
					result.Tags = append(result.Tags, v.Name)
				}
			}

			return mcp.NewToolResultStructuredOnly(result), nil
		},
	)

	httpServer := server.NewStreamableHTTPServer(s)

	go func() {
		defer close(done)

		err := httpServer.Start(c.addr)
		if err != nil {
			c.logger.ErrorContext(
				ctx,
				"failed serve mcp",
				slog.String("error", err.Error()),
			)
		}
	}()

	go func() {
		<-ctx.Done()

		sCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		err := httpServer.Shutdown(sCtx)
		if err != nil {
			c.logger.ErrorContext(
				ctx,
				"failed shutdown mcp",
				slog.String("error", err.Error()),
			)
		}
	}()

	return done, nil
}

func (c *Controller) Name() string {
	return "mcp server"
}
