package text

import (
	"strings"
	"unicode"
)

var modelAcronyms = []string{
	"ai",
	"gpt",
	"oss",
	"tars",
	"kat",
	"glm",
}

var providers = map[string]string{
	"z.ai":         "Z.AI",
	"openai":       "OpenAI",
	"google":       "Google",
	"anthropic":    "Anthropic",
	"microsoft":    "Microsoft",
	"bytedance":    "ByteDance",
	"liquidai":     "LiquidAI",
	"xiaomi":       "Xiaomi",
	"cohere":       "Cohere",
	"qwen":         "Qwen",
	"minimax":      "MiniMax",
	"aionlabs":     "AionLabs",
	"nousresearch": "NousResearch",
	"wizardlm":     "WizardLM",
	"kwaipilot":    "Kwaipilot",
	"xai":          "xAI",
	"eleutherai":   "EleutherAI",
	"essentialai":  "EssentialAI",
	"sao10k":       "Sao10k",
	"undi95":       "Undi95",
}

var weirdProviders = map[string]string{
	"x-ai/":      "xAI: ",
	"z.ai:":      "Z.AI: ",
	"z-ai/":      "Z.AI: ",
	"arcee-ai":   "ArceeAI",
	"stepfun-ai": "StepFun",
	"aion-labs":  "AionLabs",
}

// FormatModelTitle turns a raw model identifier (e.g. "openai/gpt-4o-mini")
// into a human-readable title, applying provider-specific display names and
// title-casing the remaining segments.
func FormatModelTitle(model string) string {
	for p := range weirdProviders {
		if strings.HasPrefix(model, p) {
			model = strings.ReplaceAll(model, p, weirdProviders[p])
		}
	}

	model = strings.ReplaceAll(model, "/", ": ")
	model = strings.ReplaceAll(model, "_", " ")
	model = strings.ReplaceAll(model, ":thinking", " Thinking")

	wrap := []string{":free", ":exacto", ":extended", ":latest"}
	foundSuffix := false
	for _, w := range wrap {
		if strings.HasSuffix(model, w) {
			// remove ":" and wrap in ()
			model = model[:len(model)-len(w)] + " (" + strings.TrimPrefix(w, ":") + ")"
			foundSuffix = true
		}
	}

	// ollama model handling, they contain <model>:<size>
	if !foundSuffix {
		index := strings.Index(model, ":")
		if index > -1 && model[index+1] != ' ' {
			runes := []rune(model)
			runes[index] = ' '
			model = string(runes)
		}
	}

	// ollama
	if strings.HasSuffix(model, "-cloud") {
		model = model[:len(model)-6] + " Cloud"
	}

	model = replaceModelsWithNumbersAndDates(model)
	model = splitPart(model)

	parts := strings.Fields(model)
	for i, part := range parts {
		parts[i] = normalizePart(part, i == 0)
	}

	final := strings.Join(parts, " ")
	// replace multiple spaces into one

	final = strings.ReplaceAll(final, "  ", " ")

	for _, quant := range quantization {
		for find, replace := range quant {
			if strings.HasSuffix(final, " "+find) {
				final = final[:len(final)-len(find)-1] + " " + replace
				break
			}
		}
	}

	return final
}

// order matters
var quantization = []map[string]string{
	{"Q4 K M": "Q4_K_M"},
	{"Q8 K M": "Q8_K_M"},
	{"K L": "K_L"},
	{"K M": "K_M"},
	{"K S": "K_S"},
	{"Q1 K": "Q1_K"},
	{"Q2 K": "Q2_K"},
	{"Q4 K": "Q4_K"},
	{"Q8 K": "Q8_K"},
	{"Q1 0": "Q1_0"},
	{"Q2 0": "Q2_0"},
	{"Q4 0": "Q4_0"},
	{"Q8 0": "Q8_0"},
}

// check if string ends with -yyyymmdd
func hasDateWithoutDashesSuffix(s string) bool {
	if len(s) < 9 {
		return false
	}

	suffix := s[len(s)-9:]
	if len(suffix) != 9 {
		return false
	}

	// first character must be a hyphen
	if suffix[0] != '-' {
		return false
	}

	// remaining 8 characters must all be digits
	for i := 1; i < 9; i++ {
		if !unicode.IsDigit(rune(suffix[i])) {
			return false
		}
	}

	return true
}

func hasDateSuffix(s string) bool {
	if len(s) < 11 {
		return false
	}

	suffix := s[len(s)-11:]

	// check format: -yyyy-mm-dd where d is digit
	if len(suffix) != 11 {
		return false
	}

	// first character must be a hyphen
	if suffix[0] != '-' {
		return false
	}

	for i := 1; i < 11; i++ {
		if i == 5 || i == 8 {
			if suffix[i] != '-' {
				return false
			}
		} else {
			if !unicode.IsDigit(rune(suffix[i])) {
				return false
			}
		}
	}

	return true
}

// func replaceModelsWithSizeSuffix(model string) string {
//
// }

func replaceModelsWithNumbersAndDates(model string) string {
	parts := strings.Fields(model)

	// fmt.Println("---- FormatModelTitle", parts, "----")

	for i, part := range parts {
		if hasDateSuffix(part) {
			parts[i] = part[:len(part)-11] + " " + part[len(part)-10:]
		} else if hasDateWithoutDashesSuffix(part) {
			parts[i] = part[:len(part)-9] + " " + part[len(part)-8:]
		}
	}

	// fmt.Println("---- FormatModelTitle", parts, "----")
	model = strings.Join(parts, " ")

	return model
}

func normalizePart(part string, isFirst bool) string {
	// fmt.Println("---- normalizePart", part, "----")
	// replace dashes with spaces, except between digits
	// part = splitPart(part)

	// split by spaces again and process each word
	words := strings.Fields(part)
	for i, word := range words {
		words[i] = normalizeWord(word, isFirst && i == 0)
	}

	// fmt.Println("---- /normalizePart", "----")

	return strings.Join(words, " ")
}

func splitPart(s string) string {
	// fmt.Println("---- splitPart", s, "----")
	foundChar := false
	runes := []rune(s)
	parenDepth := 0

	for i := 0; i < len(runes); i++ {
		switch runes[i] {
		case '(':
			parenDepth++
		case ')':
			parenDepth--
		}

		if unicode.IsLetter(rune(s[i])) {
			foundChar = true
		}
		switch runes[i] {
		case '-':
			if foundChar && parenDepth == 0 {
				runes[i] = ' '
			}
			foundChar = false
		case ' ':
			foundChar = false
		}
	}

	// fmt.Println("---- !splitPart", string(runes), "----")

	// Reverse pass
	foundChar = false
	parenDepth = 0

	for i := len(runes) - 1; i >= 0; i-- {
		switch runes[i] {
		case ')':
			parenDepth++
		case '(':
			parenDepth--
		}

		if unicode.IsLetter(runes[i]) {
			foundChar = true
		} else if runes[i] == '-' {
			if foundChar && parenDepth == 0 {
				runes[i] = ' '
			}
			foundChar = false
		} else if runes[i] == ' ' || runes[i] == '(' || runes[i] == ')' {
			foundChar = false
		} else if !unicode.IsDigit(runes[i]) {
			foundChar = false
		}
	}

	// fmt.Println("---- /splitPart", string(runes), "----")

	return string(runes)
}

func normalizeWord(word string, isFirstWord bool) string {
	// fmt.Println("---- normalizeWord", word, "----")
	// defer // fmt.Println("---- /normalizeWord", "----")

	lower := strings.ToLower(word)

	// check if it's a provider name (only for first word)
	if isFirstWord {
		// remove trailing colon for provider check
		checkWord := strings.TrimSuffix(lower, ":")
		if providerName, ok := providers[checkWord]; ok {
			if strings.HasSuffix(lower, ":") {
				return providerName + ":"
			}
			return providerName
		}
	}

	// check if it's an acronym
	for _, acronym := range modelAcronyms {
		if lower == acronym {
			return strings.ToUpper(word)
		}
	}

	if strings.HasSuffix(word, "it") {
		return "Instruct"
	}

	// check if it's "v<digit>" pattern
	if len(word) >= 2 && (word[0] == 'v' || word[0] == 'V') && unicode.IsDigit(rune(word[1])) {
		return "v" + word[1:]
	}

	// check if it's "<any_number_of_digits>b" pattern, uppercase b
	if isBothDigitsAndLetters(word) {
		if ok, m := isSpecificModel(word); ok {
			return m
		}
		if ok, m := isMixtureOfExperts(word); ok {
			return m
		}
		return strings.ToUpper(word)
	}
	// if word[len(word)-1] == 'b' {
	// return word[:len(word)-1] + "B"
	// }

	// dates in brackets or without "(xxxx-xx-xx"), "(xxxx-xx)", "xxxx-xx-xx", "xxxx-xx"
	// should be kept or wrapped in brackets
	if matched, flip := isDatePattern(word); matched {
		if strings.HasPrefix(word, "(") && strings.HasSuffix(word, ")") {
			word = strings.TrimPrefix(strings.TrimSuffix(word, ")"), "(")
			return "(" + flipDate(word, flip) + ")"
		}
		return "(" + flipDate(word, flip) + ")"
	}

	// title case: capitalize first letter
	if len(word) > 0 {
		if ok, m := isSpecificModel(word); ok {
			return m
		}

		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])
		return string(runes)
	}

	return word
}

// returns true and size of an expert model with lowercase x
func isMixtureOfExperts(s string) (bool, string) {
	// 8x22B or 4X120b, 15x8B => 8x22B, 4X120B, 15x8B
	runes := []rune(s)
	numberOfExperts := 0
	count := len(runes)

	// Count leading digits
	for i := range count {
		if !unicode.IsDigit(runes[i]) {
			break
		}
		numberOfExperts++
	}

	// Check bounds and for 'x' or 'X'
	if numberOfExperts == 0 || numberOfExperts >= count {
		return false, ""
	}

	if unicode.ToLower(runes[numberOfExperts]) != 'x' {
		return false, ""
	}

	// Extract expert size digits
	startIdx := numberOfExperts + 1
	endIdx := startIdx
	for endIdx < count && unicode.IsDigit(runes[endIdx]) {
		endIdx++
	}
	expertSize := string(runes[startIdx:endIdx])

	// Must have digits after 'x'
	if len(expertSize) == 0 {
		return false, ""
	}

	// Build result with lowercase 'x' and uppercase suffix
	result := string(runes[:numberOfExperts]) + "x" + expertSize
	suffixStart := startIdx + len(expertSize)
	if suffixStart < count {
		result += strings.ToUpper(string(runes[suffixStart:]))
	}

	return true, result
}

func isBothDigitsAndLetters(s string) bool {
	foundDigit := false
	foundLetter := false
	for _, r := range s {
		// allow decimals in names
		if r == '.' {
			continue
		}
		if unicode.IsDigit(r) {
			foundDigit = true
		} else if unicode.IsLetter(r) {
			foundLetter = true
		}
	}
	return foundDigit && foundLetter
}

// [A-Z][A-Z][A-Z]+
// [A-Z]{3,}
var modelsWithNumberSuffix = map[string]string{
	"qwen":     "Qwen",
	"lfm":      "LFM",
	"o":        "o",
	"internvl": "InternVL",
	"step":     "Step",
	"llama":    "Llama",
	"medllama": "Medllama",
	// "gemma3n":  "Gemma 3n",
	"gemma":   "Gemma",
	"phi":     "Phi",
	"dolphin": "Dolphin",
	// OLMO
	"olmo":      "Olmo",
	"smollm":    "SmolLM",
	"small":     "Small",
	"starcoder": "StarCoder",
	"falcon":    "Falcon",
	"granite":   "Granite",
	"wizardlm":  "WizardLM",
	"hermes":    "Hermes",
	"codegeex":  "CodeGeeX",
	"stablelm":  "StableLM",
	"exaone":    "ExaOne",
	"internlm":  "InternLM",
	"tulu":      "Tulu",
	"guard":     "Guard",
	"orca":      "Orca",
	"platypus":  "Platypus",
	"sailor":    "Sailor",
	// "qwen2.5vl": "QWEN2.5 VL",
}

var modelsWithNumberPrefix = map[string]string{
	"n": "n",
	"o": "o",
}

var modelsUppercase = []string{
	"lfm",
	"ai",
}

var modelsExactCase = []string{
	"OpenGVLab",
	"SorcererLM",
	"WizardLM",
	"ChatGPT",
	"Gemma3n",
	"SmolLM",
}

func isSpecificModel(s string) (bool, string) {
	// fmt.Println("isSpecificModel", s)
	lower := strings.ToLower(s)
	for k, v := range modelsWithNumberSuffix {
		if len(s) == len(k)+1 {
			if lower[:len(k)] == k && unicode.IsDigit(rune(lower[len(k)])) {
				return true, v + string(lower[len(k)])
			}
		}
		// support decimal suffixes 1.5, 3.5, etc
		if len(s) == len(k)+3 {
			if lower[:len(k)] == k && unicode.IsDigit(rune(lower[len(k)])) && lower[len(k)+1] == '.' && unicode.IsDigit(rune(lower[len(k)+2])) {
				return true, v + lower[len(k):]
			}
		}
	}
	for k, v := range modelsWithNumberPrefix {
		if len(s) == len(k)+1 {
			if unicode.IsDigit(rune(lower[0])) && lower[1:len(k)+1] == k {
				return true, string(lower[0]) + v
			}
		}
	}
	for _, v := range modelsUppercase {
		if lower == v {
			return true, strings.ToUpper(v)
		}
	}

	for _, v := range modelsExactCase {
		if strings.EqualFold(lower, v) {
			return true, v
		}
	}

	return false, ""
}

func flipDate(date string, flip bool) string {
	parts := strings.Split(date, "-")
	if len(parts) == 1 && !flip && len(date) == 8 {
		// 20241022 => 2024-10-22
		return date[0:4] + "-" + date[4:6] + "-" + date[6:8]
	}
	if flip {
		return parts[1] + "-" + parts[0]
	}
	return strings.Join(parts, "-")
}

func isDatePattern(word string) (bool, bool) {
	stripped := strings.TrimPrefix(strings.TrimSuffix(word, ")"), "(")

	// check if it contains only digits and dashes
	hasDigit := false
	hasDash := false

	for _, r := range stripped {
		if unicode.IsDigit(r) {
			hasDigit = true
		} else if r == '-' {
			hasDash = true
		} else {
			return false, false
		}
	}

	if hasDigit && !hasDash {
		if len(stripped) == 8 {
			return true, false
		}
	}

	// must have both digits and dashes, and match date-like patterns
	if !hasDigit || !hasDash {
		return false, false
	}

	// check for date-like patterns: xxxx-xx-xx or xxxx-xx
	parts := strings.Split(stripped, "-")
	if len(parts) == 2 || len(parts) == 3 {
		// First part should be 4 digits (year)
		if len(parts[0]) == 4 && len(parts[1]) == 2 {
			return true, false
		}
		// flipped
		if len(parts) == 2 && len(parts[0]) == 2 && len(parts[1]) == 4 {
			return true, true
		}
	}

	return false, false
}
