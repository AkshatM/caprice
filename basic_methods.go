package caprice

// Generate `n` random integers between `min` and `max`.
// If `replacement` is true, pick random numbers with replacement. Default is false.
// We do not support base selection, since it is easy to format into the base of your choice from base 10
func (rng trueRNG) GenerateIntegers(n, min, max int, replacement bool) ([]int, Error) {

	result, err := rng.GenerateIntegersRaw(n, min, max, replacement)

	if err.Message != "" {
		return []int{}, err
	}

	data, _ := result.Content().([]interface{})

	intArray := make([]int, len(data))
	for i, num := range data {
		intArray[i] = int(num.(float64))
	}

	return intArray, Error{}
}

// Generate `n` random decimal fractions with precision upto `decimalPlaces`.
// If `replacement` is true, pick random numbers with replacement. Default is false.
func (rng trueRNG) GenerateDecimalFractions(n, decimalPlaces int, replacement bool) ([]float64, Error) {

	result, err := rng.GenerateDecimalFractionsRaw(n, decimalPlaces, replacement)

	if err.Message != "" {
		return []float64{}, err
	}

	data, _ := result.Content().([]interface{})

	floatArray := make([]float64, len(data))
	for i, num := range data {
		floatArray[i] = float64(num.(float64))
	}

	return floatArray, Error{}
}

// Generate `n` Gaussians from a disribution with mean `mean` and stdev `standardDeviation`, returned with
// at most `significantDigits` sig. digits.
func (rng trueRNG) GenerateGaussians(n int, mean, standardDeviation float64, significantDigits int) ([]float64, Error) {

	result, err := rng.GenerateGaussiansRaw(n, mean, standardDeviation, significantDigits)

	if err.Message != "" {
		return []float64{}, err
	}

	data, _ := result.Content().([]interface{})

	floatArray := make([]float64, len(data))
	for i, num := range data {
		floatArray[i] = float64(num.(float64))
	}

	return floatArray, Error{}
}

// Generate `n` random strings with precision upto `decimalPlaces`.
// If `replacement` is true, pick random numbers with replacement. Default is false.
func (rng trueRNG) GenerateStrings(n, length int, characters string, replacement bool) ([]string, Error) {

	result, err := rng.GenerateStringsRaw(n, length, characters, replacement)

	if err.Message != "" {
		return []string{}, err
	}

	data, _ := result.Content().([]interface{})

	stringArray := make([]string, len(data))
	for i, string_ := range data {
		stringArray[i] = string(string_.(string))
	}

	return stringArray, Error{}
}

// Generate `n` random strings with precision upto `decimalPlaces`.
func (rng trueRNG) GenerateUUIDs(n int) ([]string, Error) {

	result, err := rng.GenerateUUIDsRaw(n)

	if err.Message != "" {
		return []string{}, err
	}

	data, _ := result.Content().([]interface{})

	stringArray := make([]string, len(data))
	for i, string_ := range data {
		stringArray[i] = string(string_.(string))
	}

	return stringArray, Error{}
}

// Generate `n` random blobs of length `size`, formatted in `format` (either base64 or hex)
func (rng trueRNG) GenerateBlobs(n, size int, format string) ([]string, Error) {

	result, err := rng.GenerateBlobsRaw(n, size, format)

	if err.Message != "" {
		return []string{}, err
	}

	data, _ := result.Content().([]interface{})

	stringArray := make([]string, len(data))
	for i, string_ := range data {
		stringArray[i] = string(string_.(string))
	}

	return stringArray, Error{}
}

// Get information about current usage as a formatted Status struct.
func (rng trueRNG) GetUsage() (Status, Error) {
	body := StatusReq{ApiKey: rng.apiKey}
	status, err := Request("getUsage", body)
	if err.Message != "" {
		return Status{}, err
	}
	return status.Content().(Status), err
}
