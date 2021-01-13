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

type Withdraw struct {
	Class          string    `json:"$class,omitempty"`
	FromAddress    string    `json:"from_address,omitempty"`
	WithdrawAmount float32   `json:"withdraw_amount,omitempty"`
	TransactionId  string    `json:"transactionId,omitempty"`
	Timestamp      time.Time `json:"timestamp,omitempty"`
}

// AddressService handles communication with the addresses related
// methods of the Bluebarricade Core API - Version 2.
type WithdrawService Service

func (s *WithdrawService) List(ctx context.Context, query *Pagination) ([]*Withdraw, *http.Response, error) {
	uri := fmt.Sprintf("transferMoneyFromExternal")

	var responseStruct []*Withdraw
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func (s *WithdrawService) Get(ctx context.Context, id string) (*Withdraw, *http.Response, error) {
	uri := fmt.Sprintf("transferMoneyFromExternal/%s", id)

	var responseStruct *Withdraw
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func GetWithdraws(baseURL string) ([]*Withdraw, error) {
	client := NewClient1(baseURL, nil)
	query := &Pagination{Limit: 1}
	responseStruct, _, err := client.Withdraws.List(context.Background(), query)
	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}

	return responseStruct, nil
}

func GetWithdraw(baseURL string, address string) (*Withdraw, error) {

	fmt.Println("GetAddress", address)
	client := NewClient1(baseURL, nil)
	responseStruct, _, err := client.Withdraws.Get(context.Background(), address)
	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}

	return responseStruct, nil
}
