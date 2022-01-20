package work

type responseQueue struct {
	responseType string
	orchestrator IOrchestrator
}

func NewResponseQueue(responseType string, orchestrator IOrchestrator) *responseQueue {
	return &responseQueue{
		responseType: responseType,
		orchestrator: orchestrator,
	}
}
