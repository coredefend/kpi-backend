package channel

type channel struct {
	name string
	quantity int
	cost float64	
}

func New(name string, quantity int, cost float64) (*channel, error) {
	return &channel{
		name: name,
		quantity: quantity,
		cost: cost,
	}, nil
}
