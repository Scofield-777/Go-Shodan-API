package shodan

//Defined a URL as a constant value
const BaseURL = "https://api.shodan.io"

//Define a Client struct, used for maintaining your API token across requests.
type Client struct{
	apiKey string
}

//Defined a New() helper function, that takes API token as input that creates and returns an 
//initialized Client Instance
func New(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}
