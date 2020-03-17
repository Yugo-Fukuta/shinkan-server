package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/AirouTUS/shinkan-server/pkg/model"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// Driver名
const (
	driverName = "mysql"
	engineName = "InnoDB"

	tableCircles          = "circles"
	tableCircleCategories = "circle_categories"
	tableCircleTypes      = "circle_types"
	tableCircleImages     = "circle_images"

	// 中間テーブル
	tableCirclesCircleTypes = "circles_circle_types"
)

type ShinkanDatabase struct {
	Map *gorp.DbMap
}

func NewDatabase(user, password, host, port, database string) *ShinkanDatabase {
	db, err := sql.Open(driverName,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database))
	if err != nil {
		log.Fatal(err)
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: engineName, Encoding: "UTF8"}}
	return &ShinkanDatabase{Map: dbMap}
}

func (db *ShinkanDatabase) ListCategory(input ListCategoryInput) ([]*model.Category, error) {
	if err := input.validate(); err != nil {
		return nil, err
	}
	var m CategoryList
	_, err := db.Map.Select(&m, fmt.Sprintf("SELECT * FROM %s ORDER BY id ASC", tableCircleCategories))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return m.category(), nil
}

func (db *ShinkanDatabase) GetCircle(input GetCircleInput) ([]*model.Circle, error) {
	if err := input.validate(); err != nil {
		return nil, errors.WithStack(err)
	}
	var m CircleList
	_, err := db.Map.Select(&m, fmt.Sprintf(
		`SELECT 
				circles.id, 
				circles.name, 
				circles.about,
				circles.catch_copy,
				circles.description,
				circles.circle_category_id,
				circles.email,
				circles.twitter,
				circles.url,
				circles.eyecatch,
				circle_types.id AS type_id,
				circle_types.name AS type_name
			FROM 
				%s 
			LEFT JOIN 
				%s
			ON
				%s.circle_id = %s.id
			LEFT JOIN
				%s
			ON 
				%s.circle_type_id = %s.id
			WHERE 
				%s.id = ?
			ORDER BY
				%s.id
			ASC`,
		tableCircles, tableCirclesCircleTypes, tableCirclesCircleTypes, tableCircles, tableCircleTypes, tableCirclesCircleTypes, tableCircleTypes, tableCircles, tableCircles),
		input.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return m.circles(), nil
}

func (db *ShinkanDatabase) ListCircle(input ListCircleInput) ([]*model.Circle, error) {
	if err := input.validate(); err != nil {
		return nil, errors.WithStack(err)
	}

	var m CircleList

	{
		if len(input.CategoryID) > 0 {
			valueStrings := make([]string, 0, len(input.CategoryID))
			valueArgs := make([]interface{}, 0, len(input.CategoryID))
			for _, v := range input.CategoryID {
				valueStrings = append(valueStrings, "circle_category_id = ?")
				valueArgs = append(valueArgs, v)
			}

			_, err := db.Map.Select(&m, fmt.Sprintf(
				`SELECT 
						circles.id, 
						circles.name, 
						circles.about,
						circles.catch_copy,
						circles.description,
						circles.circle_category_id,
						circles.email,
						circles.twitter,
						circles.url,
						circles.eyecatch,
						circle_types.id AS type_id,
						circle_types.name AS type_name
				FROM 
						%s 
				LEFT JOIN 
						%s
				ON
						%s.circle_id = %s.id
				LEFT JOIN
						%s
				ON 
						%s.circle_type_id = %s.id
				WHERE 
						%s
				ORDER BY
						%s.id
				ASC`,
				tableCircles, tableCirclesCircleTypes, tableCirclesCircleTypes, tableCircles, tableCircleTypes, tableCirclesCircleTypes, tableCircleTypes,
				strings.Join(valueStrings, " OR "), tableCircles), valueArgs...)
			if err != nil {
				return nil, errors.WithStack(err)
			}
		} else {
			_, err := db.Map.Select(&m, fmt.Sprintf(
				`SELECT 
						circles.id, 
						circles.name, 
						circles.about,
						circles.catch_copy,
						circles.description,
						circles.circle_category_id,
						circles.email,
						circles.twitter,
						circles.url,
						circles.eyecatch,
						circle_types.id AS type_id,
						circle_types.name AS type_name
				FROM 
						%s 
				LEFT JOIN 
						%s
				ON
						%s.circle_id = %s.id
				LEFT JOIN
						%s
				ON 
						%s.circle_type_id = %s.id
				ORDER BY
						%s.id
				ASC`,
				tableCircles, tableCirclesCircleTypes, tableCirclesCircleTypes, tableCircles, tableCircleTypes, tableCirclesCircleTypes, tableCircleTypes, tableCircles))
			if err != nil {
				return nil, errors.WithStack(err)
			}

		}
	}

	return m.circles(), nil
}
