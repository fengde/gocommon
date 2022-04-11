package canalx

import (
	"fmt"
	"github.com/go-mysql-org/go-mysql/canal"
)

type Slave struct {
	c *canal.Canal
}

func newSlave(flavor string, serverID int64, host string, port int64, user, password string) (*Slave, error) {
	c, err := canal.NewCanal(&canal.Config{
		Flavor: flavor,
		ServerID: uint32(serverID),
		Addr: fmt.Sprintf("%s:%d", host, port),
		User: user,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return &Slave{
		c: c,
	}, nil
}

// NewSlave 创建mysql binlog slave
// user 需要授权可读binlog
// grant replication slave on *.* to user@'%' identified by "password";
// flush privileges;
func NewSlave(serverID int64, host string, port int64, user, password string) (*Slave, error) {
	return newSlave("mysql", serverID, host, port, user, password)
}

func NewMariadbSlave(serverID int64, host string, port int64, user, password string) (*Slave, error) {
	return newSlave("mariadb", serverID, host, port, user, password)
}

// Start 开始订阅binlog
// callback 自定义的处理函数，jsonData描述变动的内容，
// 如果action=insert, jsonData返回完整的插入数据[{"id": 1, "name": "fedel", "age": 19}]
// 如果action=update, jsonData返回完整的修改数据[{"id": 1, "age": 32}, {"id": 2, "age": 1}]
// 如果action=delete, jsonDaa返回的删除id [{"id": 1}, {"id": 2}, {"id": 3}]
func (p *Slave) Start(callback func(table string, action Action, jsonData string)) error {
	pos, err := p.c.GetMasterPos()
	if err != nil {
		return err
	}
	p.c.SetEventHandler(NewEventHandler(callback))
	return p.c.RunFrom(pos)
}

// Close 关闭订阅
func (p *Slave) Close() {
	p.c.Close()
}