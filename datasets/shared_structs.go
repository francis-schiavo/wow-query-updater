package datasets

import (
	"regexp"
	"strconv"
)

type LocalizedField struct {
	EnUS string `json:"en_US"`
	EsMX string `json:"es_MX"`
	PtBR string `json:"pt_BR"`
	DeDE string `json:"de_DE"`
	EnGB string `json:"en_GB"`
	EsES string `json:"es_ES"`
	FrFR string `json:"fr_FR"`
	ItIT string `json:"it_IT"`
	RuRU string `json:"ru_RU"`
	KoKR string `json:"ko_KR"`
	ZhTW string `json:"zh_TW"`
	ZhCN string `json:"zh_CN"`
}

type Color struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
	A int `json:"a"`
}

type LocalizedDisplayString struct {
	DisplayString LocalizedField `json:"display_string"`
	Color         Color          `json:"color,omitempty"`
}

type GenderLocalizedField struct {
	Male   LocalizedField `json:"male"`
	Female LocalizedField `json:"female"`
}

type Href struct {
	Href string `json:"href"`
}

type Self struct {
	Self Href `json:"self"`
}

type Key struct {
	Key Href `json:"key" pg:"-"`
}

type SelfReference struct {
	Links Self `json:"_links" pg:"-"`
}

// Record IDs

type IIdentifiable interface {
	GetID() int
}

type Identifiable struct {
	ID int `json:"id" pg:",pk,notnull,on_delete:RESTRICT,on_update:CASCADE,use_zero"`
}

func (i Identifiable) GetID() int {
	return i.ID
}

type HrefIdentifiable struct {
	Href
}

func (h HrefIdentifiable) GetID() int {
	regexID, _ := regexp.Compile("/([0-9]+)\\?")
	id, _ := strconv.Atoi(regexID.FindStringSubmatch(h.Href.Href)[1])
	return id
}

type Identifiables []Identifiable
type HrefIdentifiables []HrefIdentifiable

// Generic assets

type Media struct {
	tableName struct{} `pg:",discard_unknown_columns"`
	Identifiable
	Key
}

type Asset struct {
	Key   string `json:"key" pg:",pk,notnull,on_delete:RESTRICT, on_update: CASCADE"`
	Value string `json:"value"`
}

type Assets []Asset