package handler_test

import (
	"bindParameters/handler"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetArrayParamPath(t *testing.T) {
	cases := []struct {
		query      string
		expected   string
		statusCode int
	}{
		{"/hoge/array/param?params[]=1", "[1]", 200},
		{"/hoge/array/param?params[]=1&params[]=2", "[1,2]", 200},
	}

	for _, c := range cases {
		// Engine不要のため無視
		res := httptest.NewRecorder()
		gc, _ := gin.CreateTestContext(res)

		h := handler.NewHandler()
		req := httptest.NewRequest("GET", c.query, nil)
		gc.Request = req

		h.GetArrayParamPath(gc)

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
