package metrics

type MetricProvider struct{}

func errorLabel(ok bool) string {
	if ok {
		return OkLabelValue
	}

	return ErrorLabelValue
}
