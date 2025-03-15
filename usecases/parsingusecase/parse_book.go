package parsingusecase

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/pkg"
)

func (uc *UseCase) ParseBook(ctx context.Context, agentID uuid.UUID, book core.Book) error {
	agentCtx, agentCancel := context.WithTimeout(ctx, uc.parseBookTimeout)
	defer agentCancel()

	if book.OriginURL == nil {
		return fmt.Errorf("missing url")
	}

	info, err := uc.agentSystem.BookParse(agentCtx, agentID, *book.OriginURL)
	if err != nil {
		return fmt.Errorf("agent parse: %w", err)
	}

	if len(info.Attributes) > 0 {
		attributes := make(map[string][]string, 7)

		for _, attr := range info.Attributes {
			attributes[attr.Code] = attr.Values
		}

		err = uc.storage.UpdateOriginAttributes(ctx, book.ID, attributes)
		if err != nil {
			return fmt.Errorf("update original attributes: %w", err)
		}

		if uc.autoRemap {
			remaps, err := uc.storage.AttributeRemaps(ctx)
			if err != nil {
				return fmt.Errorf("storage: get attributes remaps: %w", err)
			}

			remaper := core.NewAttributeRemaper(remaps, uc.remapToLower)
			attributes = remaper.Remap(attributes)
		}

		if len(attributes) > 0 {
			err = uc.storage.UpdateAttributes(ctx, book.ID, attributes)
			if err != nil {
				return fmt.Errorf("update attributes: %w", err)
			}
		}
	}

	err = uc.storage.UpdateBookPages(ctx, book.ID, pkg.Map(info.Pages, func(p agentmodel.AgentBookDetailsPagesItem) core.Page {
		return core.Page{
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

	book.Name = info.Name
	book.AttributesParsed = true
	book.PageCount = info.PageCount

	err = uc.storage.UpdateBook(ctx, book)
	if err != nil {
		return fmt.Errorf("update book: %w", err)
	}

	return nil
}
