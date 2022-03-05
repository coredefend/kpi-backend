package metric_test

import (
	"testing"
	"github.com/coredefend/kpi-backend/internal/metric"
)

func TestMetric(t *testing.T) {
	m, _ := metric.NewMetric(10, 0, 0, 0)
	if m == nil {
		t.Errorf("should not be nil")
	}
	if m.Responses() != 10 {
		t.Errorf("must equal 10")
	}
	m.SetGrossProfit(10000.00)
	if m.GrossProfit() != 10000.00 {
		t.Errorf("values should match")
	}
	
}
