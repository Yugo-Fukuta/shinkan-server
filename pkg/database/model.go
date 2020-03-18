package database

import (
	"github.com/AirouTUS/shinkan-server/pkg/model"
)

type CategoryList []Category

type Category struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func (m CategoryList) category() []*model.Category {
	result := make([]*model.Category, 0, len(m))
	for _, v := range m {
		content := model.Category{
			ID:   v.ID,
			Name: v.Name,
		}
		result = append(result, &content)
	}
	return result
}

type CircleList []Circle

func (m CircleList) circles() []*model.Circle {
	result := make([]*model.Circle, 0, len(m))
	for _, v := range m {
		content := model.Circle(v)
		result = append(result, &content)
	}
	return result
}

type Circle struct {
	ID          int     `db:"id"`
	Name        string  `db:"name"`
	About       string  `db:"about"`
	CatchCopy   string  `db:"catch_copy"`
	Description string  `db:"description"`
	CategoryID  int     `db:"circle_category_id"`
	Email       string  `db:"email"`
	Twitter     string  `db:"twitter"`
	URL         string  `db:"url"`
	EyeCatch    string  `db:"eyecatch"`
	TypeID      *int    `db:"type_id"`
	TypeName    *string `db:"type_name"`
}

type GetCircle struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	About       string `db:"about"`
	CatchCopy   string `db:"catch_copy"`
	Description string `db:"description"`
	CategoryID  int    `db:"circle_category_id"`
	Email       string `db:"email"`
	Twitter     string `db:"twitter"`
	URL         string `db:"url"`
	EyeCatch    string `db:"eyecatch"`
}

type CircleType struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type CircleImage struct {
	URL string `db:"url"`
}

func (m GetCircle) circle(ct []CircleType, ci []CircleImage) *model.GetCircle {
	var result model.GetCircle
	result.ID = m.ID
	result.Name = m.Name
	result.About = m.About
	result.CatchCopy = m.CatchCopy
	result.Description = m.Description
	result.CategoryID = m.CategoryID
	result.Email = m.Email
	result.Twitter = m.Twitter
	result.URL = m.URL
	result.EyeCatch = m.EyeCatch

	result.Types = make([]model.CircleType, 0, len(ct))
	result.Images = make([]model.CircleImages, 0, len(ci))
	for _, v := range ct {
		content := model.CircleType{
			ID:   v.ID,
			Name: v.Name,
		}
		result.Types = append(result.Types, content)
	}
	for _, v := range ci {
		content := model.CircleImages{
			URL: v.URL,
		}
		result.Images = append(result.Images, content)
	}
	return &result
}
