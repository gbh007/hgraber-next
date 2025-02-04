package external

import (
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/pkg"
)

func BookToEntity(raw Book) (core.BookContainer, error) {
	var err error

	book := core.Book{
		Name:             raw.Name,
		PageCount:        raw.PageCount,
		AttributesParsed: raw.AttributesParsed,
		CreateAt:         raw.CreateAt,
	}

	if raw.OriginURL != "" {
		u, err := url.Parse(raw.OriginURL)
		if err != nil {
			return core.BookContainer{}, fmt.Errorf("parse book url: %w", err)
		}

		book.OriginURL = u
	}

	labels := make([]core.BookLabel, 0, raw.PageCount+1)

	for _, l := range raw.Labels {
		labels = append(labels, core.BookLabel{
			Name:     l.Name,
			Value:    l.Value,
			CreateAt: l.CreateAt,
		})
	}

	pages := make([]core.Page, len(raw.Pages))

	for i, p := range raw.Pages {
		var u *url.URL
		if p.OriginURL != "" {
			u, err = url.Parse(p.OriginURL)
			if err != nil {
				return core.BookContainer{}, fmt.Errorf("parse page (%d) url: %w", p.PageNumber, err)
			}
		}

		pages[i] = core.Page{
			PageNumber: p.PageNumber,
			Ext:        p.Ext,
			OriginURL:  u,
			CreateAt:   p.CreateAt,
			Downloaded: p.Downloaded,
			LoadAt:     p.LoadAt,
		}

		for _, l := range p.Labels {
			labels = append(labels, core.BookLabel{
				PageNumber: p.PageNumber,
				Name:       l.Name,
				Value:      l.Value,
				CreateAt:   l.CreateAt,
			})
		}
	}

	b := core.BookContainer{
		Book:   book,
		Labels: labels,
		Attributes: pkg.SliceToMap(raw.Attributes, func(attr Attribute) (string, []string) {
			return attr.Code, attr.Values
		}),
		Pages: pages,
	}

	return b, nil
}
