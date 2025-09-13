package generatorcore

import "github.com/grafana/grafana-foundation-sdk/go/dashboard"

func GreenTrashHold() *dashboard.ThresholdsConfigBuilder {
	return dashboard.
		NewThresholdsConfigBuilder().
		Steps(GreenSteps)
}
