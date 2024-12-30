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
		Books: pkg.Map(bookList.Books, func(b entities.BookToWeb) serverAPI.BookShortInfo {
			return serverAPI.BookShortInfo{
				ID:                b.Book.ID,
				Created:           b.Book.CreateAt,
				PreviewURL:        c.getPagePreview(b.PreviewPage),
				ParsedName:        b.ParsedName(),
				Name:              b.Book.Name,
				ParsedPage:        b.ParsedPages,
				PageCount:         b.Book.PageCount,
				PageLoadedPercent: b.PageDownloadPercent(),
				Tags:              b.Tags,
				HasMoreTags:       b.HasMoreTags,
			}
		}),
		Pages: pkg.Map(bookList.Pages, func(v int) serverAPI.APIBookListPostOKPagesItem {
			return serverAPI.APIBookListPostOKPagesItem{
				Value:       v,
				IsCurrent:   v == req.Page.Value,
				IsSeparator: v == -1,
			}
		}),
		Count: bookList.Count,
	}, nil
}

func convertAPIBookFilter(req serverAPI.BookFilter) entities.BookFilter {
	filter := entities.BookFilter{}

	if req.Count.IsSet() {
		filter.FillLimits(req.Page.Value, req.Count.Value)
	}

	filter.Desc = req.SortDesc.Value

	if req.SortField.IsSet() {
		switch req.SortField.Value {
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

	if req.DeleteStatus.IsSet() {
		switch req.DeleteStatus.Value {
		case serverAPI.BookFilterDeleteStatusAll:
		case serverAPI.BookFilterDeleteStatusOnly:
			filter.ShowDeleted = entities.BookFilterShowTypeOnly
		case serverAPI.BookFilterDeleteStatusExcept:
			filter.ShowDeleted = entities.BookFilterShowTypeExcept
		}
	}

	if req.VerifyStatus.IsSet() {
		switch req.VerifyStatus.Value {
		case serverAPI.BookFilterVerifyStatusAll:
		case serverAPI.BookFilterVerifyStatusOnly:
			filter.ShowVerified = entities.BookFilterShowTypeOnly
		case serverAPI.BookFilterVerifyStatusExcept:
			filter.ShowVerified = entities.BookFilterShowTypeExcept
		}
	}

	if req.DownloadStatus.IsSet() {
		switch req.DownloadStatus.Value {
		case serverAPI.BookFilterDownloadStatusAll:
		case serverAPI.BookFilterDownloadStatusOnly:
			filter.ShowDownloaded = entities.BookFilterShowTypeOnly
		case serverAPI.BookFilterDownloadStatusExcept:
			filter.ShowDownloaded = entities.BookFilterShowTypeExcept
		}
	}

	filter.Fields.Name = req.Filter.Value.Name.Value
	filter.From = req.From.Value
	filter.To = req.To.Value

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
