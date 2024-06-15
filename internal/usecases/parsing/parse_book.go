package parsing

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (uc *UseCase) ParseBook(ctx context.Context, agentID uuid.UUID, book entities.Book) error {
	agentCtx, agentCancel := context.WithTimeout(ctx, ParseBookAgentTimeout)
	defer agentCancel()

	if book.OriginURL == nil {
		return fmt.Errorf("missing url")
	}

	info, err := uc.agentSystem.BookParse(agentCtx, agentID, *book.OriginURL)
	if err != nil {
		return fmt.Errorf("agent parse: %w", err)
	}

	for _, attr := range info.Attributes {
		err = uc.storage.UpdateAttribute(ctx, book.ID, attr.Code, attr.Values)
		if err != nil {
			return fmt.Errorf("update attribute (%s): %w", attr.Code, err)
		}
	}

	err = uc.storage.UpdateBookPages(ctx, book.ID, pkg.Map(info.Pages, func(p entities.AgentBookDetailsPagesItem) entities.Page {
		return entities.Page{
			BookID:     book.ID,
			PageNumber: p.PageNumber,
			Ext:        path.Ext(p.Filename),
			OriginURL:  &p.URL,
			CreateAt:   time.Now(),
		}
	}))
	if err != nil {
		return fmt.Errorf("update pages: %w", err)
	}

	book.AttributesParsed = true
	book.PageCount = info.PageCount

	err = uc.storage.UpdateBook(ctx, book)
	if err != nil {
		return fmt.Errorf("update book: %w", err)
	}

	return nil
}
