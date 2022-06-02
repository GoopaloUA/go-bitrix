package users

type User struct {
	ID                 string      `json:"ID"`
	Active             bool        `json:"ACTIVE"`
	DateRegister       string      `json:"DATE_REGISTER"`
	Email              string      `json:"EMAIL"`
	IsOnline           string      `json:"IS_ONLINE"`
	LastActivityDate   struct{}    `json:"LAST_ACTIVITY_DATE"`
	LastLogin          string      `json:"LAST_LOGIN"`
	LastName           string      `json:"LAST_NAME"`
	Name               string      `json:"NAME"`
	PersonalBirthday   string      `json:"PERSONAL_BIRTHDAY"`
	PersonalCity       string      `json:"PERSONAL_CITY"`
	PersonalCountry    string      `json:"PERSONAL_COUNTRY"`
	PersonalFax        string      `json:"PERSONAL_FAX"`
	PersonalGender     string      `json:"PERSONAL_GENDER"`
	PersonalIcq        string      `json:"PERSONAL_ICQ"`
	PersonalMailbox    string      `json:"PERSONAL_MAILBOX"`
	PersonalMobile     string      `json:"PERSONAL_MOBILE"`
	PersonalNotes      string      `json:"PERSONAL_NOTES"`
	PersonalPager      string      `json:"PERSONAL_PAGER"`
	PersonalPhone      string      `json:"PERSONAL_PHONE"`
	PersonalPhoto      string      `json:"PERSONAL_PHOTO"`
	PersonalProfession string      `json:"PERSONAL_PROFESSION"`
	PersonalState      string      `json:"PERSONAL_STATE"`
	PersonalStreet     string      `json:"PERSONAL_STREET"`
	PersonalWww        string      `json:"PERSONAL_WWW"`
	PersonalZip        string      `json:"PERSONAL_ZIP"`
	SecondName         string      `json:"SECOND_NAME"`
	TimestampX         struct{}    `json:"TIMESTAMP_X"`
	TimeZone           string      `json:"TIME_ZONE"`
	TimeZoneOffset     string      `json:"TIME_ZONE_OFFSET"`
	Title              string      `json:"TITLE"`
	UfDepartment       []int64     `json:"UF_DEPARTMENT"`
	UfDistrict         string      `json:"UF_DISTRICT"`
	UfEmploymentDate   string      `json:"UF_EMPLOYMENT_DATE"`
	UfFacebook         string      `json:"UF_FACEBOOK"`
	UfInterests        string      `json:"UF_INTERESTS"`
	UfLinkedin         string      `json:"UF_LINKEDIN"`
	UfPhoneInner       string      `json:"UF_PHONE_INNER"`
	UfSkills           string      `json:"UF_SKILLS"`
	UfSkype            string      `json:"UF_SKYPE"`
	UfSkypeLink        string      `json:"UF_SKYPE_LINK"`
	UfTimeman          interface{} `json:"UF_TIMEMAN"`
	UfTwitter          string      `json:"UF_TWITTER"`
	UfWebSites         string      `json:"UF_WEB_SITES"`
	UfXing             string      `json:"UF_XING"`
	UfZoom             string      `json:"UF_ZOOM"`
	UserType           string      `json:"USER_TYPE"`
	WorkCity           string      `json:"WORK_CITY"`
	WorkCompany        string      `json:"WORK_COMPANY"`
	WorkCountry        string      `json:"WORK_COUNTRY"`
	WorkDepartment     string      `json:"WORK_DEPARTMENT"`
	WorkFax            string      `json:"WORK_FAX"`
	WorkLogo           string      `json:"WORK_LOGO"`
	WorkMailbox        string      `json:"WORK_MAILBOX"`
	WorkNotes          string      `json:"WORK_NOTES"`
	WorkPager          string      `json:"WORK_PAGER"`
	WorkPhone          string      `json:"WORK_PHONE"`
	WorkPosition       string      `json:"WORK_POSITION"`
	WorkProfile        string      `json:"WORK_PROFILE"`
	WorkState          string      `json:"WORK_STATE"`
	WorkStreet         string      `json:"WORK_STREET"`
	WorkWww            string      `json:"WORK_WWW"`
	WorkZip            string      `json:"WORK_ZIP"`
	XMLID              string      `json:"XML_ID"`
}