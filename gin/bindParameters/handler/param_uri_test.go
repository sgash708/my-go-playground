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
		res := httptest.NewRecorder()
		gc, _ := gin.CreateTestContext(res)

		h := handler.NewHandler()
		req, _ := http.NewRequest("GET", c.query, nil)
		gc.Request = req

		h.GetParamUriPath(gc)

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
