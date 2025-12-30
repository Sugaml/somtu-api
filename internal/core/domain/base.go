package domain

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"
)

type Map map[string]interface{}

func (m Map) ToMap() map[string]interface{} {
	return map[string]interface{}(m)
}

type BaseModel struct {
	ID        string     `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

type ListRequest struct {
	Page          int    `json:"page" form:"page" example:"1"`
	Size          int    `json:"size" form:"size" example:"10"`
	SortColumn    string `json:"sort_column" form:"sort_column"`
	SortDirection string `json:"sort_direction" form:"sort_direction"`
	Query         string `json:"query" form:"query"`
	StartDate     string `json:"start_date" form:"start_date" example:"2025-01-01"`
	EndDate       string `json:"end_date" form:"end_date" example:"2025-02-01"`
}

func (req *ListRequest) Prepare() {
	if req.Page < 1 {
		req.Page = 1
	}
	if req.Size < 1 {
		req.Size = 10
	}
	if req.SortColumn == "" {
		req.SortColumn = "created_at"
	}
	if req.SortDirection == "" {
		req.SortDirection = "asc"
	}
	endDate := time.Now()
	startDate := endDate.AddDate(0, -1, 0)
	if req.StartDate != "" {
		startDate, _ = time.Parse("2006-01-02", req.StartDate)
	}
	if req.EndDate != "" {
		endDate, _ = time.Parse("2006-01-02", req.EndDate)
	}
	req.StartDate = startDate.Format("2006-01-02")
	req.EndDate = endDate.Format("2006-01-02") + " 23:59:59.999"
}

func IsValidName(str string) error {
	nameRegex := `^[\p{L}]+([\p{L}\s'-]*[\p{L}])?$`
	validName := regexp.MustCompile(nameRegex)
	if !validName.MatchString(str) {
		return fmt.Errorf("%s is an invalid name", str)
	}
	return nil
}

// Response is a generic function to convert a struct to a response type.
func Convert[I, O any](input *I) *O {
	var output O
	bytes, err := json.Marshal(input)
	if err != nil {
		fmt.Printf("Error in marshal :: %v", err)
		return &output
	}
	err = json.Unmarshal(bytes, &output)
	if err != nil {
		fmt.Printf("Error in unmarshal :: %v", err)
		return &output
	}
	return &output
}

func ConvertToJson(input any) []byte {
	bytes, err := json.Marshal(input)
	if err != nil {
		fmt.Printf("Error in convertToJson :: %v", err)
		return []byte{}
	}
	return bytes
}

func ConvertFromJson[T any](input []byte) T {
	var output T
	err := json.Unmarshal(input, &output)
	if err != nil {
		fmt.Printf("Error in convertFromJson :: %v", err)
		return output
	}
	return output
}
