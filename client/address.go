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
)

type Address struct {
	Class            string  `json:"$class,omitempty"`
	Address          string  `json:"address,omitempty"`
	Currency         string  `json:"currency,omitempty"`
	Availablebalance float32 `json:"available_balance,omitempty"`
	LockBalance      float32 `json:"lock_balance,omitempty"`
}

// AddressService handles communication with the addresses related
// methods of the Bluebarricade Core API - Version 2.
type AddressService Service

func (s *AddressService) Get(ctx context.Context, id string) (*Address, *http.Response, error) {
	uri := fmt.Sprintf("Address/%s", id)

	var responseStruct *Address
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func GetAddress(baseURL string, address string) (*Address, error) {
	fmt.Println("GetAddress", address)
	client := NewClient1(baseURL, nil)
	responseStruct, _, err := client.Addresses.Get(context.Background(), address)
	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}

	return responseStruct, nil
}
