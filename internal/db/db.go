package db

import (
	"fmt"
	"notes-goo-api/internal/config"
)

type Dbs map[string]interface{}

// InitDb: Initial database for routes
func InitDb(c *config.Config) (map[string]interface{}, error) {
	dbs := make(map[string]interface{})
	if c.Dbs != nil {
		for _, cDetail := range c.Dbs {
			fmt.Printf("[DB-Connect] %s\n", cDetail.Name)
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
	}
	return dbs, nil
}
