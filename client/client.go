// This file is part of Bluebarricade Go Client.
//
// (c) Bluebarricade  <info@Bluebarricade.com>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/joho/godotenv"
)

const (
	defaultBaseURL = "http://153.112.160.162:3000/api/"
)

type Client struct {
	client *http.Client

	BaseURL *url.URL

	common Service

	Addresses      *AddressService
	Deposits       *DepositService
	LockBalances   *LockBalanceService
	UnlockBalances *UnlockBalanceService
	Transfers      *TransferService
	Withdraws      *WithdrawService

	Delegates    *DelegatesService
	Node         *NodeService
	Peers        *PeersService
	Transactions *TransactionsService
	Votes        *VotesService
	Wallets      *WalletsService
}

type Service struct {
	client *Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	envBaseUrl := os.Getenv("BASE_URL")
	if envBaseUrl != "" {
		baseURL, _ = url.Parse(defaultBaseURL)
	}

	c := &Client{client: httpClient, BaseURL: baseURL}
	c.common.client = c

	// c.Blocks = (*BlocksService)(&c.common)
	c.Deposits = (*DepositService)(&c.common)
	c.Addresses = (*AddressService)(&c.common)
	c.LockBalances = (*LockBalanceService)(&c.common)
	c.UnlockBalances = (*UnlockBalanceService)(&c.common)
	c.Transfers = (*TransferService)(&c.common)
	c.Withdraws = (*WithdrawService)(&c.common)

	c.Delegates = (*DelegatesService)(&c.common)
	c.Node = (*NodeService)(&c.common)
	c.Peers = (*PeersService)(&c.common)
	c.Transactions = (*TransactionsService)(&c.common)
	c.Votes = (*VotesService)(&c.common)
	c.Wallets = (*WalletsService)(&c.common)

	return c
}

func NewClient1(rawURL string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(rawURL)

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// envBaseUrl := os.Getenv("BASE_URL")
	// if envBaseUrl != "" {
	// 	baseURL, _ = url.Parse(defaultBaseURL)
	// }

	c := &Client{client: httpClient, BaseURL: baseURL}
	c.common.client = c

	// c.Blocks = (*BlocksService)(&c.common)
	c.Deposits = (*DepositService)(&c.common)
	c.Addresses = (*AddressService)(&c.common)
	c.LockBalances = (*LockBalanceService)(&c.common)
	c.UnlockBalances = (*UnlockBalanceService)(&c.common)
	c.Transfers = (*TransferService)(&c.common)
	c.Withdraws = (*WithdrawService)(&c.common)

	c.Delegates = (*DelegatesService)(&c.common)
	c.Node = (*NodeService)(&c.common)
	c.Peers = (*PeersService)(&c.common)
	c.Transactions = (*TransactionsService)(&c.common)
	c.Votes = (*VotesService)(&c.common)
	c.Wallets = (*WalletsService)(&c.common)

	return c
}

func (c *Client) SendRequest(ctx context.Context, method string, urlStr string, queryString interface{}, body interface{}, model interface{}) (*http.Response, error) {
	// Create a new HTTP request
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		if method == "POST" {
			json, _ := json.Marshal(body)
			buf = bytes.NewBuffer(json)
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		return nil, err
	}

	if queryString != nil {
		switch v := queryString.(type) {
		case *Pagination:
			if v.Page == 0 {
				v.Page = 1
			}
		}

		params, _ := query.Values(queryString)

		req.URL.RawQuery = params.Encode()
	}

	req.Header.Set("Content-Type", "application/json")

	// Execute the previously created HTTP request
	resp, err := c.client.Do(req)

	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = url.String()
				return nil, e
			}
		}

		return nil, err
	}

	// Map the JSON response to a struct
	if w, ok := model.(io.Writer); ok {
		_, _ = io.Copy(w, resp.Body)
	} else {
		decErr := json.NewDecoder(resp.Body).Decode(model)

		if decErr == io.EOF {
			decErr = nil
		}

		if decErr != nil {
			err = decErr
		}
	}

	defer resp.Body.Close()

	return resp, err
}
