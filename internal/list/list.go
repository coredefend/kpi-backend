package list

import(
	"github.com/coredefend/kpi-backend/internal/channel"
)

type list struct {
	channels []channel.Channel
}

func NewList(c []channel.Channel) *list {
	return &list{channels: c}
}
