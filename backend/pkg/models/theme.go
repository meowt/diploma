package models

type ThemeHTTP struct {
	DefaultModel
	CreatorId   uint   `json:"creator_id"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

type ThemeUsecase struct {
	DefaultModel
	CreatorId   uint
	Url         string
	Description string
}

type ThemeDB struct {
	DBDefaultModel
	CreatorId   uint
	Url         string
	Preview     []string
	Description string
}

func (th *ThemeHTTP) ToUsecase() *ThemeUsecase {
	return &ThemeUsecase{
		DefaultModel: th.DefaultModel,
		CreatorId:    th.CreatorId,
		Url:          th.Url,
		Description:  th.Description,
	}
}

func (th *ThemeDB) ToUsecase() *ThemeUsecase {
	return &ThemeUsecase{
		DefaultModel: th.ToDefaultModel(),
		CreatorId:    th.CreatorId,
		Url:          th.Url,
		Description:  th.Description,
	}
}

func (th *ThemeUsecase) ToDB() *ThemeDB {
	return &ThemeDB{
		DBDefaultModel: th.ToDBDefaultModel(),
		CreatorId:      th.CreatorId,
		Url:            th.Url,
		Description:    th.Description,
	}
}
