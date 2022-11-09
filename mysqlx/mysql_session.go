package mysqlx

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/fengde/gocommon/errorx"
	"github.com/go-xorm/xorm"
)

type Session struct {
	session *xorm.Session
}

// NewSession 新建session
func NewSession(cluster *Cluster) *Session {
	return &Session{session: cluster.engineGroup.NewSession()}
}

// Insert 插入数据
func (p *Session) Insert(table string, data map[string]interface{}) (int64, error) {
	var columns []string
	var args []interface{}
	var places []string
	for k, v := range data {
		columns = append(columns, columnStandard(k))
		args = append(args, v)
		places = append(places, "?")
	}

	query := fmt.Sprintf(`insert into %s(%s) values(%s)`, table, strings.Join(columns, ", "), strings.Join(places, ", "))

	result, err := p.Exec(query, args...)
	if err != nil {
		return 0, errorx.WithStack(err)
	}

	return result.LastInsertId()
}

// Query 结构化查询记录组，structSlicePtr 传结构体数组指针: tag-- xorm
func (p *Session) Query(query string, args []interface{}, structSlicePtr interface{}) error {
	return p.session.SQL(query, args...).Find(structSlicePtr)
}

// QueryOne 结构化查询单个结果，beanPtr 传非数组指针
func (p *Session) QueryOne(query string, args []interface{}, beanPtr interface{}) (bool, error) {
	return p.session.SQL(query, args...).Get(beanPtr)
}

// Update 更新，返回更新的记录数
func (p *Session) Update(table string, set, where map[string]interface{}) (int64, error) {
	if len(set) == 0 || len(where) == 0 {
		return 0, errorx.New("set,where不允许为空")
	}
	var sets []string
	var args []interface{}
	var wheres []string
	for k, v := range set {
		sets = append(sets, columnStandard(k)+`=?`)
		args = append(args, v)
	}
	for k, v := range where {
		if fmt.Sprintf("%v=%v", k, v) == "1=1" {
			wheres = append(wheres, "1=1")
			continue
		}
		wheres = append(wheres, columnStandard(k)+`=?`)
		args = append(args, v)
	}

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE %s`, table, strings.Join(sets, ", "), strings.Join(wheres, " and "))

	result, err := p.Exec(query, args...)
	if err != nil {
		return 0, errorx.WithStack(err)
	}

	return result.RowsAffected()
}

// UpdateByID 根据id更新
func (p *Session) UpdateByID(table string, set map[string]interface{}, id int64) (int64, error) {
	return p.Update(table, set, map[string]interface{}{
		"id": id,
	})
}

// Delete 删除，返回删除的记录数
func (p *Session) Delete(table string, where map[string]interface{}) (int64, error) {
	if len(where) == 0 {
		return 0, errorx.New("where不允许为空")
	}
	var wheres []string
	var args []interface{}
	for k, v := range where {
		if fmt.Sprintf("%v=%v", k, v) == "1=1" {
			wheres = append(wheres, "1=1")
			continue
		}
		wheres = append(wheres, columnStandard(k)+`=?`)
		args = append(args, v)
	}

	query := fmt.Sprintf(`DELETE FROM %s WHERE %s`, table, strings.Join(wheres, " AND "))

	result, err := p.Exec(query, args...)
	if err != nil {
		return 0, errorx.WithStack(err)
	}

	return result.RowsAffected()
}

// DeleteByID 根据id删除
func (p *Session) DeleteByID(table string, id int64) (int64, error) {
	return p.Delete(table, map[string]interface{}{
		"id": id,
	})
}

// Exec 执行复杂SQL
func (p *Session) Exec(query string, args ...interface{}) (sql.Result, error) {

	return p.session.Exec(append([]interface{}{query}, args...)...)
}

// DoTransaction 执行事务
func (p *Session) DoTransaction(fn func(session *Session) error) error {
	if err := p.session.Begin(); err != nil {
		return errorx.WithStack(err)
	}
	defer p.session.Rollback()

	if err := fn(p); err != nil {
		return err
	}

	return p.session.Commit()
}

// Close 关闭
func (p *Session) Close() {
	p.session.Close()
}
