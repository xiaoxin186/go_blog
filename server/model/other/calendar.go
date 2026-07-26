package other

type Calendar struct {
	Date         string `json:"date"`
	LunarDate    string `json:"lunarDate"`
	Ganzhi       string `json:"ganzhi"`
	Zodiac       string `json:"zodiac"`
	DayOfYear    string `json:"dayOfYear"`
	SolarTerm    string `json:"solarTerm"`
	Auspicious   string `json:"auspicious"`
	Inauspicious string `json:"inauspicious"`
}
