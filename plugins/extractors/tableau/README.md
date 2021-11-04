# tableau

## Usage
```yaml
source:
  type: tableau
  config:
    host: http://localhost:3000
    version: 3.13
    sitename: testdev550928
    username: meteor_user
    password: xxxxxxxxxx
```


## Inputs

| Key | Value | Example | Description |    |
| :-- | :---- | :------ | :---------- | :- |
| `host` | `string` | `https://localhost:3000`         | The host at which tableau is running | *required* |
| `version` | `string` | `3.13`     | The version of [Tableau REST API](https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_concepts_versions.htm) | *required* |
| `username` | `string` | `meteor_user` | Username/email to access the tableau | *required* |
| `password` | `string` | `xxxxxxxxxx` | Password for the tableau | *required* |
| `sitename` | `string` | `testdev550928` | The name of your tableau site | *required* |

## Outputs

| Field | Sample Value |
| :---- | :---- |
| `resource.urn` | `tableau::{project_name}/workbook/{workbook_id}` |
| `resource.name` | `workbook_name` |
| `resource.service` | `tableau` |
| `description` | `workbook description` |
| `charts` | [][Chart](#chart) |

### Chart

| Field | Sample Value |
| :---- | :---- |
| `urn` | `tableau::{project_name}/view/{view_id}`             |
| `source` | `tableau` |
| `dashboard_urn` | `tableau::{project_name}/workbook/{workbook_id}` |
| `dashboard_source` | `tableau` |

## Contributing

Refer to the [contribution guidelines](../../../docs/contribute/guide.md#adding-a-new-extractor) for information on contributing to this module.
