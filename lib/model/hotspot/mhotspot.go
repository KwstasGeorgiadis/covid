package mhotspot

type CompareHotspotData struct {
	Country string    `json:"country"`
	Data    []float64 `json:"data"`
}

type Hotspot struct {
	MostCases    CompareHotspotData `json:"mostCases"`
	SecondCases  CompareHotspotData `json:"secondCases"`
	ThirdCases   CompareHotspotData `json:"thirdCases"`
	MostDeaths   CompareHotspotData `json:"mostDeaths"`
	SecondDeaths CompareHotspotData `json:"secondDeaths"`
	ThirdDeaths  CompareHotspotData `json:"thirdDeaths"`
}
