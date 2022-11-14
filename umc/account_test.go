package umc

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func Test_getAcctDestUrl(t *testing.T) {
	accountid := "2000018"
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
