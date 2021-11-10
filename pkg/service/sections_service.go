package service

import (
	"github.com/6156-DonaldDuck/sections/pkg/db"
	"github.com/6156-DonaldDuck/sections/pkg/model"
	log "github.com/sirupsen/logrus"
)

func ListSections() ([]model.Section, error) {
	var sections []model.Section
	result := db.DbConn.Find(&sections)
	if result.Error != nil {
		log.Errorf("[service.ListSections] error occurred while listing sections, err=%v\n", result.Error)
	} else {
		log.Infof("[service.ListSections] successfully listed sections, rows affected=%v\n", result.RowsAffected)
	}
	return sections, result.Error
}

func GetSectionById(sectionId uint) (model.Section, error) {
	section := model.Section{}
	result := db.DbConn.First(&section, sectionId)
	if result.Error != nil {
		log.Errorf("[service.GetSectionById] error occurred while getting section with id %v, err=%v\n", sectionId, result.Error)
	} else {
		log.Infof("[service.GetSectionById] successfully got section with id %v, rows affected=%v\n", sectionId, result.RowsAffected)
	}
	return section, result.Error
}

func CreateSection(section model.Section) (uint, error) {
	result := db.DbConn.Create(&section)
	if result.Error != nil {
		log.Errorf("[service.CreateSection] error occurred while creating section, err=%v\n", result.Error)
		return 0, result.Error
	}
	log.Infof("[service.CreateSection] successfully created the section with id=%v\n", section.ID)
	return section.ID, nil
}

func DeleteSectionById(sectionId uint) error {
	result := db.DbConn.Delete(&model.Section{}, sectionId)
	if result.Error != nil {
		log.Errorf("[service.DeleteSectionById] error occurred while deleting section %v, err=%v\n", sectionId, result.Error)
		return result.Error
	}
	log.Infof("[service.DeleteSectionById] successfully deleted section with id %v, rows affected=%v\n", sectionId, result.RowsAffected)
	return nil
}
