package models

import (
	"database/sql"
	"time"
)

type DBDefaultModel struct {
	Id         uint         `json:"id"`
	Created_At sql.NullTime `json:"created_at"`
	Updated_At sql.NullTime `json:"updated_at"`
	Deleted_At sql.NullTime `json:"deleted_at"`
}

type DefaultModel struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

type Pagination struct {
	Page     uint `json:"page"`
	PageSize uint `json:"page_size"`
}

func (model *DefaultModel) ToDBDefaultModel() DBDefaultModel {
	return DBDefaultModel{
		Id:         model.Id,
		Created_At: sql.NullTime{Time: model.CreatedAt},
		Updated_At: sql.NullTime{Time: model.UpdatedAt},
		Deleted_At: sql.NullTime{Time: model.DeletedAt},
	}
}

func (model *DefaultModel) ToAnotherDefaultModel() DefaultModel {
	return DefaultModel{
		Id:        model.Id,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: model.DeletedAt,
	}
}

func (model *DBDefaultModel) ToDefaultModel() DefaultModel {
	Default := DefaultModel{
		Id: model.Id,
	}
	if model.Created_At.Valid {
		Default.CreatedAt = model.Created_At.Time
	}
	if model.Updated_At.Valid {
		Default.UpdatedAt = model.Updated_At.Time
	}
	if model.Deleted_At.Valid {
		Default.DeletedAt = model.Deleted_At.Time
	}
	return Default
}
