//revive:disable:file-length-limit
package external

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

const (
	HG4IDLabel     = "hg4:id"
	HG4RatingLabel = "hg4:rating"
)

type HG4Title struct {
	ID      int       `json:"id"`
	Created time.Time `json:"created"`
	URL     string    `json:"url"`

	Pages []HG4Page    `json:"pages"`
	Data  HG4TitleInfo `json:"info"`
}

type HG4TitleInfo struct {
	Parsed     HG4TitleInfoParsed `json:"parsed,omitempty"`
	Name       string             `json:"name,omitempty"`
	Rate       int                `json:"rate,omitempty"` // 3.2.0+
	Tags       []string           `json:"tags,omitempty"`
	Authors    []string           `json:"authors,omitempty"`
	Characters []string           `json:"characters,omitempty"`
	Languages  []string           `json:"languages,omitempty"`
	Categories []string           `json:"categories,omitempty"`
	Parodies   []string           `json:"parodies,omitempty"`
	Groups     []string           `json:"groups,omitempty"`
}

type HG4TitleInfoParsed struct {
	Name       bool `json:"name,omitempty"`
	Page       bool `json:"page,omitempty"`
	Tags       bool `json:"tags,omitempty"`
	Authors    bool `json:"authors,omitempty"`
	Characters bool `json:"characters,omitempty"`
	Languages  bool `json:"languages,omitempty"`
	Categories bool `json:"categories,omitempty"`
	Parodies   bool `json:"parodies,omitempty"`
	Groups     bool `json:"groups,omitempty"`
}

type HG4Page struct {
	TitleID    int       `json:"title_id"`    // 3.3.0+
	PageNumber int       `json:"page_number"` // 3.3.0+
	URL        string    `json:"url"`
	URLtoView  string    `json:"url_to_view"` // 3.3.0 - 4.0.0
	Ext        string    `json:"ext"`
	Success    bool      `json:"success"`
	LoadedAt   time.Time `json:"loaded_at"`
	Rate       int       `json:"rate,omitempty"` // 3.2.0+
}

// info.txt
func HG4ParseInfoTXT(body io.Reader) (Info, bool, error) {
	rawData, err := io.ReadAll(body)
	if err != nil {
		return Info{}, false, fmt.Errorf("read body: %w", err)
	}

	info := Info{
		Version: CurrentVersion,
		Meta: Meta{
			Exported:       time.Now().UTC(),
			ServiceName:    "hgraber",
			ServiceVersion: "v0.0.0",
		},
	}

	found := false

	const v1_0_0 = "v1.0.0"

	for _, row := range strings.Split(string(rawData), "\n") {
		switch {
		case strings.HasPrefix(row, "URL:"):
			info.Data.OriginURL = cleanupPrefix(row, "URL:")
			found = true // Во всех форматах есть адрес источника

		case strings.HasPrefix(row, "name:"):
			info.Data.Name = cleanupPrefix(row, "name:")

		case strings.HasPrefix(row, "pages:"):
			info.Data.PageCount, _ = strconv.Atoi(cleanupPrefix(row, "pages:"))

		case strings.HasPrefix(row, "NAME:"):
			info.Data.Name = cleanupPrefix(row, "NAME:")
			info.Meta.ServiceVersion = v1_0_0

		case strings.HasPrefix(row, "PAGE-COUNT:"):
			info.Data.PageCount, _ = strconv.Atoi(cleanupPrefix(row, "PAGE-COUNT:"))
			info.Meta.ServiceVersion = v1_0_0

		case strings.HasPrefix(row, "INNER-ID:"):
			info.Data.Labels = append(info.Data.Labels, Label{
				Name:     HG4IDLabel,
				Value:    cleanupPrefix(row, "INNER-ID:"),
				CreateAt: time.Now().UTC(),
			})
			info.Meta.ServiceVersion = v1_0_0
		}
	}

	return info, found, nil
}

func cleanupPrefix(s, prefix string) string {
	return strings.TrimSpace(strings.TrimPrefix(s, prefix))
}

// data.json
//
//nolint:cyclop,funlen // будет исправлено позднее
func HG4ParseDataJSON(body io.Reader) (Info, bool, error) {
	info := Info{
		Version: CurrentVersion,
		Meta: Meta{
			Exported:       time.Now().UTC(),
			ServiceName:    "hgraber",
			ServiceVersion: "v3.0.0",
		},
	}

	found3_2 := false
	found3_3 := false

	rawInfo := HG4Title{}

	err := json.NewDecoder(body).Decode(&rawInfo)
	if err != nil {
		return Info{}, false, fmt.Errorf("decode body: %w", err)
	}

	// Считаем что без ID данных не может быть
	// FIXME: проверить что не было версии с другими параметрами но без ID
	if rawInfo.ID == 0 {
		return Info{}, false, nil
	}

	// Пока считаем что если хоть 1 был, то значит есть все
	attrParsed := rawInfo.Data.Parsed.Tags ||
		rawInfo.Data.Parsed.Authors ||
		rawInfo.Data.Parsed.Characters ||
		rawInfo.Data.Parsed.Languages ||
		rawInfo.Data.Parsed.Categories ||
		rawInfo.Data.Parsed.Parodies ||
		rawInfo.Data.Parsed.Groups ||
		len(rawInfo.Data.Tags) > 0 ||
		len(rawInfo.Data.Authors) > 0 ||
		len(rawInfo.Data.Characters) > 0 ||
		len(rawInfo.Data.Languages) > 0 ||
		len(rawInfo.Data.Categories) > 0 ||
		len(rawInfo.Data.Parodies) > 0 ||
		len(rawInfo.Data.Groups) > 0

	info.Data = Book{
		Name:             rawInfo.Data.Name,
		OriginURL:        rawInfo.URL,
		PageCount:        len(rawInfo.Pages),
		CreateAt:         rawInfo.Created,
		AttributesParsed: attrParsed,
		Labels: []Label{
			{
				Name:     HG4IDLabel,
				Value:    strconv.Itoa(rawInfo.ID),
				CreateAt: time.Now().UTC(),
			},
		},
	}

	if len(rawInfo.Data.Tags) > 0 {
		info.Data.Attributes = append(info.Data.Attributes, Attribute{
			Code:   AttributeCodeTag,
			Values: rawInfo.Data.Tags,
		})
	}

	if len(rawInfo.Data.Authors) > 0 {
		info.Data.Attributes = append(info.Data.Attributes, Attribute{
			Code:   AttributeCodeAuthor,
			Values: rawInfo.Data.Authors,
		})
	}

	if len(rawInfo.Data.Characters) > 0 {
		info.Data.Attributes = append(info.Data.Attributes, Attribute{
			Code:   AttributeCodeCharacter,
			Values: rawInfo.Data.Characters,
		})
	}

	if len(rawInfo.Data.Languages) > 0 {
		info.Data.Attributes = append(info.Data.Attributes, Attribute{
			Code:   AttributeCodeLanguage,
			Values: rawInfo.Data.Languages,
		})
	}

	if len(rawInfo.Data.Categories) > 0 {
		info.Data.Attributes = append(info.Data.Attributes, Attribute{
			Code:   AttributeCodeCategory,
			Values: rawInfo.Data.Categories,
		})
	}

	if len(rawInfo.Data.Parodies) > 0 {
		info.Data.Attributes = append(info.Data.Attributes, Attribute{
			Code:   AttributeCodeParody,
			Values: rawInfo.Data.Parodies,
		})
	}

	if len(rawInfo.Data.Groups) > 0 {
		info.Data.Attributes = append(info.Data.Attributes, Attribute{
			Code:   AttributeCodeGroup,
			Values: rawInfo.Data.Groups,
		})
	}

	if rawInfo.Data.Rate > 0 {
		found3_2 = true

		info.Data.Labels = append(info.Data.Labels, Label{
			Name:     HG4RatingLabel,
			Value:    strconv.Itoa(rawInfo.Data.Rate),
			CreateAt: time.Now().UTC(),
		})
	}

	for i, rawPage := range rawInfo.Pages {
		page := Page{
			PageNumber: i + 1,
			Ext:        rawPage.Ext,
			OriginURL:  rawPage.URL,
			CreateAt:   rawInfo.Created,
			Downloaded: rawPage.Success,
			LoadAt:     rawPage.LoadedAt,
		}

		if !strings.HasPrefix(page.Ext, ".") {
			page.Ext = "." + page.Ext
		}

		if rawPage.PageNumber > 0 {
			found3_3 = true
			page.PageNumber = rawPage.PageNumber
		}

		if rawPage.Rate > 0 {
			found3_2 = true

			page.Labels = append(page.Labels, Label{
				Name:     HG4RatingLabel,
				Value:    strconv.Itoa(rawPage.Rate),
				CreateAt: time.Now().UTC(),
			})
		}

		info.Data.Pages = append(info.Data.Pages, page)
	}

	switch {
	case found3_3:
		info.Meta.ServiceVersion = "v3.3.0"
	case found3_2:
		info.Meta.ServiceVersion = "v3.2.0"
	}

	return info, false, nil
}
