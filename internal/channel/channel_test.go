package channel_test

import (
	"testing"
	"github.com/coredefend/kpi-backend/internal/channel"
)

func TestNewChannel(t *testing.T) {
	c, err := channel.New("postcards", 100, 50.50)
	if err != nil {
		t.Errorf("should not have error")
	}
	if c == nil {
		t.Errorf("should not be nil")
	}
}
