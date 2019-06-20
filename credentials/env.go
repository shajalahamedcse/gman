package credentials

import "os"

const (
	// AccessKey is the name of the AWS access key environment variable.
	AccessKey = "AWS_ACCESS_KEY_ID"

	// SecretKey is the name of the AWS secret key environment variable.
	SecretKey = "AWS_SECRET_ACCESS_KEY"
)

// EnvParse returns the current access and secret key evironment variables.
func EnvReading() (string, string) {
	return os.Getenv(AccessKey), os.Getenv(SecretKey)
}
