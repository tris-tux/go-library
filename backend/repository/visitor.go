package repository

import (
	"log"

	"github.com/tris-tux/go-library/backend/schema"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//VisitorRepository is contract what visitorRepository can do to db
type VisitorRepository interface {
	InsertVisitor(visitor schema.Visitor) schema.Visitor
	UpdateVisitor(visitor schema.Visitor) schema.Visitor
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) schema.Visitor
	ProfileVisitor(visitorID string) schema.Visitor
}

type visitorConnection struct {
	connection *gorm.DB
}

//NewVisitorRepository is creates a new instance of VisitorRepository
func NewVisitorRepository(db *gorm.DB) VisitorRepository {
	return &visitorConnection{
		connection: db,
	}
}

func (db *visitorConnection) InsertVisitor(visitor schema.Visitor) schema.Visitor {
	visitor.Password = hashAndSalt([]byte(visitor.Password))
	db.connection.Save(&visitor)
	return visitor
}

func (db *visitorConnection) UpdateVisitor(visitor schema.Visitor) schema.Visitor {
	if visitor.Password != "" {
		visitor.Password = hashAndSalt([]byte(visitor.Password))
	} else {
		var tempVisitor schema.Visitor
		db.connection.Find(&tempVisitor, visitor.NoIdentitas)
		visitor.Password = tempVisitor.Password
	}

	db.connection.Save(&visitor)
	return visitor
}

func (db *visitorConnection) VerifyCredential(email string, password string) interface{} {
	var visitor schema.Visitor
	res := db.connection.Where("email = ?", email).Take(&visitor)
	if res.Error == nil {
		return visitor
	}
	return nil
}

func (db *visitorConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var visitor schema.Visitor
	return db.connection.Where("email = ?", email).Take(&visitor)
}

func (db *visitorConnection) FindByEmail(email string) schema.Visitor {
	var visitor schema.Visitor
	db.connection.Where("email = ?", email).Take(&visitor)
	return visitor
}

func (db *visitorConnection) ProfileVisitor(visitorID string) schema.Visitor {
	var visitor schema.Visitor
	db.connection.Preload("Books").Preload("Books.Visitor").Find(&visitor, visitorID)
	return visitor
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
