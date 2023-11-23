package models

type Request struct {
	Event           string `json:"ev"`
	EventType       string `json:"et"`
	AppID           string `json:"id"`
	UserID          string `json:"uid"`
	MessageID       string `json:"mid"`
	PageTitle       string `json:"t"`
	PageURL         string `json:"p"`
	BrowserLanguage string `json:"l"`
	ScreenSize      string `json:"sc"`
	Attributes      []struct {
		Key   string `json:"atrk"`
		Value string `json:"atrv"`
		Types string `json:"atrt"`
	} `json:"attributes"`
	UserTraits []struct {
		Key   string `json:"utrk"`
		Value string `json:"utrv"`
		Types string `json:"utrt"`
	} `json:"userTraits"`
}
type Attribute struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}
type Converted struct {
	Event           string               `json:"event"`
	EventType       string               `json:"event_type"`
	AppID           string               `json:"app_id"`
	UserID          string               `json:"user_id"`
	MessageID       string               `json:"message_id"`
	PageTitle       string               `json:"page_title"`
	PageURL         string               `json:"page_url"`
	BrowserLanguage string               `json:"browser_language"`
	ScreenSize      string               `json:"screen_size"`
	Attributes      map[string]Attribute `json:"attributes"`
	UserTraits      map[string]Attribute `json:"traits"`
}
