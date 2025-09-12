package generator

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func (g Generator) Build() (dashboard.Dashboard, error) {
	builder := dashboard.
		NewDashboardBuilder("[Experimental] HG next board").
		Uid(g.uid).
		Timezone("Asia/Krasnoyarsk"). // TODO: в конфиг
		Time("now-6h", "now").        // TODO: в конфиг
		WeekStart("monday").          // TODO: в конфиг
		Refresh("1m").                // TODO: в конфиг
		Tooltip(dashboard.DashboardCursorSyncCrosshair)

	g.WithPanels(builder)
	g.WithVariables(builder)
	g.WithTagAndAnotation(builder)

	d, err := builder.Build()
	if err != nil {
		return dashboard.Dashboard{}, fmt.Errorf("build: %w", err)
	}

	return d, nil
}
