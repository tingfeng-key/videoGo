package structs

type Types struct {
	Id int64
	Name string
	typeStr string `xorm:"CHAR(50) NOTNULL default('se_video')"`
	Sort int8 `xorm:"TINYINT NOTNULL DEFAULT 0"`//排序
	Status int8 `xorm:"TINYINT NOTNULL DEFAULT 1"`
}