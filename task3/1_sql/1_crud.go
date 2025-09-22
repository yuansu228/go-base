package main

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name  string
	Age   int
	Grade string
}

func crud(db *gorm.DB) {
	// db.AutoMigrate(&Student{})

	// stu := Student{Name: "张三", Age: 20, Grade: "三年级"}
	// result := db.Create(&stu)
	// fmt.Println(result.Error, result.RowsAffected)

	// var stus []Student
	// db.Debug().Where("age > ?", 18).Find(&stus)
	// fmt.Println(stus)

	// db.Debug().Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")

	db.Debug().Where("age < ?", 15).Delete(&Student{})
}
