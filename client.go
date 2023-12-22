package pintoto

import (
	"go.dtapp.net/golog"
	"math"
	"strconv"
)

// ClientConfig 实例配置
type ClientConfig struct {
	AppKey    string
	AppSecret string
}

// Client 实例
type Client struct {
	config struct {
		appKey    string
		appSecret string
	}
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.appKey = config.AppKey
	c.config.appSecret = config.AppSecret

	return c, nil
}

func (c *Client) GradeToFloat64(i interface{}) float64 {
	switch v := i.(type) {
	case string:
		float, _ := strconv.ParseFloat(v, 64)
		return float
	case float64:
		return v
	case int64:
		return float64(v) / math.Pow10(0)
	default:
		return 0
	}
}
