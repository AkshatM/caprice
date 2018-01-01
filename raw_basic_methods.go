package caprice

// Generate `n` random integers between `min` and `max`, but return the raw JSON from the API as a formatted Result
// struct. If `replacement` is true, pick random numbers with replacement. Default is false.
// We do not support base selection, since it is easy to format into the base of your choice from base 10
func (rng trueRNG) GenerateIntegersRaw(n, min, max int, replacement bool) (Result, Error) {

	body := IntegersReq{ApiKey: rng.apiKey, N: n, Min: min, Max: max, Replacement: replacement}
	result, err := Request("generateIntegers", body)
	if err.Message != "" {
		return Result{}, err
	}
	return result.(Result), err
}

// Generate `n` random decimal fractions with precision upto `decimalPlaces`, but return raw JSON from the API
// as a formatted Result struct. If `replacement` is true, pick random numbers with replacement. Default is false.
func (rng trueRNG) GenerateDecimalFractionsRaw(n, decimalPlaces int, replacement bool) (Result, Error) {

	body := DecimalFractionsReq{ApiKey: rng.apiKey, N: n, DecimalPlaces: decimalPlaces, Replacement: replacement}
	result, err := Request("generateDecimalFractions", body)
	if err.Message != "" {
		return Result{}, err
	}
	return result.(Result), err
}

// Generate `n` Gaussians from a disribution with mean `mean` and stdev `standardDeviation`, returned with
// at most `significantDigits` sig. digits, but return the raw JSON response as a formatted Result struct
func (rng trueRNG) GenerateGaussiansRaw(n int, mean, standardDeviation float64, significantDigits int) (Result, Error) {

	body := GaussiansReq{ApiKey: rng.apiKey, N: n, Mean: mean, StandardDeviation: standardDeviation,
		SignificantDigits: significantDigits}
	result, err := Request("generateGaussians", body)
	if err.Message != "" {
		return Result{}, err
	}
	return result.(Result), err
}

// Generate `n` random strings with precision upto `decimalPlaces`, but return the raw JSON response as a
// formatted Result struct
func (rng trueRNG) GenerateStringsRaw(n, length int, characters string, replacement bool) (Result, Error) {

	body := StringsReq{ApiKey: rng.apiKey, N: n, Length: length, Characters: characters,
		Replacement: replacement}
	result, err := Request("generateStrings", body)
	if err.Message != "" {
		return Result{}, err
	}
	return result.(Result), err
}

// Generate `n` random strings with precision upto `decimalPlaces`, but return the raw JSON response as a
// formatted Result struct
func (rng trueRNG) GenerateUUIDsRaw(n int) (Result, Error) {

	body := UUIDsReq{ApiKey: rng.apiKey, N: n}
	result, err := Request("generateUUIDs", body)
	if err.Message != "" {
		return Result{}, err
	}
	return result.(Result), err
}

// Generate `n` random blobs of length `size`, formatted in `format` (either base64 or hex), but return the
// raw JSON response as a formatted Result struct
func (rng trueRNG) GenerateBlobsRaw(n, size int, format string) (Result, Error) {

	body := BlobsReq{ApiKey: rng.apiKey, N: n, Size: size, Format: format}
	result, err := Request("generateBlobs", body)
	if err.Message != "" {
		return Result{}, err
	}
	return result.(Result), err
}
