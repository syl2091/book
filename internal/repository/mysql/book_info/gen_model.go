package book_info

import "time"

// BookInfo
//
//go:generate gormgen -structs BookInfo -input .
type BookInfo struct {
	Id           int32     //
	Name         string    //
	Author       string    //
	Publish      string    //
	ISBN         string    //
	Introduction string    //
	Language     string    //
	Price        float64   //
	Pubdate      time.Time `gorm:"time"`
	ClassId      int32     //
	Pressmark    int32     //
	State        int32     //
	IsDeleted    int32     // 是否删除 1:是  -1:否
	CreatedAt    time.Time `gorm:"time"` // 创建时间
	CreatedUser  string    // 创建人
	UpdatedAt    time.Time `gorm:"time"` // 更新时间
	UpdatedUser  string    // 更新人
}
