package main

type events []event

type event struct {
	EndDateUTC   string `json:"endDateUTC"`
	EventID      int    `json:"eventID"`
	Timezone     string `json:"timezone"`
	Type         string `json:"type"`
	StartDateUTC string `json:"startDateUTC"`
	Name         string `json:"name"`
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

type links struct {
	Donate string `json:"donate,omitempty"`
	Stream string `json:"stream,omitempty"`
	Page   string `json:"page,omitempty"`
}
