package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKeyEmpty(t *testing.T) {
	testHeader := http.Header{}
	testHeader.Add("Authorization", " ")
	got, gotError := GetAPIKey(testHeader)
	want, wantError := " ", errors.New("no authorization header included")
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v - %v, got: %v - %v", want, wantError.Error(), got, gotError.Error)
	}

}
