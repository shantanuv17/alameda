package notifier

import (
	"github.com/containers-ai/alameda/datahub/pkg/notifier/metrics"
)

type Management struct {
	notifiers []metrics.AlertInterface
}

func NewManagement() *Management {
	management := Management{}
	management.notifiers = make([]metrics.AlertInterface, 0)
	return &management
}

func (c *Management) AddMetrics(metrics []metrics.AlertInterface) {
	for _, m := range metrics {
		c.notifiers = append(c.notifiers, m)
	}
}

func (c *Management) GetMetrics(name, category string) metrics.AlertInterface {
	for _, m := range c.notifiers {
		if m.GetName() == name && m.GetCategory() == category {
			return m
		}
	}
	return nil
}

func (c *Management) GetAllMetrics() []metrics.AlertInterface {
	return c.notifiers
}
