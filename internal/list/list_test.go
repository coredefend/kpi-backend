package list_test

import (
	"testing"

	"github.com/coredefend/kpi-backend/internal/list"
	"github.com/coredefend/kpi-backend/internal/channel"
)

func TestList(t *testing.T) {
	postcard, _ := channel.New("Postcard", 1279, 626.71)
	dfd := list.NewList([]channel.Channel{*postcard})
	if dfd == nil {
		t.Errorf("should not be nil")
	}
}
