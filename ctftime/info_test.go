package ctftime

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalCtfInfo(t *testing.T) {
	testJSON := `[{"organizers": [{"id": 11918, "name": "SECCON CTF"}], "onsite": true, "finish": "2018-02-18T09:00:00+00:00", "description": "", "weight": 0.00, "title": "SECCON 2017 Final Japan competition", "url": "http://2017.seccon.jp/", "is_votable_now": false, "restrictions": "Open", "format": "Jeopardy", "start": "2018-02-17T00:00:00+00:00", "participants": 11, "ctftime_url": "https://ctftime.org/event/511/", "location": "", "live_feed": "", "public_votable": true, "duration": {"hours": 9, "days": 1}, "logo": "", "format_id": 1, "id": 511, "ctf_id": 92}, {"organizers": [{"id": 37721, "name": "TAMUctf"}], "onsite": false, "finish": "2018-02-26T00:00:00+00:00", "description": "Texas A&M University CTF event. \r\n\r\nThis CTF is geared more towards university and high school students. \r\n\r\nWe are also adding a series of new secure coding challenges where you will be asked to find and fix vulnerabilities in code given to you.", "weight": 0.00, "title": "TAMUctf 18", "url": "https://ctf.tamu.edu/", "is_votable_now": false, "restrictions": "Open", "format": "Jeopardy", "start": "2018-02-17T00:00:00+00:00", "participants": 61, "ctftime_url": "https://ctftime.org/event/559/", "location": "", "live_feed": "", "public_votable": true, "duration": {"hours": 0, "days": 9}, "logo": "https://ctftime.org/media/events/tamuctf.png", "format_id": 1, "id": 559, "ctf_id": 198}, {"organizers": [{"id": 11918, "name": "SECCON CTF"}], "onsite": true, "finish": "2018-02-19T00:00:00+00:00", "description": "", "weight": 0.00, "title": "SECCON 2017 Final International competition", "url": "http://2017.seccon.jp/", "is_votable_now": false, "restrictions": "Open", "format": "Jeopardy", "start": "2018-02-18T00:00:00+00:00", "participants": 8, "ctftime_url": "https://ctftime.org/event/510/", "location": "", "live_feed": "", "public_votable": true, "duration": {"hours": 0, "days": 1}, "logo": "", "format_id": 1, "id": 510, "ctf_id": 92}, {"organizers": [{"id": 51003, "name": "CyberStart"}], "onsite": false, "finish": "2018-02-25T23:59:00+00:00", "description": "Registration open until 2/16/2018\r\nGirlsGoCyberStart is a CTF created to let girls \"try out\u201d cybersecurity to see what it's all about. CyberStart is a series of interactive challenges that explore exciting topics such as cryptography, penetration testing and digital forensics.  Girls form teams of up to 4 members , then collaborate in person or remotely as the competition is completely online.\r\n\r\nWho can play: females enrolled in 9th- 12th grade at a public or private school or the homeschool equivalent.   High school must be located in one of these states or territories: American Samoa, Colorado, Connecticut, Delaware, Hawaii, Indiana, Iowa, Maine, Maryland, Mississippi, Nevada, New Jersey, North Carolina, Texas, Vermont, West Virginia, Wyoming .", "weight": 0.00, "title": "GirlsGoCyberStart 2018", "url": "http://www.girlsgocyberstart.com/", "is_votable_now": false, "restrictions": "Open", "format": "Jeopardy", "start": "2018-02-20T14:00:00+00:00", "participants": 3, "ctftime_url": "https://ctftime.org/event/572/", "location": "", "live_feed": "", "public_votable": true, "duration": {"hours": 9, "days": 5}, "logo": "", "format_id": 1, "id": 572, "ctf_id": 238}, {"organizers": [{"id": 32118, "name": "NeverLAN"}], "onsite": false, "finish": "2018-02-27T00:00:00+00:00", "description": "A Middle School focused CTF, \r\nEveryone is free to join in the fun, but Middle school students are judged separately for prizes", "weight": 0.00, "title": "NeverLAN CTF 2018", "url": "http://neverlanctf.com/", "is_votable_now": false, "restrictions": "Open", "format": "Jeopardy", "start": "2018-02-23T19:00:00+00:00", "participants": 14, "ctftime_url": "https://ctftime.org/event/569/", "location": "", "live_feed": "", "public_votable": true, "duration": {"hours": 5, "days": 3}, "logo": "https://ctftime.org/media/events/Neverlanctflogo.jpg", "format_id": 1, "id": 569, "ctf_id": 187}]`
	var info []CtfInfo
	json.Unmarshal([]byte(testJSON), &info)

	type testCase struct {
		Name string
		Test string
		Want string
	}
	testCases := []testCase{
		{
			Name: "test marshal title",
			Test: info[0].Title,
			Want: "SECCON 2017 Final Japan competition",
		},
		{
			Name: "test marshal URL",
			Test: info[0].URL,
			Want: "http://2017.seccon.jp/",
		},
		{
			Name: "test marshal CtftimeUrl",
			Test: info[0].CtftimeURL,
			Want: "https://ctftime.org/event/511/",
		},
	}

	for _, test := range testCases {
		if test.Test != test.Want {
			t.Errorf("erro ctf info unmarshal error (%v) :%v != %v", test.Name, test.Test, test.Want)
		}
	}
}
