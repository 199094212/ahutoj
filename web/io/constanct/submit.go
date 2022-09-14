package constanct

type OJResult string
type LANG int64

const (
	OJ_AC      OJResult = "AC"
	OJ_WA      OJResult = "WA"
	OJ_TLE     OJResult = "TLE"
	OJ_MLE     OJResult = "MLE"
	OJ_RE      OJResult = "RE"
	OJ_PE      OJResult = "PE"
	OJ_OLE     OJResult = "OLE"
	OJ_CE      OJResult = "CE"
	OJ_JUDGE   OJResult = "JUDGEING"
	OJ_REJUDGE OJResult = "REJUDGEING"
)
const (
	C       LANG = 1
	CPP     LANG = 2
	CPP11   LANG = 3
	CPP17   LANG = 4
	JAVA    LANG = 5
	PYTHON3 LANG = 6
)

const (
	ISPRIVATE int64 = 0
	ISPUBLIC  int64 = 1
)