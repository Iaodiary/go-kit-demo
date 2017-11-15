package adservice

import "context"

//BannerService should supply interface to get banner/banners by size,website, group_id
//And it's depend what param client send
//condition1 : (group_id, size, en)
//condition2 : (website, size, en)
//condition3 : (UUID, website, size, en, )
//condition3 : (UUID, website, size, en, )
//condition3 : (UUID, website, size, en,website_title )
//UUID should be get from transport like:
//layer-http cookie or get param
//gobuffer param token
//thrift param token
//Service will save all loading log and click log for specific UUID
//Service will receive all stats and callback for all adv User
//Service will save all tags for user
//Banner tags, landing tags, attach to User tags
//Or service should be split to 2 part or 3 parts
//One is loading banner and calculate what banner to display
//stats how much time user stay on the page and focus on Banners (should be heartbeat)
//Client must get user whether active, and sent heartbeat every 3 seconds
//It can help us to calculate put how much banners on it, and banner switch interval
type BannerService interface {
	GetBanner(ctx context.Context)
	GetBanners(ctx context.Context)
}

//BannerRequest convert request to struct BannerRequest
type BannerRequest struct {
	GroupID int    `json:"group_id"`
	Size    string `json:"size"`
	Lang    string `json:"lang"`
}

//BannersRequest is
type BannersRequest struct {
}

type bannerService struct{}

func (s bannerService) GetBanner(ctx context.Context) {

}

func (s bannerService) getPopularBanner(ctx context.Context) {

}

//BannerLog url, banner_id
