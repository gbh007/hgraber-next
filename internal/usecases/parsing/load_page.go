package parsing

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

func (uc *UseCase) DownloadPage(ctx context.Context, agentID uuid.UUID, page entities.PageForDownload) error {
	if page.BookURL == nil || page.ImageURL == nil {
		return fmt.Errorf("invalid page")
	}

	body, err := uc.agentSystem.PageLoad(ctx, agentID, entities.AgentPageURL{
		BookURL:  *page.BookURL,
		ImageURL: *page.ImageURL,
	})
	if err != nil {
		return fmt.Errorf("agent load: %w", err)
	}

	fileID := uuid.Must(uuid.NewV7())

	err = uc.fileStorage.Create(ctx, fileID, body)
	if err != nil {
		return fmt.Errorf("store file (%s): %w", fileID.String(), err)
	}

	err = uc.storage.NewFile(ctx, entities.File{
		ID:       fileID,
		Filename: fmt.Sprintf("%d%s", page.PageNumber, page.Ext),
		Ext:      page.Ext,
		CreateAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("create file in db (%s): %w", fileID.String(), err)
	}

	err = uc.storage.UpdatePageDownloaded(ctx, page.BookID, page.PageNumber, true, fileID)
	if err != nil {
		return fmt.Errorf("update page with file (%s): %w", fileID.String(), err)
	}

	return nil
}
