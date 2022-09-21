package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	url := "https://api.pinata.cloud/pinning/pinFileToIPFS"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open("./1.png")
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("file", filepath.Base("./1.png"))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}
	_ = writer.WriteField("pinataOptions", "{\"cidVersion\": 1}")
	_ = writer.WriteField("pinataMetadata", "{\"name\": \"MyFile\", \"keyvalues\": {\"company\": \"Pinata\"}}")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySW5mb3JtYXRpb24iOnsiaWQiOiIzMTY1ZmI0Zi1mYTAzLTRmNGYtYmU0YS1hZGFkZGVhYzEyYjYiLCJlbWFpbCI6Imxvbmd4Ym95aGlAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsInBpbl9wb2xpY3kiOnsicmVnaW9ucyI6W3siaWQiOiJGUkExIiwiZGVzaXJlZFJlcGxpY2F0aW9uQ291bnQiOjF9LHsiaWQiOiJOWUMxIiwiZGVzaXJlZFJlcGxpY2F0aW9uQ291bnQiOjF9XSwidmVyc2lvbiI6MX0sIm1mYV9lbmFibGVkIjpmYWxzZSwic3RhdHVzIjoiQUNUSVZFIn0sImF1dGhlbnRpY2F0aW9uVHlwZSI6InNjb3BlZEtleSIsInNjb3BlZEtleUtleSI6ImJhOTMxYWY0ZWI5YjA2M2M3M2YxIiwic2NvcGVkS2V5U2VjcmV0IjoiZDNjZjEwNGNjNDg0Nzg5OGI2NWJiYzdkYTFkOTVjZjM4MjdiMDg1OGNkMTIwOWQ0YWFmMTYzYWZjZGJhNjRlMCIsImlhdCI6MTY2MzY3OTk3MH0.laMIpj0y5jgyfM4T_Qhpv71ea98ivs4jRSHkJteq3n8")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
