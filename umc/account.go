package umc

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func Account(id *string) *acctResponse {
	acctUrl := acctDestUrl(id)
	resp := opGET(acctUrl)
	return unmarshalAcct(resp)
}
func unmarshalAcct(resp *http.Response) *acctResponse {
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
	}
	var acctData *acctResponse
	json.Unmarshal(data, &acctData)
	return acctData
}
func acctDestUrl(id *string) *url.URL {

	host := destHost()
	uri := umcURI(host)
	r := "accounts"

	respType := "json"
	opUrl, err := url.Parse(*constructURL(&uri, &r, &respType, id))
	if err != nil {
		log.Println(err)
	}
	return opUrl
}
