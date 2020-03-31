package models

type Dog struct {
	Name string `json:"name"`
	Breed string `json:"breed"`
	Type string `json:"type"`
}

type Breed struct {
	Weight struct {
		Imperial string `json:"imperial"`
		Metric   string `json:"metric"`
	} `json:"weight"`
	Height struct {
		Imperial string `json:"imperial"`
		Metric   string `json:"metric"`
	} `json:"height"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	BredFor     string `json:"bred_for"`
	BreedGroup  string `json:"breed_group,omitempty"`
	LifeSpan    string `json:"life_span"`
	Temperament string `json:"temperament"`
	Origin      string `json:"origin"`
	CountryCode string `json:"country_code,omitempty"`
}

func (Dog) CollectionName() string {
	return "Dog"
}

func (Breed) CollectionName() string {
	return "Breed"
}