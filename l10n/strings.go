package l10n

// Strings returns a translation set that will take any term and return its
// translation.
func Strings(lang string) map[string]string {
	switch lang {
	case "ar":
		return phrasesAR
	case "cs":
		return phrasesCS
	case "da":
		return phrasesDA
	case "de":
		return phrasesDE
	case "el":
		return phrasesEL
	case "eo":
		return phrasesEO
	case "es":
		return phrasesES
	case "eu":
		return phrasesEU
	case "fa":
		return phrasesFA
	case "fr":
		return phrasesFR
	case "gl":
		return phrasesGL
	case "he":
		return phrasesHE
	case "hu":
		return phrasesHU
	case "it":
		return phrasesIT
	case "ja":
		return phrasesJA
	case "ko":
		return phrasesKO
	case "lt":
		return phrasesLT
	case "mk":
		return phrasesMK
	case "nl":
		return phrasesNL
	case "pl":
		return phrasesPL
	case "pt":
		return phrasesPT
	case "ro":
		return phrasesRO
	case "ru":
		return phrasesRU
	case "sk":
		return phrasesSK
	case "sv":
		return phrasesSV
	case "tg":
		return phrasesTG
	case "tr":
		return phrasesTR
	case "zh":
		return phrasesZH
	default:
		return phrases
	}
}
