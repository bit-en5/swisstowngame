package swisspost

const (
	apiURL  = "https://swisspost.opendatasoft.com/api/v2/"
	dataset = "catalog/datasets/politische-gemeinden_v2/records"
	dSelect = "select=bfsnr%2Cgemeindename%2Ckanton"
	dWhere  = "where=bfsnr%3D"
	dLimit  = "limit=10"
	dOffset = "offset=0"
	dTZ     = "timezone=UTC"
	dKey    = "apiKey="
)

type Swisspost struct {
	apiKey string
}

// New constructs a Swisspost object
func New(apiKey string) Swisspost {
	if apiKey == "" {
		apiKey = getSecretAPIKey()
	}

	return Swisspost{
		apiKey: apiKey,
	}
}

type Response struct {
	Records []Records `json:"records"`
}

type Records struct {
	Record Record `json:"record"`
}

type Record struct {
	Fields Fields `json:"fields"`
}

type Fields struct {
	GemeindeName string `json:"gemeindename"`
	Kanton       string `json:"kanton"`
}
