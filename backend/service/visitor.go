package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/tris-tux/go-library/backend/repository"
	"github.com/tris-tux/go-library/backend/schema"
)

//VisitorService is a contract.....
type VisitorService interface {
	Update(visitor schema.VisitorUpdateDTO) schema.User
	Profile(visitorID string) schema.User
}

type visitorService struct {
	visitorRepository repository.VisitorRepository
}

//NewVisitorService creates a new instance of VisitorService
func NewVisitorService(visitorRepo repository.VisitorRepository) VisitorService {
	return &visitorService{
		visitorRepository: visitorRepo,
	}
}

func (s *visitorService) Update(visitor schema.VisitorUpdateDTO) schema.User {
	visitorToUpdate := schema.User{}
	err := smapping.FillStruct(&visitorToUpdate, smapping.MapFields(&visitor))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedVisitor := s.visitorRepository.UpdateVisitor(visitorToUpdate)
	return updatedVisitor
}

func (s *visitorService) Profile(visitorID string) schema.User {
	return s.visitorRepository.ProfileVisitor(visitorID)
}
