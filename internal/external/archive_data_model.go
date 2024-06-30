package external

import (
	"time"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

type Info struct {
	Version string `json:"version"`
	Meta    Meta   `json:"meta"`
	Data    Book   `json:"data"`
}

type Meta struct {
	Exported       time.Time `json:"exported"`
	ServiceVersion string    `json:"service_version,omitempty"`
	ServiceName    string    `json:"service_name,omitempty"`
}

type Book struct {
	Name             string      `json:"name"`
	OriginURL        string      `json:"origin_url,omitempty"`
	PageCount        int         `json:"page_count"`
	CreateAt         time.Time   `json:"create_at"`
	AttributesParsed bool        `json:"attributes_parsed"`
	Attributes       []Attribute `json:"attributes,omitempty"`
	Pages            []Page      `json:"pages,omitempty"`
	Labels           []Label     `json:"labels,omitempty"`
}

type Page struct {
	PageNumber int       `json:"page_number"`
	Ext        string    `json:"ext"`
	OriginURL  string    `json:"origin_url,omitempty"`
	CreateAt   time.Time `json:"create_at"`
	Downloaded bool      `json:"downloaded,omitempty"`
	LoadAt     time.Time `json:"load_at,omitempty"`
	Labels     []Label   `json:"labels,omitempty"`
}

type Label struct {
	Name     string    `json:"name"`
	Value    string    `json:"value"`
	CreateAt time.Time `json:"create_at"`
}

type Attribute struct {
	Code   string   `json:"code"`
	Values []string `json:"values"`
}

func BookFromEntity(raw entities.BookFull) Book {
	labels := make(map[int][]Label, raw.PageCount+1)

	for _, l := range raw.Labels {
		labels[l.PageNumber] = append(labels[l.PageNumber], Label{
			Name:     l.Name,
			Value:    l.Value,
			CreateAt: l.CreateAt,
		})
	}

	b := Book{
		Name:             raw.Name,
		PageCount:        raw.PageCount,
		CreateAt:         raw.CreateAt,
		AttributesParsed: raw.AttributesParsed,
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

	if raw.OriginURL != nil {
		b.OriginURL = raw.OriginURL.String()
	}

	b.Labels = append(b.Labels, Label{
		Name:     "hg5-id",
		Value:    raw.ID.String(),
		CreateAt: raw.CreateAt,
	})

	return b
}

func Convert(raw entities.BookFull) Info {
	return Info{
		Version: "1.0.0",
		Meta: Meta{
			Exported:    time.Now().UTC(),
			ServiceName: "hgraber next",
		},
		Data: BookFromEntity(raw),
	}
}
