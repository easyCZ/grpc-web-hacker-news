package firego

import (
	"fmt"
	"strconv"
	"strings"
)

// StartAt creates a new Firebase reference with the
// requested StartAt configuration. The value that is passed in
// is automatically escaped if it is a string value.
// Numeric strings are automatically converted to numbers.
//
//    StartAt(7)        // -> startAt=7
//    StartAt("7")      // -> startAt=7
//    StartAt("foo")    // -> startAt="foo"
//    StartAt(`"foo"`)  // -> startAt="foo"
//
// Reference https://firebase.google.com/docs/database/rest/retrieve-data#section-rest-filtering
func (fb *Firebase) StartAt(value string) *Firebase {
	c := fb.copy()
	if value != "" {
		c.params.Set(startAtParam, escapeString(value))
	} else {
		c.params.Del(startAtParam)
	}
	return c
}

// StartAtValue creates a new Firebase reference with the
// requested StartAt configuration. The value that is passed in
// is automatically escaped if it is a string value.
// Numeric strings are preserved as strings.
//
//    StartAtValue(7)        // -> startAt=7
//    StartAtValue("7")      // -> startAt="7"
//    StartAtValue("foo")    // -> startAt="foo"
//    StartAtValue(`"foo"`)  // -> startAt="foo"
//
// Reference https://firebase.google.com/docs/database/rest/retrieve-data#section-rest-filtering
func (fb *Firebase) StartAtValue(value interface{}) *Firebase {
	c := fb.copy()
	if value != "" {
		c.params.Set(startAtParam, escapeParameter(value))
	} else {
		c.params.Del(startAtParam)
	}
	return c
}

// EndAt creates a new Firebase reference with the
// requested EndAt configuration. The value that is passed in
// is automatically escaped if it is a string value.
// Numeric strings are automatically converted to numbers.
//
//    EndAt(7)        // -> endAt=7
//    EndAt("7")      // -> endAt=7
//    EndAt("foo")    // -> endAt="foo"
//    EndAt(`"foo"`)  // -> endAt="foo"
//
// Reference https://firebase.google.com/docs/database/rest/retrieve-data#section-rest-filtering
func (fb *Firebase) EndAt(value string) *Firebase {
	c := fb.copy()
	if value != "" {
		c.params.Set(endAtParam, escapeString(value))
	} else {
		c.params.Del(endAtParam)
	}
	return c
}

// EndAtValue creates a new Firebase reference with the
// requested EndAt configuration. The value that is passed in
// is automatically escaped if it is a string value.
// Numeric strings are preserved as strings.
//
//    EndAtValue(7)        // -> endAt=7
//    EndAtValue("7")      // -> endAt="7"
//    EndAtValue("foo")    // -> endAt="foo"
//    EndAtValue(`"foo"`)  // -> endAt="foo"
//
// Reference https://firebase.google.com/docs/database/rest/retrieve-data#section-rest-filtering
func (fb *Firebase) EndAtValue(value interface{}) *Firebase {
	c := fb.copy()
	if value != "" {
		c.params.Set(endAtParam, escapeParameter(value))
	} else {
		c.params.Del(endAtParam)
	}
	return c
}

// OrderBy creates a new Firebase reference with the
// requested OrderBy configuration. The value that is passed in
// is automatically escaped if it is a string value.
//
//    OrderBy("foo")   // -> orderBy="foo"
//    OrderBy(`"foo"`) // -> orderBy="foo"
//    OrderBy("$key")  // -> orderBy="$key"
//
// Reference https://firebase.google.com/docs/database/rest/retrieve-data#orderby
func (fb *Firebase) OrderBy(value string) *Firebase {
	c := fb.copy()
	if value != "" {
		c.params.Set(orderByParam, escapeString(value))
	} else {
		c.params.Del(orderByParam)
	}
	return c
}

// EqualTo sends the query string equalTo so that one can find nodes with
// exactly matching values. The value that is passed in is automatically escaped
// if it is a string value.
// Numeric strings are automatically converted to numbers.
//
//    EqualTo(7)        // -> equalTo=7
//    EqualTo("7")      // -> equalTo=7
//    EqualTo("foo")    // -> equalTo="foo"
//    EqualTo(`"foo"`)  // -> equalTo="foo"
//
// Reference https://firebase.google.com/docs/database/rest/retrieve-data#section-rest-filtering
func (fb *Firebase) EqualTo(value string) *Firebase {
	c := fb.copy()
	if value != "" {
		c.params.Set(equalToParam, escapeString(value))
	} else {
		c.params.Del(equalToParam)
	}
	return c
}

// EqualToValue sends the query string equalTo so that one can find nodes with
// exactly matching values. The value that is passed in is automatically escaped
// if it is a string value.
// Numeric strings are preserved as strings.
//
//    EqualToValue(7)        // -> equalTo=7
//    EqualToValue("7")      // -> equalTo="7"
//    EqualToValue("foo")    // -> equalTo="foo"
//    EqualToValue(`"foo"`)  // -> equalTo="foo"
//
// Reference https://firebase.google.com/docs/database/rest/retrieve-data#section-rest-filtering
func (fb *Firebase) EqualToValue(value interface{}) *Firebase {
	c := fb.copy()
	if value != "" {
		c.params.Set(equalToParam, escapeParameter(value))
	} else {
		c.params.Del(equalToParam)
	}
	return c
}

func escapeString(s string) string {
	_, errNotInt := strconv.ParseInt(s, 10, 64)
	_, errNotBool := strconv.ParseBool(s)
	if errNotInt == nil || errNotBool == nil {
		// we shouldn't escape bools or ints
		return s
	}
	return fmt.Sprintf(`%q`, strings.Trim(s, `"`))
}

func escapeParameter(s interface{}) string {
	switch s.(type) {
	case string:
		return fmt.Sprintf(`%q`, strings.Trim(s.(string), `"`))
	default:
		return fmt.Sprintf(`%v`, s)
	}
}

// LimitToFirst creates a new Firebase reference with the
// requested limitToFirst configuration.
//
// Reference https://firebase.google.com/docs/database/rest/retrieve-data#limit-queries
func (fb *Firebase) LimitToFirst(value int64) *Firebase {
	c := fb.copy()
	if value > 0 {
		c.params.Set(limitToFirstParam, strconv.FormatInt(value, 10))
	} else {
		c.params.Del(limitToFirstParam)
	}
	return c
}

// LimitToLast creates a new Firebase reference with the
// requested limitToLast configuration.
//
// Reference https://firebase.google.com/docs/database/rest/retrieve-data#limit-queries
func (fb *Firebase) LimitToLast(value int64) *Firebase {
	c := fb.copy()
	if value > 0 {
		c.params.Set(limitToLastParam, strconv.FormatInt(value, 10))
	} else {
		c.params.Del(limitToLastParam)
	}
	return c
}

// Shallow limits the depth of the data returned when calling Value.
// If the data at the location is a JSON primitive (string, number or boolean),
// its value will be returned. If the data is a JSON object, the values
// for each key will be truncated to true.
//
// Reference https://firebase.google.com/docs/database/rest/retrieve-data#shallow
func (fb *Firebase) Shallow(v bool) {
	if v {
		fb.params.Set(shallowParam, "true")
	} else {
		fb.params.Del(shallowParam)
	}
}

// IncludePriority determines whether or not to ask Firebase
// for the values priority. By default, the priority is not returned.
//
// Reference https://www.firebase.com/docs/rest/api/#section-param-format
func (fb *Firebase) IncludePriority(v bool) {
	if v {
		fb.params.Set(formatParam, formatVal)
	} else {
		fb.params.Del(formatParam)
	}
}
