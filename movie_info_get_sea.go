package pintoto

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
)

type GetSeatResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		SeatData struct {
			Restrictions int            `json:"restrictions"`
			Seats        []GetSeatSeats `json:"seats"`
		} `json:"seatData"`
	} `json:"data"`
	Success bool `json:"success"`
}

type GetSeatSeats struct {
	Area       string `json:"area"`       // 本座位所在的区域，根据场次排期接口的 scheduleArea 字段， 可得到当前座位的分区价格
	ColumnNo   string `json:"columnNo"`   // 列
	Lovestatus int    `json:"lovestatus"` // 0为非情侣座；1为情侣座左；2为情侣座右
	RowNo      string `json:"rowNo"`      // 行
	SeatId     string `json:"seatId"`     // 座位标识符，锁座位和秒出票的时候需要用到
	SeatNo     string `json:"seatNo"`     // 座位名
	Status     string `json:"status"`     // N可售，LK不可售
}

type GetSeatResult struct {
	Result GetSeatResponse    // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newGetSeatResult(result GetSeatResponse, body []byte, http gorequest.Response) *GetSeatResult {
	return &GetSeatResult{Result: result, Body: body, Http: http}
}

// GetSeat 座位 https://www.showdoc.com.cn/1154868044931571/5866824368760475
func (c *Client) GetSeat(ctx context.Context, showId string, notMustParams ...gorequest.Params) (*GetSeatResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("showId", showId)
	// 请求
	request, err := c.request(ctx, apiUrl+"/movieapi/movie-info/get-seat", params)
	if err != nil {
		return newGetSeatResult(GetSeatResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response GetSeatResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newGetSeatResult(response, request.ResponseBody, request), err
}
