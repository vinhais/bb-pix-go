package bb_pix_go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (bb *BB) CriarCob(pix Pix) (string, error) {
	token, err := bb.GetAccessToken()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/%s/%s?%s=%s", bb.Environment, "/cob", pix.TxID, DevAppKey, bb.DeveloperKey)

	// transformando a struct Pix em JSON
	data, err := json.Marshal(pix)
	if err != nil {
		return "", err
	}

	// criando a requisição PUT para fazer a cobrança
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		return "", ErrHttpRequestFoundation
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	return string(body), nil
}
