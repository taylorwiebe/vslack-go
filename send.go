package vslack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (v *VSlack) send() error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", v.IncomingWebhookURI, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	client := &http.Client{}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Failed to send slack message with response code %d", resp.StatusCode)
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}
