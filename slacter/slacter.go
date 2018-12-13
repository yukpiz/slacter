package slacter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Slacter struct {
	config *Config
}

type Config struct {
	Token    string
	Channel  string
	UserName string
	IconURL  string
}

type SlackResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

func New(config *Config) *Slacter {
	return &Slacter{config: config}
}

func (s *Slacter) Write(p []byte) (int, error) {
	err := s.postSlack(p)
	if err != nil {
		fmt.Printf("Fatal: %s\n", err)
	}
	return len(p), err
}

func (s *Slacter) postSlack(p []byte) error {
	req, err := http.NewRequest(http.MethodPost, "https://slack.com/api/chat.postMessage", nil)
	if err != nil {
		return err
	}

	v := url.Values{}
	v.Add("token", s.config.Token)
	v.Add("channel", s.config.Channel)
	v.Add("icon_url", s.config.IconURL)
	v.Add("username", s.config.UserName)
	v.Add("link_names", "true")
	v.Add("text", string(p))

	req.Header.Set("Content-Type", "application/json")
	req.URL.RawQuery = v.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf(string(body))
	}

	sres := &SlackResponse{}
	err = json.Unmarshal(body, sres)
	if err != nil {
		return err
	}

	if !sres.Ok {
		return fmt.Errorf(sres.Error)
	}
	return nil
}
