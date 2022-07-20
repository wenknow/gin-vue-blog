package httpreq

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var client = &http.Client{Timeout: 5 * time.Second}

func PostForm(urlPath string, data map[string]string) (string, error) {

	postData := url.Values{}

	for k, v := range data {
		postData.Set(k, v)
	}

	resp, err := client.PostForm(urlPath, postData)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result), nil
}

func PostFormByHeader(urlPath string, data map[string]string, header map[string]string) (res map[string]string, err error) {

	postData := url.Values{}

	for k, v := range data {
		postData.Set(k, v)
	}
	bytesData := strings.NewReader(postData.Encode())
	req, _ := http.NewRequest("POST", urlPath, bytesData)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &res)
	return
}

func PostJson(urlPath string, data map[string]string) (res map[string]string, err error) {

	bytesData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", urlPath, bytes.NewReader(bytesData))

	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &res)
	return res, err
}

func PostJsonByHeader(urlPath string, data map[string]string, header map[string]string) (res map[string]string, err error) {

	bytesData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", urlPath, bytes.NewReader(bytesData))

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &res)
	return res, err
}

//发送get请求
func Get(urlPath string) (res map[string]string, err error) {

	resp, err := http.Get(urlPath)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &res)
	return res, err
}

//发送get请求 并把参数拼接到url上
func GetByParam(urlPath string, param map[string]string) (res map[string]string, err error) {

	params := url.Values{}
	Url, _ := url.Parse(urlPath)

	for k, v := range param {
		params.Set(k, v)
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath = Url.String()

	resp, err := http.Get(urlPath)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &res)
	return res, err
}

//发送get请求和请求头
func GetByHeader(urlPath string, header map[string]string) ([]byte, error) {

	req, _ := http.NewRequest("GET", urlPath, nil)

	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, err
}
