package podiogo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	authInfo      AuthMessage
	authenticated = false
	httpClient    *http.Client
)

func init() {
	httpClient = &http.Client{}
}

func Authenticate(appId int, grantType, appToken, clientId, redirectUri, clientSecret string) bool {
	requestTemplate := "https://podio.com/oauth/token?grant_type=%v&app_id=%v&app_token=%v&client_id=%v&redirect_uri=%v&client_secret=%v"
	request := fmt.Sprintf(requestTemplate, grantType, appId, appToken, clientId, redirectUri, clientSecret)

	req, err := http.NewRequest("POST", request, nil)
	if err != nil {
		fmt.Println(err)
		return false
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if resp.StatusCode != 200 {
		fmt.Println("Malformed request")
		return false
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var authResponse AuthMessage
	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		fmt.Println(err)
		return false
	}

	authInfo = authResponse
	authenticated = true
	return true
}
