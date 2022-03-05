package myMusic

type song struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var mySongs = []song{
	{ID: "1", Title: "백예린 - PoPo", Artist: "Yerin Baek", Price: 56.99},
	{ID: "2", Title: "ทน", Artist: "SPRITE x GUYGEEGEE", Price: 39.99},
}

// Return my songs
func GetMySongs() []song {
	return mySongs
}
