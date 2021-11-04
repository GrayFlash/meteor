package tableau

import (
	"testing"
	"time"

	"github.com/odpf/meteor/models/odpf/assets/common"
	"github.com/odpf/meteor/models/odpf/assets/facets"
	"github.com/stretchr/testify/assert"
)

func TestBuildTimestamp(t *testing.T) {
	t.Run("build timestamp to time.Time data type success", func(t *testing.T) {
		bm := BaseModel{
			CreatedAtString: "2021-11-01T11:32:35Z",
			UpdatedAtString: "2021-11-02T06:07:18Z",
		}

		ca, ua, err := buildTimestamps(bm)

		assert.Equal(t, time.Date(2021, time.November, 1, 11, 32, 35, 0, time.UTC), ca)
		assert.Equal(t, time.Date(2021, time.November, 2, 6, 7, 18, 0, time.UTC), ua)
		assert.Nil(t, err)
	})

	t.Run("build timestamp failed when createdAt is invalid", func(t *testing.T) {
		bm := BaseModel{
			CreatedAtString: "20211101T113235Z",
			UpdatedAtString: "2021-11-02T06:07:18Z",
		}

		ca, ua, err := buildTimestamps(bm)

		var timeError *time.ParseError
		assert.ErrorAs(t, err, &timeError)
		assert.Equal(t, time.Time{}, ca)
		assert.Equal(t, time.Time{}, ua)
	})

	t.Run("build timestamp failed when updatedAt is invalid", func(t *testing.T) {
		bm := BaseModel{
			CreatedAtString: "2021-11-01T11:32:35Z",
			UpdatedAtString: "20211102T060718Z",
		}

		ca, ua, err := buildTimestamps(bm)

		var timeError *time.ParseError
		assert.ErrorAs(t, err, &timeError)
		assert.Equal(t, time.Date(2021, time.November, 1, 11, 32, 35, 0, time.UTC), ca)
		assert.Equal(t, time.Time{}, ua)
	})

}

func TestBuildLineage(t *testing.T) {

	wbc := []WorkbookConnection{
		{
			Datasource: BaseIdentity{
				ID:   "connection_id1",
				Name: "Data Source 1",
			},
			Type:          "excel-direct",
			EmbedPassword: false,
		}, {
			Datasource: BaseIdentity{
				ID:   "connection_id2",
				Name: "Data Source 2",
			},
			Type:          "postgres",
			EmbedPassword: false,
		}}

	expectedLineage := &facets.Lineage{
		Upstreams: []*common.Resource{
			{
				Urn:  "excel-direct::connection_id1",
				Name: "Data Source 1",
			},
			{
				Urn:  "postgres::connection_id2",
				Name: "Data Source 2",
			},
		},
	}

	assert.Equal(t, expectedLineage, buildLineage(wbc))
}
