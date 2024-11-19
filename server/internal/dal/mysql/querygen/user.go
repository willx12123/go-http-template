package querygen

import (
	"gorm.io/gen"
)

type UserQuery interface {
	// FindByEmail
	//
	// SELECT * FROM @@table WHERE email = @email LIMIT 1
	FindByEmail(email string) (*gen.T, error)
}
