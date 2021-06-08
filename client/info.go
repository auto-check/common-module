package client

import (
	"auth/model"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"runtime/debug"
)

const oauthInfoBasicURL = "https://developer-api.dsmkr.com/v1/info/basic"

type oauthInfoBasicRequest struct {}
type oauthInfoBasicResponse struct {
	Name string `json"name"`
	Gcn string `json:"gcn"`
	Email string `json:email`
}

func GetOauthInfo(oauthToken string) (*model.Student, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", oauthInfoBasicURL, nil)
	if err != nil {
		log.Errorf(fmt.Sprintf("Error %s\n%s", err, debug.Stack()))
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+oauthToken)
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf(fmt.Sprintf("Error %s\n%s", err, debug.Stack()))
		return nil, err
	}
	defer resp.Body.Close()

	var infoRespBody oauthInfoBasicResponse
	err = json.NewDecoder(resp.Body).Decode(&infoRespBody)
	if err != nil && err != io.EOF {
		log.Errorf("Error %s\n%s", err, debug.Stack())
		return nil, err
	}

	s := &model.Student{
		Name: infoRespBody.Name,
		Email: infoRespBody.Email,
		Gcn: infoRespBody.Gcn,
	}
	return s, nil
}