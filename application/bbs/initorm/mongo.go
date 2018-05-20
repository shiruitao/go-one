package initorm

import (
	"gopkg.in/mgo.v2"

	"github.com/astaxie/beego"
)

type Mongodb struct {
	S *mgo.Session
	D *mgo.Database
	C *mgo.Collection
}

func ConnectMongo(collection string) (M Mongodb) {
	var err error
	mongoURL := beego.AppConfig.String("mongo::url")
	url := mongoURL + "/" + "bbs"
	M.S, err = mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	M.S.SetMode(mgo.Monotonic, true)

	M.D = M.S.DB("bbs")
	M.C = M.D.C(collection)
	return
}
