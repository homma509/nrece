package model

// ReceiptItemInfo レセプト明細インターフェース
type ReceiptItemInfo interface {
	AddComment(co CO)
}

// Receipt レセプト
type Receipt struct {
	IR           IR                      // 医療機関情報レコード
	receiptItems map[uint32]*ReceiptItem // レセプト明細
}

// ReceiptItem レセプト番号指定のレセプト明細の取得
func (r *Receipt) ReceiptItem(key uint32) *ReceiptItem {
	if r.receiptItems == nil {
		r.receiptItems = map[uint32]*ReceiptItem{}
	}
	if _, ok := r.receiptItems[key]; !ok {
		r.receiptItems[key] = &ReceiptItem{
			SYInfos: []*SYInfo{},
			SIInfos: []*SIInfo{},
			COInfos: []*COInfo{},
		}
	}
	return r.receiptItems[key]
}

// ReceiptItems レセプト明細リストの取得
func (r *Receipt) ReceiptItems() map[uint32]*ReceiptItem {
	return r.receiptItems
}

// ReceiptItem レセプト明細
type ReceiptItem struct {
	RE      RE        // レセプト共通レコード
	SYInfos []*SYInfo // 傷病名情報
	SIInfos []*SIInfo // 診療行為情報
	COInfos []*COInfo // コメント情報
}

// SYInfo 傷病めい情報
type SYInfo struct {
	SY       // 診療行為レコード
	COs []CO // コメントレコード
}

// NewSYInfo 傷病名情報の生成
func NewSYInfo(sy SY) *SYInfo {
	return &SYInfo{
		SY:  sy,
		COs: []CO{},
	}
}

// AddComment コメントレコードの追加
func (in *SYInfo) AddComment(co CO) {
	in.COs = append(in.COs, co)
}

// SIInfo 診療行為情報
type SIInfo struct {
	SI       // 診療行為レコード
	COs []CO // コメントレコード
}

// NewSIInfo 診療行為情報の生成
func NewSIInfo(si SI) *SIInfo {
	return &SIInfo{
		SI:  si,
		COs: []CO{},
	}
}

// AddComment コメントレコードの追加
func (in *SIInfo) AddComment(co CO) {
	in.COs = append(in.COs, co)
}

// COInfo コメント情報
type COInfo struct {
	CO       // コメントレコード
	COs []CO // コメントレコード
}

// NewCOInfo コメント情報の生成
func NewCOInfo(co CO) *COInfo {
	return &COInfo{
		CO:  co,
		COs: []CO{},
	}
}

// AddComment コメントレコードの追加
func (in *COInfo) AddComment(co CO) {
	in.COs = append(in.COs, co)
}
