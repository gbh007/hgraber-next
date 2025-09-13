package generatorcore

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/units"
	promcog "github.com/grafana/promql-builder/go/cog"
	"github.com/grafana/promql-builder/go/promql"
)

const (
	MetricVariableName = "metrics"
	MetricVariableType = "prometheus"

	LogsVariableName = "logs"
	LogsVariableType = "loki"

	ServiceVariableName = "service_name"

	DeltaVariableName    = "deltaInterval"
	DeltaVariableCurrent = "4h"

	RateIntervalVar = "$__rate_interval" // TODO: возможно где-то есть готовый

	LegendAuto = "__auto"
	LELabel    = "le"

	Q95 = 0.95

	UnitShort      = units.Short
	UnitBytes      = units.BytesIEC
	UnitPercent0_1 = units.PercentUnit
	UnitSecond     = units.Seconds
	UnitRPS        = units.RequestsPerSecond
)

const (
	PanelSizeHalf = iota
	PanelSizeQuarter
	PanelSizeQuarterSlim
	PanelSizeThird
	PanelSizeFull
)

var (
	MetricDatasource = dashboard.DataSourceRef{
		Type: StrToPtr(MetricVariableType),
		Uid:  StrToPtr(NameToVarDS(MetricVariableName)),
	}
	LogsDatasource = dashboard.DataSourceRef{
		Type: StrToPtr(LogsVariableType),
		Uid:  StrToPtr(NameToVarDS(LogsVariableName)),
	}
	DeltaVariableValues = []string{"1m", "5m", "10m", "30m", "1h", "4h", "8h", "1d", "7d"}
	GreenSteps          = []dashboard.Threshold{
		{
			Color: "green",
		},
	}
	ServiceFilterPromQL = []promcog.Builder[promql.LabelSelector]{
		promql.
			NewLabelSelectorBuilder().
			Name("service_name"). // TODO: в константу?
			Operator(promql.LabelMatchingOperatorMatchRegexp).
			Value(NameToVar(ServiceVariableName)),
	}
)

type PanelSize byte

func WithPanelSize[T interface {
	Height(h uint32) T
	Span(w uint32) T
}](data T, size PanelSize) T {
	var h, w uint32

	switch size {
	case PanelSizeQuarterSlim:
		h = 3
		w = 6
	case PanelSizeQuarter:
		h = 6
		w = 6
	case PanelSizeThird:
		h = 9
		w = 8
	case PanelSizeHalf:
		h = 9
		w = 12
	case PanelSizeFull:
		h = 12
		w = 24
	}

	return data.Height(h).Span(w)
}
