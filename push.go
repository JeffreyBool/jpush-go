package jpush

import (
	"context"
	"net/http"
)

// Push 消息推送
func (c *Client) Push(ctx context.Context, payload *Payload, callback PushResultHandle) error {
	job := c.jobPool.Get().(*pushJob)
	job.Reset(ctx, payload, callback)
	c.queue.Push(job)
	return nil
}

// PushValidate 先校验，再推送
func (c *Client) PushValidate(ctx context.Context, payload *Payload, callback PushResultHandle) error {
	resp, err := pushRequest(ctx, c.opts, "/v3/push/validate", http.MethodPost, payload.Reader())
	if err != nil {
		return err
	}
	defer resp.Close()

	return c.Push(ctx, payload, callback)
}
