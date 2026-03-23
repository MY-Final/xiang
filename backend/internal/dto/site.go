package dto

type SaveSiteSettingRequest struct {
	SiteName     string `json:"site_name"`
	SiteSubtitle string `json:"site_subtitle"`
	Logo         string `json:"logo"`
	AboutText    string `json:"about_text"`
	ThemeColor   string `json:"theme_color"`
}
