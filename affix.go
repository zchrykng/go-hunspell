package hunspell

import (
	"bufio"
	"os"
	"strings"
)

type AffixType string

const (
	SetChar         AffixType = "SET"
	TryChar         AffixType = "TRY"
	EncodingConvert AffixType = "ICONV"
	NoSuggest       AffixType = "NOSUGGEST"
	CompoundMin     AffixType = "COMPOUNDMIN"
	OnlyInCompound  AffixType = "ONLYINCOMPOUND"
	CompoundRule    AffixType = "COMPOUNDRULE"
	WordChars       AffixType = "WORDCHARS"
	Prefix          AffixType = "PFX"
	Suffix          AffixType = "SFX"
	Replacement     AffixType = "REP"
)

type AffixList struct {
	lines []string
}

func NewAffixList(affPath string) *AffixList {
	al := &AffixList{
		lines: []string{},
	}
	file, err := os.Open(affPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		al.lines = append(al.lines, line)

		if line[0] == '#' {
			continue
		}

		tokens := strings.Split(line, " ")

	}

	return al
}

type AffixManager struct {
}

func NewAffixManager(affPath string, hshMngs []*HashManager, key string) *AffixManager {
	return &AffixManager{}
}

func (a *AffixManager) GetTryString() string {

}

func (a *AffixManager) GetEncoding() string {

}

func (a *AffixManager) GetLangNum() int {

}

func (a *AffixManager) GetUTF8() bool {

}

func (a *AffixManager) GetComplexPrefixes() string {

}

func (a *AffixManager) GetBreaktable() string {

}
