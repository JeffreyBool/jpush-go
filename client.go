package jpush

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/LyricTian/queue"
)

// NewClient 创建推送客户端实例
func NewClient(maxThread int, opts ...Option) *Client {
	o := defaultOptions
	for _, opt := range opts {
		opt(&o)
	}

	cli := &Client{
		opts:      &o,
		queue:     queue.NewListQueue(maxThread),
		cidClient: NewCIDClient(o.cidCount, opts...),
	}

	cli.jobPool = &sync.Pool{
		New: func() interface{} {
			return newPushJob(cli.opts, cli.queue, cli.cidClient)
		},
	}
	cli.queue.Run()

	return cli
}

// Client 推送客户端
type Client struct {
	opts      *options
	queue     queue.Queuer
	cidClient *CIDClient
	jobPool   *sync.Pool
}

// Terminate 终止客户端
func (c *Client) Terminate() {
	c.queue.Terminate()
}

// GetPushID 获取推送ID
func (c *Client) GetPushID(ctx context.Context) (string, error) {
	return c.cidClient.GetPushID(ctx)
}

// GetScheduleID 获取定时ID
func (c *Client) GetScheduleID(ctx context.Context) (string, error) {
	return c.cidClient.GetScheduleID(ctx)
}

// PushResult 推送响应结果
type PushResult struct {
	SendNO string `json:"sendno"`
	MsgID  string `json:"msg_id"`
}

func (r *PushResult) String() string {
	buf, _ := json.Marshal(r)
	return string(buf)
}

// PushResultHandle 异步响应结果
type PushResultHandle func(context.Context, *PushResult, error)
