package crnews

import (
	"encoding/json"

	applogger "github.com/junkd0g/covid/lib/applogger"
	mnews "github.com/junkd0g/covid/lib/model/news"
	news "github.com/junkd0g/covid/lib/news"
	structs "github.com/junkd0g/covid/lib/structs"
)

//Perform used in the /new/ endpoint's handle to return
//	the structs.ArticlesData struct as a json response by calling
//	news.GetNews() which get and return news about covid 19
//
//	Array of all articles for covid 19
//
//	In this JSON format
//  {
//    "data": [
//        {
//           "title": "UGC ने जारी किए नंबर, पूछ सकते हैं एडमिशन से लेकर एग्जाम तक अपने सवाल",
//           "description": "कोरोना संक्रमण के चलते देश भर में  लॉकडाउन को एक बार फिर बढ़ा दिया गया है.",
//           "url": "https://aajtak.intoday.in/education/story/ugc-direct-numbers-for-queries-helpline-for-ug-pg-students-tedu-1-1192009.html",
//           "urlToImage": "https://smedia2.intoday.in/aajtak/images/stories/092019/3_1589811689_618x347.jpeg",
//           "publishedAt": "2020-05-18T15:14:14Z",
//           "content": "UGC helpline:"
//       },
//       {
//           "title": "Karen who can't believe she has to wear a mask to enter a supermarket confronts store manager",
//           "description": "A woman who called herself Shelley Lewis acted rude and arrogant toward Gelson's supermarket employees",
//           "url": "https://boingboing.net/2020/05/18/karen-who-cant-believe-she-h.html",
//           "urlToImage": "https://i1.wp.com/media.boingboing.net/wp-content/uploads/2020/05/mask-1.jpg?fit=700%2C503&ssl=1",
//           "publishedAt": "2020-05-18T15:12:19Z",
//           "content": ""
//        }
//		]
//	}
//
//	@return array of bytes of the json object
//	@return int http code status
func Perform() ([]byte, int) {

	newsDashboard, err := news.GetNews()
	if err != nil {
		applogger.Log("ERROR", "crnews", "Perform", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(newsDashboard)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "crnews", "Perform", err.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	return jsonBody, 200
}

//PerformVaccineNews used in the /news/vaccine endpoint's handle to return
//	the structs.ArticlesData struct as a json response by calling
//	news.GetNews() which get and return news about covid 19 vaccine news
//
//	Array of all articles
//
//	In this JSON format
//  {
//    "data": [
//        {
//           "title": "UGC ने जारी किए नंबर, पूछ सकते हैं एडमिशन से लेकर एग्जाम तक अपने सवाल",
//           "description": "कोरोना संक्रमण के चलते देश भर में  लॉकडाउन को एक बार फिर बढ़ा दिया गया है.",
//           "url": "https://aajtak.intoday.in/education/story/ugc-direct-numbers-for-queries-helpline-for-ug-pg-students-tedu-1-1192009.html",
//           "urlToImage": "https://smedia2.intoday.in/aajtak/images/stories/092019/3_1589811689_618x347.jpeg",
//           "publishedAt": "2020-05-18T15:14:14Z",
//           "content": "UGC helpline:"
//       },
//       {
//           "title": "Karen who can't believe she has to wear a mask to enter a supermarket confronts store manager",
//           "description": "A woman who called herself Shelley Lewis acted rude and arrogant toward Gelson's supermarket employees",
//           "url": "https://boingboing.net/2020/05/18/karen-who-cant-believe-she-h.html",
//           "urlToImage": "https://i1.wp.com/media.boingboing.net/wp-content/uploads/2020/05/mask-1.jpg?fit=700%2C503&ssl=1",
//           "publishedAt": "2020-05-18T15:12:19Z",
//           "content": ""
//        }
//		]
//	}
//
//	@return array of bytes of the json object
//	@return int http code status
func PerformVaccineNews() ([]byte, int) {

	newsDashboard, err := news.GetVaccineNews()
	if err != nil {
		applogger.Log("ERROR", "crnews", "PerformVaccineNews", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(newsDashboard)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "crnews", "PerformVaccineNews", err.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	return jsonBody, 200
}

//PerformTreatmentNews used in the /news/treatment endpoint's handle to return
//	the structs.ArticlesData struct as a json response by calling
//	news.GetNews() which get and return news about covid 19 treatment news
//
//	Array of all articles
//
//	In this JSON format
//  {
//    "data": [
//        {
//           "title": "UGC ने जारी किए नंबर, पूछ सकते हैं एडमिशन से लेकर एग्जाम तक अपने सवाल",
//           "description": "कोरोना संक्रमण के चलते देश भर में  लॉकडाउन को एक बार फिर बढ़ा दिया गया है.",
//           "url": "https://aajtak.intoday.in/education/story/ugc-direct-numbers-for-queries-helpline-for-ug-pg-students-tedu-1-1192009.html",
//           "urlToImage": "https://smedia2.intoday.in/aajtak/images/stories/092019/3_1589811689_618x347.jpeg",
//           "publishedAt": "2020-05-18T15:14:14Z",
//           "content": "UGC helpline:"
//       },
//       {
//           "title": "Karen who can't believe she has to wear a mask to enter a supermarket confronts store manager",
//           "description": "A woman who called herself Shelley Lewis acted rude and arrogant toward Gelson's supermarket employees",
//           "url": "https://boingboing.net/2020/05/18/karen-who-cant-believe-she-h.html",
//           "urlToImage": "https://i1.wp.com/media.boingboing.net/wp-content/uploads/2020/05/mask-1.jpg?fit=700%2C503&ssl=1",
//           "publishedAt": "2020-05-18T15:12:19Z",
//           "content": ""
//        }
//		]
//	}
//
//	@return array of bytes of the json object
//	@return int http code status
func PerformTreatmentNews() ([]byte, int) {

	newsDashboard, err := news.GetVaccineNews()
	if err != nil {
		applogger.Log("ERROR", "crnews", "PerformTreatmentNews", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(newsDashboard)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "crnews", "PerformTreatmentNews", err.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	return jsonBody, 200
}

func PerformAll() ([]byte, int) {

	generalNews, err := news.GetNews()
	if err != nil {
		applogger.Log("ERROR", "crnews", "PerformAll", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	newsTreatment, errNewsTreatment := news.GetTreatmentNews()
	if errNewsTreatment != nil {
		applogger.Log("ERROR", "crnews", "PerformAll", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: errNewsTreatment.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	newsVaccine, errNewsVaccine := news.GetVaccineNews()
	if errNewsVaccine != nil {
		applogger.Log("ERROR", "crnews", "PerformAll", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: errNewsVaccine.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	var allArticlesData mnews.AllArticlesData
	allArticlesData.NewsArticles = generalNews
	allArticlesData.TreatmentArticles = newsTreatment
	allArticlesData.VaccineArticles = newsVaccine

	jsonBody, jsonBodyErr := json.Marshal(allArticlesData)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "crnews", "PerformAll", err.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	return jsonBody, 200
}
