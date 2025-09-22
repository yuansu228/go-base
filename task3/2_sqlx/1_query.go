package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

/*
CREATE TABLE employees (
    id int NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'Primary Key',
    name  VARCHAR(100) COMMENT 'name',
    department  VARCHAR(100) COMMENT 'department',
    salary  FLOAT(19,2) COMMENT 'salary'
) DEFAULT CHARSET UTF8 COMMENT 'employees';

insert into employees(name,department,salary) values('张三','技术部',5000);
insert into employees(name,department,salary) values('李四','财务部',6000);
insert into employees(name,department,salary) values('王五','人事部',3000);
*/

type Employee struct {
	Name       string
	Department string
	Salary     float32
}

func query(db *sqlx.DB) {
	employees := []Employee{}
	db.Select(&employees, "select name,department from employees where department = ?", "技术部")
	fmt.Println(employees)

	err := db.Select(&employees, "select name,department,salary from employees order by salary desc")

	if err == nil {
		fmt.Println(employees[0])
	}
}
