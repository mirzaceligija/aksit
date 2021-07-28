package repository

import (
	"github.com/mirzaceligija/askit/entity"
	"gorm.io/gorm"
)

type AnswerRepository interface {
	InsertAnswer(q entity.Answer) entity.Answer
	UpdateAnswer(q entity.Answer) entity.Answer
	DeleteAnswer(q entity.Answer)
	FindAnswerByID(answerID uint64) entity.Answer
}

type answerConnection struct {
	connection *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) AnswerRepository {
	return &answerConnection{
		connection: db,
	}
}

func (db *answerConnection) InsertAnswer(a entity.Answer) entity.Answer {
	db.connection.Save(&a)
	db.connection.Preload("User").Find(&a)
	return a
}

func (db *answerConnection) UpdateAnswer(a entity.Answer) entity.Answer {
	db.connection.Save(&a)
	db.connection.Preload("User").Find(&a)
	return a
}

func (db *answerConnection) DeleteAnswer(a entity.Answer) {
	db.connection.Delete(&a)
}

func (db *answerConnection) FindAnswerByID(answerID uint64) entity.Answer {
	var answer entity.Answer
	db.connection.Find(&answer, answerID)
	return answer
}
