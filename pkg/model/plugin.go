package model

type Plugin struct {
	ID           int      `gorm:"column:id" db:"id" json:"id" form:"id"`
	Domain       string   `gorm:"column:domain" db:"domain" json:"domain" form:"domain"`
	Name         string   `gorm:"column:name" db:"name" json:"name" form:"name"`
	Description  string   `gorm:"column:description" db:"description" json:"description" form:"description"`
	AuthType     string   `gorm:"column:auth_type" db:"auth_type" json:"auth_type" form:"auth_type"`
	LogoURL      string   `gorm:"column:logo_url" db:"logo_url" json:"logo_url" form:"logo_url"`
	ContactEmail string   `gorm:"column:contact_email" db:"contact_email" json:"contact_email" form:"contact_email"`
	Organization string   `gorm:"column:organization" db:"organization" json:"organization" form:"organization"`
	APIType      string   `gorm:"column:api_type" db:"api_type" json:"api_type" form:"api_type"`
	APIURL       string   `gorm:"column:api_url" db:"api_url" json:"api_url" form:"api_url"`
	Label        []string `gorm:"column:label" db:"label" json:"label" form:"label"`
	State        string   `gorm:"column:state" db:"state" json:"state" form:"state"`
	InstallNum   int      `gorm:"column:install_num" db:"install_num" json:"install_num" form:"install_num"`
	Score        int      `gorm:"column:score" db:"score" json:"score" form:"score"`
	Heat         int      `gorm:"column:heat" db:"heat" json:"heat" form:"heat"`
	CreatedAt    string   `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt    string   `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
}

type OrganizationType string

const (
	// OrganizationTypeUser is this plugin user organization type.
	OrganizationTypeUser OrganizationType = "user"
	// OrganizationTypeTeam is this plugin team organization type.
	OrganizationTypeTeam OrganizationType = "team"
)

func (ot OrganizationType) String() string {
	return string(ot)
}

type StateType string

const (
	// StateTypePublished is this plugin publish state.
	StateTypePublished StateType = "published"
	// StateTypeUnPublished is this plugin unpublish state.
	StateTypeUnPublished StateType = "unpublished"
)

func (st StateType) String() string {
	return string(st)
}
