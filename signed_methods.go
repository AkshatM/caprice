package caprice

import "encoding/json"

// Generate `n` random integers between `min` and `max`.
// If `replacement` is true, pick random numbers with replacement. Default is false.
// We do not support base selection, since it is easy to format into the base of your choice from base 10
func (rng trueRNG) GenerateSignedIntegers(n, min, max int, replacement bool) (SignedIntegerData, Error) {

	body := IntegersReq{ApiKey: rng.apiKey, N: n, Min: min, Max: max, Replacement: replacement}
	result, err := SignedRequest("generateSignedIntegers", body)

	if err.Message != "" {
		return SignedIntegerData{}, err
	}

	signedResult, _ := result.Content().(SignedResult)
	randomData := Random{}
	json.Unmarshal(signedResult.Raw, &randomData)

	data := randomData.Data
	intArray := make([]int, len(data))
	for i, num := range data {
		intArray[i] = int(num.(float64))
	}

	return SignedIntegerData{
		Data:         intArray,
		Raw:          signedResult.Raw,
		HashedApiKey: randomData.HashedApiKey,
		SerialNumber: randomData.SerialNumber,
		Signature:    signedResult.Signature,
	}, Error{}
}

// Generate `n` random decimal fractions with precision upto `decimalPlaces`.
// If `replacement` is true, pick random numbers with replacement. Default is false.
func (rng trueRNG) GenerateSignedDecimalFractions(n, decimalPlaces int, replacement bool) (SignedFloatData, Error) {

	body := DecimalFractionsReq{ApiKey: rng.apiKey, N: n, DecimalPlaces: decimalPlaces, Replacement: replacement}
	result, err := SignedRequest("generateSignedDecimalFractions", body)

	if err.Message != "" {
		return SignedFloatData{}, err
	}

	signedResult, _ := result.Content().(SignedResult)

	randomData := Random{}
	json.Unmarshal(signedResult.Raw, &randomData)

	data := randomData.Data
	floatArray := make([]float64, len(data))
	for i, num := range data {
		floatArray[i] = float64(num.(float64))
	}

	return SignedFloatData{
		Data:         floatArray,
		Raw:          signedResult.Raw,
		HashedApiKey: randomData.HashedApiKey,
		SerialNumber: randomData.SerialNumber,
		Signature:    signedResult.Signature,
	}, Error{}
}

// Generate `n` Gaussians from a distribution with mean `mean` and stdev `standardDeviation`, returned with
// at most `significantDigits` sig. digits.
// If `replacement` is true, pick random numbers with replacement. Default is false.
func (rng trueRNG) GenerateSignedGaussians(n int, mean, standardDeviation float64, significantDigits int) (SignedFloatData, Error) {

	body := GaussiansReq{ApiKey: rng.apiKey, N: n, Mean: mean, StandardDeviation: standardDeviation,
		SignificantDigits: significantDigits}
	result, err := SignedRequest("generateSignedGaussians", body)

	if err.Message != "" {
		return SignedFloatData{}, err
	}

	signedResult, _ := result.Content().(SignedResult)
	randomData := Random{}
	json.Unmarshal(signedResult.Raw, &randomData)

	data := randomData.Data
	floatArray := make([]float64, len(data))
	for i, num := range data {
		floatArray[i] = float64(num.(float64))
	}

	return SignedFloatData{
		Data:         floatArray,
		Raw:          signedResult.Raw,
		HashedApiKey: randomData.HashedApiKey,
		SerialNumber: randomData.SerialNumber,
		Signature:    signedResult.Signature,
	}, Error{}
}

// Generate `n` random strings with precision upto `decimalPlaces`.
func (rng trueRNG) GenerateSignedStrings(n, length int, characters string, replacement bool) (SignedStringData, Error) {

	body := StringsReq{ApiKey: rng.apiKey, N: n, Length: length, Characters: characters,
		Replacement: replacement}
	result, err := SignedRequest("generateSignedStrings", body)

	if err.Message != "" {
		return SignedStringData{}, err
	}

	signedResult, _ := result.Content().(SignedResult)
	randomData := Random{}
	json.Unmarshal(signedResult.Raw, &randomData)

	data := randomData.Data
	stringArray := make([]string, len(data))
	for i, string_ := range data {
		stringArray[i] = string(string_.(string))
	}

	return SignedStringData{
		Data:         stringArray,
		Raw:          signedResult.Raw,
		HashedApiKey: randomData.HashedApiKey,
		SerialNumber: randomData.SerialNumber,
		Signature:    signedResult.Signature,
	}, Error{}
}

// Generate `n` random strings with precision upto `decimalPlaces`.
func (rng trueRNG) GenerateSignedUUIDs(n int) (SignedStringData, Error) {

	body := UUIDsReq{ApiKey: rng.apiKey, N: n}
	result, err := SignedRequest("generateSignedUUIDs", body)

	if err.Message != "" {
		return SignedStringData{}, err
	}

	signedResult, _ := result.Content().(SignedResult)
	randomData := Random{}
	json.Unmarshal(signedResult.Raw, &randomData)

	data := randomData.Data
	stringArray := make([]string, len(data))
	for i, string_ := range data {
		stringArray[i] = string(string_.(string))
	}

	return SignedStringData{
		Data:         stringArray,
		Raw:          signedResult.Raw,
		HashedApiKey: randomData.HashedApiKey,
		SerialNumber: randomData.SerialNumber,
		Signature:    signedResult.Signature,
	}, Error{}
}

// Generate `n` random blobs of length `size`, formatted in `format` (either base64 or hex)
func (rng trueRNG) GenerateSignedBlobs(n, size int, format string) (SignedStringData, Error) {

	body := BlobsReq{ApiKey: rng.apiKey, N: n, Size: size, Format: format}
	result, err := SignedRequest("generateSignedBlobs", body)

	if err.Message != "" {
		return SignedStringData{}, err
	}

	signedResult, _ := result.Content().(SignedResult)
	randomData := Random{}
	json.Unmarshal(signedResult.Raw, &randomData)

	data := randomData.Data
	stringArray := make([]string, len(data))
	for i, string_ := range data {
		stringArray[i] = string(string_.(string))
	}

	return SignedStringData{
		Data:         stringArray,
		Raw:          signedResult.Raw,
		HashedApiKey: randomData.HashedApiKey,
		SerialNumber: randomData.SerialNumber,
		Signature:    signedResult.Signature,
	}, Error{}
}

// Currently broken - API consistently claims authenticity is false.
//
// This method verifies that received random data actually originates from RANDOM.org, given a raw `random`
// JSON that is exactly what is given to you by Signed<Int|Float|String>Data.Raw and a `signature`, also contained
// in Signed<Int|Float|String>Data.Signature.
func (rng trueRNG) VerifySignature(random json.RawMessage, signature string) (bool, Error) {

	// it turns out json.RawMessage does not survive across multiple marshalings. Our random data is
	// interpreted as a byte array, and its base64 encoding is sent along. To prevent this, we unmarshal
	// into an object first, and then have _request encode it into JSON again.
	var object map[string]interface{}
	json.Unmarshal(random, &object)

	body := VerifySignatureReq{Raw: object, Signature: signature}
	result, err := SignedRequest("verifySignature", body)

	if err.Message != "" {
		return false, err
	}

	return result.(VerifiedSignature).Authenticity, Error{}

}
