package mysqlx

import (
	"testing"
	"time"
)

func TestSession_Delete(t *testing.T) {
	client, err := NewCluster([]string{"rw_lcuser_qa:RsfNC*8xz2^hU@(9.135.108.198:3306)/tools?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"}, time.Minute, 1, 1)
	if err != nil {
		t.Fatal(err)
		return
	}
	if _, err := client.Delete(`alarm_namespace`, map[string]interface{}{
		"1": 1,
	}); err != nil {
		t.Fatal(err)
	}
}
