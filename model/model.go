package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"unicode"
	"reflect"
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
		return nil, err
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
		fields = append(fields, field)
		values = append(values, value)
		zanwei = append(zanwei, "?")
	}
	
	var sql string
	sql += "insert into " + model.tablename + "(" + strings.Join(fields, ",") + ") values (" + strings.Join(zanwei, ",") + ")"

	stmt, _ := Db.Prepare(sql)
	result, err := stmt.Exec(values...)
	if err != nil {
		log.Println(err)		
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
		fields = append(fields, field)
		values = append(values, value)
	}
	for field, value := range where {
		fieldsWhere = append(fieldsWhere, field)
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
		log.Println(err)		
	}
	rowsAffected, err := result.RowsAffected()
	return rowsAffected, err
}
