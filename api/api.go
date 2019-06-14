package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/iesreza/gutil/log"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"strconv"
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
	Data       map[string]string
	Headers    map[string]string
	Files      map[string]string
	Result     []byte
	Error      error
	Response   *http.Response
	StatusCode int
	Status     string
	debug      bool
}

func New(apiUrl string) *API {
	obj := API{
		Url:     apiUrl,
		Headers: make(map[string]string),
		Data:    make(map[string]string),
		Files:   make(map[string]string),
		method:  "POST",
	}

	return &obj
}

func (api *API) Method(method string) *API {
	api.method = method
	return api
}

func (api *API) Debug() *API {
	api.debug = true
	return api
}

func (api *API) Param(key, value string) *API {
	api.Data[key] = value
	return api
}

func (api *API) SetInterface(key string, v interface{}) {
	switch t := v.(type) {
	case int:
		api.Param(key, strconv.Itoa(t))
	case int32:
		api.Param(key, string(int64(t)))
	case int64:
		api.Param(key, string(t))
	case float64:
		api.Param(key, strconv.FormatFloat(t, 'f', 6, 64))
	case float32:
		api.Param(key, fmt.Sprintf("%f", t))
	case string:
		api.Param(key, t)
	case bool:
		api.Param(key, strconv.FormatBool(t))
	default:
		b, _ := json.Marshal(t)
		api.Param(key, string(b))
	}
}

func (api *API) Set(data interface{}) *API {
	ref := reflect.ValueOf(data)

	switch ref.Kind() {
	case reflect.Map:
		for _, e := range ref.MapKeys() {
			v := ref.MapIndex(e)
			api.SetInterface(e.String(), v.Interface())
		}
	case reflect.Struct:
		keys := reflect.Indirect(ref)
		for i := 0; i < ref.NumField(); i++ {
			api.SetInterface(keys.Type().Field(i).Name, ref.Field(i).Interface())
		}
	}

	return api

}

func (api *API) Attach(key, path string) *API {
	api.Files[key] = path
	return api
}

func (api *API) Header(key, value string) *API {
	api.Headers[key] = value
	return api
}

func (a *API) Fresh() *API {
	freshApi := New(a.Url)
	freshApi.method = a.method
	freshApi.Headers = a.Headers

	return freshApi
}

func (api *API) Call() *API {
	client := &http.Client{}
	if api.debug {
		log.Notice("API:%s", api.Url)
	}
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	for key, path := range api.Files {

		file, err := os.Open(path)
		if err != nil {
			log.Error("Unable to open file %s", err)
			api.Error = err
			return api
		}
		fileContents, err := ioutil.ReadAll(file)
		if err != nil {
			log.Error("Unable to read file %s", err)
			api.Error = err
			return api
		}
		fi, err := file.Stat()
		if err != nil {
			log.Error("Unable to get stat of file %s", err)
			api.Error = err
			return api
		}
		file.Close()

		part, err := writer.CreateFormFile(key, fi.Name())
		if err != nil {
			log.Error("Unable to write the field %s", err)
			api.Error = err
			return api
		}
		part.Write(fileContents)

	}

	for key, val := range api.Data {
		writer.WriteField(key, val)
	}
	err := writer.Close()
	if err != nil {
		log.Error("Unable to close writer %s", err)
		api.Error = err
		return api
	}

	req, err := http.NewRequest(api.method, api.Url, body)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	for key, val := range api.Headers {
		req.Header.Set(key, val)
	}

	api.Response, api.Error = client.Do(req)

	//api.Response, api.Error = client.PostForm(api.Url, api.Data)
	if api.Error != nil {
		api.Error = err
		log.Error(api.Error.Error())
	}
	defer api.Response.Body.Close()

	api.StatusCode = api.Response.StatusCode
	api.Status = api.Response.Status

	api.Result, _ = ioutil.ReadAll(api.Response.Body)
	if api.debug {
		log.Notice("Response:\n" + string(api.Result))
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

func (api *API) Scan(obj interface{}) error {
	err := json.Unmarshal(api.Result, &obj)
	return err
}
