package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/iesreza/gutil/configuration"
	"github.com/iesreza/gutil/hashmap"
	"github.com/iesreza/gutil/linkedlist"
	"github.com/iesreza/gutil/logger"
	"github.com/iesreza/gutil/path"
	"github.com/iesreza/gutil/rsa"
	"github.com/iesreza/gutil/str"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"time"
)

var log = logger.New()

// ConfigVersion returns the version, this should be incremented every time the config changes
var ConfigVersion = "1.0.0"

type Config struct {
	Version          string
	IntegerValue     int
	StringValue      string
	StringArrayValue []string
}

type Session struct {
	Expire int64
	ID     string
}

type Test struct {
	Timestamp int64
	Receipt   string
	Serial    string
}

type Response struct {
	Serial  string
	Payload string
}

func main() {

	t := Test{
		Timestamp: time.Now().Unix(),
		Serial:    "386cd23f9dc568bcd88bb284f066d4f25372592c",
	}

	publicKeyString := `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA37mV1a5mH2lQRrXm3z9G
TXiIq542Xy8xuD5V/bmxsWiT0foLtQhNdsKc1jOKJ7YT0Fl7vBY7lO/gluUvKXf+
V3T90XA2GpP3RiQaLWMgJzYowDnCyBktLUqzSg6YL8Y4OAbfK3nOISxBfuBopKap
UND4vgrbSjYZUQLo01G1NRwCyrIwo1LGrmdo3tJR4G/Te/g0E99H3k8w8g+Cdu3m
6NSTXoGPPm8hJCTz5rFeF92fYk3WrseaWpMghX4GoAF+2wJowjjHBxTqFQ/iEHk0
9sc6HZWAYm8U/qWb7CmT3kfZ9beqcIai32WzwhcLeDvLfJwqPDeyWSCiYO+w7Is6
TTO2r35SgbzIwxoXaegBpgl3bOAfkPVfybtAMMCQbKF9DQaQAUya3XXu6LjotwIB
ifesMlD3OsDzXQhjeRcdgPUlEbrqueYghg3evaa3RxfTvJeaYp0Hoo4bgQYIXZis
1GaqeExWV11m8RBqR16LbgpD5K/f1zhAALWn98iFSqwvxEK9ReQfutWv+7ZvHsIu
5ywAkzgRepWSUYqcyFhXh7dxAWYajK9QOYYoNnELKy9U/6GeGc3fuxwKaeO0fJU1
7oUWofpdYDfK4cZc337tyW7QSu/4ik+DP5UtsLblU1ocCJqmwV8Xw7czmJs9sffi
zk7Dgn86J2K2mcDOAZijaVcCAwEAAQ==
-----END PUBLIC KEY-----`
	publicKey := rsa.ParsePublicKey(publicKeyString)

	encrypted := publicKey.Encrypt(t)
	encryptedString := base64.StdEncoding.EncodeToString(encrypted)

	vals := url.Values{}
	vals.Add("Serial", "386cd23f9dc568bcd88bb284f066d4f25372592c")
	vals.Add("Token", encryptedString)
	vals.Add("Subject", "This is a test mail")
	f, _ := path.File("./test.pdf").Content()
	vals.Add("Receipt", "reza@ies-italia.it")
	vals.Add("Body", "test email")
	vals.Add("Attachment", f)

	email, _ := path.File("./email.html").Content()
	data := map[string]string{
		"Serial":  "386cd23f9dc568bcd88bb284f066d4f25372592c",
		"Token":   encryptedString,
		"Subject": "Test email with attachment",
		"Receipt": "reza@ies-italia.it",
		"Body":    email,
	}

	//response, err := http.PostForm("http://192.168.1.175:8010/command/", vals)
	response, err := UploadFile("http://192.168.1.175:8010/command/", data, "Attachment", "C:/Users/mreza/go/src/gutil/test.pdf")
	if err != nil {
		log.ErrorF("Unable to connect %s", err)
		return
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	fmt.Println(buf.String())
	return
	//
	clink := hashmap.CLink()
	clink.InsertNode("google.com", 12)
	clink.InsertNode("goo.com", 14)
	i, e := clink.Find("goo.com")
	fmt.Println(i, (*i).(int), e)

	i, e = clink.Find("google.com")
	fmt.Println(i, (*i).(int), e)

	i, e = clink.Find("gle.com")
	fmt.Println(i, (*i).(int), e)
	return

	var list = linkedlist.List{}

	list.SetMatchFunc(func(needle interface{}, el interface{}) bool {
		return needle.(Session).ID == el.(Session).ID
	})

	list.PushOnce(Session{1558378380, "abcd1"})
	list.PushOnce(Session{1568378380, "abcd2"})
	list.PushOnce(Session{1538378380, "abcd3"})
	list.PushOnce(Session{1533378380, "abcd4"})
	fmt.Println(list.String())

	var v int64
	v = time.Now().Unix()
	list.RemoveFunc(v, func(needle interface{}, el interface{}) bool {

		return needle.(int64) < el.(Session).Expire
	})

	fmt.Println(list.String())

	return
	var config = Config{}

	configurator := configuration.GetInstance(&config, ConfigVersion)
	configurator.App = "gutil"
	configurator.Load()

	configurator.Set("IntegerValue", 25)
	config.IntegerValue = 40
	err = configurator.Update()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(config.IntegerValue)

	fmt.Println(str.S(1.00000045).Trim("5").Quote().ReplaceAll("0", "9"))

	//logExample()
	//fileExample()

}

func logExample() {

	//Logger
	log.Critical("This is Critical!")
	log.Debug("This is Debug!")
	log.DebugF("Here are some numbers: %d %d %f", 10, -3, 3.14)
	// Give the Warning
	log.Warning("This is Warning!")
	log.WarningF("This is Warning!")
	// Show the error
	log.Error("This is Error!")
	log.ErrorF("This is Error!")
	// Notice
	log.Notice("This is Notice!")
	log.NoticeF("%s %s", "This", "is Notice!")
	// Show the info
	log.Info("This is Info!")
	log.InfoF("This is %s!", "Info")

	log.SetLogLevel(logger.ErrorLevel)

	log.SetFormat("[%{module}] [%{level}] %{message}")
	log.Warning("This is Warning!") // output: "[test] [WARNING] This is Warning!"
	// Also you can set your format as default format for all new loggers
	logger.SetDefaultFormat("%{message}")

}

func UploadFile(uri string, params map[string]string, paramName, path string) (*http.Response, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fi.Name())
	if err != nil {
		return nil, err
	}
	part.Write(fileContents)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	//req.Header.Set("Content-Type","multipart/form-data")
	client := &http.Client{}
	return client.Do(req)

}
