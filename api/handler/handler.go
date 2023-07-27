package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func DoRequest(url string, method, auth string, body interface{}) ([]byte, error) {

	data, err := json.Marshal(&body)

	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	// request.Header.Add("Accept-Language", "uz")
	// request.Header.Add("Accept-Encoding", "gzip, deflate, br")
	// request.Header.Add("Accept", "application/json")
	// request.Header.Add("Connection", "keep-alive")
	// request.Header.Add("Content-Length", "165")
	// request.Header.Add("Content-Type", "application/json")

	request.Header.Add("Accept-Language", "uz")
	request.Header.Add("Content-Type", "application/json")

	if auth != "" {
		request.Header.Set("Authorization", auth)
	}

	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respByte, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return respByte, nil
}

func DoRequest2(tokin string, data string) string {
	url := "https://eticket.railway.uz/api/v1/orders/create"

	// Request payload
	payload := []byte(data)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	// Set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokin)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ""
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return ""
	}

	// Print the response body
	// fmt.Println(string(body))
	return string(body)
}

func GetOrder(tokin string, id string) string {
	// var url2 string
	url := id
	// for _, val := range url {
	// 	if val != '"' {
	// 		url2 =url2 + string(val)
	// 	}
	// }
	fmt.Println("id .....------->", url)

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	// Set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", tokin)

	// Send the GET request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ""
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return ""
	}

	// Print the response body
	fmt.Println(string(body))
	return string(body)
}

func GetListPay(url string, method, auth string, body interface{}) ([]byte, error) {

	data, err := json.Marshal(&body)

	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept-Language", "uz")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Device-Type", "BROWSER")

	if auth != "" {
		request.Header.Set("Authorization", auth)
	}

	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respByte, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return respByte, nil
}
