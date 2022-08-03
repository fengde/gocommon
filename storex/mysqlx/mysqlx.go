package mysqlx

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"xorm.io/core"

	"github.com/fengde/gocommon/logx"

	"github.com/fengde/gocommon/safex"

	"github.com/fengde/gocommon/errorx"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-xorm/xorm"
)

var NotExistError = errorx.New("记录不存在")

type Cluster struct {
	engineGroup *xorm.EngineGroup
}

// NewCluster 创建集群对象，同样适用单个DB实例情况
func NewCluster(dataSourceNames []string, connMaxLifetime time.Duration, maxOpenConns, maxIdleConns int, closeShowSQL ...bool) (*Cluster, error) {
	eg, err := xorm.NewEngineGroup("mysql", dataSourceNames, xorm.LeastConnPolicy())
	if err != nil {
		return nil, errorx.WithStack(err)
	}

	eg.SetConnMaxLifetime(connMaxLifetime)
	eg.SetMaxOpenConns(maxOpenConns)
	eg.SetMaxIdleConns(maxIdleConns)

	if !(len(closeShowSQL) > 0 && closeShowSQL[0]) {
		eg.ShowSQL(true)
		eg.ShowExecTime(true)
	}

	eg.SetLogLevel(core.LOG_INFO)

	if err := eg.Ping(); err != nil {
		return nil, errorx.WithStack(err)
	}

	safex.Go(func() {
		checkSecond := time.Second * 30
		for {
			time.Sleep(checkSecond)
			if err := eg.Ping(); err != nil {
				logx.Error(err)
			}
		}
	})

	return &Cluster{
		engineGroup: eg,
	}, nil
}

func (p *Cluster) session() *Session {
	return NewSession(p)
}

// Insert 插入数据
func (p *Cluster) Insert(table string, data map[string]interface{}) (int64, error) {
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
func (p *Cluster) Query(query string, args []interface{}, structSlicePtr interface{}) error {
	session := p.session()
	defer session.Close()

	return session.Query(query, args, structSlicePtr)
}

// QueryOne 结构化查询单个结果，beanPtr 传非数组指针
func (p *Cluster) QueryOne(query string, args []interface{}, beanPtr interface{}) (bool, error) {
	session := p.session()
	defer session.Close()

	return session.QueryOne(query, args, beanPtr)
}

// Update 更新，返回更新的记录数
func (p *Cluster) Update(table string, set, where map[string]interface{}) (int64, error) {
	session := p.session()
	defer session.Close()

	return session.Update(table, set, where)
}

// UpdateByID 根据id进行更新，返回更新的记录数
func (p *Cluster) UpdateByID(table string, set map[string]interface{}, id int64) (int64, error) {
	session := p.session()
	defer session.Close()

	return session.Update(table, set, map[string]interface{}{
		"id": id,
	})
}

// Delete 删除，返回删除的记录数
func (p *Cluster) Delete(table string, where map[string]interface{}) (int64, error) {
	session := p.session()
	defer session.Close()

	return session.Delete(table, where)
}

// DeleteByID 根据id进行删除
func (p *Cluster) DeleteByID(table string, id int64) (int64, error) {
	session := p.session()
	defer session.Close()

	return session.Delete(table, map[string]interface{}{
		"id": id,
	})
}

// Exec 执行复杂SQL
func (p *Cluster) Exec(query string, args ...interface{}) (sql.Result, error) {
	session := p.session()
	defer session.Close()

	return session.Exec(query, args...)
}

// DoTransaction 执行事务
func (p *Cluster) DoTransaction(fn func(session *Session) error) error {
	session := p.session()
	defer session.Close()

	return session.DoTransaction(fn)
}
