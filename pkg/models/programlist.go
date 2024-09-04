package models

import (
	"time"

	"github.com/google/uuid"
)

type Program struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	NameEn    string    `json:"nameEn"`
	IsPublic  bool      `json:"isPublic"`
	ProjectID uuid.UUID `json:"projectId"`
}

type TextInfo struct {
	Lang       int    `json:"lang"`
	Name       string `json:"name"`
	NamePlural string `json:"namePlural"`
}

type SessionSubject struct {
	ID        int        `json:"id"`
	Color     string     `json:"color"`
	TextInfos []TextInfo `json:"textInfos"`
	Childs    []string   `json:"childs"`
}

type SessionType struct {
	ID        int        `json:"id"`
	Color     string     `json:"color"`
	TextInfos []TextInfo `json:"textInfos"`
	Childs    []string   `json:"childs"`
}

type SessionZone struct {
	ID        int        `json:"id"`
	TextInfos []TextInfo `json:"textInfos"`
}

type SessionSpeakerGroup struct {
	ID        int        `json:"id"`
	Color     string     `json:"color"`
	TextInfos []TextInfo `json:"textInfos"`
	Childs    []string   `json:"childs"`
}

type SpeakerTextInfo struct {
	Lang       int    `json:"lang"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	LastName   string `json:"lastName"`
	Prefix     string `json:"prefix"`
	Position   string `json:"position"`
	Biography  string `json:"biography"`
}

type SessionSpeakerAttribute struct {
	ID         int        `json:"id"`
	SessionID  int        `json:"sessionId"`
	OnlineMode int        `json:"onlineMode"`
	TextInfos  []TextInfo `json:"textInfos"`
}

type SessionSpeaker struct {
	ID                int                       `json:"id"`
	Photo             string                    `json:"photo"`
	NotShowInSpeakers bool                      `json:"notShowInSpeakers"`
	IsImportant       bool                      `json:"isImportant"`
	ImportantSort     int                       `json:"importantSort"`
	IsOnline          bool                      `json:"isOnline"`
	AdditionalInfo    string                    `json:"additionalInfo"`
	TextInfos         []SpeakerTextInfo         `json:"textInfos"`
	SessionAttributes []SessionSpeakerAttribute `json:"sessionAttributes"`
}

type SessionTextInfo struct {
	Lang                int    `json:"lang"`
	Title               string `json:"title"`
	Description         string `json:"description"`
	Organizer           string `json:"organizer"`
	OrganizerContact    string `json:"organizerContact"`
	DescriptionFileName string `json:"descriptionFileName"`
}

type SessionSpeakerGroupSpeakers struct {
	ID                 int `json:"id"`
	SessionAttributeID int `json:"sessionAttributeId"`
}

type SessionSpeakerGroupInfo struct {
	ID       int                           `json:"id"`
	Speakers []SessionSpeakerGroupSpeakers `json:"speakers"`
}

type Session struct {
	ID                   int                       `json:"id"`
	Date                 time.Time                 `json:"date"`
	StartTime            string                    `json:"startTime"`
	EndTime              string                    `json:"endTime"`
	SubjectID            int                       `json:"subjectId"`
	TypeID               int                       `json:"typeId"`
	ZoneID               int                       `json:"zoneId"`
	BroadcastStreamUrl   string                    `json:"broadcastStreamUrl"`
	BroadcastStreamUrlEn string                    `json:"broadcastStreamUrlEn"`
	BroadcastSiteUrl     string                    `json:"broadcastSiteUrl"`
	BroadcastSiteUrlEn   string                    `json:"broadcastSiteUrlEn"`
	SpeakerGroups        []SessionSpeakerGroupInfo `json:"speakerGroups"`
	TextInfos            []SessionTextInfo         `json:"textInfos"`
	ClosedEvent          bool                      `json:"closedEvent"`
}

type ProgramData struct {
	Program              Program               `json:"program"`
	SessionSubjects      []SessionSubject      `json:"sessionSubjects"`
	SessionTypes         []SessionType         `json:"sessionTypes"`
	SessionZones         []SessionZone         `json:"sessionZones"`
	SessionSpeakerGroups []SessionSpeakerGroup `json:"sessionSpeakerGroups"`
	SessionSpeakers      []SessionSpeaker      `json:"sessionSpeakers"`
	Sessions             []Session             `json:"sessions"`
}
