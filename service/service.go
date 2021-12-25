package service

import (
	"errors"
	"jawaban/constants/en"
	"log"
)

type service struct {
	repository Repository
}


func NewService ( repo Repository ) Service {
	return &service{
		repo,
	}
}


func (_u service) InsertArea(param1 int64, param2 int64, _type string) (err error) {
	err = _u.repository.InsertArea(param1, param2, "persegi")
	if err != nil {
		log.Println(err.Error())
		err = errors.New(en.ErrorDatabase)
		return err
	}
	return nil
}