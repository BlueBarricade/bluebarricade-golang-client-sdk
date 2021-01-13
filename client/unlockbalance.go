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

type UnlockBalance struct {
	Class               string    `json:"$class,omitempty"`
	ToAddress           string    `json:"to_address,omitempty"`
	TransactionId       string    `json:"transactionId,omitempty"`
	UnlockBalanceAmount float32   `json:"unlock_amount,omitempty"`
	Timestamp           time.Time `json:"timestamp,omitempty"`
}

// AddressService handles communication with the addresses related
// methods of the Bluebarricade Core API - Version 2.
type UnlockBalanceService Service

func (s *UnlockBalanceService) List(ctx context.Context, query *Pagination) ([]*UnlockBalance, *http.Response, error) {
	uri := fmt.Sprintf("depositMoneyFromExternal")

	var responseStruct []*UnlockBalance
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func (s *UnlockBalanceService) Get(ctx context.Context, id string) (*UnlockBalance, *http.Response, error) {
	uri := fmt.Sprintf("depositMoneyFromExternal/%s", id)

	var responseStruct *UnlockBalance
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func GetUnlockBalances(baseURL string) ([]*UnlockBalance, error) {
	client := NewClient1(baseURL, nil)
	query := &Pagination{Limit: 1}
	responseStruct, _, err := client.UnlockBalances.List(context.Background(), query)
	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}

	return responseStruct, nil
}

func GetUnlockBalance(baseURL string, address string) (*UnlockBalance, error) {

	client := NewClient1(baseURL, nil)
	responseStruct, _, err := client.UnlockBalances.Get(context.Background(), address)
	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}

	return responseStruct, nil
}
