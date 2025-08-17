package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *MassloadController) APIMassloadInfoListPost(ctx context.Context, req *serverapi.APIMassloadInfoListPostReq) (serverapi.APIMassloadInfoListPostRes, error) {
	filter := massloadmodel.Filter{}

	filter.Desc = req.Sort.Value.Desc.Value

	if req.Sort.Value.Field.Set {
		switch req.Sort.Value.Field.Value {
		case serverapi.APIMassloadInfoListPostReqSortFieldID:
			filter.OrderBy = massloadmodel.FilterOrderByID
		case serverapi.APIMassloadInfoListPostReqSortFieldName:
			filter.OrderBy = massloadmodel.FilterOrderByName
		case serverapi.APIMassloadInfoListPostReqSortFieldPageSize:
			filter.OrderBy = massloadmodel.FilterOrderByPageSize
		case serverapi.APIMassloadInfoListPostReqSortFieldFileSize:
			filter.OrderBy = massloadmodel.FilterOrderByFileSize
		default:
			filter.OrderBy = massloadmodel.FilterOrderByID
		}
	}

	filter.Fields.Name = req.Filter.Value.Name.Value
	filter.Fields.ExternalLink = req.Filter.Value.ExternalLink.Value
	filter.Fields.Flags = req.Filter.Value.Flags

	filter.Fields.Attributes = make([]massloadmodel.FilterAttribute, 0, len(req.Filter.Value.Attributes))

	for _, attr := range req.Filter.Value.Attributes {
		attrFilter := massloadmodel.FilterAttribute{
			Code:   attr.Code,
			Values: attr.Values,
		}

		switch attr.Type {
		case serverapi.APIMassloadInfoListPostReqFilterAttributesItemTypeLike:
			attrFilter.Type = massloadmodel.FilterAttributeTypeLike

		case serverapi.APIMassloadInfoListPostReqFilterAttributesItemTypeIn:
			attrFilter.Type = massloadmodel.FilterAttributeTypeIn

		default:
			continue // Скипаем неизвестный тип если появится.
		}

		filter.Fields.Attributes = append(filter.Fields.Attributes, attrFilter)
	}

	mls, err := c.massloadUseCases.Massloads(ctx, filter)
	if err != nil {
		return &serverapi.APIMassloadInfoListPostInternalServerError{
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIMassloadInfoListPostOK{
		Massloads: pkg.Map(mls, convertMassloadInfo),
	}, nil
}
