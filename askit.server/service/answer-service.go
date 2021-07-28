package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/mirzaceligija/askit/dto"
	"github.com/mirzaceligija/askit/entity"
	"github.com/mirzaceligija/askit/repository"
)

type AnswerService interface {
	Insert(a dto.AnswerCreateDTO) entity.Answer
	Update(a dto.AnswerUpdateDTO) entity.Answer
	Delete(a entity.Answer)
	FindByID(answerID uint64) entity.Answer
	IsAllowedToEdit(userID string, answerID uint64) bool
}

type answerService struct {
	answerRepository repository.AnswerRepository
}

func NewAnswerService(answerRepo repository.AnswerRepository) AnswerService {
	return &answerService{
		answerRepository: answerRepo,
	}
}

func (service *answerService) Insert(a dto.AnswerCreateDTO) entity.Answer {
	answer := entity.Answer{}
	log.Println("ovo je iz servisa", a)
	err := smapping.FillStruct(&answer, smapping.MapFields(&a))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.answerRepository.InsertAnswer(answer)
	return res
}

func (service *answerService) Update(a dto.AnswerUpdateDTO) entity.Answer {
	answer := entity.Answer{}
	err := smapping.FillStruct(&answer, smapping.MapFields(&a))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.answerRepository.UpdateAnswer(answer)
	return res
}

func (service *answerService) Delete(a entity.Answer) {
	service.answerRepository.DeleteAnswer(a)
}

func (service *answerService) FindByID(answerID uint64) entity.Answer {
	return service.answerRepository.FindAnswerByID(answerID)
}

func (service *answerService) IsAllowedToEdit(userID string, answerID uint64) bool {
	a := service.answerRepository.FindAnswerByID(answerID)
	id := fmt.Sprintf("%v", a.UserID)
	return userID == id
}
