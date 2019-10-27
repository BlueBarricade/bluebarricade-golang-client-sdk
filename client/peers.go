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

// PeersService handles communication with the peers related
// methods of the Bluebarricade Core API - Version 2.
type PeersService Service

// Get all peers.
func (s *PeersService) List(ctx context.Context, query *Pagination) (*Peers, *http.Response, error) {
	var responseStruct *Peers
	resp, err := s.client.SendRequest(ctx, "GET", "peers", query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a peer by the given IP address.
func (s *PeersService) Get(ctx context.Context, ip string) (*GetPeer, *http.Response, error) {
	uri := fmt.Sprintf("peers/%v", ip)

	var responseStruct *GetPeer
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}
