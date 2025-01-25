package test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var baseURL = "http://localhost:3000/api/v1"

func CallAPI(t *testing.T, url string, method string, body []byte, token *string) (*[]byte, int, error) {
	// Create HTTP Request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")
	if token != nil {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *token))
	}

	// Call API
	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)

	// Read Response Body
	bodyResponse, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	assert.NoError(t, err)

	return &bodyResponse, resp.StatusCode, nil
}
