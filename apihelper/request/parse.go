package request

import (
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
)

const queryTagName = "query"

//ParseJSON parses request body to given destination
//dest field should be a pointer.
func ParseJSON(req *http.Request, dest interface{}) error {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(dest)
}

//ParseQuery parses queries of url into the given destination
//dest field should be a pointer.
func ParseQuery(req *http.Request, dest interface{}) error {
	queryParser(req.URL, dest)
	return nil
}

//queryParser parses URL's query parameters according to the `query` tag
func queryParser(url *url.URL, dest interface{}) {
	destVal := reflect.ValueOf(dest).Elem()

	for i := 0; i < destVal.NumField(); i++ {
		tagVal := destVal.Type().Field(i).Tag.Get(queryTagName)

		if tagVal == "-" {
			continue
		}

		urlVal := url.Query().Get(tagVal)

		switch destVal.Type().Field(i).Type.Kind() {
		case reflect.String: // string is enough for this case
			destVal.Field(i).SetString(urlVal)
		}
	}

}
