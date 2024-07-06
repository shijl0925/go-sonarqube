package duplications

type ShowRequest struct {
	Key string `url:"key"` // File key
}

type ShowResponse struct {
	Duplications []Duplication `json:"duplications,omitempty"`
	Files        map[string]struct {
		Key         string `json:"key,omitempty"`
		Name        string `json:"name,omitempty"`
		projectName string `json:"projectName,omitempty"`
	} `json:"files,omitempty"`
}

type Duplication struct {
	Blocks []Block `json:"blocks,omitempty"`
}

type Block struct {
	Ref  string `json:"_ref,omitempty"`
	From int    `json:"from,omitempty"`
	Size int    `json:"size,omitempty"`
}
