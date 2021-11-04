//go:build integration
// +build integration

package tableau_test

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/v2/recorder"
	"github.com/odpf/meteor/models"
	"github.com/odpf/meteor/plugins"
	"github.com/odpf/meteor/plugins/extractors/tableau"
	"github.com/odpf/meteor/test/mocks"
	testutils "github.com/odpf/meteor/test/utils"
	"github.com/stretchr/testify/assert"
)

var (
	host     = "https://server.tableau.com"
	version  = "3.13"
	sitename = "testdev550928"
	username = "user@meteor"
	password = "password"
)

func TestInit(t *testing.T) {
	t.Run("should return error for invalid config", func(t *testing.T) {
		err := tableau.New(testutils.Logger).Init(context.TODO(), map[string]interface{}{
			"host": "invalid_host",
		})

		assert.Equal(t, plugins.InvalidConfigError{}, err)
	})
}

func TestExtract(t *testing.T) {
	t.Run("should return dashboard model", func(t *testing.T) {
		r, err := recorder.New("fixtures/extract_dashboard_model")
		if err != nil {
			log.Fatal(err)
		}
		defer r.Stop()

		ctx := context.TODO()
		extr := tableau.New(testutils.Logger,
			tableau.WithHTTPClient(&http.Client{
				Transport: r,
			}))
		err = extr.Init(ctx, map[string]interface{}{
			"host":     host,
			"version":  version,
			"sitename": sitename,
			"username": username,
			"password": password,
		})
		if err != nil {
			t.Fatal(err)
		}

		emitter := mocks.NewEmitter()
		err = extr.Extract(ctx, emitter.Push)
		assert.NoError(t, err)

		records := emitter.Get()
		var actuals []models.Metadata
		for _, r := range records {
			actuals = append(actuals, r.Data())
		}

		assertJSONString(t, expectedDashboardsInJSONString, actuals)
	})
}

func assertJSONString(t *testing.T, expected string, actual interface{}) {
	actualBytes, err := json.Marshal(actual)
	if err != nil {
		t.Fatal(err)
	}
	var prettyPrintActualBytes bytes.Buffer
	err = json.Indent(&prettyPrintActualBytes, []byte(actualBytes), "", "\t")
	assert.Nil(t, err)
	var parsedJSON bytes.Buffer

	err = json.Indent(&parsedJSON, []byte(expected), "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, string(parsedJSON.Bytes()), string(prettyPrintActualBytes.Bytes()))
}

var expectedDashboardsInJSONString = `
[
  {
    "resource": {
      "urn": "tableau::Samples/workbook/969428bf-0bde-4c73-8efe-04c37ec7c2ba",
      "name": "Regional",
      "service": "tableau"
    },
    "charts": [
      {
        "urn": "tableau::Samples/view/224d5a29-8a40-4b33-9ba5-9af0bab5b661",
        "source": "tableau",
        "dashboard_urn": "tableau::Samples/workbook/969428bf-0bde-4c73-8efe-04c37ec7c2ba",
        "lineage": {
          "upstreams": [
            {
              "urn": "hyper::0bfc6543-86ab-40b1-a8b7-d73a00de7d55",
              "name": "Stocks"
            },
            {
              "urn": "hyper::4a1302a6-d445-41bb-b5c9-a201b466846a",
              "name": "Economy"
            },
            {
              "urn": "hyper::82d82617-48fe-40f1-875a-d9aa4d1e1b86",
              "name": "Flight"
            },
            {
              "urn": "hyper::e5ded210-1699-42c8-82db-22f0bb6b16ec",
              "name": "Global Temperatures"
            },
            {
              "urn": "hyper::082fa857-f964-4c29-8672-6dcfc9596efa",
              "name": "Education"
            },
            {
              "urn": "hyper::57ec96ca-8e0b-4a53-a0d1-2714e13748bf",
              "name": "Obesity"
            }
          ]
        },
        "properties": {
          "attributes": {
            "content_url": "Regional/sheets/Obesity",
            "id": "224d5a29-8a40-4b33-9ba5-9af0bab5b661",
            "name": "Obesity",
            "tag": [],
            "view_url_name": "Obesity"
          }
        },
        "timestamps": {
          "create_time": {
            "seconds": 1635766351
          },
          "update_time": {
            "seconds": 1635766351
          }
        }
      },
      {
        "urn": "tableau::Samples/view/a139e8b6-03c8-4989-b37b-a7f7e6e82415",
        "source": "tableau",
        "dashboard_urn": "tableau::Samples/workbook/969428bf-0bde-4c73-8efe-04c37ec7c2ba",
        "lineage": {
          "upstreams": [
            {
              "urn": "hyper::0bfc6543-86ab-40b1-a8b7-d73a00de7d55",
              "name": "Stocks"
            },
            {
              "urn": "hyper::4a1302a6-d445-41bb-b5c9-a201b466846a",
              "name": "Economy"
            },
            {
              "urn": "hyper::82d82617-48fe-40f1-875a-d9aa4d1e1b86",
              "name": "Flight"
            },
            {
              "urn": "hyper::e5ded210-1699-42c8-82db-22f0bb6b16ec",
              "name": "Global Temperatures"
            },
            {
              "urn": "hyper::082fa857-f964-4c29-8672-6dcfc9596efa",
              "name": "Education"
            },
            {
              "urn": "hyper::57ec96ca-8e0b-4a53-a0d1-2714e13748bf",
              "name": "Obesity"
            }
          ]
        },
        "properties": {
          "attributes": {
            "content_url": "Regional/sheets/College",
            "id": "a139e8b6-03c8-4989-b37b-a7f7e6e82415",
            "name": "College",
            "tag": [],
            "view_url_name": "College"
          }
        },
        "timestamps": {
          "create_time": {
            "seconds": 1635766351
          },
          "update_time": {
            "seconds": 1635766351
          }
        }
      },
      {
        "urn": "tableau::Samples/view/bd362967-df94-45ec-b723-8c562abc8ccb",
        "source": "tableau",
        "dashboard_urn": "tableau::Samples/workbook/969428bf-0bde-4c73-8efe-04c37ec7c2ba",
        "lineage": {
          "upstreams": [
            {
              "urn": "hyper::0bfc6543-86ab-40b1-a8b7-d73a00de7d55",
              "name": "Stocks"
            },
            {
              "urn": "hyper::4a1302a6-d445-41bb-b5c9-a201b466846a",
              "name": "Economy"
            },
            {
              "urn": "hyper::82d82617-48fe-40f1-875a-d9aa4d1e1b86",
              "name": "Flight"
            },
            {
              "urn": "hyper::e5ded210-1699-42c8-82db-22f0bb6b16ec",
              "name": "Global Temperatures"
            },
            {
              "urn": "hyper::082fa857-f964-4c29-8672-6dcfc9596efa",
              "name": "Education"
            },
            {
              "urn": "hyper::57ec96ca-8e0b-4a53-a0d1-2714e13748bf",
              "name": "Obesity"
            }
          ]
        },
        "properties": {
          "attributes": {
            "content_url": "Regional/sheets/GlobalTemperatures",
            "id": "bd362967-df94-45ec-b723-8c562abc8ccb",
            "name": "Global Temperatures",
            "tag": [],
            "view_url_name": "GlobalTemperatures"
          }
        },
        "timestamps": {
          "create_time": {
            "seconds": 1635766351
          },
          "update_time": {
            "seconds": 1635766351
          }
        }
      },
      {
        "urn": "tableau::Samples/view/c7a57c38-6dd8-49f0-a905-b76f41b4b33e",
        "source": "tableau",
        "dashboard_urn": "tableau::Samples/workbook/969428bf-0bde-4c73-8efe-04c37ec7c2ba",
        "lineage": {
          "upstreams": [
            {
              "urn": "hyper::0bfc6543-86ab-40b1-a8b7-d73a00de7d55",
              "name": "Stocks"
            },
            {
              "urn": "hyper::4a1302a6-d445-41bb-b5c9-a201b466846a",
              "name": "Economy"
            },
            {
              "urn": "hyper::82d82617-48fe-40f1-875a-d9aa4d1e1b86",
              "name": "Flight"
            },
            {
              "urn": "hyper::e5ded210-1699-42c8-82db-22f0bb6b16ec",
              "name": "Global Temperatures"
            },
            {
              "urn": "hyper::082fa857-f964-4c29-8672-6dcfc9596efa",
              "name": "Education"
            },
            {
              "urn": "hyper::57ec96ca-8e0b-4a53-a0d1-2714e13748bf",
              "name": "Obesity"
            }
          ]
        },
        "properties": {
          "attributes": {
            "content_url": "Regional/sheets/FlightDelays",
            "id": "c7a57c38-6dd8-49f0-a905-b76f41b4b33e",
            "name": "Flight Delays",
            "tag": [],
            "view_url_name": "FlightDelays"
          }
        },
        "timestamps": {
          "create_time": {
            "seconds": 1635766351
          },
          "update_time": {
            "seconds": 1635766351
          }
        }
      },
      {
        "urn": "tableau::Samples/view/662b2993-69e8-4588-b57c-51c94ebdebd7",
        "source": "tableau",
        "dashboard_urn": "tableau::Samples/workbook/969428bf-0bde-4c73-8efe-04c37ec7c2ba",
        "lineage": {
          "upstreams": [
            {
              "urn": "hyper::0bfc6543-86ab-40b1-a8b7-d73a00de7d55",
              "name": "Stocks"
            },
            {
              "urn": "hyper::4a1302a6-d445-41bb-b5c9-a201b466846a",
              "name": "Economy"
            },
            {
              "urn": "hyper::82d82617-48fe-40f1-875a-d9aa4d1e1b86",
              "name": "Flight"
            },
            {
              "urn": "hyper::e5ded210-1699-42c8-82db-22f0bb6b16ec",
              "name": "Global Temperatures"
            },
            {
              "urn": "hyper::082fa857-f964-4c29-8672-6dcfc9596efa",
              "name": "Education"
            },
            {
              "urn": "hyper::57ec96ca-8e0b-4a53-a0d1-2714e13748bf",
              "name": "Obesity"
            }
          ]
        },
        "properties": {
          "attributes": {
            "content_url": "Regional/sheets/Economy",
            "id": "662b2993-69e8-4588-b57c-51c94ebdebd7",
            "name": "Economy",
            "tag": [],
            "view_url_name": "Economy"
          }
        },
        "timestamps": {
          "create_time": {
            "seconds": 1635766351
          },
          "update_time": {
            "seconds": 1635766351
          }
        }
      },
      {
        "urn": "tableau::Samples/view/b1b3da27-2f97-4add-87fd-c743bf497783",
        "source": "tableau",
        "dashboard_urn": "tableau::Samples/workbook/969428bf-0bde-4c73-8efe-04c37ec7c2ba",
        "lineage": {
          "upstreams": [
            {
              "urn": "hyper::0bfc6543-86ab-40b1-a8b7-d73a00de7d55",
              "name": "Stocks"
            },
            {
              "urn": "hyper::4a1302a6-d445-41bb-b5c9-a201b466846a",
              "name": "Economy"
            },
            {
              "urn": "hyper::82d82617-48fe-40f1-875a-d9aa4d1e1b86",
              "name": "Flight"
            },
            {
              "urn": "hyper::e5ded210-1699-42c8-82db-22f0bb6b16ec",
              "name": "Global Temperatures"
            },
            {
              "urn": "hyper::082fa857-f964-4c29-8672-6dcfc9596efa",
              "name": "Education"
            },
            {
              "urn": "hyper::57ec96ca-8e0b-4a53-a0d1-2714e13748bf",
              "name": "Obesity"
            }
          ]
        },
        "properties": {
          "attributes": {
            "content_url": "Regional/sheets/Stocks",
            "id": "b1b3da27-2f97-4add-87fd-c743bf497783",
            "name": "Stocks",
            "tag": [],
            "view_url_name": "Stocks"
          }
        },
        "timestamps": {
          "create_time": {
            "seconds": 1635766351
          },
          "update_time": {
            "seconds": 1635766351
          }
        }
      }
    ],
    "properties": {
      "attributes": {
        "content_url": "Regional",
        "id": "969428bf-0bde-4c73-8efe-04c37ec7c2ba",
        "location_id": "46b2914d-6d2c-4cc5-ba2a-643840c036c6",
        "location_name": "Samples",
        "location_type": "Project",
        "owner_id": "2aed1e7b-975b-430f-ab9d-a574bb88f245",
        "owner_name": "user@meteor",
        "project_id": "46b2914d-6d2c-4cc5-ba2a-643840c036c6",
        "project_name": "Samples",
        "tag": [],
        "webpage_url": "https://server.tableau.com/#/site/testdev550928/workbooks/88352"
      }
    },
    "timestamps": {
      "create_time": {
        "seconds": 1635766350
      },
      "update_time": {
        "seconds": 1635766351
      }
    }
  },
  {
    "resource": {
      "urn": "tableau::Samples/workbook/a22c1975-13a4-4b66-88e1-d68757e2dc97",
      "name": "Superstore",
      "service": "tableau"
    },
    "description": "A sample superstore",
    "charts": [
      {
        "urn": "tableau::Samples/view/43807016-73ff-4152-a822-3cd300d9618a",
        "source": "tableau",
        "dashboard_urn": "tableau::Samples/workbook/a22c1975-13a4-4b66-88e1-d68757e2dc97",
        "lineage": {
          "upstreams": [
            {
              "urn": "textscan::8ac2912e-066c-4ef6-9a40-ce5a1c2d159c",
              "name": "Sales Commission"
            },
            {
              "urn": "excel-direct::6f34315e-b37c-4cad-8650-c32562c92ed2",
              "name": "Sales Target (US)"
            },
            {
              "urn": "excel-direct::969e3732-a7f0-4563-8fec-ff1ad7c44470",
              "name": "Sample - Superstore"
            }
          ]
        },
        "properties": {
          "attributes": {
            "content_url": "Superstore/sheets/Overview",
            "id": "43807016-73ff-4152-a822-3cd300d9618a",
            "name": "Overview",
            "tag": [
              "forecast",
              "storage"
            ],
            "view_url_name": "Overview"
          }
        },
        "timestamps": {
          "create_time": {
            "seconds": 1635766355
          },
          "update_time": {
            "seconds": 1635766355
          }
        }
      },
      {
        "urn": "tableau::Samples/view/30df0a02-42a4-4899-b0ba-9b101b96b29b",
        "source": "tableau",
        "dashboard_urn": "tableau::Samples/workbook/a22c1975-13a4-4b66-88e1-d68757e2dc97",
        "lineage": {
          "upstreams": [
            {
              "urn": "textscan::8ac2912e-066c-4ef6-9a40-ce5a1c2d159c",
              "name": "Sales Commission"
            },
            {
              "urn": "excel-direct::6f34315e-b37c-4cad-8650-c32562c92ed2",
              "name": "Sales Target (US)"
            },
            {
              "urn": "excel-direct::969e3732-a7f0-4563-8fec-ff1ad7c44470",
              "name": "Sample - Superstore"
            }
          ]
        },
        "properties": {
          "attributes": {
            "content_url": "Superstore/sheets/Product",
            "id": "30df0a02-42a4-4899-b0ba-9b101b96b29b",
            "name": "Product",
            "tag": [
              "forecast",
              "storage"
            ],
            "view_url_name": "Product"
          }
        },
        "timestamps": {
          "create_time": {
            "seconds": 1635766355
          },
          "update_time": {
            "seconds": 1635766355
          }
        }
      },
      {
        "urn": "tableau::Samples/view/ae74c34a-8e87-4843-b595-b3e75e7dd8db",
        "source": "tableau",
        "dashboard_urn": "tableau::Samples/workbook/a22c1975-13a4-4b66-88e1-d68757e2dc97",
        "lineage": {
          "upstreams": [
            {
              "urn": "textscan::8ac2912e-066c-4ef6-9a40-ce5a1c2d159c",
              "name": "Sales Commission"
            },
            {
              "urn": "excel-direct::6f34315e-b37c-4cad-8650-c32562c92ed2",
              "name": "Sales Target (US)"
            },
            {
              "urn": "excel-direct::969e3732-a7f0-4563-8fec-ff1ad7c44470",
              "name": "Sample - Superstore"
            }
          ]
        },
        "properties": {
          "attributes": {
            "content_url": "Superstore/sheets/Customers",
            "id": "ae74c34a-8e87-4843-b595-b3e75e7dd8db",
            "name": "Customers",
            "tag": [
              "forecast",
              "storage"
            ],
            "view_url_name": "Customers"
          }
        },
        "timestamps": {
          "create_time": {
            "seconds": 1635766355
          },
          "update_time": {
            "seconds": 1635766355
          }
        }
      },
      {
        "urn": "tableau::Samples/view/ca492b1a-adc6-4d49-a2c5-f0028d22b5f8",
        "source": "tableau",
        "dashboard_urn": "tableau::Samples/workbook/a22c1975-13a4-4b66-88e1-d68757e2dc97",
        "lineage": {
          "upstreams": [
            {
              "urn": "textscan::8ac2912e-066c-4ef6-9a40-ce5a1c2d159c",
              "name": "Sales Commission"
            },
            {
              "urn": "excel-direct::6f34315e-b37c-4cad-8650-c32562c92ed2",
              "name": "Sales Target (US)"
            },
            {
              "urn": "excel-direct::969e3732-a7f0-4563-8fec-ff1ad7c44470",
              "name": "Sample - Superstore"
            }
          ]
        },
        "properties": {
          "attributes": {
            "content_url": "Superstore/sheets/Shipping",
            "id": "ca492b1a-adc6-4d49-a2c5-f0028d22b5f8",
            "name": "Shipping",
            "tag": [
              "forecast",
              "storage"
            ],
            "view_url_name": "Shipping"
          }
        },
        "timestamps": {
          "create_time": {
            "seconds": 1635766355
          },
          "update_time": {
            "seconds": 1635766355
          }
        }
      },
      {
        "urn": "tableau::Samples/view/b3e230a3-8f06-4ef5-9bca-e533c8d7aac1",
        "source": "tableau",
        "dashboard_urn": "tableau::Samples/workbook/a22c1975-13a4-4b66-88e1-d68757e2dc97",
        "lineage": {
          "upstreams": [
            {
              "urn": "textscan::8ac2912e-066c-4ef6-9a40-ce5a1c2d159c",
              "name": "Sales Commission"
            },
            {
              "urn": "excel-direct::6f34315e-b37c-4cad-8650-c32562c92ed2",
              "name": "Sales Target (US)"
            },
            {
              "urn": "excel-direct::969e3732-a7f0-4563-8fec-ff1ad7c44470",
              "name": "Sample - Superstore"
            }
          ]
        },
        "properties": {
          "attributes": {
            "content_url": "Superstore/sheets/Performance",
            "id": "b3e230a3-8f06-4ef5-9bca-e533c8d7aac1",
            "name": "Performance",
            "tag": [
              "forecast",
              "storage"
            ],
            "view_url_name": "Performance"
          }
        },
        "timestamps": {
          "create_time": {
            "seconds": 1635766355
          },
          "update_time": {
            "seconds": 1635766355
          }
        }
      },
      {
        "urn": "tableau::Samples/view/4a2aac89-b7fe-4412-86f2-e4312495862d",
        "source": "tableau",
        "dashboard_urn": "tableau::Samples/workbook/a22c1975-13a4-4b66-88e1-d68757e2dc97",
        "lineage": {
          "upstreams": [
            {
              "urn": "textscan::8ac2912e-066c-4ef6-9a40-ce5a1c2d159c",
              "name": "Sales Commission"
            },
            {
              "urn": "excel-direct::6f34315e-b37c-4cad-8650-c32562c92ed2",
              "name": "Sales Target (US)"
            },
            {
              "urn": "excel-direct::969e3732-a7f0-4563-8fec-ff1ad7c44470",
              "name": "Sample - Superstore"
            }
          ]
        },
        "properties": {
          "attributes": {
            "content_url": "Superstore/sheets/CommissionModel",
            "id": "4a2aac89-b7fe-4412-86f2-e4312495862d",
            "name": "Commission Model",
            "tag": [
              "forecast",
              "storage"
            ],
            "view_url_name": "CommissionModel"
          }
        },
        "timestamps": {
          "create_time": {
            "seconds": 1635766355
          },
          "update_time": {
            "seconds": 1635766355
          }
        }
      },
      {
        "urn": "tableau::Samples/view/752b51c0-b404-4e68-abad-204f0e10cefe",
        "source": "tableau",
        "dashboard_urn": "tableau::Samples/workbook/a22c1975-13a4-4b66-88e1-d68757e2dc97",
        "lineage": {
          "upstreams": [
            {
              "urn": "textscan::8ac2912e-066c-4ef6-9a40-ce5a1c2d159c",
              "name": "Sales Commission"
            },
            {
              "urn": "excel-direct::6f34315e-b37c-4cad-8650-c32562c92ed2",
              "name": "Sales Target (US)"
            },
            {
              "urn": "excel-direct::969e3732-a7f0-4563-8fec-ff1ad7c44470",
              "name": "Sample - Superstore"
            }
          ]
        },
        "properties": {
          "attributes": {
            "content_url": "Superstore/sheets/OrderDetails",
            "id": "752b51c0-b404-4e68-abad-204f0e10cefe",
            "name": "Order Details",
            "tag": [
              "forecast",
              "storage"
            ],
            "view_url_name": "OrderDetails"
          }
        },
        "timestamps": {
          "create_time": {
            "seconds": 1635766355
          },
          "update_time": {
            "seconds": 1635766355
          }
        }
      },
      {
        "urn": "tableau::Samples/view/6a9f4625-565e-47b9-b7dd-159a27f8b46f",
        "source": "tableau",
        "dashboard_urn": "tableau::Samples/workbook/a22c1975-13a4-4b66-88e1-d68757e2dc97",
        "lineage": {
          "upstreams": [
            {
              "urn": "textscan::8ac2912e-066c-4ef6-9a40-ce5a1c2d159c",
              "name": "Sales Commission"
            },
            {
              "urn": "excel-direct::6f34315e-b37c-4cad-8650-c32562c92ed2",
              "name": "Sales Target (US)"
            },
            {
              "urn": "excel-direct::969e3732-a7f0-4563-8fec-ff1ad7c44470",
              "name": "Sample - Superstore"
            }
          ]
        },
        "properties": {
          "attributes": {
            "content_url": "Superstore/sheets/Forecast",
            "id": "6a9f4625-565e-47b9-b7dd-159a27f8b46f",
            "name": "Forecast",
            "tag": [
              "forecast",
              "storage"
            ],
            "view_url_name": "Forecast"
          }
        },
        "timestamps": {
          "create_time": {
            "seconds": 1635766355
          },
          "update_time": {
            "seconds": 1635766355
          }
        }
      },
      {
        "urn": "tableau::Samples/view/998155ce-e41d-4262-9738-b056ea4b202d",
        "source": "tableau",
        "dashboard_urn": "tableau::Samples/workbook/a22c1975-13a4-4b66-88e1-d68757e2dc97",
        "lineage": {
          "upstreams": [
            {
              "urn": "textscan::8ac2912e-066c-4ef6-9a40-ce5a1c2d159c",
              "name": "Sales Commission"
            },
            {
              "urn": "excel-direct::6f34315e-b37c-4cad-8650-c32562c92ed2",
              "name": "Sales Target (US)"
            },
            {
              "urn": "excel-direct::969e3732-a7f0-4563-8fec-ff1ad7c44470",
              "name": "Sample - Superstore"
            }
          ]
        },
        "properties": {
          "attributes": {
            "content_url": "Superstore/sheets/WhatIfForecast",
            "id": "998155ce-e41d-4262-9738-b056ea4b202d",
            "name": "What If Forecast",
            "tag": [
              "forecast",
              "storage"
            ],
            "view_url_name": "WhatIfForecast"
          }
        },
        "timestamps": {
          "create_time": {
            "seconds": 1635766355
          },
          "update_time": {
            "seconds": 1635766355
          }
        }
      }
    ],
    "properties": {
      "attributes": {
        "content_url": "Superstore",
        "id": "a22c1975-13a4-4b66-88e1-d68757e2dc97",
        "location_id": "46b2914d-6d2c-4cc5-ba2a-643840c036c6",
        "location_name": "Samples",
        "location_type": "Project",
        "owner_id": "2aed1e7b-975b-430f-ab9d-a574bb88f245",
        "owner_name": "user@meteor",
        "project_id": "46b2914d-6d2c-4cc5-ba2a-643840c036c6",
        "project_name": "Samples",
        "tag": [
          "forecast",
          "storage"
        ],
        "webpage_url": "https://server.tableau.com/#/site/testdev550928/workbooks/88353"
      }
    },
    "timestamps": {
      "create_time": {
        "seconds": 1635766355
      },
      "update_time": {
        "seconds": 1635833238
      }
    }
  },
  {
    "resource": {
      "urn": "tableau::test-meteor/workbook/37040233-8f4c-4b17-aecf-5c4be36dbf27",
      "name": "InMail Engagement",
      "service": "tableau"
    },
    "charts": [
      {
        "urn": "tableau::test-meteor/view/aea8019f-b29a-4e0f-b5d8-a96fc52188c1",
        "source": "tableau",
        "dashboard_urn": "tableau::test-meteor/workbook/37040233-8f4c-4b17-aecf-5c4be36dbf27",
        "lineage": {
          "upstreams": [
            {
              "urn": "webdata-direct:linkedin-snap::1bf138c3-1e6d-4afa-82c0-5ccd387be2bd",
              "name": "Sales Navigator Connection"
            }
          ]
        },
        "properties": {
          "attributes": {
            "content_url": "InMailEngagement/sheets/InMailEngagement",
            "id": "aea8019f-b29a-4e0f-b5d8-a96fc52188c1",
            "name": "InMail Engagement",
            "tag": [],
            "view_url_name": "InMailEngagement"
          }
        },
        "timestamps": {
          "create_time": {
            "seconds": 1635854252
          },
          "update_time": {
            "seconds": 1635854252
          }
        }
      }
    ],
    "properties": {
      "attributes": {
        "content_url": "InMailEngagement",
        "id": "37040233-8f4c-4b17-aecf-5c4be36dbf27",
        "location_id": "c95e1f64-1c3c-4ddf-8999-047a54e2af44",
        "location_name": "test-meteor",
        "location_type": "Project",
        "owner_id": "2aed1e7b-975b-430f-ab9d-a574bb88f245",
        "owner_name": "user@meteor",
        "project_id": "c95e1f64-1c3c-4ddf-8999-047a54e2af44",
        "project_name": "test-meteor",
        "tag": [],
        "webpage_url": "https://server.tableau.com/#/site/testdev550928/workbooks/88403"
      }
    },
    "timestamps": {
      "create_time": {
        "seconds": 1635854252
      },
      "update_time": {
        "seconds": 1635854253
      }
    }
  }
]`
