package tableau

import "time"

const (
	timestampFormat = "2006-01-02T15:04:05.999999Z"
)

type Tags struct {
	Tag []struct {
		Label string `json:"label"`
	} `json:"tag"`
}

type WorkbookPagination struct {
	PageNumber     string `json:"pageNumber"`
	PageSize       string `json:"pageSize"`
	TotalAvailable string `json:"totalAvailable"`
}

type Workbook struct {
	BaseModel
	ViewsResponse
	Project                BaseIdentity `json:"project"`
	Owner                  BaseIdentity `json:"owner"`
	Tags                   Tags         `json:"tags"`
	DataAccelerationConfig interface{}  `json:"dataAccelerationConfig"`
	ContentURL             string       `json:"contentUrl"`
	WebpageURL             string       `json:"webpageUrl"`
	ShowTabs               string       `json:"showTabs"`
	Size                   string       `json:"size"`
	EncryptExtracts        string       `json:"encryptExtracts"`
	DefaultViewID          string       `json:"defaultViewId"`
	Description            string       `jsob:"description"`
	Location               struct {
		BaseIdentity
		Type string `json:"type"`
	} `json:"location"`
}

type WorkbooksResponse struct {
	Pagination WorkbookPagination `json:"pagination"`
	Workbooks  struct {
		Workbook []Workbook `json:"workbook"`
	} `json:"workbooks"`
}

type View struct {
	BaseModel
	Tags        Tags   `json:"tags"`
	ContentURL  string `json:"contentUrl"`
	ViewURLName string `json:"viewUrlName"`
}

type ViewsResponse struct {
	Views struct {
		View []View `json:"view"`
	} `json:"views"`
}

type WorkbookConnection struct {
	BaseIdentity
	Datasource          BaseIdentity `json:"datasource"`
	Type                string       `json:"type"`
	EmbedPassword       bool         `json:"embedPassword"`
	ServerAddress       string       `json:"serverAddress"`
	UserName            string       `json:"userName"`
	QueryTaggingEnabled bool         `json:"queryTaggingEnabled"`
}

type BaseIdentity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BaseModel struct {
	BaseIdentity
	CreatedAtString string `json:"createdAt"`
	UpdatedAtString string `json:"updatedAt"`
}

func (m *BaseModel) CreatedAt() (time.Time, error) {
	return time.Parse(timestampFormat, m.CreatedAtString)
}

func (m *BaseModel) UpdatedAt() (time.Time, error) {
	return time.Parse(timestampFormat, m.UpdatedAtString)
}
