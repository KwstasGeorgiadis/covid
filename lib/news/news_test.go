package news

import (
	"testing"

	mnews "github.com/junkd0g/covid/lib/model/news"
	//mockredis "github.com/junkd0g/covid/test/redis"
)

type requestDataMock struct{}

var requestDataMockFunc func(url string) (mnews.ArticlesData, error)

func (u requestDataMock) requestNewsData(url string) (mnews.ArticlesData, error) {
	return requestDataMockFunc(url)
}

type requestCacheDataMock struct{}

var requestCacheDataMockFunc func(newsType string) (mnews.ArticlesData, bool, error)

func (u requestCacheDataMock) getCacheData(newsType string) (mnews.ArticlesData, bool, error) {
	return requestCacheDataMockFunc(newsType)
}

func TestNews(t *testing.T) {
	reqCacheOB = requestCacheDataMock{}
	reqDataOB = requestDataMock{}
	//redis = mockredis.MockRedisST{}

	requestDataOne := mnews.Article{
		Title:       "To the Grand Line",
		Description: "To the Grand Line To the Grand Line",
		URL:         "stats-covid.com",
		URLToImage:  "stats-covid.com/image.jpg",
		PublishedAt: "Sun, 07 Jun 2020 16:02:58 GMT",
		Source:      "stats",
		SourceURL:   "stats-covid",
	}

	/* --------------------------
		Treatmeant news testing
	----------------------------- */
	requestDataMockFunc = func(url string) (mnews.ArticlesData, error) {
		var articles mnews.ArticlesData
		var articleArr []mnews.Article
		articleArr = append(articleArr, requestDataOne)
		articleArr = append(articleArr, requestDataOne)
		articleArr = append(articleArr, requestDataOne)
		articles.Articles = articleArr
		return articles, nil
	}

	requestCacheDataMockFunc = func(newsType string) (mnews.ArticlesData, bool, error) {
		var articles mnews.ArticlesData
		var articleArr []mnews.Article
		articleArr = append(articleArr, requestDataOne)
		articles.Articles = articleArr
		return articles, true, nil
	}

	withCacheData, err := GetTreatmentNews()
	if err != nil {
		t.Fatal(err)
	}

	if len(withCacheData.Articles) != 1 {
		t.Fatal("Not using cached data")
	}

	requestCacheDataMockFunc = func(newsType string) (mnews.ArticlesData, bool, error) {
		var articles mnews.ArticlesData
		var articleArr []mnews.Article
		articleArr = append(articleArr, requestDataOne)
		articles.Articles = articleArr
		return articles, false, nil
	}

	withCacheDataFalse, err := GetTreatmentNews()
	if err != nil {
		t.Fatal(err)
	}
	if len(withCacheDataFalse.Articles) == 1 {
		t.Fatal("Using cached data instead of requesting data")
	}

	/* ------------------------
		Vaccine news testing
	---------------------------- */
	requestDataMockFunc = func(url string) (mnews.ArticlesData, error) {
		var articles mnews.ArticlesData
		var articleArr []mnews.Article
		articleArr = append(articleArr, requestDataOne)
		articleArr = append(articleArr, requestDataOne)
		articleArr = append(articleArr, requestDataOne)
		articles.Articles = articleArr
		return articles, nil
	}

	requestCacheDataMockFunc = func(newsType string) (mnews.ArticlesData, bool, error) {
		var articles mnews.ArticlesData
		var articleArr []mnews.Article
		articleArr = append(articleArr, requestDataOne)
		articleArr = append(articleArr, requestDataOne)
		articleArr = append(articleArr, requestDataOne)
		articleArr = append(articleArr, requestDataOne)
		articleArr = append(articleArr, requestDataOne)
		articleArr = append(articleArr, requestDataOne)
		articleArr = append(articleArr, requestDataOne)
		articles.Articles = articleArr
		return articles, true, nil
	}

	withVaccineCacheData, err := GetVaccineNews()
	if err != nil {
		t.Fatal(err)
	}

	if len(withVaccineCacheData.Articles) != 7 {
		t.Fatal("Not using cached data %", len(withVaccineCacheData.Articles))
	}

	requestCacheDataMockFunc = func(newsType string) (mnews.ArticlesData, bool, error) {
		var articles mnews.ArticlesData
		var articleArr []mnews.Article
		articleArr = append(articleArr, requestDataOne)
		articleArr = append(articleArr, requestDataOne)
		articles.Articles = articleArr
		return articles, false, nil
	}

	withVaccineCacheDataFalse, err := GetVaccineNews()
	if err != nil {
		t.Fatal(err)
	}
	if len(withVaccineCacheDataFalse.Articles) != 3 {
		t.Fatal("Using cached data instead of requesting data %", len(withVaccineCacheDataFalse.Articles))
	}

	/* ------------------------
		General news testing
	---------------------------- */
	requestDataMockFunc = func(url string) (mnews.ArticlesData, error) {
		var articles mnews.ArticlesData
		var articleArr []mnews.Article
		articleArr = append(articleArr, requestDataOne)
		articleArr = append(articleArr, requestDataOne)
		articleArr = append(articleArr, requestDataOne)
		articleArr = append(articleArr, requestDataOne)
		articleArr = append(articleArr, requestDataOne)
		articles.Articles = articleArr
		return articles, nil
	}

	requestCacheDataMockFunc = func(newsType string) (mnews.ArticlesData, bool, error) {
		var articles mnews.ArticlesData
		var articleArr []mnews.Article
		articleArr = append(articleArr, requestDataOne)
		articleArr = append(articleArr, requestDataOne)
		articles.Articles = articleArr
		return articles, true, nil
	}

	withGeneralCacheData, err := GetNews()
	if err != nil {
		t.Fatal(err)
	}

	if len(withGeneralCacheData.Articles) != 2 {
		t.Fatal("Not using cached data %", len(withGeneralCacheData.Articles))
	}

	requestCacheDataMockFunc = func(newsType string) (mnews.ArticlesData, bool, error) {
		var articles mnews.ArticlesData
		var articleArr []mnews.Article
		articleArr = append(articleArr, requestDataOne)
		articleArr = append(articleArr, requestDataOne)
		articleArr = append(articleArr, requestDataOne)
		articleArr = append(articleArr, requestDataOne)
		articles.Articles = articleArr
		return articles, false, nil
	}

	withGeneralCacheDataFalse, err := GetNews()
	if err != nil {
		t.Fatal(err)
	}
	if len(withGeneralCacheDataFalse.Articles) != 5 {
		t.Fatal("Using cached data instead of requesting data %", len(withVaccineCacheDataFalse.Articles))
	}
}
