package ctftime

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// CtfInfo is ...
type CtfInfo struct {
	Organizers []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"organizers"`
	Onsite        bool      `json:"onsite"`
	Finish        time.Time `json:"finish"`
	Description   string    `json:"description"`
	Weight        float64   `json:"weight"`
	Title         string    `json:"title"`
	URL           string    `json:"url"`
	IsVotableNow  bool      `json:"is_votable_now"`
	Restrictions  string    `json:"restrictions"`
	Format        string    `json:"format"`
	Start         time.Time `json:"start"`
	Participants  int       `json:"participants"`
	CtftimeURL    string    `json:"ctftime_url"`
	Location      string    `json:"location"`
	LiveFeed      string    `json:"live_feed"`
	PublicVotable bool      `json:"public_votable"`
	Duration      struct {
		Hours int `json:"hours"`
		Days  int `json:"days"`
	} `json:"duration"`
	Logo     string `json:"logo"`
	FormatID int    `json:"format_id"`
	ID       int    `json:"id"`
	CtfID    int    `json:"ctf_id"`
}

// GetInfo is ...
func GetInfo() []CtfInfo {
	t := time.Now()
	start := t.Unix()
	finish := t.Add(14 * 24 * time.Hour).Unix()
	values := url.Values{}
	values.Add("start", strconv.FormatInt(start, 10))
	values.Add("finish", strconv.FormatInt(finish, 10))
	req, err := http.NewRequest("GET", "https://ctftime.org/api/v1/events/", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var info []CtfInfo
	json.Unmarshal(body, &info)
	return info

}
