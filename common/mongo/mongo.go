package mongo

import (
	"gt-monitor/common/zap"
	"gopkg.in/mgo.v2"
	"time"
)

type MongoPool struct {
	option  *DBOption
	session *mgo.Session
}

func NewMgoPool(options *DBOption) (*MongoPool, error) {
	mango := &MongoPool{}
	mango.option = options
	zap.Get().Info("start mongo session")
	s := mango.NewSession()
	zap.Get().Info("ping mongo")
	err := s.Ping()
	return mango, err
}

func (m *MongoPool) NewSession() *mgo.Session {
	if nil == m.session {
		zap.Get().Info("start connect to mongo db url:", m.option.Addr)
		timeout := time.Duration(m.option.Timeout) * time.Microsecond
		session, err := mgo.DialWithTimeout(m.option.Addr, timeout)
		if err != nil {
			zap.Get().Error("mongo db connect error: ", err)
			panic("dial to mongo failed")
		}
		m.session = session
	}
	return m.session.Clone()
}
