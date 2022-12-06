package repository

import (
	"time"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/class"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name          string
	StartDate     time.Time
	GraduatedDate time.Time
	UserId        uint
	// Mentee []Mentee
}

//from core to model

func fromCore(dataCore class.Core) Class {
	userGorm := Class{
		Name:          dataCore.Name,
		StartDate:     dataCore.StartDate,
		GraduatedDate: dataCore.GraduatedDate,
		UserId:        dataCore.UserId,
	}
	return userGorm
}

//from model to core

func (dataModel *Class) toCore() class.Core {
	return class.Core{
		ID:            dataModel.ID,
		Name:          dataModel.Name,
		StartDate:     dataModel.StartDate,
		GraduatedDate: dataModel.GraduatedDate,
		UserId:        dataModel.UserId,
		CreatedAt:     dataModel.CreatedAt,
		UpdatedAt:     dataModel.UpdatedAt,
	}
}

func toCoreList(dataModel []Class) []class.Core {
	var dataCore []class.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
