package forms

// Create ...
type Create struct {
	URL string `form:"url" json:"url" xml:"url"  binding:"required"`
}
