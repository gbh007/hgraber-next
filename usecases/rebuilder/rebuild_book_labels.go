package rebuilder

import (
	"context"
	"time"

	"github.com/gbh007/hgraber-next/entities"
)

func (uc *UseCase) rebuildBookLabels(
	_ context.Context,
	bookToMerge entities.Book,
	sourceBook entities.Book,
	flags entities.RebuildBookRequestFlags,
	labelsFromRequest []entities.BookLabel,
	pagesInfo rebuildedPagesInfo,
) ([]entities.BookLabel, error) {
	type labelInBookKey struct {
		Name       string
		PageNumber int
	}

	newLabels := make([]entities.BookLabel, 0, len(labelsFromRequest))
	existsNewLabels := make(map[labelInBookKey]struct{}, len(labelsFromRequest))

	for _, label := range labelsFromRequest {
		newPageNumber, ok := pagesInfo.PagesRemap[label.PageNumber]
		if !ok && label.PageNumber != 0 { // Отсекаем данные которые не были замаплены или не привязаны к книге.
			continue
		}

		label.BookID = bookToMerge.ID
		label.PageNumber = newPageNumber

		newLabels = append(newLabels, label)

		existsNewLabels[labelInBookKey{
			Name:       label.Name,
			PageNumber: label.PageNumber,
		}] = struct{}{}
	}

	if !flags.SetOriginLabels {
		return newLabels, nil
	}

	for _, oldPageNumber := range pagesInfo.SourcePageNumbers {
		newPageNumber, hasRemap := pagesInfo.PagesRemap[oldPageNumber]
		if !hasRemap && oldPageNumber != 0 { // Отсекаем данные которые не были замаплены или не привязаны к книге.
			continue
		}

		_, hasOriginID := existsNewLabels[labelInBookKey{
			Name:       entities.LabelNameRebuildOriginID,
			PageNumber: newPageNumber,
		}]

		_, hasOriginName := existsNewLabels[labelInBookKey{
			Name:       entities.LabelNameRebuildOriginName,
			PageNumber: newPageNumber,
		}]

		_, hasOriginURL := existsNewLabels[labelInBookKey{
			Name:       entities.LabelNameRebuildOriginURL,
			PageNumber: newPageNumber,
		}]

		if hasOriginID || hasOriginName || hasOriginURL { // Данные уже проставлены в любом виде
			continue
		}

		newLabels = append(newLabels, entities.BookLabel{
			BookID:     bookToMerge.ID,
			PageNumber: newPageNumber,
			Name:       entities.LabelNameRebuildOriginID,
			Value:      sourceBook.ID.String(),
			CreateAt:   time.Now().UTC(),
		})

		if sourceBook.Name != "" {
			newLabels = append(newLabels, entities.BookLabel{
				BookID:     bookToMerge.ID,
				PageNumber: newPageNumber,
				Name:       entities.LabelNameRebuildOriginName,
				Value:      sourceBook.Name,
				CreateAt:   time.Now().UTC(),
			})
		}

		if sourceBook.OriginURL != nil {
			newLabels = append(newLabels, entities.BookLabel{
				BookID:     bookToMerge.ID,
				PageNumber: newPageNumber,
				Name:       entities.LabelNameRebuildOriginURL,
				Value:      sourceBook.OriginURL.String(),
				CreateAt:   time.Now().UTC(),
			})
		}

		existsNewLabels[labelInBookKey{
			Name:       entities.LabelNameRebuildOriginID,
			PageNumber: newPageNumber,
		}] = struct{}{}

		existsNewLabels[labelInBookKey{
			Name:       entities.LabelNameRebuildOriginName,
			PageNumber: newPageNumber,
		}] = struct{}{}

		existsNewLabels[labelInBookKey{
			Name:       entities.LabelNameRebuildOriginURL,
			PageNumber: newPageNumber,
		}] = struct{}{}
	}

	return newLabels, nil
}
