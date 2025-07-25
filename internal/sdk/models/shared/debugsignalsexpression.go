// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Annotations struct {
}

type Image struct {
	Alt *string `json:"alt,omitempty"`
	Src *string `json:"src,omitempty"`
}

func (o *Image) GetAlt() *string {
	if o == nil {
		return nil
	}
	return o.Alt
}

func (o *Image) GetSrc() *string {
	if o == nil {
		return nil
	}
	return o.Src
}

type Link struct {
	Href *string `json:"href,omitempty"`
	Text *string `json:"text,omitempty"`
}

func (o *Link) GetHref() *string {
	if o == nil {
		return nil
	}
	return o.Href
}

func (o *Link) GetText() *string {
	if o == nil {
		return nil
	}
	return o.Text
}

type Signal struct {
	Annotations    *Annotations `json:"annotations,omitempty"`
	Body           *string      `json:"body,omitempty"`
	ID             *string      `json:"id,omitempty"`
	Images         []Image      `json:"images,omitempty"`
	Level          *string      `json:"level,omitempty"`
	Links          []Link       `json:"links,omitempty"`
	OrganizationID *string      `json:"organization_id,omitempty"`
	Summary        *string      `json:"summary,omitempty"`
	Tags           []string     `json:"tags,omitempty"`
}

func (o *Signal) GetAnnotations() *Annotations {
	if o == nil {
		return nil
	}
	return o.Annotations
}

func (o *Signal) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *Signal) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Signal) GetImages() []Image {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *Signal) GetLevel() *string {
	if o == nil {
		return nil
	}
	return o.Level
}

func (o *Signal) GetLinks() []Link {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *Signal) GetOrganizationID() *string {
	if o == nil {
		return nil
	}
	return o.OrganizationID
}

func (o *Signal) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *Signal) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

// DebugSignalsExpression - Debug Signals expressions
type DebugSignalsExpression struct {
	// CEL expression
	Expression string `json:"expression"`
	// List of signals to evaluate rule expression against
	Signals []Signal `json:"signals"`
}

func (o *DebugSignalsExpression) GetExpression() string {
	if o == nil {
		return ""
	}
	return o.Expression
}

func (o *DebugSignalsExpression) GetSignals() []Signal {
	if o == nil {
		return []Signal{}
	}
	return o.Signals
}
