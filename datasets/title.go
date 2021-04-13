package datasets

type Title struct {
	Identifiable
	Name       LocalizedField       `json:"name"`
	GenderName GenderLocalizedField `json:"gender_name"`
}