package crnews

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type NewsExpectedResponses struct {
	Data []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
		URLToImage  string `json:"urlToImage"`
		PublishedAt string `json:"publishedAt"`
		Source      string `json:"source"`
		SourceURL   string `json:"sourceURL"`
	} `json:"data"`
}

func Test_APINews(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/news", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(NewsHandle)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var ner NewsExpectedResponses
	json.Unmarshal([]byte(rr.Body.String()), &ner)

	if len(ner.Data) <= 0 {
		t.Errorf("Missing data in the response")
	}

	if &ner.Data[0].Title == nil {
		t.Errorf("Title field is empty %s", ner.Data[0].Title)
	}

}

func Test_APINewsVaccine(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/news/vaccine", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(NewsVaccineHandle)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var ner NewsExpectedResponses
	json.Unmarshal([]byte(rr.Body.String()), &ner)

	if len(ner.Data) <= 0 {
		t.Errorf("Missing data in the response")
	}

	if &ner.Data[0].Title == nil {
		t.Errorf("Title field is empty %s", ner.Data[0].Title)
	}

}

func Test_APINewsTreatment(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/news/treatment", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(NewsTreatmentHandle)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var ner NewsExpectedResponses
	json.Unmarshal([]byte(rr.Body.String()), &ner)

	if len(ner.Data) <= 0 {
		t.Errorf("Missing data in the response")
	}

	if &ner.Data[0].Title == nil {
		t.Errorf("Title field is empty %s", ner.Data[0].Title)
	}

}

type AllNewsExpectedResponses struct {
	Vaccine struct {
		Data []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			URL         string `json:"url"`
			URLToImage  string `json:"urlToImage"`
			PublishedAt string `json:"publishedAt"`
			Source      string `json:"source"`
			SourceURL   string `json:"sourceURL"`
		} `json:"data"`
	} `json:"vaccine"`
	Treament struct {
		Data []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			URL         string `json:"url"`
			URLToImage  string `json:"urlToImage"`
			PublishedAt string `json:"publishedAt"`
			Source      string `json:"source"`
			SourceURL   string `json:"sourceURL"`
		} `json:"data"`
	} `json:"treament"`
	News struct {
		Data []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			URL         string `json:"url"`
			URLToImage  string `json:"urlToImage"`
			PublishedAt string `json:"publishedAt"`
			Source      string `json:"source"`
			SourceURL   string `json:"sourceURL"`
		} `json:"data"`
	} `json:"news"`
}

func Test_APINewsAll(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/news/all", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(NewsAllHandle)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var aner AllNewsExpectedResponses
	json.Unmarshal([]byte(rr.Body.String()), &aner)

	if len(aner.Vaccine.Data) <= 0 {
		t.Errorf("Missing vaccine news data in the response")
	}

	if len(aner.Treament.Data) <= 0 {
		t.Errorf("Missing treatment news data in the response")
	}

	if len(aner.News.Data) <= 0 {
		t.Errorf("Missing news data in the response")
	}

	if &aner.News.Data[0].Title == nil {
		t.Errorf("Title field is empty %s", aner)
	}

}
