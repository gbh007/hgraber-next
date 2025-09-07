package generatorcore

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

const (
	MetricVariableName = "metrics"
	MetricVariableType = "prometheus"

	LogsVariableName = "logs"
	LogsVariableType = "loki"

	ServiceVariableName = "service_name"

	DeltaVariableName    = "deltaInterval"
	DeltaVariableCurrent = "4h"

	ServiceFilter = `service_name=~"$` + ServiceVariableName + `"`

	UnitShort = units.Short
)

const (
	PanelSizeCommon = iota
	PanelSizeHalf
	PanelSizeSlim
	PanelSizeFull
)

var (
	MetricDatasource = dashboard.DataSourceRef{
		Type: StrToPtr(MetricVariableType),
		Uid:  StrToPtr(NameToVar(MetricVariableName)),
	}
	DeltaVariableValues = []string{"1m", "5m", "10m", "30m", "1h", "4h", "8h", "1d", "7d"}
	GreenSteps          = []dashboard.Threshold{
		{
			Color: "green",
		},
	}
)

type PanelSize byte

func WithPanelSize[T interface {
	Height(h uint32) T
	Span(w uint32) T
}](data T, size PanelSize) T {
	var h, w uint32

	switch size {
	case PanelSizeCommon:
		h = 9
		w = 12
	case PanelSizeHalf:
		h = 6
		w = 6
	case PanelSizeSlim:
		h = 3
		w = 6
	case PanelSizeFull:
		h = 12
		w = 24
	}

	return data.Height(h).Span(w)
}
