package db

import (
	"fmt"
	"notes-goo-api/internal/config"
)

type Dbs map[string]interface{}

// InitDb:
func InitDb(c *config.Config) (Dbs, error) {

	dbs := make(Dbs)

	for _, cDetail := range c.Dbs {
		fmt.Printf("### Connect DB[%s] ###\n", cDetail.Name)
		if cDetail.Type == "Gorm" {
			// gorm, err := NewGorm(cDetail)
			// if err != nil {
			// 	return nil, err
			// }
			// dbs[cDetail.Name] = gorm
		} else if cDetail.Type == "Mongo" {
			mongo, err := NewMongo(cDetail)
			if err != nil {
				return nil, err
			}
			dbs[cDetail.Name] = mongo
		}
	}

	return dbs, nil
}
