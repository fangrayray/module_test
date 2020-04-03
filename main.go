package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fangrayray/module_test/objects/taxii/collections"
	tc "github.com/fangrayray/module_test/tempconv"
	"github.com/freetaxii/libstix2/objects/bundle"
)

func main() {
	fmt.Println("==== START ====")
	baseURL := "https://service.stg.kingsman.a1q7.net/services"
	// accept := "application/vnd.oasis.taxii+json; version=2.0"
	// boundle := bundle.New()
	// boundle := getWithBasicAuth(url, accept, "")
	cols := getTaxiiCollections(fmt.Sprintf("%s/collections/", baseURL))
	var cID string
	if cols != nil && len(cols.Collections) > 0 {
		cID = cols.Collections[0].GetID()
		fmt.Println(cID)
	}
	dateAfter := "2019-07-09T00:00:00.070Z"
	itemRange := "items 0-10"
	queryURL := fmt.Sprintf("%s/collections/%s/objects/?added_after=%s", baseURL, cID, dateAfter)
	bd := getSTXII(queryURL, itemRange)
	if len(bd.Objects) > 1 {
		fmt.Println(len(bd.Objects))
	}

}

func getSTXII(apiURL string, itemRange string) *bundle.Bundle {
	var username string = "trend-magic-carpet"
	var passwd string = "4T5DAyeEtcv8WGrM"
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiURL, nil)
	req.SetBasicAuth(username, passwd)
	req.Header.Set("Accept", "application/vnd.oasis.stix+json; version=2.0")
	if itemRange != "" {
		req.Header.Set("Range", itemRange)
	}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Println(resp.StatusCode)
	}
	bd, errs := bundle.Decode(resp.Body)
	if len(errs) > 0 {
		fmt.Println(errs)
		return nil
	}

	return bd

}

func getTaxiiCollections(apiURL string) *collections.Collections {
	var username string = "trend-magic-carpet"
	var passwd string = "4T5DAyeEtcv8WGrM"
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiURL, nil)
	req.SetBasicAuth(username, passwd)

	accept := "application/vnd.oasis.taxii+json; version=2.0"
	req.Header.Set("Accept", accept)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Println(resp.StatusCode)
	}
	cols := collections.New()
	json.NewDecoder(resp.Body).Decode(cols)
	return cols
	// bodyText, err := ioutil.ReadAll(resp.Body)
	// s := string(bodyText)
	// return s
}

func tryConv() {
	c := tc.AbsoluteZeroC
	fmt.Printf("%g°K", tc.CToK(c))
	fmt.Println()
	c = tc.FreezingC
	fmt.Printf("%g°K", tc.CToK(c))
	fmt.Println()
	c = tc.BoilingC
	fmt.Printf("%g°K", tc.CToK(c))
	fmt.Println()
}

func get1n2() (int, int) {
	return 1, 2
}

func get3n4() (int, int) {
	return 3, 4
}
