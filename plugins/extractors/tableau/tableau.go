package tableau

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/odpf/meteor/models"
	"github.com/odpf/meteor/models/odpf/assets"
	"github.com/odpf/meteor/models/odpf/assets/common"
	"github.com/odpf/meteor/models/odpf/assets/facets"
	"github.com/odpf/meteor/plugins"
	"github.com/odpf/meteor/registry"
	"github.com/odpf/meteor/utils"
	"github.com/odpf/salt/log"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//go:embed README.md
var summary string

var sampleConfig = `
host: https://server.tableau.com
version: 3.13
sitename: testdev550928
username: meteor_user
password: xxxxxxxxxx
`

// Config that holds a set of configuration for tableau extractor
type Config struct {
	Host     string `mapstructure:"host" validate:"required"`
	Version  string `mapstructure:"version" validate:"required"` // float as string
	Username string `mapstructure:"username" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	Sitename string `mapstructure:"sitename" validate:"required"`
}

// Extractor manages the extraction of data
// from tableau server
type Extractor struct {
	config    Config
	logger    log.Logger
	client    *http.Client
	authToken string
	siteID    string
}

// Option provides extension abstraction to Extractor constructor
type Option func(*Extractor)

// WithHTTPClient assign custom http client to the Extractor constructor
func WithHTTPClient(cl *http.Client) Option {
	return func(e *Extractor) {
		e.client = cl
	}
}

// New returns pointer to an initialized Extractor Object
func New(logger log.Logger, opts ...Option) *Extractor {

	e := &Extractor{
		logger: logger,
		client: &http.Client{},
	}

	for _, opt := range opts {
		opt(e)
	}

	e.client.Timeout = 30 * time.Second

	return e
}

// Info returns the brief information of the extractor
func (e *Extractor) Info() plugins.Info {
	return plugins.Info{
		Description:  "Dashboard list from Tableau server",
		SampleConfig: sampleConfig,
		Summary:      summary,
		Tags:         []string{"oss", "extractor"},
	}
}

// Validate validates the configuration of the extractor
func (e *Extractor) Validate(configMap map[string]interface{}) (err error) {
	return utils.BuildConfig(configMap, &Config{})
}

func (e *Extractor) Init(ctx context.Context, configMap map[string]interface{}) (err error) {
	// build and validate config
	err = utils.BuildConfig(configMap, &e.config)
	if err != nil {
		return plugins.InvalidConfigError{}
	}

	authToken, siteID, err := e.getAuthToken()
	if err != nil {
		return errors.Wrap(err, "failed to fetch auth token")
	}
	e.authToken = authToken
	e.siteID = siteID

	return nil
}

// Extract collects metadata from the source. The metadata is collected through the out channel
func (e *Extractor) Extract(ctx context.Context, emit plugins.Emit) (err error) {
	workbooks, err := e.fetchWorkbooks()
	if err != nil {
		return errors.Wrap(err, "failed to fetch list of workbook")
	}
	for _, w := range workbooks {
		dashboard, err := e.buildDashboard(w)
		if err != nil {
			return errors.Wrap(err, "failed to fetch dashboard data")
		}
		emit(models.NewRecord(dashboard))
	}
	return nil
}

func (e *Extractor) buildDashboard(w Workbook) (data *assets.Dashboard, err error) {
	var wbConnections []WorkbookConnection
	var workbook Workbook

	var wg sync.WaitGroup
	cherr := make(chan error)
	wg.Add(2)

	go func() {
		defer wg.Done()
		wbConnections, err = e.fetchWorkbookConnections(w.ID)
		if err != nil {
			cherr <- errors.Wrapf(err, "error fetching workbook connections")
		}
	}()

	go func() {
		defer wg.Done()
		workbook, err = e.fetchWorkbookDetail(w.ID)
		if err != nil {
			cherr <- errors.Wrapf(err, "error fetching workbook detail")
		}
	}()

	go func() {
		wg.Wait()
		close(cherr)
	}()

	for err = range cherr {
		if err != nil {
			return
		}
	}

	dashboardURN := fmt.Sprintf("tableau::%s/workbook/%s", workbook.Project.Name, workbook.ID)
	charts, err := e.buildCharts(dashboardURN, workbook, wbConnections)
	if err != nil {
		err = errors.Wrapf(err, "error building charts")
		return
	}

	createdAt, updatedAt, err := buildTimestamps(workbook.BaseModel)
	if err != nil {
		err = errors.Wrapf(err, "error building dashboard timestamps")
		return
	}

	tags := []interface{}{}
	for _, tg := range workbook.Tags.Tag {
		tags = append(tags, tg.Label)
	}

	data = &assets.Dashboard{
		Resource: &common.Resource{
			Urn:     dashboardURN,
			Name:    workbook.Name,
			Service: "tableau",
		},
		Description: workbook.Description,
		Charts:      charts,
		Properties: &facets.Properties{
			Attributes: utils.TryParseMapToProto(map[string]interface{}{
				"id":            workbook.ID,
				"project_id":    workbook.Project.ID,
				"project_name":  workbook.Project.Name,
				"owner_id":      workbook.Owner.ID,
				"owner_name":    workbook.Owner.Name,
				"location_id":   workbook.Location.ID,
				"location_name": workbook.Location.Name,
				"location_type": workbook.Location.Type,
				"content_url":   workbook.ContentURL,
				"webpage_url":   workbook.WebpageURL,
				"tag":           tags,
			}),
		},
		Timestamps: &common.Timestamp{
			CreateTime: timestamppb.New(createdAt),
			UpdateTime: timestamppb.New(updatedAt),
		},
	}
	return
}

func (e *Extractor) buildCharts(dashboardURN string, workbook Workbook, wbConns []WorkbookConnection) (charts []*assets.Chart, err error) {
	for _, vw := range workbook.Views.View {

		createdAt, updatedAt, perr := buildTimestamps(vw.BaseModel)
		if perr != nil {
			err = errors.Wrapf(perr, "error building chart %s with ID %s timestamps", vw.Name, vw.ID)
			return
		}

		tags := []interface{}{}
		for _, tg := range vw.Tags.Tag {
			tags = append(tags, tg.Label)
		}

		charts = append(charts, &assets.Chart{
			Urn:          fmt.Sprintf("tableau::%s/view/%s", workbook.Project.Name, vw.ID),
			DashboardUrn: dashboardURN,
			Source:       "tableau",
			Lineage:      buildLineage(wbConns),
			Properties: &facets.Properties{
				Attributes: utils.TryParseMapToProto(map[string]interface{}{
					"id":            vw.ID,
					"name":          vw.Name,
					"content_url":   vw.ContentURL,
					"view_url_name": vw.ViewURLName,
					"tag":           tags,
				}),
			},
			Timestamps: &common.Timestamp{
				CreateTime: timestamppb.New(createdAt),
				UpdateTime: timestamppb.New(updatedAt),
			},
		})
	}
	return
}

func (e *Extractor) fetchWorkbookConnections(workbookID string) (data []WorkbookConnection, err error) {
	workbookDetailPath := fmt.Sprintf("sites/%s/workbooks/%s/connections", e.siteID, workbookID)
	workbookDetailURL := e.buildURL(workbookDetailPath)

	type WorkbookConnectionsWrapper struct {
		Connections struct {
			Connection []WorkbookConnection `json:"connection"`
		} `json:"connections"`
	}

	var wcw WorkbookConnectionsWrapper
	err = e.makeRequest(http.MethodGet, workbookDetailURL, nil, &wcw)
	if err != nil {
		return
	}

	data = wcw.Connections.Connection
	return
}

func (e *Extractor) fetchWorkbookDetail(workbookID string) (data Workbook, err error) {
	workbookDetailPath := fmt.Sprintf("sites/%s/workbooks/%s", e.siteID, workbookID)
	workbookDetailURL := e.buildURL(workbookDetailPath)

	type WorkbookWrapper struct {
		Workbook Workbook `json:"workbook"`
	}

	var ww WorkbookWrapper
	err = e.makeRequest(http.MethodGet, workbookDetailURL, nil, &ww)
	if err != nil {
		return
	}

	data = ww.Workbook
	return
}

func (e *Extractor) fetchWorkbooks() (data []Workbook, err error) {
	//! TODO: might need to change this to paginated call for batch ingestion
	workbooksPath := fmt.Sprintf("sites/%s/workbooks?pageSize=1000&pageNumber=1", e.siteID)
	workbooksURL := e.buildURL(workbooksPath)

	var workboksResponse WorkbooksResponse
	err = e.makeRequest(http.MethodGet, workbooksURL, nil, &workboksResponse)
	if err != nil {
		return nil, err
	}

	data = workboksResponse.Workbooks.Workbook
	return
}

func (e *Extractor) getAuthToken() (authToken string, siteID string, err error) {
	payload := map[string]interface{}{
		"credentials": map[string]interface{}{
			"name":     e.config.Username,
			"password": e.config.Password,
			"site": map[string]interface{}{
				"contentUrl": e.config.Sitename,
			},
		},
	}

	type responseSignIn struct {
		Credentials struct {
			Site struct {
				ID         string `json:"id"`
				ContentURL string `json:"contentUrl"`
			} `json:"site"`
			User struct {
				ID string `json:"id"`
			} `json:"user"`
			Token string `json:"token"`
		} `json:"credentials"`
	}

	var data responseSignIn
	signInURL := e.buildURL("auth/signin")
	err = e.makeRequest(http.MethodPost, signInURL, payload, &data)
	if err != nil {
		return
	}
	return data.Credentials.Token, data.Credentials.Site.ID, nil
}

func (e *Extractor) buildURL(path string) string {
	return fmt.Sprintf("%s/api/%s/%s", e.config.Host, e.config.Version, path)
}

// helper function to avoid rewriting a request
func (e *Extractor) makeRequest(method, url string, payload interface{}, data interface{}) (err error) {
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return errors.Wrap(err, "failed to encode the payload JSON")
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Tableau-Auth", e.authToken)

	res, err := e.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to generate response")
	}
	if res.StatusCode >= 300 {
		return fmt.Errorf("getting %d status code", res.StatusCode)
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read response body")
	}
	if err = json.Unmarshal(bytes, &data); err != nil {
		return errors.Wrapf(err, "failed to parse: %s", string(bytes))
	}

	return
}

// Register the extractor to catalog
func init() {
	if err := registry.Extractors.Register("tableau", func() plugins.Extractor {
		return New(plugins.GetLog())
	}); err != nil {
		panic(err)
	}
}
