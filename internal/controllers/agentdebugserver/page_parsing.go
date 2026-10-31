package agentdebugserver

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/internal/domian/archivemodel"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/labstack/echo/v4"
)

type parseRequest struct {
	Url     string `form:"u" query:"u"`
	AsZip   bool   `form:"as-zip" query:"as-zip"`
	AsMulti bool   `form:"as-multi" query:"as-multi"`
}

func (cnt *Controller) pageParsing(c echo.Context) error {
	if c.Request().Method == http.MethodGet {
		return c.Render(http.StatusOK, "parsing.html.gotmpl", nil)
	}

	req := parseRequest{}

	err := c.Bind(&req)
	if err != nil {
		return c.Render(http.StatusOK, "parsing.html.gotmpl", map[string]any{
			"error": err,
		})
	}

	u, err := url.Parse(req.Url)
	if err != nil {
		return c.Render(http.StatusOK, "parsing.html.gotmpl", map[string]any{
			"error": err,
		})
	}

	if req.AsMulti {
		urls, err := cnt.parsing.MultiHandle(c.Request().Context(), *u)
		if err != nil {
			return c.Render(http.StatusOK, "parsing.html.gotmpl", map[string]any{
				"error": err,
			})
		}

		return c.Render(http.StatusOK, "parsing.html.gotmpl", map[string]any{
			"urls": pkg.Map(urls, func(raw serveragentdto.AgentBookCheckResult) string {
				return raw.URL.String()
			}),
		})
	}

	details, err := cnt.parsing.ParseBook(c.Request().Context(), *u)
	if err != nil {
		return c.Render(http.StatusOK, "parsing.html.gotmpl", map[string]any{
			"error": err,
		})
	}

	info := archivemodel.Info{
		Version: "1.0.0",
		Meta: archivemodel.Meta{
			Exported:    time.Now().UTC(),
			ServiceName: "hgraber next agent",
		},
		Data: archivemodel.Book{
			Name:             details.Name,
			OriginURL:        details.URL.String(),
			PageCount:        details.PageCount,
			CreateAt:         time.Now(),
			AttributesParsed: true,
			Attributes: pkg.Map(details.Attributes, func(raw serveragentdto.AgentBookDetailsAttributesItem) archivemodel.Attribute {
				return archivemodel.Attribute{
					Code:   raw.Code,
					Values: raw.Values,
				}
			}),
			Pages: pkg.Map(details.Pages, func(p serveragentdto.AgentBookDetailsPagesItem) archivemodel.Page {
				return archivemodel.Page{
					PageNumber: p.PageNumber,
					Ext:        path.Ext(p.Filename),
					OriginURL:  p.URL.String(),
					CreateAt:   time.Now(),
					Downloaded: true,
					LoadAt:     time.Now(),
				}
			}),
		},
	}

	if !req.AsZip {
		data, err := json.MarshalIndent(info, "", "\t")
		if err != nil {
			return c.Render(http.StatusOK, "parsing.html.gotmpl", map[string]any{
				"error": err,
			})
		}

		return c.Render(http.StatusOK, "parsing.html.gotmpl", map[string]any{
			"info": string(data),
		})
	}

	buff := &bytes.Buffer{}

	pageUrls := pkg.SliceToMap(details.Pages, func(p serveragentdto.AgentBookDetailsPagesItem) (int, url.URL) {
		return p.PageNumber, p.URL
	})

	zipWriter := zip.NewWriter(buff)

	err = archivemodel.WriteArchive(
		c.Request().Context(),
		zipWriter,
		func(ctx context.Context, pageNumber int) (io.Reader, error) {
			u, ok := pageUrls[pageNumber]
			if !ok {
				return nil, fmt.Errorf("missing page %d", pageNumber)
			}

			return cnt.parsing.DownloadPage(ctx, details.URL, u)
		},
		info,
	)
	if err != nil {
		return c.Render(http.StatusOK, "parsing.html.gotmpl", map[string]any{
			"error": err,
		})
	}

	err = zipWriter.Close()
	if err != nil {
		return c.Render(http.StatusOK, "parsing.html.gotmpl", map[string]any{
			"error": err,
		})
	}

	return c.Stream(http.StatusOK, "application/zip", buff)
}
