package model

type CreateMovieRequest struct { ///make model out of it
	ID       uint
	Name     string
	Year     uint
	Director string
	Rating   uint
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
