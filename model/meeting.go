package model


type Meeting struct {
	Title string
	Sponsor string
	Participators []string
	Start string
	End string
}

//the example of start and end: 2017-10-20T19:00
func (meeting *Meeting) Init(title string, sponsor string, 
	pariticipators []string, start string, end string) {
	meeting.Title = title
	meeting.Sponsor = sponsor
	meeting.Participators = pariticipators
	meeting.Start = start
	meeting.End = end
}

func (meeting Meeting) GetTitle() string {
	return meeting.Title
}

func (meeting *Meeting) SetTitle(title string) {
	meeting.Title = title
}

func (meeting Meeting) GetSponsor() string {
	return meeting.Sponsor
}

func (meeting *Meeting) SetSponsor(sponsor string) {
	meeting.Sponsor = sponsor
}

func (meeting Meeting) GetParticipators() []string {
	return meeting.Participators
}

func (meeting *Meeting) SetParticipators(pariticipators []string) {
	meeting.Participators = pariticipators
}

func (meeting Meeting) GetStart() string {
	return meeting.Start
}

func (meeting *Meeting) SetStart(start string) {
	meeting.Start = start
}

func (meeting Meeting) GetEnd() string {
	return meeting.End
}

func (meeting *Meeting) SetEnd(end string) {
	meeting.End = end
}

func (meeting Meeting) IsSponsor(username string) bool {
	if meeting.Sponsor == username {
		return true
	}
	return false
}

func (meeting Meeting) IsParticipator(username string) bool {
	for _, participator := range meeting.GetParticipators() {
		if participator == username {
			return true
		}
	}
	return false
}

//bool好像该改成err
func (meeting *Meeting) AddParticipator(username string) bool {
	if meeting.IsParticipator(username) { //判断是否已在参与者当中
		return false
	}
	meeting.SetParticipators(append(meeting.GetParticipators(), username))
	return false
}

//bool好像该改成err
func (meeting *Meeting) DeleteParticipator(username string) bool {
	for i := 0; i < len(meeting.GetParticipators()); i++ {
		if meeting.GetParticipators()[i] == username {
			meeting.SetParticipators(append(meeting.GetParticipators()[: i], meeting.GetParticipators()[i + 1:]...))
			return true
		}
	}
	return false
}

func (meeting Meeting) GetParticipatorsLength() int {
	return len(meeting.GetParticipators())
}