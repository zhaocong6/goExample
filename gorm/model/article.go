package model

type Article struct {
	Model
	User    User   `gorm:"ForeignKey:User"`
	UserId  uint64 `gorm:"index"`
	Content string `gorm:"type:text;column:article"`
}

func (*Article) TableName() string {
	return "articles"
}
