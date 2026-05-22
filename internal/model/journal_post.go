package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type StringSlice []string

func (s StringSlice) Value() (driver.Value, error) {
	if s == nil {
		return "[]", nil
	}
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

func (s *StringSlice) Scan(value any) error {
	if value == nil {
		*s = []string{}
		return nil
	}
	var data []byte
	switch v := value.(type) {
	case []byte:
		data = v
	case string:
		data = []byte(v)
	default:
		return errors.New("unsupported tags column type")
	}
	if len(data) == 0 {
		*s = []string{}
		return nil
	}
	return json.Unmarshal(data, s)
}

type JournalPost struct {
	ID        string      `gorm:"primaryKey;size:36;comment:日记ID（UUID）"`
	UserID    uint64      `gorm:"index:idx_user_date;not null;comment:所属用户ID"`
	Title     string      `gorm:"size:200;not null;comment:标题"`
	Content   string      `gorm:"type:longtext;comment:正文（Markdown）"`
	Date      string      `gorm:"size:10;index:idx_user_date;not null;comment:日记日期（YYYY-MM-DD）"`
	Mood      string      `gorm:"size:16;not null;comment:心情（great/good/calm/low）"`
	Tags      StringSlice `gorm:"type:json;comment:标签列表（JSON数组）"`
	Pinned    bool        `gorm:"not null;default:false;comment:是否置顶"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime;comment:最后编辑时间"`
	CreatedAt time.Time   `gorm:"autoCreateTime;comment:创建时间"`
}

func (JournalPost) TableName() string {
	return "journal_posts"
}
