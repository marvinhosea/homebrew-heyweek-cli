package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	return &Client{
		httpClient: client,
	}
}

func (c *Client) Request(ctx context.Context, method string) error {
	endpoint := "https://appslab.auth0.com/oauth/device/code"
	payload := strings.NewReader("client_id=VdGeQJ2V6zca1vUbr6mkwZJViXYLiHL4&scope=%7Boffline_access%7D")

	newReq, err := http.NewRequestWithContext(ctx, method, endpoint, payload)
	if err != nil {
		return err
	}

	newReq.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := c.httpClient.Do(newReq)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(res)
	fmt.Println(string(body))

	return nil
}
