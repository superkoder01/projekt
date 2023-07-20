package model

type Data struct {
	Header  Header                 `json:"Header"`
	Payload map[string]interface{} `json:"Payload"`
}

type Header struct {
	Version  string  `json:"Version"`
	Provider string  `json:"Provider"`
	Content  Content `json:"Content"`
}
type Content struct {
	Type     string `json:"Type"`
	Category string `json:"Catg"`
}
