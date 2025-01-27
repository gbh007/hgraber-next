package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIBookListPost(ctx context.Context, req *serverAPI.BookFilter) (serverAPI.APIBookListPostRes, error) {
	filter := convertAPIBookFilter(*req)

	bookList, err := c.webAPIUseCases.BookList(ctx, filter)
	if err != nil {
		return &serverAPI.APIBookListPostInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIBookListPostOK{
		Books: pkg.Map(bookList.Books, func(b entities.BookToWeb) serverAPI.APIBookListPostOKBooksItem {
			return serverAPI.APIBookListPostOKBooksItem{
				Info:              c.convertSimpleBook(b.Book, b.PreviewPage),
				PageLoadedPercent: b.PageDownloadPercent(),
				Tags:              b.Tags,
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

func convertAPIBookFilter(req serverAPI.BookFilter) entities.BookFilter {
	filter := entities.BookFilter{}

	if req.Pagination.Value.Count.IsSet() {
		filter.FillLimits(req.Pagination.Value.Page.Value, req.Pagination.Value.Count.Value)
	}

	filter.Desc = req.Sort.Value.Desc.Value

	if req.Sort.Value.Field.IsSet() {
		switch req.Sort.Value.Field.Value {
		case serverAPI.BookFilterSortFieldCreatedAt:
			filter.OrderBy = entities.BookFilterOrderByCreated
		case serverAPI.BookFilterSortFieldID:
			filter.OrderBy = entities.BookFilterOrderByID
		case serverAPI.BookFilterSortFieldName:
			filter.OrderBy = entities.BookFilterOrderByName
		case serverAPI.BookFilterSortFieldPageCount:
			filter.OrderBy = entities.BookFilterOrderByPageCount
		default:
			filter.OrderBy = entities.BookFilterOrderByCreated
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

	filter.Fields.Attributes = make([]entities.BookFilterAttribute, 0, len(req.Filter.Value.Attributes))

	for _, attr := range req.Filter.Value.Attributes {
		attrFilter := entities.BookFilterAttribute{
			Code:   attr.Code,
			Values: attr.Values,
			Count:  attr.Count.Value,
		}

		// TODO: добавить больше скипов из-за некорректных сочетаний значений.

		switch attr.Type {
		case serverAPI.BookFilterFilterAttributesItemTypeLike:
			attrFilter.Type = entities.BookFilterAttributeTypeLike

		case serverAPI.BookFilterFilterAttributesItemTypeIn:
			attrFilter.Type = entities.BookFilterAttributeTypeIn

		case serverAPI.BookFilterFilterAttributesItemTypeCountEq:
			attrFilter.Type = entities.BookFilterAttributeTypeCountEq

		case serverAPI.BookFilterFilterAttributesItemTypeCountGt:
			attrFilter.Type = entities.BookFilterAttributeTypeCountGt

		case serverAPI.BookFilterFilterAttributesItemTypeCountLt:
			attrFilter.Type = entities.BookFilterAttributeTypeCountLt

		default:
			continue // Скипаем неизвестный тип если появится.
		}

		filter.Fields.Attributes = append(filter.Fields.Attributes, attrFilter)
	}

	filter.Fields.Labels = make([]entities.BookFilterLabel, 0, len(req.Filter.Value.Labels))

	for _, label := range req.Filter.Value.Labels {
		labelFilter := entities.BookFilterLabel{
			Name:   label.Name,
			Values: label.Values,
			Count:  label.Count.Value,
		}

		// TODO: добавить больше скипов из-за некорректных сочетаний значений.

		switch label.Type {
		case serverAPI.BookFilterFilterLabelsItemTypeLike:
			labelFilter.Type = entities.BookFilterLabelTypeLike

		case serverAPI.BookFilterFilterLabelsItemTypeIn:
			labelFilter.Type = entities.BookFilterLabelTypeIn

		case serverAPI.BookFilterFilterLabelsItemTypeCountEq:
			labelFilter.Type = entities.BookFilterLabelTypeCountEq

		case serverAPI.BookFilterFilterLabelsItemTypeCountGt:
			labelFilter.Type = entities.BookFilterLabelTypeCountGt

		case serverAPI.BookFilterFilterLabelsItemTypeCountLt:
			labelFilter.Type = entities.BookFilterLabelTypeCountLt

		default:
			continue // Скипаем неизвестный тип если появится.
		}

		filter.Fields.Labels = append(filter.Fields.Labels, labelFilter)
	}

	return filter
}

func convertFlagSelector(raw serverAPI.BookFilterFlagSelector) entities.BookFilterShowType {
	switch raw {
	case serverAPI.BookFilterFlagSelectorAll:
		return entities.BookFilterShowTypeAll
	case serverAPI.BookFilterFlagSelectorOnly:
		return entities.BookFilterShowTypeOnly
	case serverAPI.BookFilterFlagSelectorExcept:
		return entities.BookFilterShowTypeExcept
	}

	return entities.BookFilterShowTypeAll
}
