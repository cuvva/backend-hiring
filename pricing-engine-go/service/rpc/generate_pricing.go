package rpc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"pricingengine"
	"pricingengine/service/app"
)

type RPC struct {
	App *app.App
}

// GeneratePricing conforms to http.HandlerFunc and handles request logic
// for the application method `GeneratePricing`.
// NOTE: As you can see, this does a lot of stuff which should be part of a
// standard request framework and is missing elemtents which you would see
// in a production system (such as authentication, validation, errors-by-design,
// etc.) Please feel free to implement certain features if you have time but do
// not over-engineer this part; we're looking for a single functional endpoint,
// not a framework!
func (rpc *RPC) GeneratePricing(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response(w, err)
		return
	}

	r.Body.Close()
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	var input *pricingengine.GeneratePricingRequest
	err = json.Unmarshal(body, &input)
	if err != nil {
		response(w, err)
		return
	}

	res, err := rpc.App.GeneratePricing(r.Context(), input)
	if err != nil {
		response(w, err)
		return
	}

	response(w, res)
}

// response writes a successful response as JSON to the client
func response(w http.ResponseWriter, res interface{}) {
	if err, ok := res.(error); ok {
		errorResponse(w, err)
		return
	}

	resBody, err := json.Marshal(res)
	if err != nil {
		errorResponse(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}

// errorResponse writes out an error to the client as plaintext
func errorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
