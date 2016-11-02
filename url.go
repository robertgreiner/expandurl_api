package main

import (
	"expandurl"
	"log"
)

type ExpandResponse struct {
	ShortURL    string `json:"short-url"`
	ExpandedURL string `json:"expanded-url"`
	Errors      error  `json:"errors"`
}

func ExpandURL(shortURL string) ExpandResponse {

	result, err := expandurl.Expand(shortURL)

	if err != nil {
		log.Println(err)
	}

	response := ExpandResponse{
		ShortURL:    shortURL,
		ExpandedURL: result,
		Errors:      err,
	}

	return response
}
