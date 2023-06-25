package redisx

import "context"

type List struct {
	client   *Client
	listName string
}

func NewList(client *Client, listName string) *List {
	return &List{
		client:   client,
		listName: listName,
	}
}

// Push 往队列插入值
func (p *List) Push(ctx context.Context, value string) error {
	return p.client.Do(ctx, "LPUSH", p.listName, value).Err()
}

// Pop 从队列读取值
func (p *List) Pop(ctx context.Context) (string, error) {
	return p.client.Do(ctx, "RPOP", p.listName).Text()
}
