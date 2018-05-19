package initorm

import (
	"time"

	"gopkg.in/mgo.v2"

	"github.com/astaxie/beego"
)

var (
	S *mgo.Session
)

func InitMongo() {
	mongoURL := beego.AppConfig.String("mongo::url")
	mongoDB := beego.AppConfig.String("mongo::datebase")
	var (
		err error
		url = mongoURL + "/" + mongoDB
	)

	S, err = mgo.DialWithTimeout(url, 5*time.Second)
	if err != nil {
		panic(err)
	}

	S.SetMode(mgo.Monotonic, true)
}