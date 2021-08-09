package cmd

import (
	"os"

	"github.com/LF-Engineering/insights-cache-projects/projectscache"
	"github.com/labstack/gommon/log"
)

const (
	METRICS_ENDPOINT_PROD = "https://metrics.lfanalytics.io/v1/projects"
	METRICS_ENDPOINT_TEST = "https://metrics.insights.test.platform.linuxfoundation.org/v1/projects"
	S3_BUCKET_PROD        = "insights-v1-prod"
	S3_BUCKET_TEST        = "insights-v1-test"
)

func Handler() {
	env := os.Getenv("ENVIRONMENT")

	s3Bucket := ""
	projectEndpoint := ""
	if env == "prod" {
		projectEndpoint = METRICS_ENDPOINT_PROD
		s3Bucket = S3_BUCKET_PROD
	} else {
		projectEndpoint = METRICS_ENDPOINT_TEST
		s3Bucket = S3_BUCKET_TEST
	}

	projects, err := projectscache.GetProjects(projectEndpoint)
	if err != nil {
		log.Fatalf("Failed to fetch projects: %+v", err)
	}

	// upload to s3
	err = projectscache.CopyToS3(projects, s3Bucket)
	if err != nil {
		log.Fatalf("Failed to upload to s3: %+v", err)
	}

}
