/*
Instance models

Capture:

- Instance list
- Moderated instances
- Relationships
*/
package models

type Moderation struct {
	Moderator *MastodonInstance
	Moderated *MastodonInstance
	Severity  string
	Comment   string
}

/*
MastodonInstance reflects an instance in the broader mastodon ecosystem

/api/v2/instance => public
*/
type MastodonInstance struct {
	Name        string
	Description string
	Domain      string `boltholdUnique:"UniqueDomain"`
	Image       string
}

func (m *MastodonInstance) GetAbout() error {
	// /api/v2/instance
	return nil
}

func (m *MastodonInstance) GetModerated() error {
	// /api/v1/instance/domain_blocks
	return nil
}

/*
type BlockType int

const (
	Limited BlockType = iota
	Suspended
)


Represents a moderated instance

/api/v1/instance/domain_blocks

type InstanceBlock struct {
	Instance *MastodonInstance
	Reason   string
	Severity BlockType
}

/*
Represent data reflecting an instance

/api/v2/instance - PUBLIC

type MastodonInstance struct {
	Name        string
	Description string
	Domain      string
	Admin       string
	Contact     string
	Image       string
	Moderated   []*InstanceBlock
	ModeratedBy []*MastodonInstance `boltholdIndex:"ModeratedBy"`
}

// GetAbout scrapes the about page to populate data
func (mi *MastodonInstance) GetAbout() error {
	return nil
}
*/
