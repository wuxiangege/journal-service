package convert

import (
	"journal-service/internal/model"
	"journal-service/internal/pkg/timefmt"
	"journal-service/internal/types"
)

func JournalPostToType(p *model.JournalPost) types.JournalPost {
	tags := []string(p.Tags)
	if tags == nil {
		tags = []string{}
	}
	return types.JournalPost{
		Id:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		Date:      p.Date,
		Mood:      p.Mood,
		Tags:      tags,
		Pinned:    p.Pinned,
		UpdatedAt: timefmt.Display(p.UpdatedAt),
	}
}

func JournalPostsToTypes(list []model.JournalPost) []types.JournalPost {
	out := make([]types.JournalPost, 0, len(list))
	for i := range list {
		out = append(out, JournalPostToType(&list[i]))
	}
	return out
}
