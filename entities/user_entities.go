package entities

import "time"

type UserEntity struct {
	ID         int              `gorm:"column:id;primaryKey" json:"id"`
	Name       string           `gorm:"column:name" json:"name"`
	Avatar     string           `gorm:"column:avatar" json:"avatar"`
	Email      string           `gorm:"column:email" json:"email"`
	Password   string           `gorm:"column:password" json:"password"`
	Role       string           `gorm:"column:role" json:"role"`
	IsVerified bool             `gorm:"column:is_verified;default:false" json:"is_verified"`
	CreatedAt  time.Time        `gorm:"column:created_at;type:TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time        `gorm:"column:updated_at;type:TIMESTAMP" json:"updated_at"`
	DeletedAt  *time.Time       `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
	UserDetail UserDetailEntity `gorm:"foreignKey:UserID"`
}

type UserDetailEntity struct {
	ID        int        `gorm:"column:id;primaryKey" json:"id"`
	UserID    int        `gorm:"column:user_id;index" json:"user_id"`
	Address   *string    `gorm:"column:address" json:"address"`
	Phone     *string    `gorm:"column:phone" json:"phone"`
	Job       *string    `gorm:"column:job" json:"job"`
	CreatedAt time.Time  `gorm:"column:created_at;type:TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

type TokenEntity struct {
	ID           int        `gorm:"primaryKey;autoIncrement" json:"id" `
	UserID       int        `gorm:"index;unique" json:"user_id" `
	User         UserEntity `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"user" `
	Token        string     `gorm:"column:token;type:varchar(255)" json:"token"`
	ExpiredToken int64      `gorm:"column:expired_token;type:bigint" json:"expired_token" `
}

func (UserEntity) TableName() string {
	return "users"
}
func (UserDetailEntity) TableName() string {
	return "user_details"
}
func (TokenEntity) TableName() string {
	return "token"
}
