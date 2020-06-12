package mnews

import (
	"encoding/xml"
)

// ReponseNews response we are getting for the news api, being used in lib/news/news.go
type ReponseNews struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Media   string   `xml:"media,attr"`
	Channel struct {
		Text          string `xml:",chardata"`
		Generator     string `xml:"generator"`
		Title         string `xml:"title"`
		Link          string `xml:"link"`
		Language      string `xml:"language"`
		WebMaster     string `xml:"webMaster"`
		Copyright     string `xml:"copyright"`
		LastBuildDate string `xml:"lastBuildDate"`
		Description   string `xml:"description"`
		Item          []struct {
			Text  string `xml:",chardata"`
			Title string `xml:"title"`
			Link  string `xml:"link"`
			GUUID struct {
				Text        string `xml:",chardata"`
				IsPermaLink string `xml:"isPermaLink,attr"`
			} `xml:"guid"`
			PubDate     string `xml:"pubDate"`
			Description string `xml:"description"`
			Source      struct {
				Text string `xml:",chardata"`
				URL  string `xml:"url,attr"`
			} `xml:"source"`
		} `xml:"item"`
	} `xml:"channel"`
}

// AllArticlesData is being used in lib/news/news.go
type AllArticlesData struct {
	VaccineArticles   ArticlesData `json:"vaccine"`
	TreatmentArticles ArticlesData `json:"treament"`
	NewsArticles      ArticlesData `json:"news"`
}

// ArticlesData is being used in lib/news/news.go
type ArticlesData struct {
	Articles []Article `json:"data"`
}

// Article is being used in lib/news/news.go
type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Source      string `json:"source"`
	SourceURL   string `json:"sourceURL"`
}
