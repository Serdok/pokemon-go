package mocking

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func MockParam(c *gin.Context, method string, params gin.Params) {
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")

	c.Params = params
}

func MockBody(c *gin.Context, method string, content interface{}) {
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonbytes))
}
