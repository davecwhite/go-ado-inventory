package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"
)

func apiURL(isGraph bool, organizationUrl string, endpoint string) string {

	base := "https://dev.azure.com/"
	if isGraph {
		base = "https://vssps.dev.azure.com/"
	}

	return base + organizationUrl + "/_apis/" + endpoint + "?api-version=7.2-preview"

}

func apiCall(name string, url string, authentication string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer "+authentication)

	secureClient := newSecureHTTPClient()

	resp, _ := secureClient.Do(req)

	defer resp.Body.Close()

	fmt.Println(name+" status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)

	responseBody := ""

	for i := 0; scanner.Scan() && i < 5; i++ {
		responseBody += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return responseBody, nil
}

func writeToFile(fileName string, data string) {
	file, err := os.Create(OUTPUT_DIRECTORY + "/" + fileName)
	if err != nil {
		println(err)
	}

	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		println(err)
	}
}

// Example of a more secure HTTP client configuration
func newSecureHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
			MaxIdleConns:    10,
			IdleConnTimeout: 120 * time.Second,
		},
	}
}
