package hunspell

type Flag int64

const (
	FlagNone Flag = iota
)

type hentry struct {
}

type HashManager struct {
	tablePtr        []*hentry
	flagMode        Flag
	complexPrefixes int
	forbiddenWord   uint
	langNum         int
	enc             string
	lang            string
	ignoreChars     string
	aliasf          [][]uint
	aliasfLen       []uint
	aliasm          [][]rune
	repTable        []replEntry
}

func NewHashManager(affPath, dPath, key string) *HashManager {
	return &HashManager{}
}
