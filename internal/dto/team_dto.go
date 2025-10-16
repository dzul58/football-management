package dto

// CreateTeamRequest represents request to create a team
type CreateTeamRequest struct {
	Name        string `json:"name" binding:"required"`
	LogoURL     string `json:"logo_url"`
	FoundedYear int    `json:"founded_year" binding:"required,min=1800,max=2100"`
	HomeAddress string `json:"home_address" binding:"required"`
	HomeCity    string `json:"home_city" binding:"required"`
}

// UpdateTeamRequest represents request to update a team
type UpdateTeamRequest struct {
	Name        string `json:"name"`
	LogoURL     string `json:"logo_url"`
	FoundedYear int    `json:"founded_year" binding:"omitempty,min=1800,max=2100"`
	HomeAddress string `json:"home_address"`
	HomeCity    string `json:"home_city"`
}

// TeamResponse represents team data in response
type TeamResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	LogoURL     string `json:"logo_url,omitempty"`
	FoundedYear int    `json:"founded_year"`
	HomeAddress string `json:"home_address"`
	HomeCity    string `json:"home_city"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
