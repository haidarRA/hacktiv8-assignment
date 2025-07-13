package models

type Source struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

var Sources = []Source{}