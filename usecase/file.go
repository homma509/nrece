package usecase

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"net/url"

	"github.com/homma509/nrece/config"
	"github.com/homma509/nrece/domain/model"
	"github.com/homma509/nrece/domain/repository"
	"github.com/homma509/nrece/domain/value"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// File ...
type File struct {
	repo repository.FileRepository
}

// NewFile ファイルユースケースの生成
func NewFile(repo repository.FileRepository) *File {
	return &File{
		repo,
	}
}

// Copy ファイルをコピーし通知
func (r *File) Copy(ctx context.Context, src string) error {
	log.Println("[info] usecase file copy", src)
	dst, err := r.copyDst(ctx, src)
	if err != nil {
		return err
	}

	err = r.repo.Copy(ctx, dst, src)
	if err != nil {
		return err
	}

	err = r.repo.Publish(ctx, dst)
	if err != nil {
		return err
	}

	return nil
}

func (r *File) copyDst(ctx context.Context, src string) (string, error) {
	log.Println("[info] usecase file dst", src)
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
			ir.FacilityID.String(),
			ir.FacilityName.String(),
			ir.InvoiceYearMonth.String(),
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
		return nil, errors.New("file is empty")
	}
	if err != nil {
		return nil, err
	}

	return ir(record)
}

func ir(record []string) (*model.IR, error) {
	log.Println("[info] usecase file ir", record)

	// 0. レコード識別情報
	recordType, err := value.NewRecordType(record[0])
	if err != nil {
		return nil, fmt.Errorf("IR NewRecordType failed %v", err)
	}

	// 1. 審査支払機関
	payer, err := value.NewPayer(record[1])
	if err != nil {
		return nil, fmt.Errorf("IR NewPayer failed %v", err)
	}

	// 2. 都道府県
	prefecture, err := value.NewPrefecture(record[2])
	if err != nil {
		return nil, fmt.Errorf("IR NewPrefecture failed %v", err)
	}

	// 3. 点数表
	pointTable, err := value.NewPointTable(record[3])
	if err != nil {
		return nil, fmt.Errorf("IR NewPointTable failed %v", err)
	}

	// 4. 医療機関コード
	facilityID, err := value.NewFacilityID(record[4])
	if err != nil {
		return nil, fmt.Errorf("IR NewFacilityID failed %v", err)
	}

	// 5. 予備
	reserve, err := value.NewReserve(record[5])
	if err != nil {
		return nil, fmt.Errorf("IR NewReserve failed %v", err)
	}

	// 6. 医療機関名称
	facilityName, err := value.NewFacilityName(record[6])
	if err != nil {
		return nil, fmt.Errorf("IR NewFacilityName failed %v", err)
	}

	// 7. 請求年月
	invoiceYearMonth, err := value.NewYearMonth(record[7])
	if err != nil {
		return nil, fmt.Errorf("IR NewYearMonth failed %v", err)
	}

	// 8. マルチボリューム識別情報
	volumeID, err := value.NewVolumeID(record[8])
	if err != nil {
		return nil, fmt.Errorf("IR NewVolumeID failed %v", err)
	}

	// 9. 電話番号
	phone, err := value.NewPhone(record[9])
	if err != nil {
		return nil, fmt.Errorf("IR NewPhone failed %v", err)
	}

	return model.NewIR(
		recordType,
		payer,
		prefecture,
		pointTable,
		facilityID,
		facilityName,
		invoiceYearMonth,
		volumeID,
		model.WithReserve(reserve),
		model.WithPhone(phone),
	)

}
