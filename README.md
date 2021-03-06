# goapi

### 数据库基本操作
#### 结构体对应实际数据库字段
```
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
```
#### 每个model表根据需要实现基本操作方法 Find FindOneBy FindBy Insert InsertAll  Update Delete DeleteBy
##### Tablename方法必须实现 设置对应的表名称
```
func (admin *Admin) Tablename() (string) {
	return "admin"
}
```
##### 其他基本操作方法 Find FindOneBy FindBy Insert InsertAll  Update Delete DeleteBy
```
func (admin *Admin) Find(id int) (*Admin, error) {
	fields := admin.ConvertToDbField(admin)		
	sql := fmt.Sprintf("select %s from %s where id = ?", strings.Join(fields, ", "), admin.Tablename())
	//sql := "select * from admin where " //Scan error on column index 5, name "created_at": converting driver.Value type time.Time ("2021-02-02 00:00:00 +0000 UTC") to a int why?
	stmt, _ := Db.Prepare(sql)
	
	row := stmt.QueryRow(id)
	var err error
	if err = row.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.RoleId, &admin.Descript, &admin.IsEnabled, &admin.LastLoginIp, &admin.CreatedAt, &admin.LastLoginAt); err != nil {
		log.Println(err)
		return nil, err
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
		log.Println(err)
		return nil, err
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
```


#### 使用方法
##### 声明结构体
```
admin := &model.Admin{}
var adminModel model.Admin
```
##### find 根据主键返回row
```
admin, _ := adminModel.Find(1)
```
##### 根据条件及排序查找一行或多行
```
var where model.OrderedMap
var order model.OrderedMap
where.Set("username", "admin")
where.Set("password", "adminpassword")
order.Set("id", "desc")

admin, _ := adminModel.FindOneBy(where, order)
admin, _ := adminModel.FindBy(where, order)
```
##### 新增
```
data := map[string]interface{}{
    "username": "mash2",
    "password": "mash2",
    "created_at": "2020-02-03",
}
adminModel.Insert(data)
```
##### 批量插入
```
data1 := map[string]interface{}{
    "username": "mash33",
    "password": "mash3",
    "created_at": "2020-02-03",
}
data2 := map[string]interface{}{
    "username": "mash44",
    "password": "mash4",
    "created_at": "2020-02-03",
}
datas := []map[string]interface{}{data1, data2}
adminModel.InsertAll(datas)
```
##### 更新
```
data := map[string]interface{}{
    "username": "mash1231231",
    "password": "mash2",
    "created_at": "2020-02-03",
}
where := map[string]interface{}{
    "username": "mash2",
}
adminModel.Update(data, where)
```
##### 删除 根据主键删除数据
```
adminModel.Delete(11)
```
##### 批量删除 根据条件删除
```
var where model.OrderedMap
where.Set("password", "123")
adminModel.DeleteBy(where)
```