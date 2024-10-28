package swx

import "os"

const apiUrlEnvVarName = "SWX_API_URL"

func GetApiUrl() string {
	return os.Getenv(apiUrlEnvVarName)
}
