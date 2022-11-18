package umc

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func Meta() interface{} {
	mUrl := metaUrl()
	resp := opGET(mUrl)
	return resp
}
func Metadata() interface{} {
	mUrl := metadataUrl()
	resp := opGET(mUrl)
	return resp
}
func metaUrl() *url.URL {

	host := destHostVM()
	uri := umcURI(host)
	r := ""

	respType := "json"
	id := ""
	opUrl, err := url.Parse(*constructURL(&uri, &r, &respType, &id))
	if err != nil {
		log.Println(err)
	}
	return opUrl
}
func metadataUrl() *url.URL {

	host := destHost()
	uri := umcURI(host)
	r := ""

	respType := "json"
	id := ""
	opUrl, err := url.Parse(*constructURL(&uri, &r, &respType, &id))
	if err != nil {
		log.Println(err)
	}
	return opUrl
}
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
