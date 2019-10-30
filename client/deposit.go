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
	"time"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Deposit struct {
	Class           string         `json:"$class,omitempty"`
	ToAddress      string           `json:"to_address,omitempty"`
	TransactionId       string          `json:"transactionId,omitempty"`
	DepositAmount     float32         `json:"deposit_amount,omitempty"`
	Timestamp       time.Time    `json:"timestamp,omitempty"`
}

// AddressService handles communication with the addresses related
// methods of the Bluebarricade Core API - Version 2.
type DepositService Service

func (s *DepositService) List(ctx context.Context, query *Pagination) ([]*Deposit, *http.Response, error) {
	uri := fmt.Sprintf("depositMoneyFromExternal")

	var responseStruct []*Deposit
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func (s *DepositService) Get(ctx context.Context, id string) (*Deposit, *http.Response, error) {
	uri := fmt.Sprintf("depositMoneyFromExternal/%s", id)

	var responseStruct *Deposit
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func GetDeposits(w http.ResponseWriter, r *http.Request) {
	client := NewClient(nil)
	query := &Pagination{Limit: 1}
	responseStruct, _, err := client.Deposits.List(context.Background(), query)
	if err != nil {
		fmt.Println("Error", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseStruct)
}

func GetDeposit(w http.ResponseWriter, r *http.Request) {
	address := mux.Vars(r)["id"]

	fmt.Println("GetAddress", address)
	client := NewClient(nil)
	responseStruct, _, err := client.Deposits.Get(context.Background(), address)
	if err != nil {
		fmt.Println("Error", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseStruct)
}