package generatorcore

import "github.com/grafana/promql-builder/go/promql"

func RPSExpr(metric string, by []string) string {
	builder := promql.Sum(
		promql.Rate(
			promql.
				Vector(metric).
				Labels(ServiceFilterPromQL).
				Range(RateIntervalVar),
		),
	)

	if len(by) > 0 {
		builder.By(by)
	}

	return builder.String()
}

func DeltaExpr(metric string, by []string) string {
	builder := promql.Sum(
		promql.Delta(
			promql.
				Vector(metric).
				Labels(ServiceFilterPromQL).
				Range(RateIntervalVar),
		),
	)

	if len(by) > 0 {
		builder.By(by)
	}

	return builder.String()
}

func SumExpr(metric string, by []string) string {
	builder := promql.Sum(
		promql.
			Vector(metric).
			Labels(ServiceFilterPromQL),
	)

	if len(by) > 0 {
		builder.By(by)
	}

	return builder.String()
}
