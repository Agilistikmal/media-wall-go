package model

type TikTokSearchResponse struct {
	SugList []TikTokSearchSug
}

type TikTokSearchSug struct {
	Pos string
}
