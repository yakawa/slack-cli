package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	API_END_POINT = "https://slack.com/api/"
)

type postTokenStruct struct {
	Token string `json:"token"`
}

type postMessageStruct struct {
	Token     string `json:"token"`
	Channel   string `json:"channel"`
	Text      string `json:"text"`
	AsUser    bool   `json:"as_user"`
	LinkNames bool   `json:"link_names"`
}

func post(req *http.Request) (map[string]interface{}, error) {
	c := new(http.Client)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var jsonBody map[string]interface{}
	err = json.Unmarshal(body, &jsonBody)
	return jsonBody, nil
}

func makeJsonRequest(path string, v interface{}) (*http.Request, error) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
	}

	ep := API_END_POINT + path

	req, _ := http.NewRequest("POST", ep, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Authorization", "Bearer "+viper.GetString("token"))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return req, nil
}

//
func PostMessage(ch string, msg string) error {
	pm := postMessageStruct{
		Channel:   ch,
		Text:      msg,
		AsUser:    true,
		LinkNames: true,
	}

	req, _ := makeJsonRequest("chat.postMessage", pm)
	_, _ = post(req)

	return nil
}

//
func UploadFile(ch string, comment string, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	params := map[string]string{
		"token":    viper.GetString("token"),
		"channels": ch,
		"filename": filepath.Base(fileName),
	}
	if comment != "" {
		params["initial_comment"] = comment
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(fileName))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()

	ep := API_END_POINT + "files.upload"

	req, err := http.NewRequest("POST", ep, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	_, _ = post(req)
	return nil
}

func CheckAuth() error {
	pm := postTokenStruct{
		Token: viper.GetString("token"),
	}

	req, _ := makeJsonRequest("auth.test", pm)
	rslt, _ := post(req)

	r := rslt["ok"].(bool)
	if r == true {
		fmt.Println("Auth: ok")
	} else {
		fmt.Println("Auth: ng")
	}
	return nil
}
