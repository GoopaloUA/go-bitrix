package client

import (
	"crypto/tls"
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/nightwriter/go-bitrix/types"
	"github.com/pkg/errors"
	"gopkg.in/resty.v1"
)

type Client struct {
	client      *resty.Client
	oAuth       *OAuthData
	webhookAuth *WebhookAuthData
	Url         *url.URL
}

type OAuthData struct {
	AuthToken    string `valid:"alphanum,required"`
	RefreshToken string `valid:"alphanum,required"`
}

type WebhookAuthData struct {
	UserID int    `valid:"required"`
	Secret string `valid:"alphanum,required"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewClientWithOAuth(intranetUrl, authToken, refreshToken string) (*Client, error) {
	u, err := url.Parse(intranetUrl)
	if err != nil {
		return nil, errors.Wrap(err, "Error parsing B24 URL")
	}

	auth := &OAuthData{
		AuthToken:    authToken,
		RefreshToken: refreshToken,
	}

	_, err = govalidator.ValidateStruct(auth)
	if err != nil {
		return nil, errors.Wrap(err, "Auth params validation failed")
	}

	return &Client{
		client: resty.DefaultClient,
		Url:    u,
		oAuth:  auth,
	}, nil
}

func NewClientWithWebhookAuth(intranetUrl string, userId int, secret string) (*Client, error) {
	u, err := url.Parse(intranetUrl)
	if err != nil {
		return nil, errors.Wrap(err, "Error parsing B24 URL")
	}

	auth := &WebhookAuthData{
		UserID: userId,
		Secret: secret,
	}

	_, err = govalidator.ValidateStruct(auth)
	if err != nil {
		return nil, errors.Wrap(err, "Auth params validation failed")
	}

	return &Client{
		client:      resty.DefaultClient,
		Url:         u,
		webhookAuth: auth,
	}, nil
}

func NewEnvClientWithOauth() (*Client, error) {
	return NewClientWithOAuth(
		os.Getenv("BITRIX_URL"),
		os.Getenv("BITRIX_AUTH_TOKEN"),
		os.Getenv("BITRIX_REFRESH_TOKEN"))
}

func NewEnvClientWithWebhookAuth() (*Client, error) {

	userId, err := strconv.Atoi(os.Getenv("BITRIX_WEBHOOK_USER"))

	if err != nil {
		return nil, errors.Wrap(err, "Incorrect User ID")
	}

	return NewClientWithWebhookAuth(
		os.Getenv("BITRIX_URL"),
		userId,
		os.Getenv("BITRIX_WEBHOOK_SECRET"))
}

func (c *Client) SetInsecureSSL(v bool) {
	resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: v})
}

func (c *Client) SetDebug(v bool) {
	resty.SetDebug(v)
}

func (c *Client) DoRaw(method string, reqData interface{}, respData interface{}) (*resty.Response, error) {
	resty.SetHostURL(c.Url.String())
	//	resty.SetHeader("Accept", "application/json") // commented because of causing "fatal error: concurrent map writes" with goroutines
	req := resty.R()

	var endpoint string
	if c.webhookAuth != nil {
		endpoint = fmt.Sprintf("/rest/%d/%s/%s", c.webhookAuth.UserID, c.webhookAuth.Secret, method)
	} else {
		endpoint = fmt.Sprintf("/rest/%s", method)

		params := map[string]string{
			"auth": c.oAuth.AuthToken,
		}

		req.SetQueryParams(params)
	}

	if respData != nil {
		req.SetResult(respData)
	}

	req.SetError(&types.ResponseError{})

	// values, err := query.Values(reqData)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "Error encoding form")
	// }

	resp, err := req.
		SetBody(reqData).
		Post(endpoint)

	if err != nil {
		return nil, errors.Wrap(err, "Error posting data")
	}

	if resp.IsError() {
		error := resp.Error().(*types.ResponseError)
		return resp, errors.New(fmt.Sprintf("REST method error (%s): %s", error.Code, error.Description))
	}

	return resp, err
}

func (c *Client) Do(method string, reqData interface{}, respData interface{}) (interface{}, error) {
	resp, err := c.DoRaw(method, reqData, respData)
	return resp.Result(), err
}

func (c *Client) PaginationData(ml map[string]string, reqData interface{}, respData interface{}) (*resty.Response, error) {
	Portal := os.Getenv("BITRIX_URL")
	WebhookAuthSecret := os.Getenv("BITRIX_WEBHOOK_SECRET")
	WebhookAuthUserID, _ := strconv.Atoi(os.Getenv("BITRIX_WEBHOOK_USER"))
	Method := fmt.Sprintf("batch.json?halt:%d", 0)
	for i := 0; i < len(ml)/2; i++ {
		drn := fmt.Sprintf("data_rec_%d", i)
		pr := fmt.Sprintf("parametr_%d", i)
		buff := fmt.Sprintf("&cmd[%s]=%s%s", drn, ml[drn], ml[pr])
		Method = fmt.Sprintf(Method + buff)
	}
	webhook := fmt.Sprintf("%s/rest/%d/%s/%s", Portal, WebhookAuthUserID, WebhookAuthSecret, Method)

	req := resty.R()

	if respData != nil {
		req.SetResult(respData)
	}
	req.SetError(&types.ResponseError{})

	resp, err := req.
		SetBody(reqData).
		Post(webhook)

	if err != nil {
		return nil, errors.Wrap(err, "Error posting data")
	}

	if resp.IsError() {
		error := resp.Error().(*types.ResponseError)
		return resp, errors.New(fmt.Sprintf("REST method error (%s): %s", error.Code, error.Description))
	}

	return resp, err
}
