// Package client @author: Violet-Eva @date  : 2025/6/10 @notes :
package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/violet-eva-01/mingdaoApi/client/types"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Client[t types.T] struct {
	Method     string
	Url        string
	Proxy      *url.URL
	Size       int
	Headers    map[string]string
	WSReqBody  *types.WorkSheetRequestBody
	WSRespBody []t
	IsDebug    bool
}

func NewClient[t types.T]() *Client[t] {
	return &Client[t]{}
}

func (c *Client[t]) SetMethod(method string) *Client[t] {
	c.Method = method
	return c
}

func (c *Client[t]) SetUrl(url string) *Client[t] {
	c.Url = url
	return c
}

func (c *Client[t]) SetProxy(proxy string) *Client[t] {
	if proxy == "" {
		c.Proxy = nil
		return c
	}

	if parse, err := url.Parse(proxy); err != nil {
		log.Fatal(err)
	} else {
		c.Proxy = parse
	}

	return c
}

func (c *Client[t]) SetSize(size int) *Client[t] {
	if size != 0 {
		c.Size = size
	}
	return c
}

func (c *Client[t]) SetHeaders(headers map[string]string) *Client[t] {
	c.Headers = headers
	return c
}

func (c *Client[t]) SetRequestBody(body *types.WorkSheetRequestBody) *Client[t] {
	c.WSReqBody = body
	return c
}

func (c *Client[t]) SetResponseBody(body []t) *Client[t] {
	c.WSRespBody = append(c.WSRespBody, body...)
	return c
}

func (c *Client[t]) Request(reqBody *bytes.Buffer) (respBody []byte, err error) {
	defer func() {
		if err != nil {
			err = errors.New(fmt.Sprintf("Request error: %s", err))
		}
	}()
	var (
		req  *http.Request
		resp *http.Response
	)

	if req, err = http.NewRequest(c.Method, c.Url, reqBody); err != nil {
		return
	}

	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}

	var transport *http.Transport
	if c.Proxy != nil {
		transport = &http.Transport{Proxy: http.ProxyURL(c.Proxy)}
	} else {
		transport = &http.Transport{}
	}

	if resp, err = (&http.Client{Transport: transport}).Do(req); err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return
	}

	if respBody, err = io.ReadAll(resp.Body); err != nil {
		return
	}
	return
}

func (c *Client[t]) WorkSheetRequest() (err error) {
	if c.Size != 0 {
		c.WSReqBody.SetPageSize(c.Size)
		for i := 1; true; i++ {
			if c.IsDebug {
				fmt.Printf("[%s] start index [%d] request\n", time.Now().Local().Format(time.DateTime), i)
			}
			c.WSReqBody.SetPageIndex(i)
			if err = c.getWorkSheetResponseBody(); err != nil {
				return
			}
			if c.IsDebug {
				fmt.Printf("[%s] finish index [%d] request ,get data [%d]\n", time.Now().Local().Format(time.DateTime), i, len(c.WSRespBody))
			}
			if len(c.WSRespBody)%c.Size != 0 {
				break
			}
		}
	} else {
		if c.IsDebug {
			fmt.Printf("[%s] start\n", time.Now().Local().Format(time.DateTime))
		}
		if err = c.getWorkSheetResponseBody(); err != nil {
			return
		}
	}

	return
}

func (c *Client[t]) getWorkSheetResponseBody() (err error) {
	defer func() {
		if err != nil {
			err = errors.New(fmt.Sprintf("getWorkSheetResponseBody error: %s", err))
		}
	}()

	var (
		reqBody  *bytes.Buffer
		respData []byte
	)
	if c.WSReqBody.WorksheetId != "" {
		var marshal []byte
		if marshal, err = json.Marshal(c.WSReqBody); err != nil {
			return
		}
		reqBody = bytes.NewBuffer(marshal)
	} else {
		reqBody = nil
	}

	if respData, err = c.Request(reqBody); err != nil {
		return
	}

	var respBody types.WorkSheetResponseBody[t]
	if err = json.Unmarshal(respData, &respBody); err != nil {
		return
	}

	c.SetResponseBody(respBody.Data.Rows)
	return
}

func (c *Client[t]) WebhookRequest(data ...t) (respBody []byte, err error) {
	var marshal []byte
	if marshal, err = json.Marshal(data); err != nil {
		return
	}
	reqData := bytes.NewBuffer(marshal)
	respBody, err = c.Request(reqData)
	return
}
