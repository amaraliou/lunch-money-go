package lunchmoney

// Asset ...
type Asset struct {
	ID              int64  `json:"id"`
	TypeName        string `json:"type_name"`
	SubtypeName     string `json:"subtype_name"`
	Name            string `json:"name"`
	Balance         string `json:"balance"`
	BalanceAsOf     string `json:"balance_as_of"`
	Currency        string `json:"currency"`
	InstitutionName string `json:"institution_name"`
	CreatedAt       string `json:"created_at"`
}
