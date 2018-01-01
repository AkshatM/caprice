package caprice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// The actual URL endpoint to hit
const endpoint string = "https://api.random.org/json-rpc/1/invoke"

// Caprice's core object. Responsible for safekeeping the API key,
// as well as managing advisory delays in concurrent implementations.
type trueRNG struct {
	apiKey string
}

// A helper function that will return a new trueRNG object
func TrueRNG(apiKey string) trueRNG {
	return trueRNG{apiKey: apiKey}
}

// The outer JSON wrapper we send in our request body. It contains
// `params`, which is a JSON object containing all the method parameters.
// Typically, a struct implementing RequestShell will be populated for you.
type RequestShell struct {
	Version string      `json:"jsonrpc"`
	Params  interface{} `json:"params"`
	Method  string      `json:"method"`
	Id      int         `json:"id"`
}

// The outer JSON wrapper we receive in our request body. It contains
// `result`, which is a JSON object containing an object containing
// all the actual data we care for.
// `error` is returned from the API in lieu of `result` in the event of an error.
// 'id' and versions are constants supplied and returned by the API.
type ResponseShell struct {
	Version string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   Error           `json:"error"`
	Id      int             `json:"id"`
}

// A nested inner JSON wrapper around Random within ResponseShell.Result. It includes some
// basic diagnostic information as well. One of four types implementing Response
// interface
type Result struct {
	Random        Random `json:"random"`
	BitsUsed      int    `json:"bitsUsed"`
	BitsLeft      int    `json:"bitsLeft"`
	RequestsLeft  int    `json:"requestsLeft"`
	AdvisoryDelay int    `json:"advisoryDelay"`
}

// A nested inner JSON object within ResponseShell.Result. It includes only
// basic diagnostic information. One of four types implementing Response
// interface; returned directly as part of the GetUsage method.
type Status struct {
	Status        string `json:"status"`
	CreationTime  string `json:"creationTime"`
	BitsLeft      int    `json:"bitsLeft"`
	RequestsLeft  int    `json:"requestsLeft"`
	TotalBits     int    `json:"totalBits"`
	TotalRequests int    `json:"totalRequests"`
}

// A nested inner JSON object within ResponseShell.Result. It includes only
// basic diagnostic information. One of four types implementing Response
// interface; returned directly as part of the GetUsage method.
type VerifiedSignature struct {
	Authenticity bool `json:"authenticiity"`
}

// A nested inner JSON wrapper around Random within ResponseShell.Result. It includes some
// basic diagnostic information as well as a signature to verify the source of the data.
// One of four types implementing Response interface. Raw is an alias for Random here.
type SignedResult struct {
	Raw           json.RawMessage `json:"random"`
	Signature     string          `json:"signature"`
	BitsUsed      int             `json:"bitsUsed"`
	BitsLeft      int             `json:"bitsLeft"`
	RequestsLeft  int             `json:"requestsLeft"`
	AdvisoryDelay int             `json:"advisoryDelay"`
}

// A deeply nested inner JSON object contained inside ResponseShell.Random, which includes
// everything we asked for. For basic methods, `Data` and `CompletionTime` are returned.
// For signed methods, the values `HashedApiKey` and `SerialNumber` are also returned.
type Random struct {
	Data           []interface{} `json:"data"`
	CompletionTime string        `json:"completionTime"`
	SerialNumber   int           `json:"serialNumber"`
	HashedApiKey   string        `json:"hashedApiKey"`
}

type SignedIntegerData struct {
	Raw          json.RawMessage
	HashedApiKey string
	SerialNumber int
	Data         []int
	Signature    string
}

type SignedFloatData struct {
	Raw          json.RawMessage
	HashedApiKey string
	SerialNumber int
	Data         []float64
	Signature    string
}

type SignedStringData struct {
	Raw          json.RawMessage
	HashedApiKey string
	Data         []string
	SerialNumber int
	Signature    string
}

// An interface that all Status, Result and SignedResult implement so that they can be
// returned as valid outputs of the Request and SignedRequest functions. The expectation
// is that any struct that satisfies this interface has the actual data we would care about.
type Response interface {
	Content() interface{}
}

func (s Status) Content() interface{} {
	return s
}

func (r Result) Content() interface{} {
	return r.Random.Data
}

func (sr SignedResult) Content() interface{} {
	return sr
}

func (vs VerifiedSignature) Content() interface{} {
	return vs
}

// Core error object. Includes error message, the status code returned by the API,
// and optional data returned by the upstream API.
type Error struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    []interface{} `json:"data,omitempty"`
}

type IntegersReq struct {
	ApiKey      string `json:"apiKey"`
	N           int    `json:"n"`
	Min         int    `json:"min"`
	Max         int    `json:"max"`
	Replacement bool   `json:"replacement,omitempty"`
}

type DecimalFractionsReq struct {
	ApiKey        string `json:"apiKey"`
	N             int    `json:"n"`
	DecimalPlaces int    `json:"decimalPlaces"`
	Replacement   bool   `json:"replacement,omitempty"`
}

type GaussiansReq struct {
	ApiKey            string  `json:"apiKey"`
	N                 int     `json:"n"`
	Mean              float64 `json:"mean"`
	StandardDeviation float64 `json:"standardDeviation"`
	SignificantDigits int     `json:"significantDigits"`
}

type StringsReq struct {
	ApiKey      string `json:"apiKey"`
	N           int    `json:"n"`
	Length      int    `json:"length"`
	Characters  string `json:"characters"`
	Replacement bool   `json:"replacement,omitempty"`
}

type UUIDsReq struct {
	ApiKey string `json:"apiKey"`
	N      int    `json:"n"`
}

type BlobsReq struct {
	ApiKey string `json:"apiKey"`
	N      int    `json:"n"`
	Size   int    `json:"size"`
	Format string `json:"format"`
}

type StatusReq struct {
	ApiKey string `json:"apiKey"`
}

type VerifySignatureReq struct {
	Raw       map[string]interface{} `json:"random"`
	Signature string                 `json:"signature"`
}

func (e Error) Error() string {
	return fmt.Sprintf("Code: %d, Error: %s", e.Code, e.Message)
}

func clientError(message string) (ResponseShell, Error) {
	return ResponseShell{}, Error{
		Code:    409,
		Message: message,
	}
}

// A helper function that makes HTTP calls to RANDOM.org with our request parameters and method name.
func _request(method string, params interface{}) (ResponseShell, Error) {

	// create the JSON body for our request - ID is set to any number, doesn't matter which as API doesn't support batch notifs.
	body, err := json.Marshal(RequestShell{Version: "2.0", Params: params, Method: method, Id: 1})

	fmt.Println(string(body))
	if err != nil {
		return clientError(err.Error())
	}

	// fire off the POST request
	resp, err := http.Post(endpoint, "application/json-rpc", bytes.NewReader(body))
	if err != nil {
		return clientError(err.Error())
	}
	defer resp.Body.Close()

	// convert resp.Body into a buffer we can unmarshall from
	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return clientError(err.Error())
	}

	// handle non-successful behaviour
	if resp.StatusCode != 200 {
		errorMessage := Error{}
		json.Unmarshal(text, &errorMessage)
		return ResponseShell{}, errorMessage
	}

	// unmarshall response data
	response := ResponseShell{}
	json.Unmarshal(text, &response)

	if response.Error.Message != "" {
		return ResponseShell{}, response.Error
	}

	return response, Error{}
}

func Request(method string, params interface{}) (Response, Error) {

	response, err := _request(method, params)
	if err.Message != "" {
		return nil, err
	}

	if method == "getUsage" {
		status := Status{}
		json.Unmarshal(response.Result, &status)
		return status, Error{}
	}

	result := Result{}
	json.Unmarshal(response.Result, &result)
	return result, Error{}
}

func SignedRequest(method string, params interface{}) (Response, Error) {

	response, err := _request(method, params)
	if err.Message != "" {
		return nil, err
	}

	if method == "verifySignature" {
		verifiedSignature := VerifiedSignature{}
		json.Unmarshal(response.Result, &verifiedSignature)
		return verifiedSignature, Error{}
	}

	result := SignedResult{}
	json.Unmarshal(response.Result, &result)
	return result, Error{}
}
