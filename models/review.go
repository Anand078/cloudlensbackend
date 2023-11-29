package models

import "time"

// Architecture Review Details
type Reviews struct {
	ID           int64   `json:"id,omitempty"`
	ProjectName  string  `json:"projectname,omitempty"`
	ProjectOwner string  `json:"projectowner,omitempty"`
	Reviewer     string  `json:"reviewer,omitempty"`
	Auditor      string  `json:"auditor,omitempty"`
	StatusId     int32   `json:"statusid,omitempty"`
	StartDate    string  `json:"startdate,omitempty"`
	EndDate      string  `json:"enddate,omitempty"`
	ProjectScore float64 `json:"projectscore,omitempty"`
	IsActive     string  `json:"isactive,omitempty"`
}

type Badge struct {
	ID    int64  `json:"id,omitempty"`
	Badge string `json:"badge,omitempty"`
}

type ArbStatus struct {
	ID     int64  `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
}

type ArbStatusId struct {
	StatusId int32 `json:"statusid,omitempty"`
}

type BestPractice struct {
	Id          int32  `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}
type Pillars struct {
	ID         int64  `json:"id,omitempty"`
	PillarName string `json:"pillarname,omitempty"`
	CreatedAt  string `json:"createdat,omitempty"`
	IsActive   string `json:"isactive,omitempty"`
}

type Topics struct {
	ID        int64  `json:"id,omitempty"`
	PillarId  string `json:"pillarid,omitempty"`
	Topic     string `json:"topic,omitempty"`
	CreatedAt string `json:"createdat,omitempty"`
	IsActive  string `json:"isactive,omitempty"`
}

type ARBReviews struct {
	PillarID       int64  `json:"pillarid,omitempty"`
	PillarName     string `json:"pillarname,omitempty"`
	TopicID        int64  `json:"topicid,omitempty"`
	Topic          string `json:"topic,omitempty"`
	BestPracticeId int64  `json:"bestpracticeid,omitempty"`
	BestPractice   string `json:"bestpractice,omitempty"`
	Description    string `json:"description,omitempty"`
}

type TecMember struct {
	ID          int64     `json:"id,omitempty"`
	Member      string    `json:"member,omitempty"`
	Project     string    `json:"project,omitempty"`
	CoreSkills  string    `json:"coreskills,omitempty"`
	IsAvailable bool      `json:"isavailable,omitempty"`
	Comments    string    `json:"comments,omitempty"`
	UpdatedOn   time.Time `json:"updatedon,omitempty"`
}

type SaveTecMember struct {
	ID          int64     `json:"id,omitempty"`
	Member      string    `json:"member,omitempty"`
	Project     string    `json:"project,omitempty"`
	CoreSkills  string    `json:"coreskills,omitempty"`
	Comments    string    `json:"comments,omitempty"`
	IsAvailable uint8     `json:"isavailable,omitempty"`
	UpdatedOn   time.Time `json:"updatedon,omitempty"`
}

type AccSnap struct {
	ID                  int64     `json:"id,omitempty"`
	AccName             string    `json:"accname,omitempty"`
	Version             string    `json:"version,omitempty"`
	IndicativeTimeline  string    `json:"indicativetimeline,omitempty"`
	ResourceRequirement string    `json:"resourcerequirement,omitempty"`
	Blocker             string    `json:"blocker,omitempty"`
	Comments            string    `json:"comments,omitempty"`
	UpdatedOn           time.Time `json:"updatedon,omitempty"`
}

type TECTimeline struct {
	ID        int64     `json:"id,omitempty"`
	TECId     int64     `json:"tecid,omitempty"`
	Comments  string    `json:"comments,omitempty"`
	UpdatedOn time.Time `json:"updatedon,omitempty"`
}

type AccTimeline struct {
	ID        int64     `json:"id,omitempty"`
	AccId     int64     `json:"accid,omitempty"`
	Comments  string    `json:"comments,omitempty"`
	UpdatedOn time.Time `json:"updatedon,omitempty"`
}

type Skills struct {
	Skills string `json:"skill,omitempty"`
}
