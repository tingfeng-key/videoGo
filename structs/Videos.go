package structs

type Videos struct {
	Id int64
	TypeId int64 //类型ID
	Title string //标题
	Thumb string //缩略图
	FileUrl string //文件URL
	SourceUrl string //资源URL
	Remark string //备注
	Sort int16 //排序
	Status int8 `xorm:"DEFAULT 1"` //状态
	CreatedAt int64 `xorm:"created NULL"` //创建时间
	UpdatedAt int64 `xorm:"updated NULL"` // 更新时间
	DeletedAt int64 `xorm:"deleted NULL"` //删除时间
}