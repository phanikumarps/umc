package umc

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func Test_getAcctDestUrl(t *testing.T) {
	accountid := "2000018"
	localvmacctid := "2000019"
	type args struct {
		id *string
	}
	tests := []struct {
		name string
		args args
		want *url.URL
	}{
		// TODO: Add test cases.
		{"Account", args{id: &accountid}, getaccturl()},
		{"AccountLocalVM", args{id: &localvmacctid}, getlocalvmUrl()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAcctDestUrl(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAcctDestUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
func getaccturl() *url.URL {
	opUrl, err := url.Parse("http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')?$format=json")
	if err != nil {
		fmt.Println(err)
	}
	return opUrl
}
func getlocalvmUrl() *url.URL {
	opUrl, err := url.Parse("http://http-host:8001/j")
	if err != nil {
		fmt.Println(err)
	}
	return opUrl
}
