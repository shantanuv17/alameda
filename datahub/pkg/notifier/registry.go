package notifier

import (
	"github.com/containers-ai/alameda/datahub/pkg/notifier/metrics"
)

type Registry struct {
	notifiers []metrics.AlertInterface
}

func NewRegistry() *Registry {
	registry := Registry{}
	registry.notifiers = make([]metrics.AlertInterface, 0)
	return &registry
}

func (c *Registry) Register(metrics []metrics.AlertInterface) {
	for _, m := range metrics {
		c.notifiers = append(c.notifiers, m)
	}
}

func (c *Registry) Get(name, category string) metrics.AlertInterface {
	for _, m := range c.notifiers {
		if m.GetName() == name && m.GetCategory() == category {
			return m
		}
	}
	return nil
}

func (c *Registry) GetAll() []metrics.AlertInterface {
	return c.notifiers
}
