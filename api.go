package main

import (
	"encoding/base64"
	"fmt"
	"gutil/api"
	"gutil/path"
	"gutil/rsa"
	"time"
)

type T struct {
	Timestamp int64
	Serial    string
}

func main() {
	t := T{
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

	email, _ := path.File("./email.html").Content()
	data := map[string]string{
		"Serial":  "386cd23f9dc568bcd88bb284f066d4f25372592c",
		"Token":   encryptedString,
		"Subject": "Test email with attachment",
		"Receipt": "reza@ies-italia.it",
		"Body":    email,
	}
	_ = data

	var res map[string]string

	api.New("http://192.168.1.175:8010/command/").Set(data).Attach("Attachment", "./test.pdf").Call().Scan(&res)

	fmt.Printf("%v", res)
}
