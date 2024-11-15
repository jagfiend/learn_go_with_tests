package hello

const englishPrefix = "Hello"
const spanishPrefix = "Hola"
const frenchPrefix = "Bonjour"

func Hello(name string, lang LangEnum) string {
	if name == "" {
		name = "world"
	}

	var prefix string

	switch lang {
	case English:
		prefix = englishPrefix
	case Spanish:
		prefix = spanishPrefix
	case French:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}

	return prefix + " " + name
}
