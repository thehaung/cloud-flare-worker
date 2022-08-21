package configs

import (
	"os"
	"strconv"
)

func GetEnvironment() string {
	env := os.Getenv("ENVIRONMENT")

	if len(env) > 0 && (env == "prod" || env == "non-prod") {
		return env
	}

	return "prod"
}

func GetTimeDeadlineAPI() int {
	res, err := strconv.Atoi(os.Getenv("TIME_DEADLINE_API"))
	if err != nil {
		panic(err)
	}
	return res
}

func GetServerPort() string {
	return os.Getenv("SERVER_PORT")
}

func GetAPIUrl() string {
	return os.Getenv("API_URL")
}

func APIGetIPUrl() string {
	return os.Getenv("API_GET_IP_URL")
}

func GetCloudFlareAPIUrl() string {
	return os.Getenv("CLOUD_FLARE_BASE_URL")
}

func GetCloudFlareEmail() string {
	return os.Getenv("CLOUD_FLARE_EMAIL")
}

func GetCloudFlareAPIToken() string {
	return os.Getenv("CLOUD_FLARE_API_TOKEN")
}

func GetCloudFlareZoneID() string {
	return os.Getenv("CLOUD_FLARE_ZONE_ID")
}
