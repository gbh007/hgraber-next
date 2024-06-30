package external

import (
	"fmt"
	"net/url"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func BookToEntity(raw Book) (entities.BookFull, error) {
	var err error

	book := entities.Book{
		Name:             raw.Name,
		PageCount:        raw.PageCount,
		AttributesParsed: raw.AttributesParsed,
		CreateAt:         raw.CreateAt,
	}

	if raw.OriginURL != "" {
		u, err := url.Parse(raw.OriginURL)
		if err != nil {
			return entities.BookFull{}, fmt.Errorf("parse book url: %w", err)
		}

		book.OriginURL = u
	}

	labels := make([]entities.BookLabel, 0, raw.PageCount+1)

	for _, l := range raw.Labels {
		labels = append(labels, entities.BookLabel{
			Name:     l.Name,
			Value:    l.Value,
			CreateAt: l.CreateAt,
		})
	}

	pages := make([]entities.Page, len(raw.Pages))

	for i, p := range raw.Pages {
		var u *url.URL
		if p.OriginURL != "" {
			u, err = url.Parse(p.OriginURL)
			if err != nil {
				return entities.BookFull{}, fmt.Errorf("parse page (%d) url: %w", p.PageNumber, err)
			}
		}

		pages[i] = entities.Page{
			PageNumber: p.PageNumber,
			Ext:        p.Ext,
			OriginURL:  u,
			CreateAt:   p.CreateAt,
			Downloaded: p.Downloaded,
			LoadAt:     p.LoadAt,
		}

		for _, l := range p.Labels {
			labels = append(labels, entities.BookLabel{
				PageNumber: p.PageNumber,
				Name:       l.Name,
				Value:      l.Value,
				CreateAt:   l.CreateAt,
			})
		}
	}

	b := entities.BookFull{
		Book:   book,
		Labels: labels,
		Attributes: pkg.SliceToMap(raw.Attributes, func(attr Attribute) (string, []string) {
			return attr.Code, attr.Values
		}),
		Pages: pages,
	}

	return b, nil
}
