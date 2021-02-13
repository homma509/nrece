package usecase

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"net/url"
	"strconv"

	"github.com/homma509/nrece/config"
	"github.com/homma509/nrece/domain/model"
	"github.com/homma509/nrece/domain/repository"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// // ReceiptUsecase レセプトユースケースのインターフェース
// type ReceiptUsecase interface {
// 	Copy(ctx context.Context, s3URL string) error
// 	Store(ctx context.Context, s3URL string) error
// }

// // ReceiptFile レセプトファイルのインターフェース
// type ReceiptFile interface {
// 	GetObject(bucket, key string) (io.ReadCloser, error)
// 	CopyObject(srcBucket, srcKey, dstBucket, dstKey string) error
// }

// Receipt ...
type Receipt struct {
	repo repository.ReceiptRepository
}

// NewReceipt レセプトユースケースの生成
func NewReceipt(repo repository.ReceiptRepository) *Receipt {
	return &Receipt{
		repo,
	}
}

// Copy レセプトファイルのコピー
func (r *Receipt) Copy(ctx context.Context, src string) error {
	log.Println("[info] usecase receipt copy", src)
	dst, err := r.copyDst(ctx, src)
	if err != nil {
		return err
	}

	return r.repo.Copy(ctx, dst, src)
}

func (r *Receipt) copyDst(ctx context.Context, src string) (string, error) {
	log.Println("[info] usecase receipt dst", src)
	f, err := r.repo.Get(ctx, src)
	if err != nil {
		return "", err
	}
	defer f.Close()

	ir, err := readIR(f)
	if err != nil {
		return "", err
	}

	u := url.URL{
		Scheme: "s3",
		Host:   config.Env().BucketName(),
		Path: fmt.Sprintf("%s/%s/%s_%s.UKE",
			config.Env().BucketFolderName(),
			ir.FacilityID,
			ir.FacilityName,
			ir.InvoiceYM,
		),
	}

	return u.String(), nil
}

func readIR(f io.ReadCloser) (*model.IR, error) {
	r := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
	r.Comma = ','
	r.FieldsPerRecord = -1
	r.ReuseRecord = true

	record, err := r.Read()
	if err == io.EOF {
		return nil, errors.New("receipt file is empty")
	}
	if err != nil {
		return nil, err
	}

	return ir(record)
}

func ir(record []string) (*model.IR, error) {
	log.Println("[info] usecase receipt ir", record)
	if record[0] != model.IRRecordType {
		return nil, fmt.Errorf("ir RecordType invalid value %s", record[0])
	}
	payer, err := strconv.ParseUint(record[1], 10, 8)
	if err != nil {
		return nil, err
	}
	pointTable, err := strconv.ParseUint(record[3], 10, 8)
	if err != nil {
		return nil, err
	}

	return &model.IR{
		RecordType:    record[0],
		Payer:         uint8(payer),
		Prefecture:    record[2],
		PointTable:    uint8(pointTable),
		FacilityID:    record[4],
		Reserve:       record[5],
		FacilityName:  record[6],
		InvoiceYM:     record[7],
		MultiVolumeNo: record[8],
		Phone:         record[9],
	}, nil
}
