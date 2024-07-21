// Frabit - The next-generation database automatic operation platform
// Copyright © 2022-2024 Frabit Team
//
// Licensed under the GNU General Public License, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.gnu.org/licenses/gpl-3.0.txt
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package frabit

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-cleanhttp"
)

const Version = "2.0.19"
const UserAgent = "frabit-go-sdk/" + Version
const jsonMediaType = "application/json"

type ErrorCode string

const (
	ErrInternal          ErrorCode = "internal"
	ErrInvalid           ErrorCode = "invalid"
	ErrNotFound          ErrorCode = "not_found"
	ErrResponseMalformed ErrorCode = "response_malformed"
)

type Client struct {
	BaseURL   *url.URL
	client    *http.Client
	UserAgent string
	Token     string
	Headers   map[string]string

	// services used for communicate with the Frabit API
	Database DatabaseService
	Team     TeamService
	Agent    AgentService
}

type service struct {
	*Client
}

type ClientOption func(client *Client) error

func WithBaseURL(baseUrl string) ClientOption {
	return func(c *Client) error {
		ParseURL, err := url.Parse(baseUrl)
		if err != nil {
			return err
		}

		c.BaseURL = ParseURL
		return nil
	}
}

func WithToken(token string) ClientOption {
	return func(c *Client) error {
		c.Token = token
		return nil
	}
}

func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.UserAgent = fmt.Sprintf("%s %s", userAgent, c.UserAgent)
		return nil
	}
}

func WithRequestHeaders(headers map[string]string) ClientOption {
	return func(c *Client) error {
		for k, v := range headers {
			c.Headers[k] = v
		}
		return nil
	}
}

func NewClient(opts ...ClientOption) (*Client, error) {
	c := &Client{
		UserAgent: UserAgent,
		Headers:   make(map[string]string, 0),
		client:    cleanhttp.DefaultClient(),
	}

	// rewrite client config via opts
	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	c.Database = &databaseService{c}
	c.Team = &teamService{c}
	c.Agent = &agentService{c}

	return c, nil
}

func (c *Client) do(ctx context.Context, req *http.Request, body interface{}) error {
	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return c.handleResponse(ctx, resp, body)
}

func (c *Client) newRequest(method string, path string, body interface{}) (*http.Request, error) {
	addr, err := c.BaseURL.Parse(path)
	if err != nil {
		return nil, err
	}
	var req *http.Request
	switch method {
	case http.MethodGet:
		req, err = http.NewRequest(method, addr.String(), nil)
		if err != nil {
			return nil, err
		}
	default:
		buf := new(bytes.Buffer)
		if body != nil {
			err := json.NewDecoder(buf).Decode(body)
			if err != nil {
				return nil, err
			}
		}
		req, err = http.NewRequest(method, addr.String(), buf)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", jsonMediaType)
	}

	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

func (c *Client) handleResponse(ctx context.Context, resp *http.Response, body interface{}) error {
	out, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// check http status
	if resp.StatusCode >= 400 {
		return nil
	}

	if body == nil || resp.StatusCode == http.StatusNoContent {
		return nil
	}
	err = json.Unmarshal(out, body)
	if err != nil {
		var jsonErr *json.SyntaxError
		if errors.As(err, &jsonErr) {
			return &Error{
				msg:  "malformed response body received",
				Code: ErrResponseMalformed,
				Meta: map[string]string{
					"body":        string(out),
					"http_status": http.StatusText(resp.StatusCode),
				},
			}
		}
		return err
	}
	return nil
}

type Error struct {
	msg  string
	Code ErrorCode
	Meta map[string]string
}

func (e Error) Error() string { return e.msg }
