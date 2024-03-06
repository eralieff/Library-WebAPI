package model

type Author struct {
	Id         int    `json:"id"`
	FullName   string `json:"full_name"`
	Nickname   string `json:"nickname"`
	Speciality string `json:"speciality"`
}
