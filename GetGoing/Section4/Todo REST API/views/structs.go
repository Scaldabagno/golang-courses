package views

// Response code and body
type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

//PostRequest name and body
type PostRequest struct {
	Name string `json:"name"`
	Todo string `json:"todo"`
}
