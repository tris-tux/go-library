package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/tris-tux/go-library/backend/repository"
	"github.com/tris-tux/go-library/backend/schema"
	"golang.org/x/crypto/bcrypt"
)

//AuthService is a contract about something that this service can do
type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateVisitor(visitor schema.RegisterDTO) schema.User
	FindByEmail(email string) schema.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	visitorRepository repository.VisitorRepository
}

//NewAuthService creates a new instance of AuthService
func NewAuthService(visitorRep repository.VisitorRepository) AuthService {
	return &authService{
		visitorRepository: visitorRep,
	}
}

func (s *authService) VerifyCredential(email string, password string) interface{} {
	res := s.visitorRepository.VerifyCredential(email, password)
	if v, ok := res.(schema.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (s *authService) CreateVisitor(visitor schema.RegisterDTO) schema.User {
	visitorToCreate := schema.User{}
	err := smapping.FillStruct(&visitorToCreate, smapping.MapFields(&visitor))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := s.visitorRepository.InsertVisitor(visitorToCreate)
	return res
}

func (s *authService) FindByEmail(email string) schema.User {
	return s.visitorRepository.FindByEmail(email)
}

func (s *authService) IsDuplicateEmail(email string) bool {
	res := s.visitorRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
