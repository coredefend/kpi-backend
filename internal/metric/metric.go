package metric

type metric struct {
	responses int
	leads	int
	appointments int
	contracts int
	grossProfit float64
}

func NewMetric(responses, leads, appointments, contracts int) (*metric, error) {
	return &metric{
		responses: responses,
		leads: leads,
		appointments: appointments,
		contracts: contracts,
	}, nil
}

func (m *metric) Responses() int {
	return m.responses
}

func (m *metric) Leads() int {
	return m.leads
}

func (m *metric) Appointments() int {
	return m.appointments
}

func (m *metric) Contracts() int {
	return m.contracts
}

func (m *metric) GrossProfit() float64) {
	return m.grossProfit
}

func (m *metric) SetGrossProfit(profit float64) {
	m.grossProfit = profit	
}
