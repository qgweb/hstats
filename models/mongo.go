package models

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"github.com/astaxie/beego"
)

var (
	mgoSession *mgo.Session
)

/**
 * 公共方法，获取session，如果存在则拷贝一份
 */
func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(fmt.Sprintf("%s:%s", beego.AppConfig.String("mongodb::host"),
			beego.AppConfig.String("mongodb::port")))
		if err != nil {
			panic(err) //直接终止程序运行
		}
		mgoSession.SetMode(mgo.Monotonic, true)
	}
	//最大连接池默认为4096
	return mgoSession.Clone()
}

func SaveData(data map[string]interface{}) {
	sess := getSession()
	defer sess.Close()

	err := sess.DB("consult").C("datacenter").Insert(data)
	if err != nil {
		beego.Error(err)
	}
}