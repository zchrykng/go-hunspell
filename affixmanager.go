package hunspell

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
