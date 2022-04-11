package canalx

import (
	"github.com/fengde/gocommon/logx"
	"testing"
)

func TestNewSlave(t *testing.T) {
	slave, err := NewSlave(100, "", 3306, "", "")
	if err != nil {
		logx.Error(err)
		return
	}
	err = slave.Start(func(table string, action Action, jsonData string) {
			logx.Infof("table: %v, action: %v, json: %v", table, action, jsonData)
			// table: test.user, action: delete, json: [{\"id\":8}]
			// table: test.user, action: insert, json: [{\"age\":30,\"id\":8,\"name\":\"fedel\"}]
			// table: test.user, action: update, json: [{\"age\":32,\"id\":8}]
	})
	if err != nil {
		logx.Error(err)
		return
	}


	select {}
}
