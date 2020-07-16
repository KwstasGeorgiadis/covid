package crnews

import (
	"encoding/json"
	"net/http"
	"time"

	applogger "github.com/junkd0g/covid/lib/applogger"
	merror "github.com/junkd0g/covid/lib/model/error"
	mnews "github.com/junkd0g/covid/lib/model/news"
	news "github.com/junkd0g/covid/lib/news"
)

/*
	Get request to /api/news/all with no parameters

	Response:

	{
    "vaccine": {
        "data": [
            {
                "title": "Here's where we stand on getting a coronavirus vaccine - CNN",
                "description": "<ol><li><a href=\"https://www.cnn.com/2020/06/08/health/covid-19-vaccine-latest/index.html\" target=\"_blank\">Here's where we stand on getting a coronavirus vaccine</a>&nbsp;&nbsp;<font color=\"#6f6f6f\">CNN</font></li><li><a href=\"https://www.thelancet.com/journals/lancet/article/PIIS0140-6736(20)31252-6/fulltext\" target=\"_blank\">COVID-19 vaccine development pipeline gears up</a>&nbsp;&nbsp;<font color=\"#6f6f6f\">The Lancet</font></li><li><a href=\"https://www.healthline.com/health-news/why-companies-are-making-billions-of-covid-19-vaccine-doses-that-may-not-work\" target=\"_blank\">Why Companies Are Making Billions of COVID-19 Vaccine Doses</a>&nbsp;&nbsp;<font color=\"#6f6f6f\">Healthline</font></li><li><a href=\"https://www.weforum.org/agenda/2020/06/astrazeneca-covid19-vaccine-gates-foundation\" target=\"_blank\">Pharmaceutical company pledges 2 billion COVID-19 vaccine doses</a>&nbsp;&nbsp;<font color=\"#6f6f6f\">World Economic Forum</font></li><li><a href=\"https://www.usnews.com/news/health-news/articles/2020-06-08/experts-optimistic-in-search-for-covid-19-vaccine\" target=\"_blank\">Experts Optimistic in Search for COVID-19 Vaccine | Health News</a>&nbsp;&nbsp;<font color=\"#6f6f6f\">U.S. News & World Report</font></li><li><strong><a href=\"https://news.google.com/stories/CAAqOQgKIjNDQklTSURvSmMzUnZjbmt0TXpZd1NoTUtFUWpyNnZMMmo0QU1FWTRkbWhyVmgyeWxLQUFQAQ?oc=5\" target=\"_blank\">View Full Coverage on Google News</a></strong></li></ol>",
                "url": "https://www.cnn.com/2020/06/08/health/covid-19-vaccine-latest/index.html",
                "urlToImage": "",
                "publishedAt": "Mon, 08 Jun 2020 12:12:50 GMT",
                "source": "CNN",
                "sourceURL": "https://www.cnn.com"
            },
            {
                "title": "COVID-19 vaccine trials bring hope for many but come too late for this family - NBC News",
                "description": "<a href=\"https://www.nbcnews.com/health/health-news/covid-19-vaccine-trials-bring-hope-many-come-too-late-n1226716\" target=\"_blank\">COVID-19 vaccine trials bring hope for many but come too late for this family</a>&nbsp;&nbsp;<font color=\"#6f6f6f\">NBC News</font>",
                "url": "https://www.nbcnews.com/health/health-news/covid-19-vaccine-trials-bring-hope-many-come-too-late-n1226716",
                "urlToImage": "",
                "publishedAt": "Sun, 07 Jun 2020 16:02:58 GMT",
                "source": "NBC News",
                "sourceURL": "https://www.nbcnews.com"
            }
        ]
    },
    "treament": {
        "data": [
            {
                "title": "U.S. government’s supply of COVID-19 treatment drug, remdesivir, will run out at the end of the month - KTLA",
                "description": "<a href=\"https://ktla.com/news/coronavirus/u-s-governments-supply-of-covid-19-treatment-drug-remdesivir-will-run-out-at-the-end-of-the-month/\" target=\"_blank\">U.S. government’s supply of COVID-19 treatment drug, remdesivir, will run out at the end of the month</a>&nbsp;&nbsp;<font color=\"#6f6f6f\">KTLA</font>",
                "url": "https://ktla.com/news/coronavirus/u-s-governments-supply-of-covid-19-treatment-drug-remdesivir-will-run-out-at-the-end-of-the-month/",
                "urlToImage": "",
                "publishedAt": "Mon, 08 Jun 2020 02:39:00 GMT",
                "source": "KTLA",
                "sourceURL": "https://ktla.com"
            },
            {
                "title": "Drug Trial Planned For Synthetic Cannabinoid COVID-19 Treatment - The Fresh Toast",
                "description": "<a href=\"https://thefreshtoast.com/cannabis/drug-trial-planned-for-synthetic-cannabinoid-covid-19-treatment/\" target=\"_blank\">Drug Trial Planned For Synthetic Cannabinoid COVID-19 Treatment</a>&nbsp;&nbsp;<font color=\"#6f6f6f\">The Fresh Toast</font>",
                "url": "https://thefreshtoast.com/cannabis/drug-trial-planned-for-synthetic-cannabinoid-covid-19-treatment/",
                "urlToImage": "",
                "publishedAt": "Mon, 08 Jun 2020 13:36:15 GMT",
                "source": "The Fresh Toast",
                "sourceURL": "https://thefreshtoast.com"
            }
        ]
    },
    "news": {
        "data": [
            {
                "title": "What you need to know about the COVID-19 pandemic on 8 June - World Economic Forum",
                "description": "<a href=\"https://www.weforum.org/agenda/2020/06/covid-19-what-you-need-to-know-about-the-coronavirus-pandemic-on-8-june/\" target=\"_blank\">What you need to know about the COVID-19 pandemic on 8 June</a>&nbsp;&nbsp;<font color=\"#6f6f6f\">World Economic Forum</font>",
                "url": "https://www.weforum.org/agenda/2020/06/covid-19-what-you-need-to-know-about-the-coronavirus-pandemic-on-8-june/",
                "urlToImage": "",
                "publishedAt": "Mon, 08 Jun 2020 08:51:16 GMT",
                "source": "World Economic Forum",
                "sourceURL": "https://www.weforum.org"
            },
            {
                "title": "How Reskilling Can Soften the Economic Blow of Covid-19 - Harvard Business Review",
                "description": "<a href=\"https://hbr.org/2020/06/how-reskilling-can-soften-the-economic-blow-of-covid-19\" target=\"_blank\">How Reskilling Can Soften the Economic Blow of Covid-19</a>&nbsp;&nbsp;<font color=\"#6f6f6f\">Harvard Business Review</font>",
                "url": "https://hbr.org/2020/06/how-reskilling-can-soften-the-economic-blow-of-covid-19",
                "urlToImage": "",
                "publishedAt": "Mon, 08 Jun 2020 12:10:46 GMT",
                "source": "Harvard Business Review",
                "sourceURL": "https://hbr.org"
			}
        ]
    }
}
*/
func NewsAllHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := perform()
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "crnews", "NewsAllHandle",
		"Endpoint /api/news/all called with response JSON body "+string(jsonBody), status, elapsed)
}

func perform() ([]byte, int) {

	generalNews, err := news.GetNews()
	if err != nil {
		applogger.Log("ERROR", "crnews", "perform", err.Error())
		statsErrJSONBody, _ := json.Marshal(merror.ErrorMessage{Message: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	newsTreatment, errNewsTreatment := news.GetTreatmentNews()
	if errNewsTreatment != nil {
		applogger.Log("ERROR", "crnews", "perform", err.Error())
		statsErrJSONBody, _ := json.Marshal(merror.ErrorMessage{Message: errNewsTreatment.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	newsVaccine, errNewsVaccine := news.GetVaccineNews()
	if errNewsVaccine != nil {
		applogger.Log("ERROR", "crnews", "perform", err.Error())
		statsErrJSONBody, _ := json.Marshal(merror.ErrorMessage{Message: errNewsVaccine.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	var allArticlesData mnews.AllArticlesData
	allArticlesData.NewsArticles = generalNews
	allArticlesData.TreatmentArticles = newsTreatment
	allArticlesData.VaccineArticles = newsVaccine

	jsonBody, jsonBodyErr := json.Marshal(allArticlesData)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "crnews", "perform", err.Error())
		errorJSONBody, _ := json.Marshal(merror.ErrorMessage{Message: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	return jsonBody, 200
}
