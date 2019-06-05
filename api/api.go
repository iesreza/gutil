package api

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

type API struct {
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

func New(apiUrl string) *API {
	obj := API{
		Url:     apiUrl,
		Headers: make(map[string]string),
		Data:    make(url.Values),
	}

	return &obj
}

func (api *API) Method(method string) *API {
	api.method = method
	return api
}

func (api *API) Param(key, value string) *API {
	api.Data.Add(key, value)
	return api
}

func (api *API) Set(key, value string) *API {
	return api.Param(key, value)
}

func (api *API) File(key, path string) *API {
	s, err := path2.File(path).Content()
	if err != nil {
		log.Error("Unable to attach file %s", err)
		return api
	}
	return api.Param(key, s)
}

func (api *API) Header(key, value string) *API {
	api.Headers[key] = value
	return api
}

func (api *API) ContentType(value string) *API {
	api.Headers["ContentType"] = value
	return api
}

func (api *API) Type(t apiType) *API {
	api.t = t
	if t == JSON {
		api.ContentType("application/json")
	}
	if t == HTML || t == TEXT {
		api.ContentType("text/plain")
	}
	return api
}

func (a *API) Fresh() *API {
	freshApi := API{
		Url:     a.Url,
		method:  a.method,
		Headers: a.Headers,
		Data:    make(url.Values),
	}
	return &freshApi
}

func (api *API) Call(key, value string) *API {
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

func (api *API) String() string {
	return string(api.Result)
}

func (api *API) Bytes() []byte {
	return api.Result
}

func (api *API) JSON() (interface{}, error) {
	var obj interface{}
	err := json.Unmarshal(api.Result, &obj)
	return obj, err
}

func (api *API) Scan(obj *interface{}) (interface{}, error) {
	err := json.Unmarshal(api.Result, &obj)
	return obj, err
}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}
