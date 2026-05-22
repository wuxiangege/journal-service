package repo

import (
	"fmt"
	"strings"
	"time"

	"journal-service/internal/model"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) DB() *gorm.DB {
	return r.db
}

func (r *Repo) FindUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repo) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

type ListJournalPostsFilter struct {
	UserID     uint64
	Q          string
	Tag        string
	Mood       string
	DateFilter string
	Page       int
	PageSize   int
}

func (r *Repo) ListJournalPosts(f ListJournalPostsFilter) ([]model.JournalPost, int64, error) {
	q := r.db.Model(&model.JournalPost{}).Where("user_id = ?", f.UserID)

	if f.Mood != "" && f.Mood != "all" {
		q = q.Where("mood = ?", f.Mood)
	}
	if f.Tag != "" {
		q = q.Where("JSON_CONTAINS(tags, JSON_QUOTE(?))", f.Tag)
	}
	if f.DateFilter == "week" || f.DateFilter == "month" {
		days := 7
		if f.DateFilter == "month" {
			days = 30
		}
		from := time.Now().AddDate(0, 0, -(days - 1)).Format("2006-01-02")
		q = q.Where("date >= ?", from)
	}
	if kw := strings.TrimSpace(strings.ToLower(f.Q)); kw != "" {
		like := "%" + kw + "%"
		q = q.Where(
			"LOWER(title) LIKE ? OR LOWER(content) LIKE ? OR LOWER(mood) LIKE ? OR CAST(tags AS CHAR) LIKE ?",
			like, like, like, like,
		)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	page := f.Page
	if page < 1 {
		page = 1
	}
	pageSize := f.PageSize
	if pageSize < 1 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	var list []model.JournalPost
	err := q.Order("pinned DESC, date DESC, updated_at DESC").
		Offset(offset).Limit(pageSize).Find(&list).Error
	return list, total, err
}

func (r *Repo) GetJournalPost(userID uint64, id string) (*model.JournalPost, error) {
	var post model.JournalPost
	err := r.db.Where("user_id = ? AND id = ?", userID, id).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *Repo) CreateJournalPost(post *model.JournalPost) error {
	return r.db.Create(post).Error
}

func (r *Repo) SaveJournalPost(post *model.JournalPost) error {
	return r.db.Save(post).Error
}

func (r *Repo) DeleteJournalPost(userID uint64, id string) error {
	res := r.db.Where("user_id = ? AND id = ?", userID, id).Delete(&model.JournalPost{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *Repo) ListAllJournalPosts(userID uint64) ([]model.JournalPost, error) {
	var list []model.JournalPost
	err := r.db.Where("user_id = ?", userID).Find(&list).Error
	return list, err
}

func (r *Repo) CountJournalPostsByMonth(userID uint64, prefix string) (int64, error) {
	var count int64
	err := r.db.Model(&model.JournalPost{}).
		Where("user_id = ? AND date LIKE ?", userID, prefix+"%").
		Count(&count).Error
	return count, err
}

func MonthPrefix() string {
	now := time.Now()
	return fmt.Sprintf("%d-%02d", now.Year(), int(now.Month()))
}
