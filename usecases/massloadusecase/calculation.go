package massloadusecase

import (
	"context"
	"fmt"
	"time"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/pkg"
)

func (uc *UseCase) MassloadForUpdateCalculation(ctx context.Context) ([]massloadmodel.Massload, error) {
	mls, err := uc.storage.Massloads(ctx, massloadmodel.Filter{})
	if err != nil {
		return nil, fmt.Errorf("storage get massloads: %w", err)
	}

	return mls, nil
}

func (uc *UseCase) UpdateCalculation(ctx context.Context, ml massloadmodel.Massload) error {
	all := handledExternalURL{}

	links, err := uc.storage.MassloadExternalLinks(ctx, ml.ID)
	if err != nil {
		return fmt.Errorf("storage get external links: %w", err)
	}

	for _, link := range links {
		if !link.AutoCheck {
			continue
		}

		info, err := uc.calcExternalLink(ctx, link.URL)
		if err != nil {
			return fmt.Errorf("storage calc external link (%s): %w", link.URL.String(), err)
		}

		link.UpdatedAt = time.Now()
		link.BooksAhead = intToInt64Ptr(len(pkg.Unique(info.urlsAhead)))
		link.NewBooks = intToInt64Ptr(len(pkg.Unique(info.urlsNew)))
		link.ExistingBooks = intToInt64Ptr(len(pkg.Unique(info.urlsExisting)))

		err = uc.storage.UpdateMassloadExternalLinkCounts(ctx, ml.ID, link)
		if err != nil {
			return fmt.Errorf("storage update external link (%s): %w", link.URL.String(), err)
		}

		all.urlsAhead = append(all.urlsAhead, info.urlsAhead...)
		all.urlsNew = append(all.urlsNew, info.urlsNew...)
		all.urlsExisting = append(all.urlsExisting, info.urlsExisting...)
	}

	ml.UpdatedAt = time.Now()
	ml.BooksAhead = intToInt64Ptr(len(pkg.Unique(all.urlsAhead)))
	ml.NewBooks = intToInt64Ptr(len(pkg.Unique(all.urlsNew)))
	ml.ExistingBooks = intToInt64Ptr(len(pkg.Unique(all.urlsExisting)))

	err = uc.storage.UpdateMassloadCounts(ctx, ml)
	if err != nil {
		return fmt.Errorf("storage update massload: %w", err)
	}

	return nil
}

func intToInt64Ptr(i int) *int64 {
	v := int64(i)

	return &v
}
