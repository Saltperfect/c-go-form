package models

type Page struct {
	Title string
	Body  []byte
}

type Attributes struct {
	Label       string `json:"label"`
	Placeholder string `json:"placeholder"`
}

type Element struct {
	Type  string       `json:"type"`
	Value string       `json:"value"`
	List  []Attributes `json:"attributelist"`
}

type ElementList struct {
	Name string    `json:"name"`
	List []Element `json:"list"`
}

func (e *Element) IsRadioOrCheckbox() bool {
	return e.Type == "radio" || e.Type == "checkbox"
}

func (e *Element) IsText() bool {
	return e.Type == "text"
}

func (e *Element) IsSelect() bool {
	return e.Type == "select"
}
