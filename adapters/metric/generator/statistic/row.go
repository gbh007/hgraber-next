package statistic

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/heatmap"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metricstatistic"
)

func WithRow(builder *dashboard.DashboardBuilder) *dashboard.DashboardBuilder {
	builder.WithRow(
		dashboard.NewRowBuilder("Statistic").
			WithPanel(generatorcore.WithPanelSize(BookSizes(), generatorcore.PanelSizeThird)).
			WithPanel(generatorcore.WithPanelSize(PageSizes(), generatorcore.PanelSizeThird)).
			WithPanel(generatorcore.WithPanelSize(PageInBook(), generatorcore.PanelSizeThird)).
			WithPanel(generatorcore.WithPanelSize(PageSizesByAuthors(), generatorcore.PanelSizeThird)).
			WithPanel(generatorcore.WithPanelSize(BookByAuthors(), generatorcore.PanelSizeThird)).
			WithPanel(generatorcore.WithPanelSize(PageByAuthors(), generatorcore.PanelSizeThird)).
			Collapsed(true),
	)

	return builder
}

func BookSizes() *heatmap.PanelBuilder {
	return Heatmap(
		"Book sizes",
		metricstatistic.BookSize,
		generatorcore.UnitBytes,
	)
}

func PageSizes() *heatmap.PanelBuilder {
	return Heatmap(
		"Page sizes",
		metricstatistic.PageSize,
		generatorcore.UnitBytes,
	)
}

func PageInBook() *heatmap.PanelBuilder {
	return Heatmap(
		"Page in book",
		metricstatistic.PageInBook,
		generatorcore.UnitShort,
	)
}

func PageSizesByAuthors() *heatmap.PanelBuilder {
	return Heatmap(
		"Page sizes by authors",
		metricstatistic.PagesSizeByAuthor,
		generatorcore.UnitBytes,
	)
}

func BookByAuthors() *heatmap.PanelBuilder {
	return Heatmap(
		"Book by authors",
		metricstatistic.BookCountByAuthor,
		generatorcore.UnitShort,
	)
}

func PageByAuthors() *heatmap.PanelBuilder {
	return Heatmap(
		"Page by authors",
		metricstatistic.PagesByAuthor,
		generatorcore.UnitShort,
	)
}
