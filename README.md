# Mornin's plus controller service

## Usage:
### Metrics
**URL:** `/metrics`

**Method:** `GET`
#### Success Response
**Code:** `200 OK`

**Content examples:**

```
# HELP chicken_battery Robot battery level
# TYPE chicken_battery gauge
chicken_battery{} 44
```

### Action
**URL:** `/act`

**Method:** `POST`
#### Fields
| Name      | Type    | Description           |
| --        | --      | --                    |
| token     | String  | Token of bot          |
| action    | String  | Open \| Close \| Stop |
| highSpeed | Boolean | Optional              |

#### Success Response
**Code:** `200 OK`

### Fetch token
**URL:** `/token`

**Method:** `GET`
#### Query
| Name      | Type    | Description           |
| --        | --      | --                    |
| data      | String  | Data from qrcode      |

#### Success Response
**Code:** `200 OK`

**Content examples:**
```
{
  "token": "hogehoge"
}
```
