package credentials

import (
	"os"
	"testing"
)

func Test_EnvParse(t *testing.T) {

	access := "Hi aws"
	secret := "hope you are doing great"

	os.Setenv(AccessKey, access)
	os.Setenv(SecretKey, secret)

	parsedAccess, parsedSecret := EnvReading()

	if parsedAccess != access {
		t.Fatalf("Unexpected AWS Access Key returned, expected=%v, got=%v", access, parsedAccess)
	} else if parsedSecret != secret {
		t.Fatalf("Unexpected Secret Key returned, expected=%v, got=%v", secret, parsedSecret)
	}
}
