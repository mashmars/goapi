package model

import (
	"strings"
	"time"
	"fmt"
)

type Admin123 struct {
	Model
	Id int64				`json:"id"`
	Username string			`json:"username"`
	Password string 		`json:"password"`
	RoleId int64			`json:"role_id"`
	Descript string			`json:"descript"`
	IsEnabled int			`json:"is_enabled"`
	LastLoginIp string		`json:"last_login_ip"`
	CreatedAt time.Time 	`json:"created_at" time_format:"2020-02-01"`
	LastLoginAt time.Time	`json:"last_login_at"`
}

func (admin *Admin123) Tablename() (string) {
	return "admin"
}

func (admin *Admin123) Find(id int) (*Admin123) {
	fields := admin.ConvertToDbField(admin)		
	sql := fmt.Sprintf("select %s from %s where id = ?", strings.Join(fields, ", "), admin.Tablename())
	//sql := "select * from admin where " //Scan error on column index 5, name "created_at": converting driver.Value type time.Time ("2021-02-02 00:00:00 +0000 UTC") to a int why?
	stmt, _ := Db.Prepare(sql)
	
	row := stmt.QueryRow(id)
	var err error
	if err = row.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.RoleId, &admin.Descript, &admin.IsEnabled, &admin.LastLoginIp, &admin.CreatedAt, &admin.LastLoginAt); err != nil {
		panic(err)
	}
	
	return admin
}

func(admin *Admin123) FindOneBy(where OrderedMap, order OrderedMap) (*Admin123) {	
	fields := admin.ConvertToDbField(admin)		
	sql := fmt.Sprintf("select %s from %s ", strings.Join(fields, ", "), admin.Tablename())
	sql = admin.BuildSql(sql, where, order) + " limit 1"

	row := admin.Row(sql, where)

	//row.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.RoleId, &admin.Descript, &admin.IsEnabled, &admin.LastLoginIp)
	var err error
	if err = row.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.RoleId, &admin.Descript, &admin.IsEnabled, &admin.LastLoginIp, &admin.CreatedAt, &admin.LastLoginAt); err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	
	return admin
}

func (admin *Admin123) FindBy(where OrderedMap, order OrderedMap) ([]Admin123) {	
	fields := admin.ConvertToDbField(admin)		
	sql := fmt.Sprintf("select %s from %s ", strings.Join(fields, ", "), admin.Tablename())
	sql = admin.BuildSql(sql, where, order)

	rows := admin.Rows(sql, where)
	
	var admins []Admin123
	for rows.Next() {
		adminNew := Admin123{}		
		if err := rows.Scan(&adminNew.Id, &adminNew.Username, &adminNew.Password, &adminNew.RoleId, &adminNew.Descript, &adminNew.IsEnabled, &adminNew.LastLoginIp, &adminNew.CreatedAt, &adminNew.LastLoginAt); err != nil {
			panic(err)
		}
		admins = append(admins, adminNew)
	}
	
	return admins
}

func (admin *Admin123) FindAll() ([]Admin123) {	
	fields := admin.ConvertToDbField(admin)		
	sql := fmt.Sprintf("select %s from %s ", strings.Join(fields, ", "), admin.Tablename())
	
	stmt, _ := Db.Prepare(sql)
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	
	var admins []Admin123
	for rows.Next() {
		adminNew := Admin123{}		
		if err = rows.Scan(&adminNew.Id, &adminNew.Username, &adminNew.Password, &adminNew.RoleId, &adminNew.Descript, &adminNew.IsEnabled, &adminNew.LastLoginIp, &adminNew.CreatedAt, &adminNew.LastLoginAt); err != nil {
			panic(err)
		}
		
		admins = append(admins, adminNew)
	}
	
	return admins
}

////
func (admin *Admin123) Insert(data map[string]interface{}) (int64) {
	admin.SetTablename(admin.Tablename())
	lastInsertId := admin.Model.Insert(data)
	return lastInsertId
}
func (admin *Admin123) InsertAll(datas []map[string]interface{}) ([]int64) {
	admin.SetTablename(admin.Tablename())
	ids := admin.Model.InsertAll(datas)
	return ids
}
func (admin *Admin123) Update(data map[string]interface{}, where map[string]interface{}) (int64) {	
	admin.SetTablename(admin.Tablename())
	rowsAffected := admin.Model.Update(data, where)
	return rowsAffected
}

//
func (admin *Admin123) Delete(id int) (int64) {
	sql := fmt.Sprintf("delete from %s where id = ?", admin.Tablename())	
	stmt, _ := Db.Prepare(sql)	
	result, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	rowsAffected, _ := result.RowsAffected()
	return rowsAffected
}

func (admin *Admin123) DeleteBy(where OrderedMap) (int64) {
	sql := fmt.Sprintf("delete from %s ", admin.Tablename())
	sql = admin.BuildSql(sql, where, OrderedMap{})
	stmt, _ := Db.Prepare(sql)
	result, err := stmt.Exec(where.Values...)
	if err != nil {
		panic(err)
	}
	rowsAffected, _ := result.RowsAffected()
	return rowsAffected
}


/*
func(admin *Admin123) FindOneBy(kv map[string]string) (*Admin123, error) {
	sql := "select id, username, password, role_id, descript, is_enabled, last_login_ip, created_at, last_login_at from admin where "
	//sql := "select * from admin where " //Scan error on column index 5, name "created_at": converting driver.Value type time.Time ("2021-02-02 00:00:00 +0000 UTC") to a int why?
	values := []interface{}{}
	fields := []string{}

	for field, value := range kv {
		fields = append(fields, field)
		values = append(values, value)
	}

	sql += strings.Join(fields, " = ? and ") + " = ? limit 1";	
	stmt, _ := Db.Prepare(sql)
	row := stmt.QueryRow(values...)
	//row.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.RoleId, &admin.Descript, &admin.IsEnabled, &admin.LastLoginIp)
	var err error
	if err = row.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.RoleId, &admin.Descript, &admin.IsEnabled, &admin.LastLoginIp, &admin.CreatedAt, &admin.LastLoginAt); err != nil {
		log.Println(err)
		return nil, err
	}
	
	return admin, err
}*/
