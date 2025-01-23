// SPDX-FileCopyrightText: 2024 NOI Techpark <digital@noi.bz.it>
//
// SPDX-License-Identifier: MPL-2.0

package utilities

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"
	"strconv"

	"github.com/hashicorp/go-retryablehttp"
	"golang.org/x/oauth2"
)


type RawFilterId struct {
	Items []struct {
		Id string `json:"Id"`
	} `json:"Items"`
}

func GetAccomodationIdByRawFilter(id string, url *url.URL) (string, error) {
	newurl,err := url.Parse(fmt.Sprintf(url.String(), id))
	if err != nil {
		return "", fmt.Errorf("could not parse url: %w", err)
	}

	client := retryablehttp.NewClient()
	req,err := retryablehttp.NewRequest("GET", newurl.String(), nil)
	if err != nil {
		return "", fmt.Errorf("could not create http request: %w", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error during http request: %w", err)
	}

	defer resp.Body.Close()
	
	var rawFilterId RawFilterId

	err = json.NewDecoder(resp.Body).Decode(&rawFilterId)
	if err != nil {
		return "", fmt.Errorf("could not decode response: %w", err)
	}

	if len(rawFilterId.Items) > 0 {
		return rawFilterId.Items[0].Id, nil
	}else{
		return "",nil
	}

}

func GetAccessToken(tokenURL, username, password, clientID, clientSecret string) (*oauth2.Token, error) {
    ctx := context.Background()

    config := &oauth2.Config{
        ClientID:     clientID,
        ClientSecret: clientSecret,
        Endpoint: oauth2.Endpoint{
            TokenURL: tokenURL,
        },
    }

    token, err := config.PasswordCredentialsToken(ctx, username, password)
    if err != nil {
        return nil, fmt.Errorf("failed to get token: %w", err)
    }

    return token, nil
}

func PutContentApi(url *url.URL, token string, payload interface{}, id string) (string,error) {
    jsonData, err := json.Marshal(payload)
    if err != nil {
		return "", fmt.Errorf("could not marshal payload: %w", err)
	}	

	u := fmt.Sprintf("%s/%s", url.String(), id)
	slog.Info("PUT URL", "url", u)
	newurl, err := url.Parse(u)
	if err != nil {
		return "", fmt.Errorf("could not parse url: %w", err)
	}

	req, err := retryablehttp.NewRequest("PUT", newurl.String(), bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("could not create http request: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	client := retryablehttp.NewClient()
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error during http request: %w", err)
	}

	return strconv.Itoa(resp.StatusCode), nil   
}
	
func PostContentApi(url *url.URL, token string, payload interface{}) (string,error) {

    jsonData, err := json.Marshal(payload)
    if err != nil {
		return "", fmt.Errorf("could not marshal payload: %w", err)
	}	
    u := url

	req, err := retryablehttp.NewRequest("POST", u.String(), bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("could not create http request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	client := retryablehttp.NewClient()
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error during http request: %w", err)
	}

	return strconv.Itoa(resp.StatusCode), nil
}