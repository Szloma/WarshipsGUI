package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

//zaimmplementować statystyke skuteczności strzałów

//------------------------------------------
//game properties

type GameProperties struct {
	Token        string
	Board        [20]string
	PlayerShoots []string
	Enemy        string
	Nick         string
	Description  string
}

var gameProperties GameProperties

//------------------------------------------
//utils

func AddIfNotPresent(slice []string, value string) []string {
	if !Contains(slice, value) {
		slice = append(slice, value)
	}
	return slice
}

func Contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func getCoords() (string, error) {
	fmt.Println("Podaj kordynaty")
	waitingLoop := true
	for ok := true; ok; ok = waitingLoop {
		userInput := strings.ToLower(getInput())
		if isValidFormat(userInput) {
			if !Contains(gameProperties.PlayerShoots, userInput) {
				return userInput, nil
				waitingLoop = false
			}

		} else {
			fmt.Println("Błędne koordynaty, spróbuj jeszcze raz")
		}
	}

	return "", nil
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">")
	input, _ := reader.ReadString('\n')

	cleanedInput := strings.ReplaceAll(input, "\n", "")
	cleanedInput = strings.ReplaceAll(cleanedInput, "\r", "")
	return cleanedInput
}

func getLastFromSlice(s interface{}) string {

	key := fmt.Sprintf("%s", s)
	result := stringToSlice(key)
	return result[len(result)-1]
}

func isValidFormat(input string) bool {
	validFormat := regexp.MustCompile(`^[a-j](?:10|[1-9])$`)
	return validFormat.MatchString(input)
}

func stringToSlice(inp string) []string {
	inp = strings.Replace(inp, "[", "", -1)
	inp = strings.Replace(inp, "]", "", -1)
	s := strings.Split(inp, " ")
	return s
}

//------------------------------------------
//Main game functions

func Board() ([]string, error) {

	Loop := true
	no_tries := 0
	for ok := true; ok; ok = Loop {
		client := &DefaultHTTPClient{}

		getHeaders := map[string]string{
			"X-Auth-Token": gameProperties.Token,
		}

		//////
		resp, err := client.Get("https://go-pjatk-server.fly.dev/api/game/board", getHeaders)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("unexpected status: %d, %s", resp.StatusCode, resp.Header.Get("message"))

			return nil, fmt.Errorf("unexpected status: %d, %s", resp.StatusCode, resp.Header.Get("message"))
		}
		if resp.StatusCode == http.StatusOK {
			Loop = false
		}
		var data map[string]interface{}

		err = json.Unmarshal([]byte(body), &data)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal response")
		}
		key := fmt.Sprintf("%s", data["board"])
		result := stringToSlice(key)
		if len(result) != 20 {
			fmt.Printf("Not enough pieces")
			return nil, fmt.Errorf("Not enough pieces")
		}
		fmt.Println("%d", len(result))

		if no_tries == 3 {
			Loop = false
			return nil, fmt.Errorf("Couldn't retrieve board")
		}
		no_tries += 1

		return result, nil

	}
	return nil, fmt.Errorf("Couldn't retrieve board")

}

func customBoard() ([20]string, error) {
	return [20]string{
		"A1",
		"A3",
		"B9",
		"C7",
		"D1",
		"D2",
		"D3",
		"D4",
		"D7",
		"E7",
		"F1",
		"F2",
		"F3",
		"F5",
		"G5",
		"G8",
		"G9",
		"I4",
		"J4",
		"J8",
	}, nil
}

func Fire(coord string) (string, error) {

	Loop := true
	no_tries := 0
	for ok := true; ok; ok = Loop {
		client := &DefaultHTTPClient{}

		GameBoard := map[string]interface{}{
			"coord": coord,
		}
		b, err := json.Marshal(GameBoard)
		if err != nil {
			fmt.Println(err)
			return "", fmt.Errorf("failed to marshall: ", err)
		}
		jsonBody := []byte(b)
		fmt.Println("body////")
		fmt.Printf((string(jsonBody)))

		getHeaders := map[string]string{
			"X-Auth-Token": gameProperties.Token,
		}

		postResponse, err := client.Post("https://go-pjatk-server.fly.dev/api/game/fire", "application/json", jsonBody, getHeaders)
		if err != nil {
			fmt.Println("POST request failed", err)
			fmt.Errorf("post request failed", err)
		}
		if postResponse.StatusCode != http.StatusOK {
			fmt.Printf("unexpected status: %d, %s", postResponse.StatusCode, postResponse.Header.Get("message"))
			fmt.Errorf("unexpected status: %d, %s", postResponse.StatusCode, postResponse.Header.Get("message"))
		}
		if postResponse.StatusCode == http.StatusOK {
			Loop = false
		}

		defer postResponse.Body.Close()
		fmt.Println("POST ResponseBody:")
		fmt.Println(string(postResponse.Header.Get("X-Auth-Token")))
		token := postResponse.Header.Get("X-Auth-Token")
		if len(token) == 0 {
			fmt.Errorf("cannot obtain token")
		}
		// Reading response body
		postResponseBody, err := io.ReadAll(postResponse.Body)
		if err != nil {
			// fmt.Println("Failed to read POST response body:", err)
			fmt.Errorf("Failed to read POST response body", err)
		}

		var data map[string]interface{}

		err = json.Unmarshal([]byte(postResponseBody), &data)
		if err != nil {
			return "", fmt.Errorf("failed to unmarshal response")
		}
		result := fmt.Sprintf("%s", data["result"])
		//fmt.Println(result)

		if no_tries == 3 {
			Loop = false
			return "", fmt.Errorf("couldn't perform fire request")
		}
		no_tries += 1

		return result, nil

	}
	return "", nil

}

func DeleteGame() (*StatusResponse, error) {
	Loop := true
	no_tries := 0
	for ok := true; ok; ok = Loop {

		client := &DefaultHTTPClient{}

		getHeaders := map[string]string{
			"X-Auth-Token": gameProperties.Token,
		}

		//////
		resp, err := client.Delete("https://go-pjatk-server.fly.dev/api/game/abandon", getHeaders)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var data map[string]interface{}

		err = json.Unmarshal([]byte(body), &data)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal response")
		}

		if no_tries == 3 {
			Loop = false
			return nil, fmt.Errorf("couldn't retrieve lobby data")
		}
		no_tries += 1

		if resp.StatusCode == 200 {
			Loop = false
		}
		return &StatusResponse{
			StatusCode: resp.StatusCode,
			Body:       data,
		}, nil

	}
	return nil, fmt.Errorf("couldn't retrieve lobby data")
}

func InitGame() error {

	Loop := true
	no_tries := 0
	for ok := true; ok; ok = Loop {

		client := &DefaultHTTPClient{}
		getHeaders := map[string]string{
			"accept": "application/json",
		}
		GameBoard := map[string]interface{}{
			"coords":      gameProperties.Board,
			"desc":        gameProperties.Description,
			"nick":        gameProperties.Nick,
			"target_nick": "",
			"wpbot":       true,
		}
		b, err := json.Marshal(GameBoard)
		if err != nil {
			fmt.Println(err)
			fmt.Errorf("failed to marshall: ", err)
		}
		jsonBody := []byte(b)

		postResponse, err := client.Post("https://go-pjatk-server.fly.dev/api/game", "application/json", jsonBody, getHeaders)
		if err != nil {
			fmt.Println("POST request failed", err)
			fmt.Errorf("post request failed", err)
		}
		if postResponse.StatusCode != http.StatusOK {
			fmt.Printf("unexpected status: %d, %s", postResponse.StatusCode, postResponse.Header.Get("message"))
			fmt.Errorf("unexpected status: %d, %s", postResponse.StatusCode, postResponse.Header.Get("message"))
		}
		if postResponse.StatusCode == http.StatusOK {
			Loop = false
		}

		defer postResponse.Body.Close()
		fmt.Println("POST ResponseBody:")
		fmt.Println(string(postResponse.Header.Get("X-Auth-Token")))
		token := postResponse.Header.Get("X-Auth-Token")
		if len(token) == 0 {
			fmt.Errorf("cannot obtain token")
		}
		// Reading response body
		postResponseBody, err := io.ReadAll(postResponse.Body)
		if err != nil {
			// fmt.Println("Failed to read POST response body:", err)
			fmt.Errorf("Failed to read POST response body", err)
		}
		fmt.Println("POST Response:")
		fmt.Println(string(postResponseBody))
		gameProperties.Token = token

		if no_tries == 3 {
			Loop = false
			return fmt.Errorf("couldn't initialize game")
		}
		no_tries += 1
	}
	return nil

}

func PlayerList() (*StatusResponse, error) {

	Loop := true
	no_tries := 0
	for ok := true; ok; ok = Loop {

		client := &DefaultHTTPClient{}

		getHeaders := map[string]string{}

		//////
		resp, err := client.Get("https://go-pjatk-server.fly.dev/api/lobby", getHeaders)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		if len(body) == 2 {
			Loop = false
			return nil, nil
		} else {
			var data map[string]interface{}

			err = json.Unmarshal([]byte(body), &data)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshall list of available players")
			}

			if no_tries == 3 {
				Loop = false
				return nil, fmt.Errorf("couldn't get list of available players")
			}

			if resp.StatusCode == http.StatusOK {
				Loop = false
			}
			no_tries += 1
			return &StatusResponse{
				StatusCode: resp.StatusCode,
				Body:       data,
			}, nil
		}

	}
	return nil, fmt.Errorf("couldn't get list of available players")
}

func getStats() (*StatusResponse, error) {

	Loop := true
	no_tries := 0
	for ok := true; ok; ok = Loop {

		client := &DefaultHTTPClient{}

		getHeaders := map[string]string{
			"X-Auth-Token": gameProperties.Token,
		}

		//////
		resp, err := client.Get("https://go-pjatk-server.fly.dev/api/stats", getHeaders)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var data map[string]interface{}

		err = json.Unmarshal([]byte(body), &data)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal response")
		}

		if no_tries == 3 {
			Loop = false
			return nil, fmt.Errorf("couldn't retrieve lobby data")
		}
		no_tries += 1

		if resp.StatusCode == 200 {
			Loop = false
		}
		return &StatusResponse{
			StatusCode: resp.StatusCode,
			Body:       data,
		}, nil

	}
	return nil, fmt.Errorf("couldn't retrieve lobby data")
}

func getLobby() (*StatusResponse, error) {

	Loop := true
	no_tries := 0
	for ok := true; ok; ok = Loop {

		client := &DefaultHTTPClient{}

		getHeaders := map[string]string{}

		//////
		resp, err := client.Get("https://go-pjatk-server.fly.dev/api/lobby", getHeaders)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var data map[string]interface{}

		err = json.Unmarshal([]byte(body), &data)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal response")
		}

		if no_tries == 3 {
			Loop = false
			return nil, fmt.Errorf("couldn't retrieve lobby data")
		}
		no_tries += 1

		if resp.StatusCode == 200 {
			Loop = false
		}
		return &StatusResponse{
			StatusCode: resp.StatusCode,
			Body:       data,
		}, nil

	}
	return nil, fmt.Errorf("couldn't retrieve lobby data")
}

func GameDescription() (*StatusResponse, error) {

	Loop := true
	no_tries := 0
	for ok := true; ok; ok = Loop {

		client := &DefaultHTTPClient{}

		getHeaders := map[string]string{
			"X-Auth-Token": gameProperties.Token,
		}

		//////
		resp, err := client.Get("https://go-pjatk-server.fly.dev/api/game/desc", getHeaders)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var data map[string]interface{}

		err = json.Unmarshal([]byte(body), &data)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal description response")
		}

		if no_tries == 3 {
			Loop = false
			return nil, fmt.Errorf("couldn't get game description")
		}

		if resp.StatusCode == http.StatusOK {
			Loop = false
		}
		no_tries += 1
		return &StatusResponse{
			StatusCode: resp.StatusCode,
			Body:       data,
		}, nil

	}
	return nil, fmt.Errorf("couldn't get game description")
}

func Status() (*StatusResponse, error) {

	Loop := true
	no_tries := 0
	for ok := true; ok; ok = Loop {

		client := &DefaultHTTPClient{}

		getHeaders := map[string]string{
			"X-Auth-Token": gameProperties.Token,
		}

		//////
		resp, err := client.Get("https://go-pjatk-server.fly.dev/api/game", getHeaders)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var data map[string]interface{}

		err = json.Unmarshal([]byte(body), &data)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal response")
		}

		if no_tries == 3 {
			Loop = false
			return nil, fmt.Errorf("couldn't initialize game")
		}

		if resp.StatusCode == http.StatusOK {
			Loop = false
		}
		no_tries += 1
		return &StatusResponse{
			StatusCode: resp.StatusCode,
			Body:       data,
		}, nil

	}
	return nil, fmt.Errorf("couldn't get status")

}

// ---------------------
// http client and methods
type DefaultHTTPClient struct{}
type StatusResponse struct {
	StatusCode int
	Body       map[string]interface{}
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func (c *DefaultHTTPClient) Get(url string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return http.DefaultClient.Do(req)
}

func (c *DefaultHTTPClient) Post(url string, bodyType string, body []byte, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", bodyType)
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return http.DefaultClient.Do(req)
}
func (c *DefaultHTTPClient) Delete(url string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return http.DefaultClient.Do(req)
}
