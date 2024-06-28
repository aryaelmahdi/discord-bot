package domain

type Users struct {
	ID       uint   `gorm:"type:int;autoIncrement;primaryKey" json:"articleID"`
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(50);not null" json:"password"`
	Class    string `gorm:"type:varchar(10);not null" json:"class"`
}
