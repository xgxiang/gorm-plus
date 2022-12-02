package gormplus

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	GormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
	}
}

type Test1 struct {
	gorm.Model
	Code  string
	Price uint
}

func TestSave(t *testing.T) {
	test1 := Test1{Code: "D455", Price: 200}
	resultDb := Insert(&test1)
	fmt.Println(resultDb)
	fmt.Println(test1)
}

func TestSaveMigrate(t *testing.T) {
	test1 := Test1{Code: "D455", Price: 100}
	resultDb, err := InsertMigrate(&test1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultDb)
	fmt.Println(test1)
}

func TestBatchSave(t *testing.T) {
	test1 := Test1{Code: "D466", Price: 100}
	test2 := Test1{Code: "D466", Price: 100}
	resultDb := InsertBatch(&test1, &test2)
	fmt.Println(resultDb)
	fmt.Println(test1)
	fmt.Println(test2)
}

func TestSaveBatchMigrate(t *testing.T) {
	test1 := Test1{Code: "D477", Price: 100}
	test2 := Test1{Code: "D477", Price: 100}
	resultDb, err := InsertBatchMigrate(&test1, &test2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultDb)
	fmt.Println(test1)
	fmt.Println(test2)
}

func TestDeleteById(t *testing.T) {
	resultDb := DeleteById[Test1](1)
	fmt.Println(resultDb)
}

func TestDeleteByIds(t *testing.T) {
	resultDb := DeleteByIds[Test1](4, 5)
	fmt.Println(resultDb)
}

func TestDelete(t *testing.T) {
	q := Query[Test1]{}
	q.Eq("code", "D1").Eq("price", 100)
	resultDb := Delete(&q)
	fmt.Println(resultDb)
}

func TestUpdateById(t *testing.T) {
	test1 := Test1{Code: "777"}
	resultDb := UpdateById(6, &test1)
	fmt.Println(resultDb)
}

func TestUpdate(t *testing.T) {
	q := Query[Test1]{}
	q.Eq("code", "D42").Eq("price", 100)
	test1 := Test1{Code: "888"}
	resultDb := Update(&q, &test1)
	fmt.Println(resultDb)
}

func TestSelectById(t *testing.T) {
	db, result := SelectById[Test1](1)
	fmt.Println(db)
	fmt.Println(result)
}

func TestSelectByIds(t *testing.T) {
	db, result := SelectByIds[Test1](1, 2)
	fmt.Println(db)
	fmt.Println(result)
}

func TestSelectOne(t *testing.T) {
	q := Query[Test1]{}
	q.Eq("code", "F42").Eq("price", 200).Select("code", "price")
	db, result := SelectOne(&q)
	fmt.Println(db)
	fmt.Println(result)
}

func TestSelectList(t *testing.T) {
	q := Query[Test1]{}
	q.Group("price", "code")
	db, result := SelectList(&q)
	fmt.Println(db.RowsAffected)
	for _, v := range result {
		marshal, _ := json.Marshal(v)
		fmt.Println(string(marshal))
	}
}

func TestSelectCount(t *testing.T) {
	q := Query[Test1]{}
	q.Eq("price", 100)
	db, count := SelectCount(&q)
	fmt.Println(db)
	fmt.Println(count)
}