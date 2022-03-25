package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetHogePath(t *testing.T) {
	cases := []struct {
		query      string
		expected   string
		statusCode int
	}{
		{"/hoge/path?params[]=1", "[1]", 200},
		{"/hoge/path?params[]=1&params[]=2", "[1,2]", 200},
	}

	for _, c := range cases {
		// Engine不要のため無視
		res := httptest.NewRecorder()
		gc, _ := gin.CreateTestContext(res)

		req := httptest.NewRequest("GET", c.query, nil)
		gc.Request = req

		GetHogePath(gc)

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
