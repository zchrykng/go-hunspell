package hunspell

type Hunspell struct {
	hashManagers      []*HashManager
	affixManager      *AffixManager
	suggestionManager *SuggestionManager
	affixPath         string
	encoding          string
	languageNum       int
	complexprefixes   bool
	wordbreak         []string
}

// LIBHUNSPELL_DLL_EXPORTED Hunhandle* Hunspell_create(const char* affpath,
// const char* dpath);
// NewHunspell
func NewHunspell(affPath, dPath string) *Hunspell {

	return &Hunspell{}
}

// LIBHUNSPELL_DLL_EXPORTED Hunhandle* Hunspell_create_key(const char* affpath,
// const char* dpath,
// const char* key);
// NewHunspellKey
func NewHunspellKey(affPath, dPath, key string) *Hunspell {
	csconv := nil

	// first set up the hash manager
	hshMngs := []*HashManager{NewHashManager(affPath, dPath, key)}

	// next set up the affix manager
	// it needs access to the hash manager lookup methods
	affxMng := NewAffixManager(affPath, hshMngs, key)

	tryString := affxMng.GetTryString()
	encoding := affxMng.GetEncoding()
	langnum := affxMng.GetLangNum()

	complexPrefixes := affxMng.GetComplexPrefixes()
	wordbreak := affxMng.GetBreaktable()

	suggestionManager := NewSuggestionManager(tryString, MaxSuggestions, affxMng)

	h := &Hunspell{
		hashManagers:      hshMngs,
		affixManager:      affxMng,
		suggestionManager: suggestionManager,
		affixPath:         affPath,
		encoding:          encoding,
		languageNum:       langnum,
		complexprefixes:   complexPrefixes,
		wordbreak:         wordbreak,
	}

	return h
}

// LIBHUNSPELL_DLL_EXPORTED void Hunspell_destroy(Hunhandle* pHunspell);
// Destroy
func (h *Hunspell) Destroy() error {
	return nil
}

// /* load extra dictionaries (only dic files)
//   - output: 0 = additional dictionary slots available, 1 = slots are now full*/
//
// LIBHUNSPELL_DLL_EXPORTED int Hunspell_add_dic(Hunhandle* pHunspell,
// const char* dpath);
// LoadDic loads extra dictionaries (only dic files)
func (h *Hunspell) LoadDic(dPath string) error {

	return nil
}

// /* spell(word) - spellcheck word
//   - output: false = bad word, true = good word
//     *
//   - plus output:
//   - info: information bit array, fields:
//   - SPELL_COMPOUND  = a compound word
//   - SPELL_FORBIDDEN = an explicit forbidden word
//   - root: root (stem), when input is a word with affix(es)
//     */
//
// bool spell(const std::string& word, int* info = NULL, std::string* root = NULL);
// SpellCheck word
func (h *Hunspell) SpellCheck(word string) (bool, error) {

	return true, nil
}

// /* suggest(suggestions, word) - search suggestions
//   - input: pointer to an array of strings pointer and the (bad) word
//   - array of strings pointer (here *slst) may not be initialized
//   - output: number of suggestions in string array, and suggestions in
//   - a newly allocated array of strings (*slts will be NULL when number
//   - of suggestion equals 0.)
//     */
//
// std::vector<std::string> suggest(const std::string& word);
// Suggest spelling corrections word
func (h *Hunspell) Suggest(word string) ([]string, error) {

	return []string{}, nil
}

// /* Suggest words from suffix rules
//   - suffix_suggest(suggestions, root_word)
//   - input: pointer to an array of strings pointer and the  word
//   - array of strings pointer (here *slst) may not be initialized
//   - output: number of suggestions in string array, and suggestions in
//   - a newly allocated array of strings (*slts will be NULL when number
//   - of suggestion equals 0.)
//     */
//     std::vector<std::string> suffix_suggest(const std::string& root_word);
//
// SuggestSuffix
func (h *Hunspell) SuggestSuffix(word string) ([]string, error) {

	return []string{}, nil
}

// LIBHUNSPELL_DLL_EXPORTED char* Hunspell_get_dic_encoding(Hunhandle* pHunspell);
// DictionaryEncoding
func (h *Hunspell) DictionaryEncoding() (Encoding, error) {

	return UTF8, nil
}

// Morphological functions

// /* analyze(result, word) - morphological analysis of the word */
//
//	std::vector<std::string> analyze(const std::string& word);
//
// Analyze
func (h *Hunspell) Analyze(slst []string, word string) (int, error) {

	return 0, nil
}

//	/* stem(word) - stemmer function */
//	std::vector<std::string> stem(const std::string& word);
//
// Stem
func (h *Hunspell) Stem(slst []string, word string) (int, error) {

	return 0, nil
}

// /* stem(analysis, n) - get stems from a morph. analysis
//   - example:
//   - char ** result, result2;
//   - int n1 = analyze(&result, "words");
//   - int n2 = stem(&result2, result, n1);
//     */
//     std::vector<std::string> stem(const std::vector<std::string>& morph);
//
// StemFromMorph
func (h *Hunspell) StemFromMorph(morph []string) ([]string, error) {

	return []string{}, nil
}

// /* generate(result, word, word2) - morphological generation by example(s) */
//
//	std::vector<std::string> generate(const std::string& word, const std::string& word2);
//
// Generate
func (h *Hunspell) Generate(word, word2 string) ([]string, error) {

	return []string{}, nil
}

// /* generate(result, word, desc, n) - generation by morph. description(s)
//   - example:
//   - char ** result;
//   - char * affix = "is:plural"; // description depends from dictionaries, too
//   - int n = generate(&result, "word", &affix, 1);
//   - for (int i = 0; i < n; i++) printf("%s\n", result[i]);
//     */
//     std::vector<std::string> generate(const std::string& word, const std::vector<std::string>& pl);
//
// GenerateFromAffix
func (h *Hunspell) GenerateFromAffix(word string, affixes []string) ([]string, error) {
	return []string{}, nil
}

// Run-time dictionary modification methods

// int add(const std::string& word);
//
//	int add_with_flags(const std::string& word, const std::string& flags, const std::string& desc = NULL);
//
// AddWord to the run-time dictionary
func (h *Hunspell) AddWord(word string) error {
	return nil
}

// LIBHUNSPELL_DLL_EXPORTED int Hunspell_add_with_flags(Hunhandle* pHunspell,
// const char* word, const char* flags, const char* desc);
// AddWordWithFlags
func (h *Hunspell) AddWordWithFlags(word string, flags []string, desc string) error {
	return nil
}

// /* add word to the run-time dictionary with affix flags of
// * the example (a dictionary word): Hunspell will recognize
// * affixed forms of the new word, too.
// */
// LIBHUNSPELL_DLL_EXPORTED int Hunspell_add_with_affix(Hunhandle* pHunspell,
// const char* word,
// const char* example);
func (h *Hunspell) AddWordWithAffix(word string, examples string) error {
	return nil
}

// /* remove word from the run-time dictionary */
// LIBHUNSPELL_DLL_EXPORTED int Hunspell_remove(Hunhandle* pHunspell,
// const char* word);
// RemoveWord
func (h *Hunspell) RemoveWord(word string) error {
	return nil
}

// /* free suggestion lists */
// LIBHUNSPELL_DLL_EXPORTED void Hunspell_free_list(Hunhandle* pHunspell,
// char*** slst,
// int n);
// FreeSuggestionLists
func (h *Hunspell) FreeSuggestionLists(slst []string, n int) {

}

//	/* get extra word characters definied in affix file for tokenization */
//	const char* get_wordchars() const;
//	const std::string& get_wordchars_cpp() const;
//	const std::vector<w_char>& get_wordchars_utf16() const;
//
// GetWordChars
func (h *Hunspell) GetWordChars() (string, error) {
	return "", nil
}
