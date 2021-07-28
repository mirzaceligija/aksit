package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/mirzaceligija/askit/dto"
	"github.com/mirzaceligija/askit/entity"
	"github.com/mirzaceligija/askit/repository"
)

type QuestionService interface {
	Insert(q dto.QuestionCreateDTO) entity.Question
	Update(q dto.QuestionUpdateDTO) entity.Question
	Delete(q entity.Question)
	All(page uint64, orderBy string) []entity.Question
	FindByID(questionID uint64) entity.Question
	FindByUserID(questionID uint64) []entity.Question
	IsAllowedToEdit(userID string, questionID uint64) bool
}

type questionService struct {
	questionRepository repository.QuestionRepository
}

func NewQuestionService(questionRepo repository.QuestionRepository) QuestionService {
	return &questionService{
		questionRepository: questionRepo,
	}
}

func (service *questionService) Insert(q dto.QuestionCreateDTO) entity.Question {
	question := entity.Question{}
	err := smapping.FillStruct(&question, smapping.MapFields(&q))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	log.Println("Ovo je mapped question", question)
	res := service.questionRepository.InsertQuestion(question)
	return res
}

func (service *questionService) Update(q dto.QuestionUpdateDTO) entity.Question {
	question := entity.Question{}
	log.Println("ovo je service obj", q)
	err := smapping.FillStruct(&question, smapping.MapFields(&q))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	log.Println("ovo je MAPPED service obj", question)
	res := service.questionRepository.UpdateQuestion(question)
	return res
}

func (service *questionService) Delete(q entity.Question) {
	service.questionRepository.DeleteQuestion(q)
}

func (service *questionService) All(page uint64, orderBy string) []entity.Question {
	return service.questionRepository.AllQuestion(page, orderBy)
}

func (service *questionService) FindByID(questionID uint64) entity.Question {
	return service.questionRepository.FindQuestionByID(questionID)
}

func (service *questionService) FindByUserID(userID uint64) []entity.Question {
	return service.questionRepository.FindQuestionByUserID(userID)
}

func (service *questionService) IsAllowedToEdit(userID string, questionID uint64) bool {
	q := service.questionRepository.FindQuestionByID(questionID)
	id := fmt.Sprintf("%v", q.UserID)
	return userID == id
}
