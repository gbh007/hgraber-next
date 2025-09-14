package metric

func New() (*MetricProvider, error) {
	return &MetricProvider{}, nil
}

type MetricProvider struct{}
