package definition

type Communication struct {
	queueName string
}

func NewCommunicationDefinition(queueName string) *Communication {
	return &Communication{queueName: queueName}
}
