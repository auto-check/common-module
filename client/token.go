package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"runtime/debug"
)

const oauthTokenURL = "https://developer-api.dsmkr.com/dsmauth/token"

type oauthTokenRequest struct {
	ClientID string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code string	`json:"code"`
}

type oauthTokenResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GetOauthAccessToken(code string) (string, error){
	var reqByte []byte
	reqBody := bytes.NewBuffer(reqByte)
	json.NewEncoder(reqBody).Encode(oauthTokenRequest{
		ClientID: os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Code: code,
	})

	resp, err := http.Post(oauthTokenURL, "application/json", reqBody)
	if err != nil {
		log.Errorf(fmt.Sprintf("Error %s\n%s", err, debug.Stack()))
		return "", err
	}
	defer resp.Body.Close()

	var tokenRespBody oauthTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&tokenRespBody)
	if err != nil && err != io.EOF {
		log.Errorf(fmt.Sprintf("Error %s\n%s", err, debug.Stack()))
		return "", err
	}

	return  tokenRespBody.AccessToken, nil
}
