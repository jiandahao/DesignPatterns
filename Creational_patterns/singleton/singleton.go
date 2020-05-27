package singleton

import "sync"

//type Singleton interface {
//	func GetSingleton()
//}

type DatabaseSingleton struct {
	mux sync.Mutex
	dbInstance *Database
}

type Database struct {
	User string
	Password string
}

func newDatabase() *Database {
	return &Database{
		User:"jiandahao",
		Password:"1234567",
	}
}

func (d *DatabaseSingleton) GetInstance() *Database {
	if d.dbInstance == nil {
		d.mux.Lock()
		defer d.mux.Unlock()
		// 确保在该线程等待解锁时，其他线程没有初始化该实例。
		if d.dbInstance == nil {
			d.dbInstance = newDatabase()
		}
	}
	return d.dbInstance
}
