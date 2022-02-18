package models

type ReviewModel struct {
	Id int
	Title string
}
func MockReivews() []*ReviewModel{
	return []*ReviewModel{
		{Id:2000,Title:"测试评论1"},
		{Id:2002,Title:"测试评论2"},
	}
}