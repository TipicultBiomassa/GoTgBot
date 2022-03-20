package telegram

import (
	"net/http"
	"net/url"
	"path"
	"strconv"

	"lib/e"
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

func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
	defer func() { err = e.WrapIfErr(errMsg, err) }()
	const errMsg = "Can't do request"
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)

	defer func() { _ = resp.Body.Close() }
}

func (c *Client) SendMessage() {

}
