package main

type Trustee struct {
	OrganisationNumber int    `json:"organisation_number"`
	TrusteeName        string `json:"trustee_name"`
	TrusteeID          int    `json:"trustee_id"`
}
type Charity struct {
	OrganisationNumber        int       `json:"organisation_number"`
	RegCharityNumber          int       `json:"reg_charity_number"`
	GroupSubsidSuffix         int       `json:"group_subsid_suffix"`
	CharityName               string    `json:"charity_name"`
	CharityType               string    `json:"charity_type"`
	Insolvent                 bool      `json:"insolvent"`
	InAdministration          bool      `json:"in_administration"`
	PrevExceptedInd           bool      `json:"prev_excepted_ind"`
	CifCdfInd                 any       `json:"cif_cdf_ind"`
	CioDissolutionInd         bool      `json:"cio_dissolution_ind"`
	InterimManagerInd         any       `json:"interim_manager_ind"`
	DateOfInterimManagerAppt  any       `json:"date_of_interim_manager_appt"`
	RegStatus                 string    `json:"reg_status"`
	DateOfRegistration        string    `json:"date_of_registration"`
	DateOfRemoval             any       `json:"date_of_removal"`
	LatestAccFinYearStartDate string    `json:"latest_acc_fin_year_start_date"`
	LatestAccFinYearEndDate   string    `json:"latest_acc_fin_year_end_date"`
	LatestIncome              float64   `json:"latest_income"`
	LatestExpenditure         float64   `json:"latest_expenditure"`
	AddressLineOne            string    `json:"address_line_one"`
	AddressLineTwo            string    `json:"address_line_two"`
	AddressLineThree          any       `json:"address_line_three"`
	AddressLineFour           any       `json:"address_line_four"`
	AddressLineFive           string    `json:"address_line_five"`
	AddressPostCode           string    `json:"address_post_code"`
	Phone                     string    `json:"phone"`
	Email                     string    `json:"email"`
	Web                       string    `json:"web"`
	CharityCoRegNumber        string    `json:"charity_co_reg_number"`
	ReportingStatus           string    `json:"reporting_status"`
	RemovalReason             any       `json:"removal_reason"`
	CioInd                    bool      `json:"cio_ind"`
	LastModifiedTime          string    `json:"last_modified_time"`
	TrusteeNames              []Trustee `json:"trustee_names"`
	WhoWhatWhere              []struct {
		ClassificationCode string `json:"classification_code"`
		ClassificationType string `json:"classification_type"`
		ClassificationDesc string `json:"classification_desc"`
	} `json:"who_what_where"`
	CharityAoOCountryContinent []struct {
		Country   string `json:"country"`
		Continent string `json:"continent"`
	} `json:"CharityAoOCountryContinent"`
	CharityAoOLocalAuthority []any `json:"CharityAoOLocalAuthority"`
	CharityAoORegion         []struct {
		Region string `json:"region"`
	} `json:"CharityAoORegion"`
	OtherNames []struct {
		OtherName string `json:"other_name"`
		NameType  string `json:"name_type"`
	} `json:"other_names"`
	ConstituencyName []struct {
		ConstituencyName string `json:"constituency_name"`
	} `json:"constituency_name"`
}
