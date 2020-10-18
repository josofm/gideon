// +build unit

package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/josofm/gideon/api"
	"github.com/josofm/gideon/controller"

	"github.com/stretchr/testify/assert"
)

func TestUpAPI(t *testing.T) {
	c := controller.Controller{}
	api := api.NewApi(c)

	r, err := http.NewRequest("GET", "/up", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.Up)

	handler.ServeHTTP(rr, r)

	assert.Nil(t, rr.Code, "e ai?")

}
