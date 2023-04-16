package presenter

import (

)

type SearchResquest struct {
	Username string `json:"search"`
}

type SearchResponse struct {
	Match []LogInResponse
}
