package umc

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func GetAccount(id *string) *acctResponse {
	acctUrl := getAcctDestUrl(id)
	resp := onPremGET(acctUrl)
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
func getAcctDestUrl(id *string) *url.URL {

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
