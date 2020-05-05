package dto

import (
	"hexample.com/src/oleander/datastore/domain"
)

type DatastoreDTO struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Path     string `json:"path,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Port     int    `json:"port,omitempty"`
}

func NewDatastoreDTO(ag *domain.DatastoreAG) *DatastoreDTO {
	return &DatastoreDTO{
		ID:       ag.GetID(),
		Name:     ag.GetName(),
		Path:     ag.GetPath(),
		Hostname: ag.GetHostname(),
		Port:     ag.GetPort(),
	}
}

func NewDatastoreDTOList(agList []*domain.DatastoreAG) []*DatastoreDTO {
	dtoList := make([]*DatastoreDTO, 0)

	for _, ag := range agList{
		dtoList = append(dtoList, NewDatastoreDTO(ag))
	}

	return dtoList
}