package credentials

import (
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strings"
)

const (
	awsCredentialsFile = ".aws/credentials"
	profileNamePrefix  = "["
	profileNameSuffix  = "]"

	accessKeyPrefix = "aws_access_key_id="

	secretKeyPrefix = "aws_secret_access_key="
)

func Parse() ([]Profile, error) {
	credentials, err := readCredentials()

	if err != nil {
		return nil, err
	}

	return parse(credentials)
}

// Reads the default aws credentials file
func readCredentials() (string, error) {

	usr, err := user.Current()

	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadFile(filepath.Join(usr.HomeDir, awsCredentialsFile))

	if err != nil {
		return "", err
	}

	return string(data), err
}

func parse(s string) ([]Profile, error) {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	var profiles []Profile

	for i := 0; i <= len(lines)-3; i += 3 {

		var p Profile
		p.Name = lines[i]

		filterKey(lines[i+1], &p)
		filterKey(lines[i+2], &p)

		//ToDO profile validation
		profiles = append(profiles, p)
	}
	return profiles, nil
}

func filterKey(s string, p *Profile) {
	s = strings.Replace(s, " ", "", -1)

	if strings.HasPrefix(s, accessKeyPrefix) {

		p.AccessKey = strings.Replace(s, accessKeyPrefix, "", -1)

	} else if strings.HasPrefix(s, secretKeyPrefix) {

		p.SecretKey = strings.Replace(s, secretKeyPrefix, "", -1)

	}

}
