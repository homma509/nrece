package infra

import (
	"context"
	"errors"
	"io"
	"log"
	"net/url"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/homma509/nrece/config"
)

// S3Repository ...
type S3Repository struct {
	conn *s3.S3
}

// NewS3Repository ...
func NewS3Repository() (*S3Repository, error) {
	log.Println("[info] infra new s3 repository")
	sess, err := open()
	if err != nil {
		return nil, err
	}
	return &S3Repository{
		conn: s3.New(sess),
	}, nil
}

// Get ...
func (h *S3Repository) Get(ctx context.Context, s3url string) (io.ReadCloser, error) {
	log.Println("[info] infra s3 Get", s3url)
	u, err := toURL(s3url)
	if err != nil {
		return nil, err
	}

	out, err := h.conn.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String(u.Host),
		Key:    aws.String(strings.TrimPrefix(u.Path, "/")),
	})
	if err != nil {
		return nil, err
	}

	return out.Body, nil
}

func toURL(s3url string) (*url.URL, error) {
	log.Println("[info] infra s3 toURL", s3url)
	u, err := url.Parse(s3url)
	if err != nil {
		return nil, err
	}
	if u.Scheme != "s3" {
		return nil, errors.New("s3:// required")
	}
	return u, nil
}

// Move ...
func (h *S3Repository) Move(ctx context.Context, src, dest string) error {
	log.Println("[info] infra s3 move")
	u, err := url.Parse(dest)
	if err != nil {
		return err
	}

	_, err = h.conn.CopyObjectWithContext(ctx, &s3.CopyObjectInput{
		CopySource: aws.String(src),
		Bucket:     aws.String(u.Host),
		Key:        aws.String(strings.TrimPrefix(u.Path, "/")),
	})
	if err != nil {
		return err
	}
	return nil
}

func open() (*session.Session, error) {
	c := config.NewAWSConfig()
	log.Println("[info] infra s3 open", c)
	return session.NewSession(c)
}
