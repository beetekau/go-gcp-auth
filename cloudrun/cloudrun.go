package RUN

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/beetekau/go-gcp-auth/gcp"
	"golang.org/x/oauth2/google"
)

func Get(URL string, results interface{}) error {
	u, err := url.Parse(URL)
	if err != nil {
		return err
	}
	ctx := context.Background()
	targetAudience := u.Scheme + "://" + u.Hostname()
	credentials, err := google.FindDefaultCredentials(ctx)
	if err != nil {
		fmt.Printf("cannot get credentials: %v", err)
		os.Exit(1)
	}

	jwtSource, err := gcp.JWTAccessTokenSourceFromJSON(credentials.JSON, targetAudience)
	if err != nil {
		return err
	}

	client := gcp.NewClient(jwtSource)
	res, err := client.Get(URL)
	if err != nil {
		fmt.Printf("cannot fetch result: %v", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("cannot read response: %v", err)
		os.Exit(1)
	}
	return json.Unmarshal(body, &results)
}
