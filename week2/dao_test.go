package week2_test

import (
	"fmt"
	"go-training-camp/week2"
	"net/url"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=%s", "root", "root", "127.0.0.1", 3306, "some", url.QueryEscape("Asia/Shanghai"))
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		t.Logf("create db connection err:%v", err)
		os.Exit(-1)
	}

	_, err = week2.GetUser(db, "2")
	// 通过 errors.Cause() 获取error原始的类型
	if assert.Equal(t, gorm.ErrRecordNotFound, errors.Cause(err)) {
		t.Logf("error trace:%+v", err)
	}

	// 通过errors.Is() 作断言
	if assert.True(t, errors.Is(err, gorm.ErrRecordNotFound)) {
		t.Logf("error trace:%+v", err)
	}
}
