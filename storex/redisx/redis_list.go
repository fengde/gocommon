package redisx

type List struct {
	client *Client
	listName string
}

func NewList(client *Client, listName string) *List {
	return &List{
		client: client,
		listName: listName,
	}
}

// Push 往队列插入值
func (p *List) Push(value string) error {
	return p.client.Do("LPUSH", p.listName, value).Err()
}

// Pop 从队列读取值
func (p *List) Pop() (string, error) {
	return p.client.Do("RPOP", p.listName).Text()
}