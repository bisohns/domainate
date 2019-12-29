// Stores and loads all godaddy domain configurations

package godaddy

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"regexp"
)

const apiURL string = "https://api.ote-godaddy.com/"

// Returns the API URL to use in the query if it is overwritten in env on conf files
func getApiUrl() string {
	if viper.IsSet("api_url") {
		return viper.GetString("api_url")
	}
	return apiURL
}

// GetApiUrl returns the URL to be used for requests. It includes the api version
func GetApiUrl(version string) (string, error) {
	if v, _ := regexp.MatchString("^v[0-9]$", version); v {
		return fmt.Sprintf("%s/%s", getApiUrl(), version), nil
	}
	return apiURL, errors.New("Invalid Version specified. verssion must be of the fomr `v<int>`")
}

func init() {
	viper.SetEnvPrefix("godaddy") // will be uppercased automatically
	viper.BindEnv("api_url")
	viper.BindEnv("sso_key")
}
