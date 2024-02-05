package canalx

import (
	"github.com/fengde/gocommon/jsonx"
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/pkg/errors"
)

type Action string

const InsertAction = Action(canal.InsertAction)
const UpdateAction = Action(canal.UpdateAction)
const DeleteAction = Action(canal.DeleteAction)

type EventHandler struct {
	canal.DummyEventHandler
	callback func(table string, action Action, data string)
}

// NewEventHandler 事件处理
// callback 自定义的处理函数，jsonData描述变动的内容，
// 如果action=insert, jsonData返回完整的插入数据[{"id": 1, "name": "fedel", "age": 19}]
// 如果action=update, jsonData返回完整的修改数据[{"id": 1, "age": 32}, {"id": 2, "age": 1}]
// 如果action=delete, jsonDaa返回的删除id [{"id": 1}, {"id": 2}, {"id": 3}]
func NewEventHandler(callback func(table string, action Action, jsonData string)) *EventHandler {
	return &EventHandler{
		callback: callback,
	}
}

// OnRow 消费binlog, 目前根据业务只接收insert, update, delete事件
func (p *EventHandler) OnRow(e *canal.RowsEvent) error {
	if p.callback == nil {
		return errors.New("请NewEventHandler初始化")
	}

	table := e.Table.Schema + "." + e.Table.Name

	switch e.Action {
	case canal.InsertAction:
		var ms []map[string]interface{}
		for _, row := range e.Rows {
			m := map[string]interface{}{}
			for i, col := range e.Table.Columns {
				m[col.Name] = row[i]
			}
			ms = append(ms, m)
		}
		p.callback(table, InsertAction, jsonx.MarshalToStringNoErr(ms))
	case canal.UpdateAction:
		var ms []map[string]interface{}
		for i := 0; i < len(e.Rows); i = i + 2 {
			var m = map[string]interface{}{}
			for j, col := range e.Table.Columns {
				if col.Name == "id" {
					m[col.Name] = e.Rows[i][j]
					break
				}
			}
			for j, colValue := range e.Rows[i+1] {
				if colValue != nil {
					m[e.Table.Columns[j].Name] = colValue
				}
			}
			ms = append(ms, m)
		}
		p.callback(table, UpdateAction, jsonx.MarshalToStringNoErr(ms))
	case canal.DeleteAction:
		var ms []map[string]interface{}
		for _, row := range e.Rows {
			m := map[string]interface{}{}
			for i, col := range e.Table.Columns {
				if col.Name == "id" {
					m[col.Name] = row[i]
					break
				}
			}
			ms = append(ms, m)
		}
		p.callback(table, DeleteAction, jsonx.MarshalToStringNoErr(ms))
	}

	return nil

}
