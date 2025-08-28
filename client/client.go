// Package client @author: Violet-Eva @date  : 2025/6/10 @notes :
package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/violet-eva-01/mingdaoApi/client/types"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client[t types.T] struct {
	method     string
	url        string
	proxy      *url.URL
	headers    map[string]string
	WSReqBody  *types.WorkSheetRequestBody
	WSRespBody []t
	isDebug    bool
	isGetExpr  bool
	expr       int
}

func NewClient[t types.T]() *Client[t] {
	return &Client[t]{}
}

func (c *Client[t]) SetMethod(method string) *Client[t] {
	c.method = method
	return c
}

func (c *Client[t]) GetMethod() string {
	return c.method
}

func (c *Client[t]) SetUrl(url string) *Client[t] {
	c.url = url
	return c
}

func (c *Client[t]) GetUrl() string {
	return c.url
}

func (c *Client[t]) SetProxy(proxy string) *Client[t] {
	if proxy == "" {
		c.proxy = nil
		return c
	}

	if parse, err := url.Parse(proxy); err != nil {
		fmt.Printf("build proxy failed, err is %s", err)
	} else {
		c.proxy = parse
	}

	return c
}

func (c *Client[t]) GetProxy() string {
	return c.proxy.String()
}

func (c *Client[t]) SetHeaders(headers map[string]string) *Client[t] {
	c.headers = headers
	return c
}

func (c *Client[t]) GetHeaders() map[string]string {
	return c.headers
}

func (c *Client[t]) SetRequestBody(body *types.WorkSheetRequestBody) *Client[t] {
	c.WSReqBody = body
	return c
}

func (c *Client[t]) GetRequestBody() *types.WorkSheetRequestBody {
	return c.WSReqBody
}

func (c *Client[t]) SetResponseBody(body []t) *Client[t] {
	c.WSRespBody = append(c.WSRespBody, body...)
	return c
}

func (c *Client[t]) GetResponseBody() []t {
	return c.WSRespBody
}

func (c *Client[t]) Debug() *Client[t] {
	c.isDebug = true
	return c
}

func (c *Client[t]) GetExpr(expr int) *Client[t] {
	c.isGetExpr = true
	c.expr = expr
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

	if req, err = http.NewRequest(c.method, c.url, reqBody); err != nil {
		return
	}

	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	var transport *http.Transport
	if c.proxy != nil {
		transport = &http.Transport{Proxy: http.ProxyURL(c.proxy)}
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
	if c.isGetExpr {
		if c.isDebug {
			fmt.Printf("[%s] start\n", time.Now().Local().Format(time.DateTime))
		}
		c.WSReqBody.PageSize = c.expr
		if _, err = c.getWorkSheetResponseBody(); err != nil {
			return
		}
	} else {
		for i := 1; true; i++ {
			var again bool
			if c.isDebug {
				fmt.Printf("[%s] start index [%d] request\n", time.Now().Local().Format(time.DateTime), i)
			}
			c.WSReqBody.SetPageIndex(i)
			if again, err = c.getWorkSheetResponseBody(); err != nil {
				return
			}
			if c.isDebug {
				fmt.Printf("[%s] finish index [%d] request ,get data [%d]\n", time.Now().Local().Format(time.DateTime), i, len(c.WSRespBody))
			}
			if !again {
				break
			}
		}
	}

	return
}

func (c *Client[t]) getWorkSheetResponseBody() (again bool, err error) {
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

	if len(respBody.Data.Rows) > 0 && len(respBody.Data.Rows)%c.WSReqBody.PageSize == 0 {
		again = true
	} else {
		again = false
	}
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
