package generator

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func (g Generator) Build() (dashboard.Dashboard, error) {
	builder := dashboard.NewDashboardBuilder("[Experimental] HG next board").
		Uid(g.uid)

	g.WithPanels(builder)
	g.WithVariables(builder)

	d, err := builder.Build()
	if err != nil {
		return dashboard.Dashboard{}, fmt.Errorf("build: %w", err)
	}

	return d, nil
}
