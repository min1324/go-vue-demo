package model

import (
	"log"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	// 创建uuid
	//uuid "github.com/satori/go.uuid"
)

type Post struct {
	ID         uuid.UUID `json:"id" gorm:"type:char(36);primary key"`
	UserId     uint      `json:"user_id" gorm:"not null"`
	CategoryId uint      `json:"category_id" gorm:"not null"` // gorm:"foreignkey:Category"`
	Category   *Category
	Tittle     string `json:"tittle" gorm:"type:varchar(20);not null"`
	HeadImg    string `json:"head_img"`
	Content    string `json:"content" gorm:"type:text;not null"`
	CreatedAt  Time   `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt  Time   `json:"updated_at" gorm:"type:timestamp"`
	DeletedAt  Time   `json:"-" gorm:"index" `
}

// 自动外键
// gorm 外键规范为 名字+Id.
// CategoryId,Category
// CategoryId 是 Category 的外键
// 也可以手动指定外键 `gorm:"foreignkey:name"`

var u1 = uuid.Must(uuid.NewV4())

func (p *Post) BeforeCreate(tx *gorm.DB) error {
	log.Println("BeforeCreate")
	uuid, err := uuid.NewV4()
	p.ID = uuid
	log.Println("BeforeCreate", uuid.String())
	return err
}
