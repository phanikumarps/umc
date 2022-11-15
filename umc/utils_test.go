package umc

import (
	"log"
	"net/url"
	"reflect"
	"testing"
)

func Test_setSAPClient(t *testing.T) {
	c := "100"
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{"default", c},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sapClient(); got != tt.want {
				t.Errorf("setSAPClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_proxyURL(t *testing.T) {
	tests := []struct {
		name string
		want *url.URL
	}{
		// TODO: Add test cases.
		{"proxy url", prxUrl()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := proxyURL(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("proxyURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
func prxUrl() *url.URL {
	Url, err := url.Parse("http://connectivity-proxy.kyma-system.svc.cluster.local:20003")
	if err != nil {
		log.Println(err)
	}
	return Url
}
