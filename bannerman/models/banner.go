package models

//Banner is for banner data
type Banner struct {
	ID   int
	Name string
	URL  string
}

//GetBannerByID is func for Get Banner by ID
func (b *Banner) GetBannerByID(id int) (Banner, error) {
	banner := Banner{ID: id, Name: "test Banner", URL: "//test"}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS test.hello(world varchar(50))")

	rows, err := db.Query("INSERT INTO ")
	return banner, nil
}
