package domain

type UserEntity struct {
	ID        string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Username  string `gorm:"unique;not null;type:varchar(255)"`
	Password  string `gorm:"not null;type:varchar(255)"`
	Email     string `gorm:"unique;not null;type:varchar(255)"`
	CreatedAt string `gorm:"not null;type:timestamp;default:now()"`
	UpdatedAt string `gorm:"not null;type:timestamp"`
}

func (UserEntity) TableName() string {
	return "users"
}
