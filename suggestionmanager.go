package hunspell

type SuggestionManager struct {
}

func NewSuggestionManager(try string, maxSuggestions int, affxMng *AffixManager) *SuggestionManager {
	return &SuggestionManager{}
}
