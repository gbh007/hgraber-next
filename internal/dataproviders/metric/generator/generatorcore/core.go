package generatorcore

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/units"
	promcog "github.com/grafana/promql-builder/go/cog"
	"github.com/grafana/promql-builder/go/promql"

	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/metriccore"
)

const (
	MetricVariableName = "metrics"
	MetricVariableType = "prometheus"

	LogsVariableName             = "logs"
	LogsVariableTypeLoki         = "loki"
	LogsVariableTypeVictoriaLogs = "victoriametrics-logs-datasource"

	ServiceVariableName = "service_name"

	RateIntervalVar  = "$__rate_interval" // TODO: возможно где-то есть готовый
	IntervalVar      = "$__interval"
	RangeIntervalVar = "$__range"

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
	PanelSizeQuarterHigh
	PanelSizeThird
	PanelSizeFull
)

var (
	MetricDatasource = dashboard.DataSourceRef{
		Type: new(MetricVariableType),
		Uid:  new(NameToVarDS(MetricVariableName)),
	}
	LogsLokiDatasource = dashboard.DataSourceRef{
		Type: new(LogsVariableTypeLoki),
		Uid:  new(NameToVarDS(LogsVariableName)),
	}
	LogsVictoriaLogsDatasource = dashboard.DataSourceRef{
		Type: new(LogsVariableTypeVictoriaLogs),
		Uid:  new(NameToVarDS(LogsVariableName)),
	}
	GreenSteps = []dashboard.Threshold{
		{
			Color: "green",
		},
	}
	ServiceFilterPromQL = []promcog.Builder[promql.LabelSelector]{
		promql.
			NewLabelSelectorBuilder().
			Name(metriccore.ServiceNameLabel).
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
	case PanelSizeQuarterHigh:
		h = 9
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
