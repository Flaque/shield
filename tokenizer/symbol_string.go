package tokenizer

import "strconv"

const _Symbol_name = "FUNCLABELDOESBYCOLONDONEL_PARENR_PARENL_SQUIGR_SQUIGPLUSMINUSEQUALSASSIGNRETURNISNUMCOMMAEND"

var _Symbol_index = [...]uint8{0, 4, 9, 13, 15, 20, 24, 31, 38, 45, 52, 56, 61, 67, 73, 79, 81, 84, 89, 92}

func (i Symbol) String() string {
	if i < 0 || i >= Symbol(len(_Symbol_index)-1) {
		return "Symbol(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Symbol_name[_Symbol_index[i]:_Symbol_index[i+1]]
}
