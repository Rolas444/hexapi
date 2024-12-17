package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var client *gorm.DB

func ConnectClient(uri string) (*gorm.DB, error) {
	// if client != nil {
	// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// 	return client, ctx, cancel, nil
	// }

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// var err error
	client, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		// cancel()
		return nil, err
	}

	return client, nil
}
