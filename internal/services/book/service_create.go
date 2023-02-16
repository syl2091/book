package book

import (
	"book/internal/pkg/core"
	"book/internal/repository/mysql/book_info"
	"time"
)

type CreateBookData struct {
	Name         string    `json:"name"`         //
	Author       string    `json:"author"`       //
	Publish      string    `json:"publish"`      //
	ISBN         string    `json:"ISBN"`         //
	Introduction string    `json:"introduction"` //
	Language     string    `json:"language"`     //
	Price        float64   `json:"price"`        //
	Pubdate      time.Time `json:"pubdate"`
	ClassId      int32     `json:"classId"`   //
	Pressmark    int32     `json:"pressmark"` //
}

func (s *service) Create(ctx core.Context, data *CreateBookData) (id int32, err error) {
	model := book_info.NewModel()
	model.Name = data.Name
	model.Author = data.Author
	model.Publish = data.Publish
	model.ISBN = data.ISBN
	model.Introduction = data.Introduction
	model.Language = data.Language
	model.Price = data.Price
	model.Pubdate = data.Pubdate
	model.ClassId = data.ClassId
	model.Pressmark = data.Pressmark
	model.CreatedUser = ctx.SessionUserInfo().UserName
	model.State = 1
	model.IsDeleted = -1
	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
