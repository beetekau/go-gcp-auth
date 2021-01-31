package RUN

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/beetekau/go-gcp-auth/gcp"
	"golang.org/x/oauth2/google"
)

//Get make cloud run secure request based on default credentials
func Get(URL string, results interface{}) error {
	u, err := url.Parse(URL)
	if err != nil {
		return err
	}
	ctx := context.Background()
	targetAudience := u.Scheme + "://" + u.Hostname()
	key := os.Getenv("SERVICE_KEY_FILE")
	if key == "" {
		return errors.New("invalid SERVICE_KEY_FILE")
	}
	keyValue, err := ioutil.ReadFile(key)
	if err != nil {
		return err
	}
	credentials, err := google.CredentialsFromJSON(ctx, keyValue)
	if err != nil {
		return err
	}

	jwtSource, err := gcp.JWTAccessTokenSourceFromJSON(credentials.JSON, targetAudience)
	if err != nil {
		return err
	}

	client := gcp.NewClient(jwtSource)
	res, err := client.Get(URL)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &results)
}
