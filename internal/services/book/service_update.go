package book

import (
	"book/internal/pkg/core"
	"book/internal/repository/mysql"
	"book/internal/repository/mysql/book_info"
	"time"
)

type UpdateData struct {
	Id           int32     `json:"id"`
	Name         string    `json:"name" `
	Author       string    `json:"author"`
	Publish      string    `json:"publish"`
	ISBN         string    `json:"ISBN"`
	Introduction string    `json:"introduction"`
	Language     string    `json:"language"`
	Price        float64   `json:"price"`
	Pubdate      time.Time `json:"pubdate"`
	ClassId      int32     `json:"classId"`
	Pressmark    int32     `json:"pressmark"`
}

func (s *service) Update(ctx core.Context, data *UpdateData) (err error) {
	update := map[string]interface{}{
		"name":         data.Name,
		"author":       data.Author,
		"publish":      data.Publish,
		"ISBN":         data.ISBN,
		"introduction": data.Introduction,
		"language":     data.Language,
		"price":        data.Price,
		"pubdate":      data.Pubdate,
		"class_id":     data.ClassId,
		"pressmark":    data.Pressmark,
		"updated_user": ctx.SessionUserInfo().UserName,
	}
	qb := book_info.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, data.Id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), update)
	if err != nil {
		return err
	}
	return
}
