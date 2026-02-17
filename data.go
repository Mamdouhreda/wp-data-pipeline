package main
import "encoding/json"

type Rendered struct{
	Rendered string `json:"rendered"`
}

//get the image url
type ImageFile string
func (I *ImageFile) UnmarshalJSON(data []byte) error {
	var img struct {
		File string `json:"file"`
	}
	if err := json.Unmarshal(data, &img); err != nil {
		return err
	}
	*I = ImageFile(img.File)
	return nil
}
type Post struct {
	ID      int    `json:"id"`
	  Link    string `json:"link"`
	Title   Rendered `json:"title"`
	// Content Rendered `json:"content"`
	Image   ImageFile `json:"featured_image"`
}
