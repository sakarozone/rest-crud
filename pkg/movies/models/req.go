package model

type CreateMovieRequest struct {
	ID       uint   `json:"ID" binding:"required"`
	Name     string `json:"Name" binding:"required"`
	Year     uint   `json:"Year" binding:"required"`
	Director string `json:"Director" binding:"required"`
	Rating   uint   `json:"Rating" binding:"required"`
}

func (r CreateMovieRequest) ToMovie() MovieTable {
	return MovieTable{
		ID:       r.ID,
		Name:     r.Name,
		Year:     r.Year,
		Director: r.Director,
		Rating:   r.Rating,
	}
}
