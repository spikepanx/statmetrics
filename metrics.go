package statmetrics

import "github.com/prometheus/client_golang/prometheus"

func NewCounterVec(ns, sub, name, help string, labels []string) *prometheus.CounterVec {
	return prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: name,
		Subsystem: sub,
		Name:      name,
		Help:      help,
	}, labels)
}

func NewGaugeVec(ns, sub, name, help string, labels []string) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: name,
		Subsystem: sub,
		Name:      name,
		Help:      help,
	}, labels)
}
