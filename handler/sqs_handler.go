package handler

import sqsclient "github.com/inaciogu/go-sqs-consumer/client"

// SQSHandler is responsible for running the SQS clients concurrently
type SQSHandler struct {
	Clients []sqsclient.SQSClientInterface
}

func New(clients []sqsclient.SQSClientInterface) *SQSHandler {
	return &SQSHandler{
		Clients: clients,
	}
}

func (h *SQSHandler) Run() {
	for _, client := range h.Clients {
		go client.Poll()
	}

	select {}
}
