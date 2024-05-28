package funcs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetAll(url string, token string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "Buku didapatkan", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Success", err
	}

	result := string(body)

	return result, nil
}

func GetBukuByID(url string, id string, token string) (string, error) {
	URL := fmt.Sprintf(url + id)
	fmt.Println(URL)

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	hasil := string(data)

	return hasil, nil
}

func CreateBuku(url string, body map[string]interface{}, token string) error {
	payload, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer resp.Body.Close()

	var hasil map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&hasil); err != nil {
		// fmt.Println("Error decoding JSON:", err)
		// fmt.Println("Response:", hasil)
		return nil
	}
	
	return nil
}

func DeleteBuku(url string, id string, token string) error {
	URL := fmt.Sprintf(url + id)

	req, err := http.NewRequest("DELETE", URL, nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	return nil
}

func UpdateBuku(url string, id string, body map[string]interface{}, token string) error {
	Payload, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error Encoding JSON: ", err)
	}

	URL := fmt.Sprintf(url + id)
	req, err := http.NewRequest("PUT", URL, bytes.NewBuffer(Payload))
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer res.Body.Close()

	return nil
}
