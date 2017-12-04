package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"gopkg.in/logger.v1"
)

var engine *xorm.Engine

type User struct {
	Id        int64
	Name      string   `xorm:"varchar(50) notnull unique 'user_name'"`
	CreatedAt JsonTime `xorm:"created"`
}

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(j).Format("2006-01-02 15:04:05") + `"`), nil
}

func main() {
	var err error

	engine, err = xorm.NewEngine("mysql", "root:123456@/test?charset=utf8")
	if err != nil {
		//todo something
	}

	//connection pool settings
	engine.SetMaxIdleConns(5)
	engine.SetMaxOpenConns(10)
	log.Info("_")
	engine.SetLogLevel(core.LOG_DEBUG)
	engine.ShowSQL(true)
	engine.ShowExecTime(true)

	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "test_")
	//	cacheMapper := core.NewCacheMapper(core.SnakeMapper{})
	engine.SetTableMapper(tbMapper)

	var user User
	isExist, err := engine.IsTableExist(&user)
	if err == nil && !isExist {
		engine.CreateTables(&user)
		engine.CreateUniques(&user)

	}
	user.Name = "liwei"

	result, err := engine.Insert(&user)
	if err != nil {
		log.Infof("ERROR %s", err)
	}
	log.Info(result)

	fmt.Println("vim-go")
}
