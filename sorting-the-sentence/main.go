package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://telesign-telesign-send-sms-verification-code-v1.p.rapidapi.com/sms-verification-code?verifyCode=1231232&phoneNumber=%2B998933433379&appName=optional"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("X-RapidAPI-Key", "b9697df1a9mshf5ebaf2cdd51281p11752djsn140388e4501b")
	req.Header.Add("X-RapidAPI-Host", "telesign-telesign-send-sms-verification-code-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
