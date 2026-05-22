package timefmt

import "time"

const DisplayLayout = "2006-01-02 15:04"

func Display(t time.Time) string {
	return t.In(time.Local).Format(DisplayLayout)
}

func Today() string {
	return time.Now().Format("2006-01-02")
}

func DefaultTitle(date, mood string) string {
	if len(date) >= 5 {
		return date[5:] + " " + moodLabel(mood) + "日记"
	}
	return "今日日记"
}

func moodLabel(mood string) string {
	switch mood {
	case "great":
		return "很好"
	case "good":
		return "不错"
	case "low":
		return "低落"
	default:
		return "平静"
	}
}
