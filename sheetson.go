package sheetson

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type (
	Client interface {
		AddRow(spreadSheetID string, sheetName string, row map[string]interface{}) error
	}

	clientImpl struct {
		apiKey string
	}
)

// NewClient returns a new Client.
func NewClient(apiKey string) Client {
	return &clientImpl{
		apiKey: apiKey,
	}
}

// AddRow adds a row to a sheet.
func (c *clientImpl) AddRow(spreadSheetID string, sheetName string, row map[string]interface{}) error {
	payloadBytes, err := json.Marshal(row)
	if err != nil {
		println(err.Error())
		return err
	}
	body := bytes.NewReader(payloadBytes)

	url := "https://api.sheetson.com/v2/sheets/" + sheetName
	auth := "Bearer " + c.apiKey

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		println(err.Error())
		return err
	}
	req.Header.Set("Authorization", auth)
	req.Header.Set("X-Spreadsheet-Id", spreadSheetID)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		println(err.Error())
		return err
	}
	defer resp.Body.Close()

	// Successful status code https://docs.sheetson.com/#successful-requests
	// Error status code https://docs.sheetson.com/status-codes/
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New(resp.Status)
	}

	return nil
}
