package main

import (
	"encoding/base64"
	"fmt"
	"github.com/iesreza/gutil/api"
	"github.com/iesreza/gutil/rsa"
	"time"
)

type T struct {
	Timestamp int64  `json:'timestamp'`
	Username  string `json:"username"`
}

func main() {
	t := T{
		Timestamp: time.Now().Unix(),
		Username:  "admin",
	}

	publicKeyString := `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAs2G3sB8yvQ+EsB69H7+k
6pctFXp4X8VBDg+y/uk588BX1lRDw+Y295PKzQ7qiOULtxcKfZhtPywGaDnJU3Je
iR/rSF5mdE8YV50JjkSDkBykuTZNE56sZdcbl3f2I745VUxQMxGG8AxjhBAllm4r
+ULgfi08zZUoPJ1yaSHrpQU2La3Az865cgoZ/sBdlo2P9OQbaSTP4uh7lztXkpdL
SNpEBT9Iw7nXfmMQ/pK2jrWqcaXIHuhEkxqA7E9A+LKq7uvdccVSAvVNHKfGx4PC
B2oGgLadXyywtNAk82n1XEz0zXhmi259p/ZIt4177wY/sYEzUn2ea3HQoRS+JyDb
9Dmn9C9ouadObybuXerfpV15WU2dne+MquIAA8DQHrqEcZJ3+7KY8OQlG0m3pFf2
pNlThsHjsWIRopQM2JYy4Auuqf3gsPUC5O7J/gciXKgd5VP49EeDhOnFsMEyHvEn
iWTgnVyHAncBeMmWypf7cJiMEoEcWs/12WF9hj15Ol0VsU4NJAEvbt+CYAluM0sA
+ERZOWtCMjDYGuIaTgFPbOHhHi6uXpBjt09rDhxVFE21WUzWMjLotaekvB/MaZ0P
s3s/NXqa8CEVRgqhCO0jBD4FpXBR35ggYnt4rfxspvzTjsYlNrtGVLkW3JeAraVg
0ghaybMj5HZLxOfysfPpHQUCAwEAAQ==
-----END PUBLIC KEY-----`
	publicKey := rsa.ParsePublicKey(publicKeyString)

	encrypted := publicKey.Encrypt(t)
	encryptedString := base64.StdEncoding.EncodeToString(encrypted)

	//email, _ := path.File("./email.html").Content()
	data := map[string]string{
		"serial":   "386cd23f9dc568bcd88bb284f066d4f25372592c",
		"token":    encryptedString,
		"amount":   "20",
		"username": "admin",
	}
	_ = data

	var res map[string]string

	x := api.New("http://192.168.1.175:8010/charge/").Set(data).Call()
	x.Scan(&res)

	fmt.Print(x)
}
