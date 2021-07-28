package repository

import (
	"log"

	"github.com/mirzaceligija/askit/entity"
	"gorm.io/gorm"
)

type QuestionRepository interface {
	InsertQuestion(q entity.Question) entity.Question
	UpdateQuestion(q entity.Question) entity.Question
	DeleteQuestion(q entity.Question)
	AllQuestion(page uint64, orderBy string) []entity.Question
	FindQuestionByID(questionID uint64) entity.Question
	FindQuestionByUserID(userID uint64) []entity.Question
}

type questionConnection struct {
	connection *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) QuestionRepository {
	return &questionConnection{
		connection: db,
	}
}

func (db *questionConnection) InsertQuestion(q entity.Question) entity.Question {
	log.Println("ovo je iz repositorya", q)
	db.connection.Save(&q)
	db.connection.Preload("User").Preload("AnswerS").Find(&q)
	return q
}

func (db *questionConnection) UpdateQuestion(q entity.Question) entity.Question {
	db.connection.Save(&q)
	db.connection.Preload("User").Preload("Answers").Find(&q)
	return q
}

func (db *questionConnection) DeleteQuestion(q entity.Question) {
	db.connection.Delete(&q)
}

func (db *questionConnection) FindQuestionByID(questionID uint64) entity.Question {
	var question entity.Question
	db.connection.Preload("User").Preload("Answers").Preload("Answers.User").Find(&question, questionID)
	return question
}

func (db *questionConnection) AllQuestion(page uint64, orderBy string) []entity.Question {
	var offset = (page - 1) * 20
	var questions []entity.Question
	db.connection.Limit(20).Offset(int(offset)).Preload("User").Preload("Answers").Order(orderBy + " desc").Find(&questions)
	return questions
}

func (db *questionConnection) FindQuestionByUserID(userID uint64) []entity.Question {
	log.Println("Ovo je user id", userID)
	var questions []entity.Question
	db.connection.Preload("User").Where("user_id", userID).Find(&questions)
	return questions
}
