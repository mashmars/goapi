package model

import (
	"strings"
	"log"
	"time"
	"fmt"
)

type Admin struct {
	Model
	Id int64
	Username string
	Password string 
	RoleId int64
	Descript string
	IsEnabled int
	LastLoginIp string
	CreatedAt time.Time `json:"created_at" time_format:"2020-02-01"`
	LastLoginAt time.Time
}

func (admin *Admin) Tablename() (string) {
	return "admin"
}

func (admin *Admin) Find(id int) (*Admin, error) {
	fields := admin.ConvertToDbField(admin)		
	sql := fmt.Sprintf("select %s from %s where id = ?", strings.Join(fields, ", "), admin.Tablename())
	//sql := "select * from admin where " //Scan error on column index 5, name "created_at": converting driver.Value type time.Time ("2021-02-02 00:00:00 +0000 UTC") to a int why?
	stmt, _ := Db.Prepare(sql)
	
	row := stmt.QueryRow(id)
	var err error
	if err = row.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.RoleId, &admin.Descript, &admin.IsEnabled, &admin.LastLoginIp, &admin.CreatedAt, &admin.LastLoginAt); err != nil {
		panic(err)
	}
	
	return admin, err
}

func(admin *Admin) FindOneBy(where OrderedMap, order OrderedMap) (*Admin, error) {	
	fields := admin.ConvertToDbField(admin)		
	sql := fmt.Sprintf("select %s from %s ", strings.Join(fields, ", "), admin.Tablename())
	sql = admin.BuildSql(sql, where, order) + " limit 1"

	row := admin.Row(sql, where)

	//row.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.RoleId, &admin.Descript, &admin.IsEnabled, &admin.LastLoginIp)
	var err error
	if err = row.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.RoleId, &admin.Descript, &admin.IsEnabled, &admin.LastLoginIp, &admin.CreatedAt, &admin.LastLoginAt); err != nil {
		panic(err)
	}
	
	return admin, err
}

func (admin *Admin) FindBy(where OrderedMap, order OrderedMap) ([]Admin, error) {	
	fields := admin.ConvertToDbField(admin)		
	sql := fmt.Sprintf("select %s from %s ", strings.Join(fields, ", "), admin.Tablename())
	sql = admin.BuildSql(sql, where, order)

	rows, err := admin.Rows(sql, where)
	
	var admins []Admin
	for rows.Next() {
		adminNew := Admin{}		
		if err = rows.Scan(&adminNew.Id, &adminNew.Username, &adminNew.Password, &adminNew.RoleId, &adminNew.Descript, &adminNew.IsEnabled, &adminNew.LastLoginIp, &adminNew.CreatedAt, &adminNew.LastLoginAt); err != nil {
			log.Println(err)
			return nil, err
		}
		admins = append(admins, adminNew)
	}
	
	return admins, err
}

func (admin *Admin) FindAll() ([]Admin, error) {	
	fields := admin.ConvertToDbField(admin)		
	sql := fmt.Sprintf("select %s from %s ", strings.Join(fields, ", "), admin.Tablename())
	
	stmt, _ := Db.Prepare(sql)
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	
	var admins []Admin
	for rows.Next() {
		adminNew := Admin{}		
		if err = rows.Scan(&adminNew.Id, &adminNew.Username, &adminNew.Password, &adminNew.RoleId, &adminNew.Descript, &adminNew.IsEnabled, &adminNew.LastLoginIp, &adminNew.CreatedAt, &adminNew.LastLoginAt); err != nil {
			panic(err)
		}
		
		admins = append(admins, adminNew)
	}
	
	return admins, err
}

////
func (admin *Admin) Insert(data map[string]interface{}) (int64, error) {
	admin.SetTablename(admin.Tablename())
	lastInsertId, err := admin.Model.Insert(data)
	return lastInsertId, err
}
func (admin *Admin) InsertAll(datas []map[string]interface{}) ([]int64) {
	admin.SetTablename(admin.Tablename())
	ids := admin.Model.InsertAll(datas)
	return ids
}
func (admin *Admin) Update(data map[string]interface{}, where map[string]interface{}) (int64, error) {	
	admin.SetTablename(admin.Tablename())
	rowsAffected, err := admin.Model.Update(data, where)
	return rowsAffected, err
}

//
func (admin *Admin) Delete(id int) (int64) {
	sql := fmt.Sprintf("delete from %s where id = ?", admin.Tablename())	
	stmt, _ := Db.Prepare(sql)	
	result, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	rowsAffected, _ := result.RowsAffected()
	return rowsAffected
}

func (admin *Admin) DeleteBy(where OrderedMap) (int64) {
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
func(admin *Admin) FindOneBy(kv map[string]string) (*Admin, error) {
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
