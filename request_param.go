package password_service

const (
	// MinTotalLen describes min total length
	MinTotalLen = 5
	// MaxTotalLen describes max total length
	MaxTotalLen = 15
	// MaxCount describes max count passwords
	MaxCount = 50
	// defCount describes default count passwords
	defCount = 1
	// defLen describes default length
	defLen = 0
)

// RequestParam describes struct for request params.
type RequestParam struct {
	LettersLen int `json:"letters_len"`
	SpecChLen  int `json:"spec_ch_len"`
	NumbersLen int `json:"numbers_len"`
	Count      int `json:"count"`
}

// DefVal set default values.
func (rp *RequestParam) DefVal() {
	if rp.LettersLen <= 0 {
		rp.LettersLen = defLen
	}

	if rp.SpecChLen <= 0 {
		rp.SpecChLen = defLen
	}

	if rp.NumbersLen <= 0 {
		rp.NumbersLen = defLen
	}

	if rp.Count <= 0 {
		rp.Count = defCount
	}
}

// IsValidReqParamsLen validates min and max length.
func IsValidReqParamsLen(rp RequestParam) bool {
	totalLen := rp.LettersLen + rp.SpecChLen + rp.NumbersLen
	return totalLen >= MinTotalLen && totalLen <= MaxTotalLen
}

// IsValidReqParamsCount validates count.
func IsValidReqParamsCount(rp RequestParam) bool {
	return rp.Count <= MaxCount
}
