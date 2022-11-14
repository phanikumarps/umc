package umc

import "testing"

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
