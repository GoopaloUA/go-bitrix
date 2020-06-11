package client

import "github.com/ikarpovich/go-bitrix/types"

func (c *Client) Batch(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("batch", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Scope(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("scope", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Methods(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("methods", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) _placements(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("_placements", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Events(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("events", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Profile(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("profile", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

