package database

import "github.com/AirouTUS/shinkan-server/internal/model"

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

type Circle struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	About       string `db:"about"`
	CatchCopy   string `db:"catch_copy"`
	Description string `db:"description"`
	CategoryID  int    `db:"circle_category_id"`
	Email       string `db:"email"`
	Twitter     string `db:"twitter"`
	URL         string `db:"url"`
}

func (m Circle) circle() *model.Circle {
	var result model.Circle
	result = model.Circle(m)
	return &result
}

type CirclesCircleTypesList []CirclesCircleTypes

type CirclesCircleTypes struct {
	CircleTypeID int    `db:"circle_type_id"`
	Name         string `db:"name"`
}

func (m CirclesCircleTypesList) circlesCircleTypes() []*model.CirclesCircleTypes {
	result := make([]*model.CirclesCircleTypes, 0, len(m))
	for _, v := range m {
		content := model.CirclesCircleTypes{
			CircleTypeID: v.CircleTypeID,
			Name:         v.Name,
		}
		result = append(result, &content)
	}
	return result
}