package firego

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// TransactionFn is used to run a transaction on a Firebase reference.
// See Firebase.Transaction for more information.
type TransactionFn func(currentSnapshot interface{}) (result interface{}, err error)

func getTransactionParams(headers http.Header, body []byte) (etag string, snapshot interface{}, err error) {
	etag = headers.Get("ETag")
	if len(etag) == 0 {
		return etag, snapshot, errors.New("no etag returned by Firebase")
	}

	if err := json.Unmarshal(body, &snapshot); err != nil {
		return etag, snapshot, fmt.Errorf("failed to unmarshal Firebase response. %s", err)
	}

	return etag, snapshot, nil
}

// Transaction runs a transaction on the data at this location. The TransactionFn parameter
// will be called, possibly multiple times, with the current data at this location.
// It is responsible for inspecting that data and specifying either the desired new data
// at the location or that the transaction should be aborted.
//
// Since the provided function may be called repeatedly for the same transaction, be extremely careful of
// any side effects that may be triggered by this method.
//
// Best practices for this method are to rely only on the data that is passed in.
func (fb *Firebase) Transaction(fn TransactionFn) error {
	// fetch etag and current value
	headers, body, err := fb.doRequest("GET", nil, withHeader("X-Firebase-ETag", "true"))
	if err != nil {
		return err
	}

	etag, snapshot, err := getTransactionParams(headers, body)
	if err != nil {
		return err
	}

	// set the error value to something non-nil so that
	// we step into the loop
	tErr := errors.New("")
	for i := 0; i < 25 && tErr != nil; i++ {
		// run transaction
		result, err := fn(snapshot)
		if err != nil {
			return nil
		}

		newBody, err := json.Marshal(result)
		if err != nil {
			return fmt.Errorf("failed to marshal transaction result. %s", err)
		}

		// attempt to update it
		headers, body, tErr = fb.doRequest("PUT", newBody, withHeader("if-match", etag))
		if tErr == nil {
			// we're good, break the loop
			break
		}

		// we failed to update, so grab the new snapshot/etag
		e, s, tErr := getTransactionParams(headers, body)
		if tErr != nil {
			return tErr
		}
		etag, snapshot = e, s
	}

	if tErr != nil {
		return fmt.Errorf("failed to run transaction. %s", tErr)
	}
	return nil
}
