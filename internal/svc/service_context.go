// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"log"
	"time"

	"journal-service/internal/config"
	"journal-service/internal/middleware"
	"journal-service/internal/model"
	"journal-service/internal/repo"

	"github.com/zeromicro/go-zero/rest"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config config.Config
	Cors   rest.Middleware
	DB     *gorm.DB
	Repo   *repo.Repo
}

func NewServiceContext(c config.Config) *ServiceContext {
	logLevel := logger.Warn
	if c.MySQL.LogSQL {
		logLevel = logger.Info
	}
	db, err := gorm.Open(mysql.Open(c.MySQL.DataSource), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		log.Fatalf("mysql connect failed: %v", err)
	}
	if err := db.AutoMigrate(&model.User{}, &model.JournalPost{}); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}

	r := repo.New(db)
	seedDefaultUser(r)
	seedDemoJournalPosts(r)

	return &ServiceContext{
		Config: c,
		Cors:   middleware.NewCorsMiddleware(c.Cors.AllowOrigins).Handle,
		DB:     db,
		Repo:   r,
	}
}

func seedDefaultUser(r *repo.Repo) {
	const username = "871240671@qq.com"
	const password = "123456"
	_, err := r.FindUserByUsername(username)
	if err == nil {
		return
	}
	if err != gorm.ErrRecordNotFound {
		log.Printf("seed user check: %v", err)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("seed user hash: %v", err)
		return
	}
	if err := r.CreateUser(&model.User{
		Username:     username,
		PasswordHash: string(hash),
	}); err != nil {
		log.Printf("seed user create: %v", err)
	}
}

func seedDemoJournalPosts(r *repo.Repo) {
	user, err := r.FindUserByUsername("871240671@qq.com")
	if err != nil {
		return
	}
	var count int64
	if err := r.DB().Model(&model.JournalPost{}).Where("user_id = ?", user.ID).Count(&count).Error; err != nil || count > 0 {
		return
	}

	demos := []model.JournalPost{
		{ID: "e1", UserID: user.ID, Title: "五月的小雨", Content: "窗外淅淅沥沥下了一整天。泡了杯茶，把书桌整理干净，突然觉得很适合写点什么。\n\n记下今天完成的几件小事，也算是对自己的交代。", Date: "2026-05-18", Mood: "calm", Tags: model.StringSlice{"生活", "随想"}, Pinned: true},
		{ID: "e2", UserID: user.ID, Title: "项目里程碑前的备忘", Content: "上线前 checklist：\n1. 回归核心流程\n2. 文案与权限再对一遍\n3. 预留回滚窗口\n\n心态放稳，一步一步来。", Date: "2026-05-16", Mood: "good", Tags: model.StringSlice{"工作"}, Pinned: false},
		{ID: "e3", UserID: user.ID, Title: "周末想去的书店", Content: "朋友推荐了一家独立书店，据说二楼有安静的阅读角。下周如果天气好，就骑车过去看看。", Date: "2026-05-12", Mood: "great", Tags: model.StringSlice{"生活", "旅行"}, Pinned: false},
		{ID: "e4", UserID: user.ID, Title: "晨跑记录", Content: "绕湖两圈，心率稳定。早餐吃了燕麦和香蕉。", Date: "2026-05-10", Mood: "good", Tags: model.StringSlice{"生活"}, Pinned: false},
		{ID: "e5", UserID: user.ID, Title: "读了一半的小说", Content: "主角终于踏上旅程，节奏偏慢但文笔舒服，准备周末读完。", Date: "2026-05-08", Mood: "calm", Tags: model.StringSlice{"随想"}, Pinned: false},
		{ID: "e6", UserID: user.ID, Title: "周报草稿", Content: "本周完成：接口联调、文档更新。下周：压测与灰度。", Date: "2026-05-06", Mood: "good", Tags: model.StringSlice{"工作"}, Pinned: false},
		{ID: "e7", UserID: user.ID, Title: "夜雨听风", Content: "没开灯，只听雨打在遮阳棚上的声音，意外地很放松。", Date: "2026-05-04", Mood: "calm", Tags: model.StringSlice{"生活", "随想"}, Pinned: false},
	}
	for i := range demos {
		t, _ := time.Parse("2006-01-02 15:04", demos[i].Date+" 12:00")
		demos[i].UpdatedAt = t
		demos[i].CreatedAt = t
		if err := r.CreateJournalPost(&demos[i]); err != nil {
			log.Printf("seed journal post %s: %v", demos[i].ID, err)
		}
	}
}
