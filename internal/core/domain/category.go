package domain

import (
	"encoding/json"
	"time"
)

type Category struct {
	BaseModel
	Name     string `gorm:"name"`
	Slug     string `gorm:"slug"`
	Weight   int    `gorm:"weight"`
	ParentID string `gorm:"parent_id"`
	Tags     string `json:"tags"`
	Labels   string `json:"labels"`
	IsActive bool   `gorm:"is_active"`
	Books    []Book `gorm:"foreignKey:CategoryID" json:"books,omitempty"`
}

type CategoryRequest struct {
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Type     string `json:"type"`
	UserID   string `json:"user_id"`
	Weight   int    `json:"weight"`
	Tags     string `json:"tags"`
	Labels   string `json:"labels"`
	IsActive bool   `json:"is_active"`
}

type ListCategoryRequest struct {
	ListRequest
	Weight   int  `form:"weight"`
	IsActive bool `form:"is_active"`
}

type CategoryUpdateRequest struct {
	Name     string `json:"name"`
	Tags     string `json:"tags"`
	Labels   string `json:"labels"`
	Weight   int    `json:"weight"`
	IsActive *bool  `json:"is_active"`
}

type CategoryResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Weight    int       `json:"weight"`
	IsActive  bool      `json:"is_active"`
}

func (r *CategoryRequest) Validate() error {
	if err := IsValidName(r.Name); err != nil {
		return err
	}
	return nil
}

func (c *Category) NewCategory(req *CategoryRequest) {
	c.Name = req.Name
	c.Labels = req.Labels
	c.Tags = req.Tags
	c.IsActive = true
}

func (r *CategoryUpdateRequest) NewUpdateRequest() Map {
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

func (c *Category) CategoryResponse() *CategoryResponse {
	data := &CategoryResponse{}
	_ = ConvertType(c, data)
	return data
}

func ConvertType(source any, target any) error {
	sbytes, err := json.Marshal(source)
	if err != nil {
		return err
	}
	return json.Unmarshal(sbytes, &target)
}
