package helpers

import (
	"bytes"
	"context"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/codingtroop/ubl-store/pkg/helpers/interfaces"
)

type s3Storer struct {
	endpoint string
	folder   string
	bucket   string
}

func NewS3Storer(endpoint string, f string, b string) interfaces.Storer {
	return &s3Storer{folder: f, bucket: b, endpoint: endpoint}
}

func (h *s3Storer) getSession(c context.Context) *session.Session {
	defaultResolver := endpoints.DefaultResolver()

	s3CustResolverFn := func(service, region string, optFns ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
		if service == "s3" && h.endpoint != "" {
			return endpoints.ResolvedEndpoint{
				URL: h.endpoint,
			}, nil
		}

		return defaultResolver.EndpointFor(service, region, optFns...)
	}
	// The session the S3 Uploader will use
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			EndpointResolver: endpoints.ResolverFunc(s3CustResolverFn),
			S3ForcePathStyle: aws.Bool(true),
		},
	}))

	return sess
}

func (h *s3Storer) Exists(c context.Context, hash string) (bool, error) {

	s3Svc := s3.New(h.getSession(c))

	_, err := s3Svc.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(h.bucket),
		Key:    aws.String(h.folder + "/" + hash),
	})

	if err == nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case "NotFound": // s3.ErrCodeNoSuchKey does not work, aws is missing this error code so we hardwire a string
				return false, nil
			default:
				return false, err
			}
		}
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

func (h *s3Storer) Read(c context.Context, uuid string) ([]byte, error) {

	s3Svc := s3.New(h.getSession(c))

	out, err := s3Svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(h.bucket),
		Key:    aws.String(h.folder + "/" + uuid),
	})

	if err != nil {
		return nil, err
	}

	defer out.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(out.Body)
	return buf.Bytes(), nil

}

func (h *s3Storer) Write(c context.Context, uuid string, data []byte) error {

	// S3 service client the Upload manager will use.
	s3Svc := s3.New(h.getSession(c))

	// Create an uploader with S3 client and default options
	uploader := s3manager.NewUploaderWithClient(s3Svc)

	key := h.folder + "/" + uuid
	r := bytes.NewReader(data)

	// Upload input parameters
	upParams := &s3manager.UploadInput{
		Bucket: &h.bucket,
		Key:    &key,
		Body:   r,
	}

	// Perform an upload.
	_, err := uploader.Upload(upParams)

	return err
}
