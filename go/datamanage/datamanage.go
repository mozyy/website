package datamanage

import (
	"database/sql"

	"github.com/mozyy/utils"

	// sql driver
	_ "github.com/go-sql-driver/mysql"
)

// private String dataKey; // 服务名称
// private String sqlCode; // sql标识
// private String sqlRemark; // SQL语句描述
// private String tableName; // 数据库表名称（与sql_code形成唯一索引）
// private String operationType; // SQL类型（DELETE:删除，UPDATE:修改，SELECT:查询，INSERT:新增）
// private String sqlInfo; // SQL语句
// private Integer operTimes; // 操作次数
// private Integer slowThreshold; // 慢查询阀值（second，毫秒）
// private Integer slowMaxSecond; // 慢查询最大时间（second，毫秒）
// private Integer isParamAgile; // 是否参数可变（0：否，1：是；针对SELECT语句）
// private String pageSql; // 分页语句（若有分页则把分页语句写入，sql_info仍需写入，默认：LIMIT #{startIndex}, #{pageSize}）
// private String orderSql; // 排序语句（若有排序则把排序语句写入，sql_info仍需写入，默认：ORDER BY #{orderField} #{orderFieldType}）
// private String creator;//创建者

// type Stmt struct {
// }

// GetDb is get database stmt
func GetDb() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.99.100:3306)/user?charset=utf8")
	utils.PanicErr(err)

	return db
}
