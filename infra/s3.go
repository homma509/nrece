package infra

import (
	"context"
	"errors"
	"fmt"
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
	s3 := &S3Repository{}
	if err := s3.open(); err != nil {
		return nil, err
	}
	return s3, nil
}

func (h *S3Repository) open() error {
	if h.conn == nil {
		c := config.NewAWSConfig()
		log.Println("[info] infra s3 open", c)

		sess, err := session.NewSession(c)
		if err != nil {
			log.Println("[error] infra s3 open", err)
			return err
		}
		h.conn = s3.New(sess)
	}
	return nil
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

// Copy ...
func (h *S3Repository) Copy(ctx context.Context, dst, src string) error {
	log.Printf("[info] infra s3 copy %s from %s", dst, src)
	dstURL, err := url.Parse(dst)
	if err != nil {
		return err
	}
	srcURL, err := url.Parse(src)
	if err != nil {
		return err
	}

	in := &s3.CopyObjectInput{
		CopySource: aws.String(fmt.Sprintf("%s%s", srcURL.Host, srcURL.Path)),
		Bucket:     aws.String(dstURL.Host),
		Key:        aws.String(strings.TrimPrefix(dstURL.Path, "/")),
	}
	log.Println("[info] infra s3 copy", in)

	if _, err := h.conn.CopyObjectWithContext(ctx, in); err != nil {
		log.Println("[error] infra s3 copy", err)
		return err
	}
	return nil
}
