package projectscache

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/gommon/log"

	jsoniter "github.com/json-iterator/go"
)

const (
	PAGE_SIZE  = 10
	s3FILENAME = "projects.json"
)

func GetProjects(url string) (projects *ProjectsData, err error) {
	method := "GET"
	currentPage := 1
	totalSize := 0

	projects = &ProjectsData{}

	log.Infof("Fetching projects from %s", url)
	for {

		var req *http.Request
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return
		}
		q := req.URL.Query()
		q.Add("pageNo", fmt.Sprintf("%d", currentPage))
		q.Add("pageSize", fmt.Sprintf("%d", PAGE_SIZE))
		req.URL.RawQuery = q.Encode()

		var resp *http.Response
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			return
		}

		var body []byte
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}

		var data ProjectsData
		err = jsoniter.Unmarshal(body, &data)
		if err != nil {
			return
		}

		if currentPage == 1 {
			totalSize = data.Total
			projects.Total = totalSize
		}

		log.Infof("(Total Projects: %d) <=> Fetched page %d (%d projects/page) ", totalSize, currentPage, PAGE_SIZE)

		projects.Projects = append(projects.Projects, data.Projects...)

		if totalSize <= (currentPage * PAGE_SIZE) {
			break
		}

		currentPage++
	}

	return projects, nil
}

func CopyToS3(projects *ProjectsData, bucket string) (err error) {

	projectsJson, err := jsoniter.Marshal(projects)
	if err != nil {
		fmt.Println(err)
	}

	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession())

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	objectBody := bytes.NewReader(projectsJson)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(s3FILENAME),
		Body:   objectBody,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	log.Infof("%s uploaded to %s", s3FILENAME, aws.StringValue(&result.Location))

	return nil

}
