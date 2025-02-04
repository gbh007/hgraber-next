package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/bff"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) APIBookListPost(ctx context.Context, req *serverAPI.BookFilter) (serverAPI.APIBookListPostRes, error) {
	filter := convertAPIBookFilter(*req)

	bookList, err := c.bffUseCases.BookList(ctx, filter)
	if err != nil {
		return &serverAPI.APIBookListPostInternalServerError{
			InnerCode: BFFUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIBookListPostOK{
		Books: pkg.Map(bookList.Books, func(b bff.BookShort) serverAPI.APIBookListPostOKBooksItem {
			return serverAPI.APIBookListPostOKBooksItem{
				Info: c.convertSimpleBook(b.Book, b.PreviewPage),
				Tags: b.Tags,
			}
		}),
		Pages: pkg.Map(bookList.Pages, func(v int) serverAPI.APIBookListPostOKPagesItem {
			return serverAPI.APIBookListPostOKPagesItem{
				Value:       v,
				IsCurrent:   v == req.Pagination.Value.Page.Value,
				IsSeparator: v == -1,
			}
		}),
		Count: bookList.Count,
	}, nil
}

func convertAPIBookFilter(req serverAPI.BookFilter) core.BookFilter {
	filter := core.BookFilter{}

	if req.Pagination.Value.Count.IsSet() {
		filter.FillLimits(req.Pagination.Value.Page.Value, req.Pagination.Value.Count.Value)
	}

	filter.Desc = req.Sort.Value.Desc.Value

	if req.Sort.Value.Field.IsSet() {
		switch req.Sort.Value.Field.Value {
		case serverAPI.BookFilterSortFieldCreatedAt:
			filter.OrderBy = core.BookFilterOrderByCreated
		case serverAPI.BookFilterSortFieldID:
			filter.OrderBy = core.BookFilterOrderByID
		case serverAPI.BookFilterSortFieldName:
			filter.OrderBy = core.BookFilterOrderByName
		case serverAPI.BookFilterSortFieldPageCount:
			filter.OrderBy = core.BookFilterOrderByPageCount
		default:
			filter.OrderBy = core.BookFilterOrderByCreated
		}
	}

	if req.Filter.Value.Flags.Value.DeleteStatus.IsSet() {
		filter.ShowDeleted = convertFlagSelector(req.Filter.Value.Flags.Value.DeleteStatus.Value)
	}

	if req.Filter.Value.Flags.Value.VerifyStatus.IsSet() {
		filter.ShowVerified = convertFlagSelector(req.Filter.Value.Flags.Value.VerifyStatus.Value)
	}

	if req.Filter.Value.Flags.Value.DownloadStatus.IsSet() {
		filter.ShowDownloaded = convertFlagSelector(req.Filter.Value.Flags.Value.DownloadStatus.Value)
	}

	if req.Filter.Value.Flags.Value.ShowRebuilded.IsSet() {
		filter.ShowRebuilded = convertFlagSelector(req.Filter.Value.Flags.Value.ShowRebuilded.Value)
	}

	if req.Filter.Value.Flags.Value.ShowWithoutPages.IsSet() {
		filter.ShowWithoutPages = convertFlagSelector(req.Filter.Value.Flags.Value.ShowWithoutPages.Value)
	}

	if req.Filter.Value.Flags.Value.ShowWithoutPreview.IsSet() {
		filter.ShowWithoutPreview = convertFlagSelector(req.Filter.Value.Flags.Value.ShowWithoutPreview.Value)
	}

	filter.Fields.Name = req.Filter.Value.Name.Value
	filter.From = req.Filter.Value.CreatedAt.Value.From.Value
	filter.To = req.Filter.Value.CreatedAt.Value.To.Value

	filter.Fields.Attributes = make([]core.BookFilterAttribute, 0, len(req.Filter.Value.Attributes))

	for _, attr := range req.Filter.Value.Attributes {
		attrFilter := core.BookFilterAttribute{
			Code:   attr.Code,
			Values: attr.Values,
			Count:  attr.Count.Value,
		}

		// TODO: добавить больше скипов из-за некорректных сочетаний значений.

		switch attr.Type {
		case serverAPI.BookFilterFilterAttributesItemTypeLike:
			attrFilter.Type = core.BookFilterAttributeTypeLike

		case serverAPI.BookFilterFilterAttributesItemTypeIn:
			attrFilter.Type = core.BookFilterAttributeTypeIn

		case serverAPI.BookFilterFilterAttributesItemTypeCountEq:
			attrFilter.Type = core.BookFilterAttributeTypeCountEq

		case serverAPI.BookFilterFilterAttributesItemTypeCountGt:
			attrFilter.Type = core.BookFilterAttributeTypeCountGt

		case serverAPI.BookFilterFilterAttributesItemTypeCountLt:
			attrFilter.Type = core.BookFilterAttributeTypeCountLt

		default:
			continue // Скипаем неизвестный тип если появится.
		}

		filter.Fields.Attributes = append(filter.Fields.Attributes, attrFilter)
	}

	filter.Fields.Labels = make([]core.BookFilterLabel, 0, len(req.Filter.Value.Labels))

	for _, label := range req.Filter.Value.Labels {
		labelFilter := core.BookFilterLabel{
			Name:   label.Name,
			Values: label.Values,
			Count:  label.Count.Value,
		}

		// TODO: добавить больше скипов из-за некорректных сочетаний значений.

		switch label.Type {
		case serverAPI.BookFilterFilterLabelsItemTypeLike:
			labelFilter.Type = core.BookFilterLabelTypeLike

		case serverAPI.BookFilterFilterLabelsItemTypeIn:
			labelFilter.Type = core.BookFilterLabelTypeIn

		case serverAPI.BookFilterFilterLabelsItemTypeCountEq:
			labelFilter.Type = core.BookFilterLabelTypeCountEq

		case serverAPI.BookFilterFilterLabelsItemTypeCountGt:
			labelFilter.Type = core.BookFilterLabelTypeCountGt

		case serverAPI.BookFilterFilterLabelsItemTypeCountLt:
			labelFilter.Type = core.BookFilterLabelTypeCountLt

		default:
			continue // Скипаем неизвестный тип если появится.
		}

		filter.Fields.Labels = append(filter.Fields.Labels, labelFilter)
	}

	return filter
}

func convertFlagSelector(raw serverAPI.BookFilterFlagSelector) core.BookFilterShowType {
	switch raw {
	case serverAPI.BookFilterFlagSelectorAll:
		return core.BookFilterShowTypeAll
	case serverAPI.BookFilterFlagSelectorOnly:
		return core.BookFilterShowTypeOnly
	case serverAPI.BookFilterFlagSelectorExcept:
		return core.BookFilterShowTypeExcept
	}

	return core.BookFilterShowTypeAll
}
