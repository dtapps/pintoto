package pintoto

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
)

type GetCityListResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		List []GetCityListResponseDataList `json:"list"`
	} `json:"data"`
	Success bool `json:"success"`
}

type GetCityListResponseDataList struct {
	PinYin     string `json:"pinYin"`     // 城市首字母
	RegionName string `json:"regionName"` // 城市名
	CityId     int    `json:"cityId"`     // 城市id
}

type GetCityListResult struct {
	Result GetCityListResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newGetCityListResult(result GetCityListResponse, body []byte, http gorequest.Response) *GetCityListResult {
	return &GetCityListResult{Result: result, Body: body, Http: http}
}

// GetCityList 城市列表
// https://www.showdoc.com.cn/1154868044931571/5865562425538244
func (c *Client) GetCityList(ctx context.Context, notMustParams ...gorequest.Params) (*GetCityListResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/movieapi/movie-info/get-city-list", params)
	if err != nil {
		return newGetCityListResult(GetCityListResponse{}, request.ResponseBody, request), err
	}
	var response GetCityListResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newGetCityListResult(response, request.ResponseBody, request), err
}
