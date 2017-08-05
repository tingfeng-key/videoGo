package config
type Database struct {
	Type string
	Dns string
}

func Db()  Database {
	db := Database{}
	db.Type = "mysql"
	db.Dns = "root:root@/video_go?charset=utf8"
	return db
}
