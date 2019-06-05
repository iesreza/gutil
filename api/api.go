package API

import (
	"encoding/json"
	"errors"
	"fmt"
	path2 "github.com/iesreza/gutil/path"
	"gutil/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type apiType int

const (
	JSON apiType = 1
	TEXT apiType = 2
	HTML apiType = 3
)

type api struct {
	Url        string
	method     string
	Data       url.Values
	Headers    map[string]string
	Result     []byte
	Error      error
	Response   *http.Response
	StatusCode int
	Status     string
	t          apiType
}

func New(url string) *api {
	obj := api{
		Url: url,
	}

	return &obj
}

func (api *api) Method(method string) *api {
	api.method = method
	return api
}

func (api *api) Param(key, value string) *api {
	api.Data.Add(key, value)
	return api
}

func (api *api) Set(key, value string) *api {
	return api.Param(key, value)
}

func (api *api) File(key, path string) *api {
	s, err := path2.File(path).Content()
	if err != nil {
		log.Error("Unable to attach file %s", err)
		return api
	}
	return api.Param(key, s)
}

func (api *api) Header(key, value string) *api {
	api.Headers[key] = value
	return api
}

func (api *api) ContentType(value string) *api {
	api.Headers["ContentType"] = value
	return api
}

func (api *api) Type(t apiType) *api {
	api.t = t
	if t == JSON {
		api.ContentType("application/json")
	}
	if t == HTML || t == TEXT {
		api.ContentType("text/plain")
	}
	return api
}

func (a *api) Fresh(key, value string) *api {
	freshApi := api{
		Url:     a.Url,
		method:  a.method,
		Headers: a.Headers,
	}
	return &freshApi
}

func (api *api) Call(key, value string) *api {
	client := http.Client{}
	req, err := http.NewRequest(api.method, api.Url, strings.NewReader(api.Data.Encode()))

	for key, val := range api.Headers {
		req.Header.Add(key, val)
	}

	api.Response, api.Error = client.Do(req)
	if api.Error != nil {
		log.Error(err.Error())
	}

	api.StatusCode = api.Response.StatusCode
	api.Status = api.Response.Status
	api.Error = err
	api.Result, _ = ioutil.ReadAll(api.Response.Body)

	if api.t == JSON || api.StatusCode == 200 {
		if !isJSON(api.String()) {
			if err == nil {
				api.Error = errors.New(fmt.Sprintf("Invalid JSON format %s", api.String()))
			}

		}
	}

	return api
}

func (api *api) String() string {
	return string(api.Result)
}

func (api *api) Bytes() []byte {
	return api.Result
}

func (api *api) JSON() (interface{}, error) {
	var obj interface{}
	err := json.Unmarshal(api.Result, &obj)
	return obj, err
}

func (api *api) Scan(obj *interface{}) (interface{}, error) {
	err := json.Unmarshal(api.Result, &obj)
	return obj, err
}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}
