package sql

import (
	"github.com/meixiaofei/flow-bpmn/qlang"
	"gorm.io/gorm"
)

func Reg(tx *gorm.DB) {
	qlang.Import("sql", map[string]interface{}{
		"Query": func(query string, args ...interface{}) (out []map[string]interface{}) {
			tx.Raw(query, args...).Scan(&out)
			return
		},
		"Count": func(query string, args ...interface{}) (count int64) {
			tx.Raw(query, args...).Count(&count)
			return
		},
		"One": func(query string, args ...interface{}) (out map[string]interface{}) {
			tx.Raw(query, args...).Limit(1).Find(&out)
			return
		},
		"Exec": func(query string, args ...interface{}) int64 {
			return tx.Exec(query, args...).RowsAffected
		},
	})
}
