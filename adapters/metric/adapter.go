package metric

import (
	"fmt"
	"time"

	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
	"github.com/gbh007/hgraber-next/adapters/metric/metricfs"
	"github.com/gbh007/hgraber-next/adapters/metric/metricserver"
	"github.com/prometheus/client_golang/prometheus"
)

type SystemType byte

const (
	UnknownSystemType = iota
	ServerSystemType
	AgentSystemType
)

type Config struct {
	ServiceName string
	Type        SystemType

	WithVersion bool
	WithFS      bool
	WithServer  bool
}

func New(cfg Config) (p *MetricProvider, err error) {
	p = &MetricProvider{
		// reg: prometheus.NewRegistry(),
		// FIXME: временная мера, пока не будет полного порта метрик
		reg: prometheus.DefaultRegisterer.(*prometheus.Registry),
	}

	defaultLabels := prometheus.Labels{}

	if cfg.ServiceName != "" {
		defaultLabels[metriccore.ServiceNameLabel] = cfg.ServiceName
	}

	switch cfg.Type {
	case ServerSystemType:
		defaultLabels[metriccore.ServiceTypeLabel] = metriccore.ServiceTypeLabelValueServer
	case AgentSystemType:
		defaultLabels[metriccore.ServiceTypeLabel] = metriccore.ServiceTypeLabelValueAgent
	case UnknownSystemType:
		defaultLabels[metriccore.ServiceTypeLabel] = metriccore.ServiceTypeLabelValueUnknown
	default:
		defaultLabels[metriccore.ServiceTypeLabel] = metriccore.ServiceTypeLabelValueUnknown
	}

	reg := prometheus.WrapRegistererWith(defaultLabels, p.reg)
	p.registerer = reg

	if cfg.WithVersion {
		err = reg.Register(metriccore.VersionInfoMetric)
		if err != nil {
			return nil, fmt.Errorf("register version info: %w", err)
		}

		// Проставляем время запуска приложения
		metriccore.VersionInfoMetric.Set(float64(time.Now().Unix()))
	}

	if cfg.WithFS {
		err = reg.Register(metricfs.ActionTime)
		if err != nil {
			return nil, fmt.Errorf("register fs action time: %w", err)
		}

		err = reg.Register(metricfs.FileTotal)
		if err != nil {
			return nil, fmt.Errorf("register fs file total: %w", err)
		}

		err = reg.Register(metricfs.FileBytes)
		if err != nil {
			return nil, fmt.Errorf("register fs file bytes: %w", err)
		}
	}

	if cfg.WithServer {
		err = reg.Register(metricserver.BookTotal)
		if err != nil {
			return nil, fmt.Errorf("register server book total: %w", err)
		}

		err = reg.Register(metricserver.PageTotal)
		if err != nil {
			return nil, fmt.Errorf("register server page total: %w", err)
		}

		err = reg.Register(metricserver.LastCollectorScrapeDuration)
		if err != nil {
			return nil, fmt.Errorf("register server collector scrape duration: %w", err)
		}
	}

	return p, nil
}

type MetricProvider struct {
	reg        *prometheus.Registry
	registerer prometheus.Registerer
}

func (p MetricProvider) Registry() *prometheus.Registry {
	return p.reg
}

func (p MetricProvider) Registerer() prometheus.Registerer {
	return p.registerer
}
