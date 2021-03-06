// This file is part of Bluebarricade Go Client.
//
// (c) Bluebarricade  <info@Bluebarricade.com>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Transfer struct {
	Class         string    `json:"$class,omitempty"`
	From          string    `json:"from,omitempty"`
	To            string    `json:"to,omitempty"`
	SendAmount    float32   `json:"send_amount,omitempty"`
	ReceiveAmount float32   `json:"receive_amount,omitempty"`
	TransactionId string    `json:"transactionId,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

type Historian struct {
	Class                string           `json:"$class,omitempty"`
	TransactionId        string           `json:"transactionId,omitempty"`
	TransactionType      string           `json:"transactionType,omitempty"`
	TransactionInvoked   string           `json:"transactionInvoked,omitempty"`
	ParticipantInvoking  string           `json:"participantInvoking,omitempty"`
	IdentityUsed         string           `json:"identityUsed,omitempty"`
	EventsEmittedList    []*EventsEmitted `json:"eventsEmitted,omitempty"`
	TransactionTimestamp time.Time        `json:"transactionTimestamp,omitempty"`
}

type EventsEmitted struct {
	Class         string    `json:"$class,omitempty"`
	ToAddress     string    `json:"to_address,omitempty"`
	DepositAmount float32   `json:"deposit_amount,omitempty"`
	EventId       string    `json:"eventId,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

// AddressService handles communication with the addresses related
// methods of the Bluebarricade Core API - Version 2.
type TransferService Service

func (s *TransferService) ListHistorian(ctx context.Context, query *Pagination) ([]*Historian, *http.Response, error) {
	uri := fmt.Sprintf("system/historian")

	var responseStruct []*Historian
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)
	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func (s *TransferService) List(ctx context.Context, query *Pagination) ([]*Transfer, *http.Response, error) {
	uri := fmt.Sprintf("transferMoneyFromExternal")

	var responseStruct []*Transfer
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func (s *TransferService) Get(ctx context.Context, id string) (*Transfer, *http.Response, error) {
	uri := fmt.Sprintf("transferMoneyFromExternal/%s", id)

	var responseStruct *Transfer
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func GetTransfers(baseURL string) ([]*Transfer, error) {
	client := NewClient1(baseURL, nil)
	query := &Pagination{Limit: 1}
	responseStruct, _, err := client.Transfers.List(context.Background(), query)
	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}

	return responseStruct, nil
}

func GetTransfer(baseURL string, address string) (*Transfer, error) {

	client := NewClient1(baseURL, nil)
	responseStruct, _, err := client.Transfers.Get(context.Background(), address)
	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}

	return responseStruct, nil
}

// GetHistorian
func GetHistorian(baseURL string) ([]*Historian, error) {
	client := NewClient1(baseURL, nil)
	query := &Pagination{Limit: 10}
	responseStruct, _, err := client.Transfers.ListHistorian(context.Background(), query)
	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}

	return responseStruct, nil
}
