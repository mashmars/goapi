package model

import (
	"database/sql"
	"database/sql/driver"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"unicode"
	"reflect"
	"time"
)

var Db *sql.DB


type Model struct{
	tablename string
}

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:mash@tcp(127.0.0.1)/sfadmin?parseTime=true")
	if err != nil {
		panic("数据库连接失败")
	}
}

//返回的时间格式转换 2021-02-02T00:00:00Z => 2021-02-02 00:00:00
const TimeFormat = "2021-02-02 00:00:00"
type LocalTime time.Time
func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return
	}
  
	// 指定解析的格式
	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = LocalTime(now)
	return
}
func (t LocalTime) MarshalJSON() ([]byte, error) {
    b := make([]byte, 0, len(TimeFormat)+2)
    b = append(b, '"')
    b = time.Time(t).AppendFormat(b, TimeFormat)
    b = append(b, '"')
    return b, nil
}
func (t LocalTime) Value() (driver.Value, error) {
    if t.String() == "0001-01-01 00:00:00" {
        return nil, nil
    }
    return []byte(time.Time(t).Format(TimeFormat)), nil
}

func (t *LocalTime) Scan(v interface{}) error {
    tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
    *t = LocalTime(tTime)
    return nil
}

func (t LocalTime) String() string {
    return time.Time(t).Format(TimeFormat)
}


//为了实现有序的处理字段值
type OrderedMap struct {
	Keys []string
	Values []interface{}
}

func (om *OrderedMap) Set(key string, value string) {
	om.Keys = append(om.Keys, key)
	om.Values = append(om.Values, value)
	log.Println(om)
	
}

func (om *OrderedMap) Delete(key string) {
	var index int 
	for i, k := range om.Keys {
		if k == key {
			index = i
		}
	}
	
	newKeys := append(om.Keys[:index], om.Keys[index+1:]...)
	om.Keys = newKeys	

	newValues := append(om.Values[:index], om.Values[index+1:]...)
	om.Values = newValues
	
}

func (om *OrderedMap) PrintAll() {
	log.Println(om)
}

////驼峰单词转下划线单词
func CamelCaseToUdnderscore(s string) string {
    var output []rune
    for i, r := range s {
        if i == 0 {
            output = append(output, unicode.ToLower(r))
        } else {
            if unicode.IsUpper(r) {
                output = append(output, '_')
            }

            output = append(output, unicode.ToLower(r))
        }
    }
    return string(output)
}
//获取结构体中字段的名称
func GetFieldName(structName interface{}) []string  {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic("Check type error not Struct")
		log.Println("Check type error not Struct")
		return nil		
	}
	fieldNum := t.NumField()
	result := make([]string,0,fieldNum)
	for i:= 0;i<fieldNum;i++ {
		if name := t.Field(i).Name; name != "Model" {
			result = append(result,CamelCaseToUdnderscore(t.Field(i).Name))	
		}		
	}
	return result
}

//
func (model *Model) SetTablename(tablename string) (string) {
	model.tablename = tablename
	return model.tablename
}

//
func (model *Model) BuildSql(sql string, where OrderedMap, order OrderedMap) (string) {
	if len(where.Keys) > 0 {
		sql += " where "
		sql += strings.Join(where.Keys, " = ? and ") + " = ? "
	}

	if len(order.Keys) > 0 {
		sql += " order by " ;
		for index, field := range order.Keys {
			sql += field + " " + order.Values[index].(string)
		}
	}	
	return sql
}

//
func (model *Model) Rows(sql string, where OrderedMap) (*sql.Rows, error) {
	stmt, _ := Db.Prepare(sql)
	rows, err := stmt.Query(where.Values...)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return rows, err
}

//
func (model *Model) Row(sql string, where OrderedMap) (*sql.Row) {
	stmt, _ := Db.Prepare(sql)
	row := stmt.QueryRow(where.Values...)
	return row
}

//
func (model *Model) ConvertToDbField(structName interface{}) []string {
	return GetFieldName(structName)
}


//insert update insertAll
func (model *Model) Insert(data map[string]interface{}) (int64, error) {
	//sql := "insert into admin(username, password) values (?, ?)"
	var values []interface{}
	var fields []string
	var zanwei []string

	for field, value := range data {
		fields = append(fields, CamelCaseToUdnderscore(field))
		values = append(values, value)
		zanwei = append(zanwei, "?")
	}
	
	var sql string
	sql += "insert into " + model.tablename + "(" + strings.Join(fields, ",") + ") values (" + strings.Join(zanwei, ",") + ")"
	
	stmt, _ := Db.Prepare(sql)
	
	result, err := stmt.Exec(values...)
	if err != nil {
		panic(err)		
	}
	
	lastInsertId, err := result.LastInsertId()
	return lastInsertId, err
}

func (model *Model) InsertAll(datas []map[string]interface{}) ([]int64) {
	var ids []int64
	for _, data := range datas {
		lastInsertId, _ := model.Insert(data)
		ids = append(ids, lastInsertId)
	}
	return ids
}

func (model *Model) Update(data map[string]interface{}, where map[string]interface{}) (int64, error) {
	//update admin set username = ?, password = ? where field = ? and field = ?
	var values []interface{}
	var fields []string
	var fieldsWhere []string

	for field, value := range data {
		fields = append(fields, CamelCaseToUdnderscore(field))
		values = append(values, value)
	}
	for field, value := range where {
		fieldsWhere = append(fieldsWhere, CamelCaseToUdnderscore(field))
		values = append(values, value)
	}
	
	var sql = "update " + model.tablename + " set "
	sql += strings.Join(fields, " = ?, ") + " = ? "

	if len(where) > 0 {
		sql += " where " + strings.Join(fieldsWhere, " = ? ") + " = ? "
	}

	stmt, _ := Db.Prepare(sql)
	result, err := stmt.Exec(values...)
	if err != nil {
		panic(err)		
	}
	rowsAffected, err := result.RowsAffected()
	return rowsAffected, err
}

//分页
func pagination() {
	/*page := 1
	pageSize := 10
	offset := (page - 1) * pageSize + 1
	pages := 123
	sql := "select * from admin order id asc" /// 11 - 20
	sql += " limit 10 10 "
	stmt, _ := Db.Prepare(sql)
	values := []interface{}{}
	values = append(values, offset, pageSize)
	
	sql = "select count(id) from admin"
	*/
}