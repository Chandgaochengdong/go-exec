package main

import (
	model2 "com.shopline.richard/exec/dal/dao/model"
	"com.shopline.richard/exec/dal/dao/query"
	"com.shopline.richard/exec/dal/model"
	"com.shopline.richard/exec/initializer/mysql"
	"context"
	"fmt"
	"gopkg.inshopline.com/commons/config"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
)

func main() {
	fmt.Println("let's begin the project!")
	// 配置读取
	// 加载配置
	config.LoadDefaults()
	dsn, found := config.Find("rdb.default.dsn")
	if !found {
		panic("missing RDB DSN config")
	}
	fmt.Println("rdb.default.dsn:" + dsn)
	mysql.Init()
	db := mysql.Db
	fmt.Sprintf("mysql DB %v", db)
	fmt.Println(db)
	/*stu := model.Student{
		Name:       "richard3",
		Age:        18,
		Sex:        "男",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	db.Create(&stu)
	db.First(&stu)
	rad := &model.Student{}
	last := &model.Student{}
	db.Take(&rad)
	db.Last(&last)
	fmt.Printf("stu:%v \n", stu)
	fmt.Printf("stu:%v \n", rad)
	fmt.Printf("stu:%v \n", last)
	*/
	//.Select("id","name", "age", "create_time", "update_time").
	//var stus []model.Student
	//prictiseGorm(db)

}

func practiseRpc() {

}

func prictiseGorm(db *gorm.DB) {
	var stu model.Student
	db.Where("id", 1).Order("age asc").Limit(10).Find(&stu)
	fmt.Printf("stu v1:%v \n", stu)
	stu.Name = "John"
	stu.Age = 41
	db.Save(&stu)

	db.Where("id", 1).Order("age asc").Limit(10).Find(&stu)

	fmt.Printf("stu v2:%v \n", stu)

	//根据条件更新
	var stuV3 model.Student
	db.Model(&stuV3).Where("id = ?", 7).Updates(model.Student{Name: "Candy", Age: 23})

	db.Where("id", 7).Find(&stuV3)
	fmt.Printf("stu v3:%v \n", stuV3)
	db.Model(&model.Student{}).Delete("id", 4)
	db.Transaction(func(tx *gorm.DB) error {
		return nil
	})
	stuV4, count, err := query.Q.Student.WithContext(context.Background()).FindByPage(0, 100)
	if err == nil && count > 0 {
		for _, stu := range stuV4 {
			fmt.Println(*stu)
		}
	}
	query.Q.Student.WithContext(context.Background()).Create(&model2.Student{
		Name: "Lisa",
		Age:  "199",
		Sex:  "女",
	})
}

func FmtSlice2String(x any) (res string) {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	switch val.Kind() {
	case reflect.Struct:
		typ := val.Type()
		//获取结构体里的名称
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			res += field.Name + ":" + FmtSlice2String(val.Field(i).Interface()) + " "
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			res += FmtSlice2String(val.Index(i).Interface())
		}
	default:
		res += fmt.Sprint(x)
	}
	return
}
