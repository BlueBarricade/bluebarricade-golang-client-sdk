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
	

	"./client"
	"github.com/gorilla/mux"
)

const (
	// baseURLPath is a non-empty Client.BaseURL path to use during tests,
	// to ensure relative URLs are used for all endpoints. See issue #752.
	baseURLPath = "/api/v1"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(baseURLPath+"/addresses/{id}", client.GetAddress).Methods("GET")
	router.HandleFunc(baseURLPath+"/deposits", client.GetDeposits).Methods("GET")
	router.HandleFunc(baseURLPath+"/deposits/{id}", client.GetDeposit).Methods("GET")
	router.HandleFunc(baseURLPath+"/lockbalances", client.GetLockBalances).Methods("GET")
	router.HandleFunc(baseURLPath+"/lockbalances/{id}", client.GetLockBalance).Methods("GET")
	router.HandleFunc(baseURLPath+"/unlockbalances", client.GetUnlockBalances).Methods("GET")
	router.HandleFunc(baseURLPath+"/unlockbalances/{id}", client.GetUnlockBalances).Methods("GET")
	router.HandleFunc(baseURLPath+"/transfers", client.GetTransfers).Methods("GET")
	router.HandleFunc(baseURLPath+"/transfers/{id}", client.GetTransfer).Methods("GET")
	router.HandleFunc(baseURLPath+"/withdraws", client.GetWithdraws).Methods("GET")
	router.HandleFunc(baseURLPath+"/withdraws/{id}", client.GetWithdraw).Methods("GET")
	log.Fatal(http.ListenAndServe(":9999", router))
}
