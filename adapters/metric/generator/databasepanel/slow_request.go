package databasepanel

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/table"
	"github.com/grafana/promql-builder/go/promql"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metricdatabase"
)

func SlowRequest() *table.PanelBuilder {
	query := promql.Div(
		promql.Sum(
			promql.
				Vector(metricdatabase.RequestDurationName+"_sum").
				Labels(generatorcore.ServiceFilterPromQL),
		).
			By([]string{metricdatabase.StmtLabelName}),
		promql.Sum(
			promql.
				Vector(metricdatabase.RequestDurationName+"_count").
				Labels(generatorcore.ServiceFilterPromQL),
		).
			By([]string{metricdatabase.StmtLabelName}),
	)

	return generatorcore.SimpleTablePanel(
		[]generatorcore.PromQLExpr{
			{
				Query:  query.String(),
				Legend: fmt.Sprintf("{{%s}}", metricdatabase.StmtLabelName),
			},
		},
		"Slow request",
		generatorcore.UnitSecond,
	)
}
