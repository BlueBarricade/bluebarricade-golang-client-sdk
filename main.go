// This file is part of Bluebarricade Go Client.
//
// (c) Bluebarricade  <info@Bluebarricade.com>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	// baseURLPath is a non-empty Client.BaseURL path to use during tests,
	// to ensure relative URLs are used for all endpoints. See issue #752.
	baseURLPath = "/api/v1"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	log.Fatal(http.ListenAndServe(":9999", router))
}
