package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	// sql driver
	_ "github.com/go-sql-driver/mysql"
	"go.yyue.dev/common/message"
	"go.yyue.dev/common/utils"
	"go.yyue.dev/datamanage/proto"
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
func GetDb(table string) *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.99.100:3306)/"+table+"?charset=utf8")
	utils.PanicErr(err)

	return db
}

func Connect(database string) (*sql.DB, error) {
	dbc := utils.GetConfig().Database
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?%s`, dbc.User, dbc.Password,
		dbc.Domain, dbc.Port, database, "charset=utf8&parseTime=true")
	fmt.Println(dsn)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	//设置连接池
	conn.SetMaxOpenConns(100)
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(10 * time.Minute)
	return conn, conn.Ping()
}

type Query struct {
	DBMap map[string]*sql.DB
}

func New() *Query {

	return &Query{make(map[string]*sql.DB)}
}

func (q *Query) Connect(ctx context.Context, req *proto.ConnectRequest, reply *message.Message) error {
	database := req.Database
	if DB, ok := q.DBMap[database]; ok && DB.Ping() == nil {
		reply.Success(fmt.Sprintf("存在已连接的数据库[%s]", database))
		return nil
	}

	DB, err := Connect(database)

	if err != nil {
		return reply.Error(fmt.Sprintf("连接数据库[%s]失败: %s", database, err.Error()))
	}
	q.DBMap[database] = DB
	reply.Success(fmt.Sprintf("连接数据库[%s]成功", database))
	return nil
}
func (q *Query) InsertHouse(ctx context.Context, req *proto.InsertHouseRequest, reply *message.Message) error {
	fmt.Println("receive insert: ")
	DB, ok := q.DBMap[req.Database]

	if !ok {
		return reply.Error(fmt.Sprintf("未连接数据库[%s]", req.Database))
	}
	fmt.Println("exists database: ", req.Database)
	if DB.Ping() != nil {
		return reply.Error(fmt.Sprintf("数据库[%s]已断开", req.Database))
	}
	fmt.Println("DB alive: ", DB)
	tx, err := DB.Begin()
	defer tx.Rollback()
	houseInfo := req.House.GetHouseInfo()
	sql := fmt.Sprintf(`INSERT INTO %s (
		house_no, url, title, sub_title, region, total_price, unit_price, room_info, room_sub, type_info, type_sub, area_info, area_sub
	) VALUES (%s, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q);`,
		"house_info",
		houseInfo.GetHouseNo(),
		houseInfo.GetUrl(),
		houseInfo.GetTitle(),
		houseInfo.GetSubTitle(),
		houseInfo.GetRegion(),
		houseInfo.GetTotalPrice(),
		houseInfo.GetUnitPrice(),
		houseInfo.GetRoomInfo(),
		houseInfo.GetRoomSub(),
		houseInfo.GetTypeInfo(),
		houseInfo.GetTypeSub(),
		houseInfo.GetAreaInfo(),
		houseInfo.GetAreaSub(),
	)
	fmt.Println("sql: ", sql)
	_, err = tx.Exec(sql)
	if err != nil {
		return err
	}
	houseBaseInfo := req.House.GetHouseBaseInfo()
	_, err = tx.Exec(
		fmt.Sprintf(`INSERT INTO %s (
			house_no, layout, floor, area_build, struct_house, area_inner, build_type, face, struct_build, decoration, elevator_ratio,
			elevator, property
		) VALUES (%s, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q);`,
			"house_base_info",
			houseBaseInfo.GetHouseNo(),
			houseBaseInfo.GetLayout(),
			houseBaseInfo.GetFloor(),
			houseBaseInfo.GetAreaBuild(),
			houseBaseInfo.GetStructHouse(),
			houseBaseInfo.GetAreaInner(),
			houseBaseInfo.GetBuildType(),
			houseBaseInfo.GetFace(),
			houseBaseInfo.GetStructBuild(),
			houseBaseInfo.GetDecoration(),
			houseBaseInfo.GetElevatorRatio(),
			houseBaseInfo.GetElevator(),
			houseBaseInfo.GetProperty(),
		))
	if err != nil {
		return err
	}
	houseTransactionInfo := req.House.GetHouseTransactionInfo()
	_, err = tx.Exec(
		fmt.Sprintf(`INSERT INTO %s (
			house_no, listing_time, trading_authority, last_transaction, housing_purposes, house_year, property_rights, mortgage_info, document_photo
		) VALUES (%s, %q, %q, %q, %q, %q, %q, %q, %q);`,
			"house_transaction_info",
			houseTransactionInfo.GetHouseNo(),
			houseTransactionInfo.GetListingTime(),
			houseTransactionInfo.GetTradingAuthority(),
			houseTransactionInfo.GetLastTransaction(),
			houseTransactionInfo.GetHousingPurposes(),
			houseTransactionInfo.GetHouseYear(),
			houseTransactionInfo.GetPropertyRights(),
			houseTransactionInfo.GetMortgageInfo(),
			houseTransactionInfo.GetDocumentPhoto(),
		))
	if err != nil {
		return err
	}
	housePics := req.House.GetHousePics()
	for _, housePic := range housePics {
		_, err = tx.Exec(
			fmt.Sprintf(`INSERT INTO %s (
				house_no, description, pic_small_url, pic_normal_url, pic_large_url
			) VALUES (%s, %q, %q, %q, %q);`,
				"house_pic",
				housePic.GetHouseNo(),
				housePic.GetDescription(),
				housePic.GetPicSmallUrl(),
				housePic.GetPicNormalUrl(),
				housePic.GetPicLargeUrl(),
			))
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	reply.Success("insert success")
	return nil
}

func execSliceInsert(DB *sql.DB, table string, ins interface{}) (message.Message, error) {
	reply := message.Message{}
	v := reflect.ValueOf(ins)
	// 剥离指针
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == reflect.Struct {
		v = reflect.ValueOf([]interface{}{v.Interface()})
	}
	// 输入只有一个元素, 且为slice, 当成多个输入
	if v.Kind() != reflect.Slice {
		return reply, reply.Error("Insert error: first argument must be a slice or struct")
	}
	// 先用第一个无素生成keys
	var (
		keys []string
		// values [][]string
		value []string
		stmt  *sql.Stmt
		err   error
		errs  []string
	)

	for i := 0; i < v.Len(); i++ {
		item := v.Index(i)
		// 剥离指针
		for item.Kind() == reflect.Ptr || item.Kind() == reflect.Interface {
			item = item.Elem()
		}
		if i == 0 {
			switch item.Kind() {
			case reflect.Struct:
				keys, value = struct2KV(item)
			}
			sql := fmt.Sprintf(`INSERT %s SET %s=?`, table, strings.Join(keys, "=?,"))
			log.Println("sql: ", sql)
			stmt, err = DB.Prepare(sql)
			if err != nil {
				reply.Error(err.Error())
				return reply, err
			}
			defer stmt.Close()
		} else {
			_, value = struct2KV(item)
		}
		values := make([]interface{}, len(value))
		for i, v := range value {
			values[i] = v
		}
		log.Println("insert data: ", values)
		_, err = stmt.Exec(values...)
		if err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return reply, reply.Error(strings.Join(errs, "\n"))
	}
	reply.Success(fmt.Sprintf("Insert %d data success", v.Len()))
	return reply, nil
}

func struct2KV(s reflect.Value) ([]string, []string) {
	var keys, values = []string{}, []string{}
	t := s.Type()
	for i := 0; i < t.NumField(); i++ {
		tf := t.Field(i)
		vf := s.Field(i)
		// 忽略非导出字段
		if tf.Anonymous || !vf.CanInterface() {
			continue
		}
		// 忽略无效, 零值字段
		if !vf.IsValid() || reflect.DeepEqual(vf.Interface(), reflect.Zero(vf.Type())) {
			continue
		}
		// 剥离指针
		for vf.Type().Kind() == reflect.Ptr {
			vf = vf.Elem()
		}
		// 支持嵌套struct , 但不是Time类型的
		if vf.Type().Kind() == reflect.Struct && tf.Type.Name() != "Time" {
			subKeys, subValues := struct2KV(vf)
			keys = append(keys, subKeys...)
			values = append(values, subValues...)
			continue
		}
		// 获取type tag
		key := strings.Split(tf.Tag.Get("db"), ",")[0]
		if key == "" {
			continue
		}
		value := format(vf)
		if value != "" {
			keys = append(keys, key)
			values = append(values, value)
		}
	}
	return keys, values
}

func format(v reflect.Value) string {
	//断言出time类型直接转unix时间戳
	if t, ok := v.Interface().(time.Time); ok {
		return fmt.Sprintf("FROM_UNIXTIME(%d)", t.Unix())
	}
	switch v.Kind() {
	case reflect.String:
		return fmt.Sprintf("%s", v.Interface())
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return fmt.Sprintf("%d", v.Interface())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", v.Interface())
	//如果是切片类型，遍历元素，递归格式化成"(, , , )"形式
	case reflect.Slice:
		var values []string
		for i := 0; i < v.Len(); i++ {
			values = append(values, format(v.Index(i)))
		}
		return fmt.Sprintf(`(%s)`, strings.Join(values, ","))
	//接口类型剥一层递归
	case reflect.Interface:
		return format(v.Elem())
	}
	return ""
}
