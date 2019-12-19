package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var (
	// Queue is the name of the queue.
	Queue string
	// Host is the name of the host to connect to.
	Host string
)

// UpdateRequest is a request model for updating prices.
type UpdateRequest struct {
	ID    int
	Name  string
	Price Price
}

// QueueResponse is a queue item response from the API.
type QueueResponse struct {
	ID           int
	Name         string
	URL          string
	Domain       string
	LastChecked  time.Time
	IsChecking   bool
	PriceHistory []Price
}

// Price is the price model for the QueueResponse.
type Price struct {
	Value     float32
	Timestamp time.Time
}

/*
GetAvailableQueues returns a list of available queues from the API.
*/
func GetAvailableQueues() ([]string, error) {
	client := &http.Client{}

	res, err := client.Get(Host + "/queues")
	if err != nil {
		return nil, err
	}

	respBody := res.Body
	defer respBody.Close()

	availableQueues := struct {
		Queues []string `json:"queues"`
	}{}
	if err := json.NewDecoder(respBody).Decode(&availableQueues); err != nil {
		return nil, err
	}

	return availableQueues.Queues, nil
}

/*
GetJobFromQueue gets the next job from the queue via the API.
*/
func GetJobFromQueue() (*QueueResponse, error) {
	client := &http.Client{}
	responseModel := QueueResponse{}

	res, err := client.Get(Host + "/queues/" + Queue + "/next")
	if err != nil {
		return nil, err
	}

	respBody := res.Body
	defer respBody.Close()

	if err := json.NewDecoder(respBody).Decode(&responseModel); err != nil {
		return nil, err
	}

	return &responseModel, nil
}

/*
UpdatePrice calls the price update API endpoint with an UpdateRequest model in the body.
*/
func UpdatePrice(update *UpdateRequest) error {
	client := &http.Client{}
	rBody, err := json.Marshal(update)
	if err != nil {
		return err
	}

	res, err := client.Post(Host+"/prices/update/"+strconv.Itoa(update.ID), "application/json", bytes.NewBuffer(rBody))
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode > 200 {
		return fmt.Errorf("could not update price, status code %s returned", strconv.Itoa(res.StatusCode))
	}

	return nil
}
