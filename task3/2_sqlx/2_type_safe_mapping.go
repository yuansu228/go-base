package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

/*
CREATE TABLE books (
    id int NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'Primary Key',
    title  VARCHAR(100) COMMENT 'title',
    author   VARCHAR(100) COMMENT 'author ',
    price   FLOAT(19,2) COMMENT 'price '
) DEFAULT CHARSET UTF8 COMMENT 'books';

insert into books(title,author,price) values('语文','张三',40);
insert into books(title,author,price) values('数学','李四',50);
insert into books(title,author,price) values('英语','王五',60);
insert into books(title,author,price) values('计算机','赵六',70);
*/

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float32 `db:"price"`
}

func typeSafeMapping(db *sqlx.DB) {
	book := Book{}
	db.Get(&book, "select id,title,author,price from books where price>?", 50)
	fmt.Println(book)
}
