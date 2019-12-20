// This file is part of Bluebarricade Go Client.
//
// (c) Bluebarricade  <info@Bluebarricade.com>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

// AddressService handles communication with the addresses related
// methods of the Bluebarricade Core API - Version 2.
type TransferService Service

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

func GetTransfers(w http.ResponseWriter, r *http.Request) {
	client := NewClient(nil)
	query := &Pagination{Limit: 1}
	responseStruct, _, err := client.Transfers.List(context.Background(), query)
	if err != nil {
		fmt.Println("Error", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseStruct)
}

func GetTransfer(w http.ResponseWriter, r *http.Request) {
	address := mux.Vars(r)["id"]

	fmt.Println("GetAddress", address)
	client := NewClient(nil)
	responseStruct, _, err := client.Transfers.Get(context.Background(), address)
	if err != nil {
		fmt.Println("Error", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseStruct)
}
