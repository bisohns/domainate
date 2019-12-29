package godaddy

import (
	//  "github.com/spf13/cobra"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

// CheckRequestBody represents the retrieved response from checking if a domain exists
type CheckRequestBody struct {
	available  bool
	definitive bool
	domain     string
	period     int
	price      int64
}

// CheckAvailability checks if a Domain name is available
func CheckAvailability(domain string) (data CheckRequestBody, err error) {
	ssoKey := viper.GetString("ssoKey")
	apiURL, err := GetApiUrl("v1")
	if err != nil {
		return data, err
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return data, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("sso-key: %s", ssoKey))
	resp, err := client.Do(req)
	if err != nil {
		return data, err
	}
	if resp.StatusCode != 200 {
		return data, fmt.Errorf("%d: %s", resp.StatusCode, resp.Status)
	}
	fmt.Println(resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return
}
