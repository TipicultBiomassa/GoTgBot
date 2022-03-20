package telegram

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func main(host string, token string) Client {
	return Client{
		host:     "",
		basePath: newPath(token),
		client:   http.Client{},
	}
}

func newPath(token string) string {
	return "bot" + token
}

func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error){
	u:= url.URL{
		Scheme: "https",
		Host: c.host
		Path: path.join(c.basePath, method),
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Can't do request: %w", err)
	}
}

func (c *Client) SendMessage() {

}
