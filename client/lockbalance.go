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

type LockBalance struct {
	Class             string    `json:"$class,omitempty"`
	Address           string    `json:"address,omitempty"`
	TransactionId     string    `json:"transactionId,omitempty"`
	LockBalanceAmount float32   `json:"lock_amount,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}

// AddressService handles communication with the addresses related
// methods of the Bluebarricade Core API - Version 2.
type LockBalanceService Service

func (s *LockBalanceService) List(ctx context.Context, query *Pagination) ([]*LockBalance, *http.Response, error) {
	uri := fmt.Sprintf("depositMoneyFromExternal")

	var responseStruct []*LockBalance
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func (s *LockBalanceService) Get(ctx context.Context, id string) (*LockBalance, *http.Response, error) {
	uri := fmt.Sprintf("depositMoneyFromExternal/%s", id)

	var responseStruct *LockBalance
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func GetLockBalances(w http.ResponseWriter, r *http.Request) {
	client := NewClient(nil)
	query := &Pagination{Limit: 1}
	responseStruct, _, err := client.LockBalances.List(context.Background(), query)
	if err != nil {
		fmt.Println("Error", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseStruct)
}

func GetLockBalance(w http.ResponseWriter, r *http.Request) {
	address := mux.Vars(r)["id"]

	fmt.Println("GetAddress", address)
	client := NewClient(nil)
	responseStruct, _, err := client.LockBalances.Get(context.Background(), address)
	if err != nil {
		fmt.Println("Error", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseStruct)
}
