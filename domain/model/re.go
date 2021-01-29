package model

const (
	// RERecordType レセプト共通レコードのレコード識別情報
	RERecordType = "RE"
)

// RE レセプト共通レコード
type RE struct {
	RecordType   string // レコード識別情報
	ReceiptNo    uint32 // レセプト番号
	ReceiptType  string // レセプト種別
	ReceiptYM    string // 診療年月
	Name         string // 氏名
	Sex          uint8  // 男女区分
	Birth        string // 生年月日
	BenefitRate  string // 給付割合
	AdmittedAt   string // 入院年月日
	WardType     string // 病棟区分
	ChargeType   string // 一部負担金・食事療養費・生活療養費標準負担額区分
	ReceiptNote  string // レセプト特記事項
	WardCount    string // 病棟数
	KarteNo      string // カルテ番号等
	DiscountUnit string // 割引点数単価
	Reserve1     string // 予備
	Reserve2     string // 予備
	Reserve3     string // 予備
	SearchNo     string // 検索番号 ※数字30桁 uint64で足りないためstring
	Reserve4     string // 予備
	InvoiceInfo  string // 請求情報
	Subject1     string // 診療科1 診療科目
	Part1        string // 診療科1 人体の部位等
	Sex1         string // 診療科1 性別等
	Treatment1   string // 診療科1 医学的処置
	Disease1     string // 診療科1 特定疾病
	Subject2     string // 診療科2 診療科目
	Part2        string // 診療科2 人体の部位等
	Sex2         string // 診療科2 性別等
	Treatment2   string // 診療科2 医学的処置
	Disease2     string // 診療科2 特定疾病
	Subject3     string // 診療科3 診療科目
	Part3        string // 診療科3 人体の部位等
	Sex3         string // 診療科3 性別等
	Treatment3   string // 診療科3 医学的処置
	Disease3     string // 診療科3 特定疾病
	Kana         string // カタナカ（氏名）
	Condition    string // 患者の状態
}
