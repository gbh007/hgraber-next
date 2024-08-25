package external

import (
	"time"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func BookFromEntity(raw entities.BookFull) Book {
	labels := make(map[int][]Label, raw.Book.PageCount+1)

	for _, l := range raw.Labels {
		labels[l.PageNumber] = append(labels[l.PageNumber], Label{
			Name:     l.Name,
			Value:    l.Value,
			CreateAt: l.CreateAt,
		})
	}

	b := Book{
		Name:             raw.Book.Name,
		PageCount:        raw.Book.PageCount,
		CreateAt:         raw.Book.CreateAt,
		AttributesParsed: raw.Book.AttributesParsed,
		Labels:           labels[0],
		Attributes: pkg.MapToSlice(raw.Attributes, func(code string, values []string) Attribute {
			return Attribute{
				Code:   code,
				Values: values,
			}
		}),
		Pages: pkg.Map(raw.Pages, func(p entities.Page) Page {
			u := ""
			if p.OriginURL != nil {
				u = p.OriginURL.String()
			}

			return Page{
				PageNumber: p.PageNumber,
				Ext:        p.Ext,
				OriginURL:  u,
				CreateAt:   p.CreateAt,
				Downloaded: p.Downloaded,
				LoadAt:     p.LoadAt,
				Labels:     labels[p.PageNumber],
			}
		}),
	}

	if raw.Book.OriginURL != nil {
		b.OriginURL = raw.Book.OriginURL.String()
	}

	b.Labels = append(b.Labels, Label{
		Name:     "hg5:id",
		Value:    raw.Book.ID.String(),
		CreateAt: raw.Book.CreateAt,
	})

	return b
}

func Convert(raw entities.BookFull) Info {
	return Info{
		Version: CurrentVersion,
		Meta: Meta{
			Exported:       time.Now().UTC(),
			ServiceName:    "hgraber next",
			ServiceVersion: "v0.0.0",
		},
		Data: BookFromEntity(raw),
	}
}
