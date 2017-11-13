package models

//Banner describe the banner
type Banner struct {
	ID       int
	GroupID  int
	Name     string
	Language string
	Size     string
	URL      string
}

//GetBannerByID 根据ID获取Banner
func GetBannerByID(id int) (Banner, error) {
	banner := Banner{}
	rows, err := db.Query("SELECT * FROM gw_adv_banner WHERE id=? LIMIT 1", id)
	defer rows.Close()
	if err == nil {
		return banner, err
	}
	rows.Scan(&banner)
	return banner, nil
}

//GetBannersBySize Get Banners By Size
func GetBannersBySize(size string) ([]*Banner, error) {
	var banners []*Banner
	rows, err := db.Query("SELECT id,group_id, name, size, url FROM gw_adv_banner  WHERE status=1 AND size=?", size)
	defer rows.Close()
	if err != nil {
		return banners, err
	}
	for rows.Next() {
		b := new(Banner)

		if err := rows.Scan(&b.ID, &b.GroupID, &b.Name, &b.Size, &b.URL); err != nil {
			return banners, err
		}
		banners = append(banners, b)
	}
	return banners, nil
}
