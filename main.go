package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

type APIResponse struct {
	Drop []string `json:"drop"`
	Pull []string `json:"pull"`
}

func oldApiResponse(appId string) (*APIResponse, error) {

	url := "https://server.apxor.com/v2/sync/" + appId + "/configs/validate?platform=android&userId=76def15fc3f353d9&actionType=rta&version=218"

	postedData := map[string]interface{}{
		"pid_mid": "",
	}

	jsonPostData, err := json.Marshal(postedData)
	if err != nil {
		fmt.Println("error occured while converting the go map into json byte slice")
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPostData))
	if err != nil {
		fmt.Println("Error creating request", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error occured making POST request", err)
		return nil, err
	}

	defer resp.Body.Close()

	if err != nil {
		println("error occured while getting the old api response")
		return nil, err
	}

	var goStructResp APIResponse

	err = json.NewDecoder(resp.Body).Decode(&goStructResp)

	if err != nil {
		println("error occured while parsing the json repsonse to go understandable struct")
		return nil, err
	}

	return &goStructResp, nil

}

func newApiResponse(appId string) (*APIResponse, error) {

	url := "http://serverg.apxor.com/v3/sync/" + appId + "/configs/validate?platform=android&userId=afe6988b-3d35-4c8c-939c-60809a93f642&actionType=rta&version=218&customerId=1234"

	postedData := map[string]interface{}{
		"pid_mid":           "",
		"installation_date": 1696425885461,
		"session":           2,
	}

	jsonPostData, err := json.Marshal(postedData)
	if err != nil {
		fmt.Println("error occured while converting the go map into json byte slice")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPostData))
	if err != nil {
		fmt.Println("Error creating request", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error occured making POST request", err)
		return nil, err
	}

	defer resp.Body.Close()

	if err != nil {
		println("error occured while getting the old api response")
		return nil, err
	}

	var goStructResp APIResponse

	err = json.NewDecoder(resp.Body).Decode(&goStructResp)

	if err != nil {
		println("error occured while parsing the json repsonse to go understandable struct")
		return nil, err
	}

	return &goStructResp, nil

}

func main() {

	resp1, err := oldApiResponse("01762e21-7b69-4f4b-8f65-6abfee40e8a2")

	if err != nil {
		fmt.Println(err)
		return
	}

	resp2, err := newApiResponse("8d199ec0-9148-4ec8-bd8c-fcce3041f882")
	if err != nil {
		fmt.Println(err)
		return
	}

	if reflect.DeepEqual(resp1, resp2) {
		fmt.Println("Old and new api responses are equal")
		return
	} else {
		fmt.Println("Old and new api responses are not equal")
		return

	}
}
