package caprice

import (
	"os"
	"testing"
)

func TestRequests(t *testing.T) {
	t.Skip("Don't run until I can get a decent testing facility")
}

func TestGenerateIntegers(t *testing.T) {

	t.Run("Test GenerateIntegers network call", func(t *testing.T) {
		if os.Getenv("APIKEY") == "" {
			t.Skip("Skipping direct network call tests because APIKEY env variable is not set")
		}
		rng := TrueRNG(os.Getenv("APIKEY"))
		numbers, err := rng.GenerateIntegers(5, 1, 10, true)
		t.Log(numbers)
		if err.Message != "" {
			t.Error(err)
		}
	})
}

func TestGenerateSignedIntegers(t *testing.T) {

	t.Run("Test GenerateSignedIntegers network call", func(t *testing.T) {
		if os.Getenv("APIKEY") == "" {
			t.Skip("Skipping direct network call tests because APIKEY env variable is not set")
		}
		rng := TrueRNG(os.Getenv("APIKEY"))
		result, err := rng.GenerateSignedIntegers(5, 1, 10, true)
		if err.Message != "" {
			t.Error(err)
		}

		t.Run("Verify signature", func(t *testing.T) {
			result, err := rng.VerifySignature(result.Raw, result.Signature)
			if err.Message != "" {
				t.Error(err)
			}
			if result == false {
				t.Error("Signature verification failed. Known issue: https://github.com/AkshatM/caprice/issues/6")
			}
		})
	})
}

func TestGenerateDecimalFractions(t *testing.T) {
	t.Run("Test GenerateDecimalFractions network call", func(t *testing.T) {
		if os.Getenv("APIKEY") == "" {
			t.Skip("Skipping direct network call tests because APIKEY env variable is not set")
		}
		rng := TrueRNG(os.Getenv("APIKEY"))
		numbers, err := rng.GenerateDecimalFractions(5, 10, true)
		t.Log(numbers)
		if err.Message != "" {
			t.Error(err)
		}
	})
}

func TestGenerateSignedDecimalFractions(t *testing.T) {
	t.Run("Test GenerateSignedDecimalFractions network call", func(t *testing.T) {
		if os.Getenv("APIKEY") == "" {
			t.Skip("Skipping direct network call tests because APIKEY env variable is not set")
		}
		rng := TrueRNG(os.Getenv("APIKEY"))
		numbers, err := rng.GenerateSignedDecimalFractions(5, 10, true)
		t.Logf("%+v", numbers)
		if err.Message != "" {
			t.Error(err)
		}
	})
}

func TestGenerateGaussians(t *testing.T) {
	t.Run("Test GenerateGaussians network call", func(t *testing.T) {
		if os.Getenv("APIKEY") == "" {
			t.Skip("Skipping direct network call tests because APIKEY env variable is not set")
		}
		rng := TrueRNG(os.Getenv("APIKEY"))
		numbers, err := rng.GenerateGaussians(10, 5, 1.4, 12)
		t.Log(numbers)
		if err.Message != "" {
			t.Error(err)
		}
	})
}

func TestGenerateSignedGaussians(t *testing.T) {
	t.Run("Test GenerateSignedGaussians network call", func(t *testing.T) {
		if os.Getenv("APIKEY") == "" {
			t.Skip("Skipping direct network call tests because APIKEY env variable is not set")
		}
		rng := TrueRNG(os.Getenv("APIKEY"))
		numbers, err := rng.GenerateSignedGaussians(10, 5, 1.4, 12)
		t.Logf("%+v", numbers)
		if err.Message != "" {
			t.Error(err)
		}
	})
}

func TestGenerateStrings(t *testing.T) {
	t.Run("Test GenerateStrings network call", func(t *testing.T) {
		if os.Getenv("APIKEY") == "" {
			t.Skip("Skipping direct network call tests because APIKEY env variable is not set")
		}
		rng := TrueRNG(os.Getenv("APIKEY"))
		strings, err := rng.GenerateStrings(10, 12, "ab%⌘", false)
		t.Log(strings)
		if err.Message != "" {
			t.Error(err)
		}
	})
}

func TestGenerateSignedStrings(t *testing.T) {
	t.Run("Test GenerateSignedStrings network call", func(t *testing.T) {
		if os.Getenv("APIKEY") == "" {
			t.Skip("Skipping direct network call tests because APIKEY env variable is not set")
		}
		rng := TrueRNG(os.Getenv("APIKEY"))
		strings, err := rng.GenerateSignedStrings(10, 12, "ab%⌘", false)
		t.Logf("%+v", strings)
		if err.Message != "" {
			t.Error(err)
		}
	})
}

func TestGenerateUUIDs(t *testing.T) {
	t.Run("Test GenerateUUIDs network call", func(t *testing.T) {
		if os.Getenv("APIKEY") == "" {
			t.Skip("Skipping direct network call tests because APIKEY env variable is not set")
		}
		rng := TrueRNG(os.Getenv("APIKEY"))
		strings, err := rng.GenerateUUIDs(10)
		t.Log(strings)
		if err.Message != "" {
			t.Error(err)
		}
	})
}

func TestGenerateSignedUUIDs(t *testing.T) {
	t.Run("Test GenerateSignedUUIDs network call", func(t *testing.T) {
		if os.Getenv("APIKEY") == "" {
			t.Skip("Skipping direct network call tests because APIKEY env variable is not set")
		}
		rng := TrueRNG(os.Getenv("APIKEY"))
		strings, err := rng.GenerateSignedUUIDs(10)
		t.Logf("%+v", strings)
		if err.Message != "" {
			t.Error(err)
		}
	})
}

func TestGenerateBlobs(t *testing.T) {
	t.Run("Test GenerateBlobs network call", func(t *testing.T) {
		if os.Getenv("APIKEY") == "" {
			t.Skip("Skipping direct network call tests because APIKEY env variable is not set")
		}
		rng := TrueRNG(os.Getenv("APIKEY"))
		strings, err := rng.GenerateBlobs(10, 8, "base64")
		t.Log(strings)
		if err.Message != "" {
			t.Error(err)
		}
	})
}

func TestGenerateSignedBlobs(t *testing.T) {
	t.Run("Test GenerateSignedBlobs network call", func(t *testing.T) {
		if os.Getenv("APIKEY") == "" {
			t.Skip("Skipping direct network call tests because APIKEY env variable is not set")
		}
		rng := TrueRNG(os.Getenv("APIKEY"))
		strings, err := rng.GenerateSignedBlobs(10, 8, "base64")
		t.Logf("%+v", strings)
		if err.Message != "" {
			t.Error(err)
		}
	})
}

func TestGetUsage(t *testing.T) {

	t.Run("Test GetUsage network call", func(t *testing.T) {
		if os.Getenv("APIKEY") == "" {
			t.Skip("Skipping direct network call tests because APIKEY env variable is not set")
		}
		rng := TrueRNG(os.Getenv("APIKEY"))
		response, err := rng.GetUsage()
		t.Logf("%+v", response)
		if err.Message != "" {
			t.Error(err)
		}
	})
}
