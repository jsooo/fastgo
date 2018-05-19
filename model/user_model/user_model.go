package user_model

const (
	USER_ALIAS = "db_user"
	USER_DB    = "db_user"
	USER_TB    = "tb_user"
)

type User struct {
	Id          string `orm:"pk;size(32)"`
	Name        string `orm:"size(32);index"`
	CreateTime  int    `orm:"index"`
	UpdateTime  int    `orm:"index"`
	IsDeleted   uint8  `orm:"index"`
	DeletedTime uint   `orm:"index"`
}

func (this *User) TableName() string {
	return USER_TB
}
