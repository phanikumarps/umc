package umc

import (
	"encoding/base64"
	"log"
	"net/http"
	"net/url"
)

func onPremGET(Url *url.URL) *http.Response {
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL()),
	}

	//adding the Transport object to the http Client
	client := &http.Client{
		Transport: transport,
	}

	//generating the HTTP GET request
	request, err := http.NewRequest("GET", Url.String(), nil)
	if err != nil {
		log.Println(err)
	}
	authorization := authSet()
	h := map[string][]string{
		"Content-Type":                     {"application/json"},
		"SAP-Connectivity-SCC-Location_ID": {""},
		"Authorization":                    {*authorization},
		"sap-client":                       {*setSAPClient()},
	}
	request.Header = h

	//calling the URL
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}
	return resp
}
func setSAPClient() *string {
	c := "100"
	return &c
}
func authSet() *string {
	a := "Basic" + " " + basicauth("ZZZODATA", "Totw@2022#")
	return &a
}
func basicauth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
func getDestHost() *string {
	h := "http://s42020"
	p := "8001"
	host := h + ":" + p
	return &host
}
func getUMCUri(host *string) *string {
	u := *host + "/sap/opu/odata/sap/ERP_ISU_UMC/"
	return &u
}
func constructURL(uri *string, resource *string, respType *string, id *string) *string {
	if len(*resource) > 0 {
		URL := *getUMCUri(getDestHost()) + *getResource(resource, id) + "?" + *getFormat(respType)
		return &URL
	} else {
		URL := *getUMCUri(getDestHost()) + *getResource(resource, id)
		return &URL
	}
}
func getFormat(formatType *string) *string {
	switch t := *formatType; t {
	case "json":
		f := "$format=" + *formatType
		return &f
	case "xml":
		f := "$format=" + *formatType
		return &f
	default:
		f := "$format=" + "json"
		return &f
	}
}
func proxyURL() *url.URL {
	Url, err := url.Parse("http://connectivity-proxy.kyma-system.svc.cluster.local:20003")
	if err != nil {
		log.Println(err)
	}
	return Url
}
