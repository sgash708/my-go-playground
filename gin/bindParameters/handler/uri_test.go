package handler_test

import (
	"bindParameters/handler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetParamUriPath(t *testing.T) {
	cases := []struct {
		query      string
		expected   string
		statusCode int
	}{
		{"/hoge/param/a1", "a1", 200},
	}

	for _, c := range cases {
		// param := gin.Param{"id", c.expected}
		// params := gin.Params{param}

		res := httptest.NewRecorder()
		gc, e := gin.CreateTestContext(res)

		req, _ := http.NewRequest("GET", c.query, nil)
		e.ServeHTTP(res, req)
		// gc.Request = req

		handler.NewHandler().GetParamUriPath(gc)

		targetStatusCode := res.Code
		if res.Code != c.statusCode {
			t.Errorf("expected: %v, input: %v", c.statusCode, targetStatusCode)
		}
		targetWord := res.Body.String()
		if targetWord != c.expected {
			t.Errorf("expected: %v, input: %v", c.expected, targetWord)
		}
	}
}
