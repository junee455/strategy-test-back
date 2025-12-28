package agent

type Agent struct{}

type AgentDescription struct {
	ID    string `json:"id,string"`
	State string `json:"state"`
}

type IAgent interface {
	GetAgentDescription() AgentDescription
}
