package massloadusecase

import (
	"context"
	"fmt"
	"time"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/pkg"
)

func (uc *UseCase) MassloadForUpdate(ctx context.Context) ([]massloadmodel.Massload, error) {
	mls, err := uc.storage.Massloads(ctx, massloadmodel.Filter{})
	if err != nil {
		return nil, fmt.Errorf("storage get massloads: %w", err)
	}

	return mls, nil
}

func (uc *UseCase) UpdateSize(ctx context.Context, ml massloadmodel.Massload) error {
	attrs, err := uc.storage.MassloadAttributes(ctx, ml.ID)
	if err != nil {
		return fmt.Errorf("storage get attributes: %w", err)
	}

	attrMap := make(map[string][]string)

	for _, attr := range attrs {
		attrMap[attr.Code] = append(attrMap[attr.Code], attr.Value)
	}

	for code, values := range attrMap {
		attrMap[code] = pkg.Unique(values)
	}

	filesInfo, err := uc.storage.AttributesFileSize(ctx, attrMap)
	if err != nil {
		return fmt.Errorf("storage get file size: %w", err)
	}

	pagesInfo, err := uc.storage.AttributesPageSize(ctx, attrMap)
	if err != nil {
		return fmt.Errorf("storage get page size: %w", err)
	}

	bookCount, err := uc.storage.BookCount(ctx, core.BookFilter{
		Fields: core.BookFilterFields{
			Attributes: pkg.MapToSlice(attrMap, func(k string, v []string) core.BookFilterAttribute {
				return core.BookFilterAttribute{
					Code:   k,
					Type:   core.BookFilterAttributeTypeIn,
					Values: v,
				}
			}),
		},
	})
	if err != nil {
		return fmt.Errorf("storage get book count: %w", err)
	}

	bookInSystem := int64(bookCount)

	err = uc.storage.UpdateMassloadSize(ctx, massloadmodel.Massload{
		ID:           ml.ID,
		PageSize:     &pagesInfo.Size,
		FileSize:     &filesInfo.Size,
		PageCount:    &pagesInfo.Count,
		FileCount:    &filesInfo.Count,
		BookInSystem: &bookInSystem,
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		return fmt.Errorf("storage update size: %w", err)
	}

	return nil
}

func (uc *UseCase) MassloadAttributesForUpdate(ctx context.Context) ([]massloadmodel.Attribute, error) {
	attrs, err := uc.storage.MassloadsAttributes(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage get massloads attributes: %w", err)
	}

	return attrs, nil
}

func (uc *UseCase) UpdateAttributesSize(ctx context.Context, attr massloadmodel.Attribute) error {
	attrMap := map[string][]string{
		attr.Code: {
			attr.Value,
		},
	}

	filesInfo, err := uc.storage.AttributesFileSize(ctx, attrMap)
	if err != nil {
		return fmt.Errorf("storage get file size: %w", err)
	}

	pagesInfo, err := uc.storage.AttributesPageSize(ctx, attrMap)
	if err != nil {
		return fmt.Errorf("storage get page size: %w", err)
	}

	bookCount, err := uc.storage.BookCount(ctx, core.BookFilter{
		Fields: core.BookFilterFields{
			Attributes: pkg.MapToSlice(attrMap, func(k string, v []string) core.BookFilterAttribute {
				return core.BookFilterAttribute{
					Code:   k,
					Type:   core.BookFilterAttributeTypeIn,
					Values: v,
				}
			}),
		},
	})
	if err != nil {
		return fmt.Errorf("storage get book count: %w", err)
	}

	bookInSystem := int64(bookCount)

	err = uc.storage.UpdateMassloadAttributeSize(ctx, massloadmodel.Attribute{
		Code:         attr.Code,
		Value:        attr.Value,
		PageSize:     &pagesInfo.Size,
		FileSize:     &filesInfo.Size,
		PageCount:    &pagesInfo.Count,
		FileCount:    &filesInfo.Count,
		BookInSystem: &bookInSystem,
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		return fmt.Errorf("storage update size: %w", err)
	}

	return nil
}
