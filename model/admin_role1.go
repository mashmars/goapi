package model

import (
	"fmt"
	"strings"
)

type AdminRole1 struct {
	Model
	Id int64			`json:"id"`
	Name string			`json:"name"`
	IsEnabled int		`json:"is_enabled"`
}


func (admin_role *AdminRole1) Tablename() string {
	return "admin_role"
}

func (admin_role *AdminRole1) FindAll() ([]AdminRole1) {
	fields := admin_role.ConvertToDbField(admin_role)
	sql := fmt.Sprintf("select %s from %s ", strings.Join(fields, ", "), admin_role.Tablename())

	stmt, _ := Db.Prepare(sql)
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var adminRoles []AdminRole1
	for rows.Next() {
		adminRole := AdminRole1{}
		if err = rows.Scan(&adminRole.Id, &adminRole.Name, &adminRole.IsEnabled); err != nil {
			panic(err)
		}

		adminRoles = append(adminRoles, adminRole)
	}

	return adminRoles
}