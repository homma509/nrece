package model

import (
	"fmt"

	"github.com/homma509/nrece/domain/value"
)

// Option ...
type Option func(*IR) error

// IR 医療機関情報レコード
type IR struct {
	RecordType       value.RecordType   // レコード識別情報
	Payer            value.Payer        // 審査支払機関
	Prefecture       value.Prefecture   // 都道府県
	PointTable       value.PointTable   // 点数表
	FacilityID       value.FacilityID   // 医療機関コード
	Reserve          value.Reserve      // 予備
	FacilityName     value.FacilityName // 医療機関名称
	InvoiceYearMonth value.YearMonth    // 請求年月
	VolumeID         value.VolumeID     // マルチボリューム識別情報
	Phone            value.Phone        // 電話番号
}

// WithReserve ...
func WithReserve(reserve value.Reserve) Option {
	return func(ir *IR) error {
		ir.Reserve = reserve
		return nil
	}
}

// WithPhone ...
func WithPhone(phone value.Phone) Option {
	return func(ir *IR) error {
		ir.Phone = phone
		return nil
	}
}

// NewIR ...
func NewIR(recordType value.RecordType,
	payer value.Payer,
	prefecture value.Prefecture,
	pointTable value.PointTable,
	facilityID value.FacilityID,
	facilityName value.FacilityName,
	invoiceYearMonth value.YearMonth,
	volumeID value.VolumeID,
	options ...Option) (*IR, error) {
	if recordType.String() != value.IR {
		return nil, fmt.Errorf("IR RecordType invalid %v", recordType)
	}
	ir := IR{
		RecordType:       recordType,
		Payer:            payer,
		Prefecture:       prefecture,
		PointTable:       pointTable,
		FacilityID:       facilityID,
		FacilityName:     facilityName,
		InvoiceYearMonth: invoiceYearMonth,
		VolumeID:         volumeID,
	}
	for _, option := range options {
		if err := option(&ir); err != nil {
			return nil, err
		}
	}
	return &ir, nil
}
