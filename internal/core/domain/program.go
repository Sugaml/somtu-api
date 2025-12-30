package domain

import (
	"time"
)

type ProgramNew struct {
	BaseModel
	FacultyID uint
	Faculty   Faculty

	Name          string `gorm:"size:100;not null"`
	DurationYears int
}

type Program struct {
	BaseModel
	Name     string `gorm:"name"`
	Slug     string `gorm:"slug"`
	Weight   int    `gorm:"weight"`
	ParentID string `gorm:"parent_id"`
	Tags     string `json:"tags"`
	Labels   string `json:"labels"`
	IsActive bool   `gorm:"is_active"`
}

type ProgramRequest struct {
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Type     string `json:"type"`
	Weight   int    `json:"weight"`
	Tags     string `json:"tags"`
	Labels   string `json:"labels"`
	IsActive bool   `json:"is_active"`
}

type ListProgramRequest struct {
	ListRequest
	Weight   int  `form:"weight"`
	IsActive bool `form:"is_active"`
}

type ProgramUpdateRequest struct {
	Name     string `json:"name"`
	Tags     string `json:"tags"`
	Labels   string `json:"labels"`
	Weight   int    `json:"weight"`
	IsActive *bool  `json:"is_active"`
}

type ProgramResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Weight    int       `json:"weight"`
	IsActive  bool      `json:"is_active"`
}

func (r *ProgramRequest) Validate() error {
	if err := IsValidName(r.Name); err != nil {
		return err
	}
	return nil
}

func (c *Program) NewProgram(req *ProgramRequest) {
	c.Name = req.Name
	c.Labels = req.Labels
	c.Tags = req.Tags
	c.IsActive = true
}

func (r *ProgramUpdateRequest) NewUpdateRequest() Map {
	mp := map[string]interface{}{}
	if r.Name != "" {
		mp["name"] = r.Name
	}
	if r.Weight != 0 {
		mp["weight"] = r.Weight
	}
	if r.Labels != "" {
		mp["labels"] = r.Labels
	}
	if r.IsActive != nil {
		mp["is_active"] = *r.IsActive
	}
	return mp
}

func (c *Program) ProgramResponse() *ProgramResponse {
	data := &ProgramResponse{}
	_ = ConvertType(c, data)
	return data
}
