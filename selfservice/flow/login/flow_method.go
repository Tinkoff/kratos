package login

import (
	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/text"
)

// swagger:model loginFlowMethod
type FlowMethod struct {
	// Method contains the methods' credentials type.
	//
	// required: true
	Method identity.CredentialsType `json:"method" db:"method"`

	// Config is the credential type's config.
	//
	// required: true
	Config *FlowMethodConfig `json:"config" db:"config"`
}

type FlowMethods map[identity.CredentialsType]*FlowMethod

// swagger:model loginFlowMethodConfig
type FlowMethodConfig struct {
	Action   string        `json:"action"`
	Method   string        `json:"method"`
	Fields   Fields        `json:"fields"`
	Messages text.Messages `json:"messages,omitempty"`
}

type Fields []Field

type Field struct {
	Name     string        `json:"name"`
	Type     string        `json:"type"`
	Pattern  string        `json:"pattern,omitempty"`
	Disabled bool          `json:"disabled,omitempty"`
	Required bool          `json:"required,omitempty"`
	Value    interface{}   `json:"value,omitempty" faker:"string"`
	Messages text.Messages `json:"messages,omitempty"`
}
