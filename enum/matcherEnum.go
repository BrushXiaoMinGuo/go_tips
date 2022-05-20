package enum

type MatchType int

const (
	MatchEqual MatchType = iota
	MatchNotEqual
	MatchRegexp
	MatchNotRegexp
)

// 不确定长度的使用...
// 可以指定下标赋值
var matchTypeStr = [...]string{
	MatchEqual:     "=",
	MatchNotEqual:  "!=",
	MatchRegexp:    "~",
	MatchNotRegexp: "!~",
}

func (m MatchType) String() string {
	if m < MatchEqual || m > MatchNotRegexp {
		return "error"
	}
	return matchTypeStr[m]
}
