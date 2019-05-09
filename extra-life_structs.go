package main

type apiConf struct {
	APIURL  string         `json:"api_url"`
	TeamID  string         `json:"team_id"`
	Between donationPeriod `json:"between"`
}

type donationPeriod struct {
	Start date `json:"start,omitempty"`
	End   date `json:"end,omitempty"`
}

type date struct {
	Year  string `json:"year"`
	Month string `json:"month"`
	Day   string `json:"day"`
}

type events []event

type event struct {
	EndDateUTC   string `json:"endDateUTC,omitempty"`
	EventID      int    `json:"eventID,omitempty"`
	Timezone     string `json:"timezone,omitempty"`
	Type         string `json:"type,omitempty"`
	StartDateUTC string `json:"startDateUTC,omitempty"`
	Name         string `json:"name,omitempty"`
}

type teams []team

type team struct {
	FundraisingGoal    float64 `json:"fundraisingGoal,omitempty"`
	IsInviteOnly       bool    `json:"isInviteOnly,omitempty"`
	CaptainDisplayName string  `json:"captainDisplayName,omitempty"`
	EventName          string  `json:"eventName,omitempty"`
	AvatarImageURL     string  `json:"avatarImageURL,omitempty"`
	Links              links   `json:"links,omitempty"`
	CreatedDateUTC     string  `json:"createdDateUTC,omitempty"`
	EventID            int     `json:"eventID,omitempty"`
	SumDonations       float64 `json:"sumDonations,omitempty"`
	TeamID             int     `json:"teamID,omitempty"`
	Name               string  `json:"name,omitempty"`
	NumDonations       int     `json:"numDonations,omitempty"`
}

type teamUsers []user

type user struct {
	DisplayName     string  `json:"displayName,omitempty"`
	FundraisingGoal float64 `json:"fundraisingGoal,omitempty"`
	EventName       string  `json:"eventName,omitempty"`
	Links           links   `json:"links,omitempty"`
	CreatedDateUTC  string  `json:"createdDateUTC,omitempty"`
	EventID         int     `json:"eventID,omitempty"`
	SumDonations    float64 `json:"sumDonations,omitempty"`
	ParticipantID   int     `json:"participantID,omitempty"`
	TeamName        string  `json:"teamName,omitempty"`
	AvatarImageURL  string  `json:"avatarImageURL,omitempty"`
	TeamID          int     `json:"teamID,omitempty"`
	IsTeamCaptain   bool    `json:"isTeamCaptain,omitempty"`
	SumPledges      float64 `json:"sumPledges,omitempty"`
	NumDonations    int     `json:"numDonations,omitempty"`
}

type donations []donation

type donation struct {
	DisplayName    string  `json:"displayName"`
	Message        string  `json:"message,omitempty"`
	ParticipantID  int     `json:"participantID"`
	Amount         float64 `json:"amount"`
	DonorID        string  `json:"donorID"`
	AvatarImageURL string  `json:"avatarImageURL"`
	CreatedDateUTC string  `json:"createdDateUTC"`
	TeamID         int     `json:"teamID"`
	DonationID     string  `json:"donationID"`
}

type links struct {
	Donate string `json:"donate,omitempty"`
	Stream string `json:"stream,omitempty"`
	Page   string `json:"page,omitempty"`
}
