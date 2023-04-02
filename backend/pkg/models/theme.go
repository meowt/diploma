package models

import "time"

type ThemeHTTP struct {
	Id          uint      `json:"id"`
	CreatorId   uint      `json:"creator_id"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ThemeCore struct {
	Id          uint
	CreatorId   uint
	Url         string
	Description string
	UpdatedAt   time.Time
}

type ThemeDB struct {
	DefaultModel
	CreatorId   uint
	Url         string
	Description string
}

func (th *ThemeHTTP) ToCore() ThemeCore {
	return ThemeCore{
		Id:          th.Id,
		CreatorId:   th.CreatorId,
		Url:         th.Url,
		Description: th.Description,
		UpdatedAt:   th.UpdatedAt,
	}
}

func (th *ThemeHTTP) FromCore(core *ThemeCore) {
	th.Id = core.Id
	th.CreatorId = core.CreatorId
	th.Url = core.Url
	th.Description = core.Description
	th.UpdatedAt = core.UpdatedAt
}

func (th *ThemeDB) ToCore() ThemeCore {
	return ThemeCore{
		Id:          th.Id,
		CreatorId:   th.CreatorId,
		Url:         th.Url,
		Description: th.Description,
		UpdatedAt:   th.DefaultModel.CreatedAt,
	}
}

func (th *ThemeDB) FromCore(core *ThemeCore) {
	th.Id = core.Id
	th.CreatorId = core.CreatorId
	th.Url = core.Url
	th.Description = core.Description
	th.DefaultModel.UpdatedAt = core.UpdatedAt
}
