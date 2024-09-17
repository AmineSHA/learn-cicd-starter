package auth

import (
	"testing"
	"reflect"
	"net/http"
)

func TestGetAPIKeyEmpty(t *testing.T) {
  got, got_error:= GetAPIKey({"Authorization"",})
  want,want_error  := {}, ErrNoAuthHeaderIncluded
  if !reflect.DeepEqual(want, got) {
	t.Fatalf("expected: %v, got: %v", want, got)
}



}
