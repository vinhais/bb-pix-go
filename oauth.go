package bb_pix_go

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type AuthScope string

const (
	ScopePixWrite  AuthScope = "pix.write"
	ScopePixRead   AuthScope = "pix.read"
	ScopeCobWrite  AuthScope = "cob.write"
	ScopeCobRead   AuthScope = "cob.read"
	ScopeCobvWrite AuthScope = "cobv.write"
	ScopeCobvRead  AuthScope = "cobv.read"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func (bb *BB) GetAccessToken(scope AuthScope) (*AccessToken, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Add("scope", string(scope))

	body := bytes.NewBufferString(data.Encode())

	// criando requisição
	req, err := http.NewRequest(http.MethodPost, OauthHomologacaoUrl, body)

	req.Header.Set("Authorization", bb.Basic)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode <= 599 {
		return nil, ErrOauthFailed
	}

	var token AccessToken
	err = json.Unmarshal(res, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
