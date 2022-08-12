package pintoto

import (
	"context"
	"encoding/json"
	"go.dtapp.net/gorequest"
)

type GetVersionResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Success bool   `json:"success"`
}

type GetVersionResult struct {
	Result GetVersionResponse // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
	Err    error              // 错误
}

func newGetVersionResult(result GetVersionResponse, body []byte, http gorequest.Response, err error) *GetVersionResult {
	return &GetVersionResult{Result: result, Body: body, Http: http, Err: err}
}

// GetVersion 获取同步版本号 https://www.showdoc.com.cn/1154868044931571/6566701084841699
func (c *Client) GetVersion(ctx context.Context) *GetVersionResult {
	request, err := c.request(ctx, apiUrl+"/movieapi/movie-info/get-version", map[string]interface{}{})
	// 定义
	var response GetVersionResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newGetVersionResult(response, request.ResponseBody, request, err)
}
