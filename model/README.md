## User
```
type User struct {
	Username string
	Password string
	Email string
	Phone string
}

func (user *User) Init(name, psw, email, phone string) 

func (user User) GetUsername() string 

func (user *User) SetUsername(username string) 

func (user User) GetPassword() string 

func (user *User) SetPassword(password string) 

func (user User) GetEmail() string 

func (user *User) SetEmail(email string) 

func (user User) GetPhone() string 

func (user *User) SetPhone(phone string) 

```


## Meeting

type Meeting struct {
	Title string
	Sponsor string
	Participators []string
	Start string
	End string
}

func (meeting *Meeting) Init(title string, sponsor string, 
	pariticipators []string, start string, end string) 

func (meeting Meeting) GetTitle() string 

func (meeting *Meeting) SetTitle(title string) 

func (meeting Meeting) GetSponsor() string 

func (meeting *Meeting) SetSponsor(sponsor string) 

func (meeting Meeting) GetParticipators() []string 

func (meeting *Meeting) SetParticipators(pariticipators []string) 

func (meeting Meeting) GetStart() string 

func (meeting *Meeting) SetStart(start string) 

func (meeting Meeting) GetEnd() string 

func (meeting *Meeting) SetEnd(end string) 

func (meeting Meeting) IsSponsor(username string) bool

func (meeting Meeting) IsParticipator(username string) bool 

//bool好像该改成err
func (meeting *Meeting) AddParticipator(username string) bool 

//bool好像该改成err
func (meeting *Meeting) DeleteParticipator(username string) bool 

func (meeting Meeting) GetParticipatorsLength() int 

```

## Date

```

type Date struct {
	Time time.Time
}

func (date *Date) Init(st string) 

func (date Date) GetYear() int 

func (date *Date) SetYear(year int) 

//待修改
func (date Date) GetMonth() time.Month 

func (date *Date) SetMonth(month int) 

func (date Date) GetDay() int

func (date *Date) SetDay(day int)

func (date Date) GetHour() int 

func (date *Date) SetHour(hour int) 

func (date Date) GetMinute() int 

func (date *Date) SetMinute(minute int) 

func (date Date)IsEqual(other Date) bool

func (date Date)IsAfter(other Date) bool 

//待修改
func (date Date)IsValid() bool 

//待修改
func (date Date)DateToString() string 

//待修改
func StringToDate(date string) (time.Time, error) 

```