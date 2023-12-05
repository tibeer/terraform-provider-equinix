package main

import (
	"strings"
    "unicode"
)

// ToCamelCase converts a string into CamelCase.
func ToCamelCase(str string) string {
    // Split the string into words using a space as the delimiter.
    words := strings.Fields(str)

    // Convert each word to title case.
    for i, word := range words {
        words[i] = strings.Title(word)
    }

    // Join the words back into a single string.
    return strings.Join(words, "")
}

// ToCamelCaseAlternative handles more complex string cases, including
// underscores, hyphens, and mixed case input.
func ToCamelCaseAlternative(str string) string {
    // Function to check if a rune is a delimiter.
    isDelimiter := func(r rune) bool {
        return r == '_' || r == '-' || unicode.IsSpace(r)
    }

    // Builder to construct the new string.
    var builder strings.Builder
    upperNext := true

    for _, r := range str {
        if isDelimiter(r) {
            upperNext = true
            continue
        }
        if upperNext {
            builder.WriteRune(unicode.ToUpper(r))
            upperNext = false
        } else {
            builder.WriteRune(unicode.ToLower(r))
        }
    }

    return builder.String()
}

// ToSnakeCase converts a string to snake_case.
func ToSnakeCase(str string) string {
    var result []rune
    for i, r := range str {
        if i > 0 && unicode.IsUpper(r) && (len(result) > 0 && result[len(result)-1] != '_') {
            result = append(result, '_')
        }
        result = append(result, unicode.ToLower(r))
    }
    return string(result)
}

// ToPascalCase converts a string to Pascal Case (first letter of each word is uppercase).
func ToPascalCase(str string) string {
    // Function to check if a rune is a delimiter.
    isDelimiter := func(r rune) bool {
        return r == '_' || r == '-' || unicode.IsSpace(r)
    }

    // Builder to construct the new string.
    var builder strings.Builder
    upperNext := true

    for _, r := range str {
        if isDelimiter(r) {
            upperNext = true
            continue
        }
        if upperNext {
            builder.WriteRune(unicode.ToUpper(r))
            upperNext = false
        } else {
            builder.WriteRune(unicode.ToLower(r))
        }
    }

    return builder.String()
}

func SDKTypeName(sdkPath string) string {
    // Split the SDK path and return the last element as the type name
    // e.g., "github.com/packethost/packngo" -> "packngo"
    parts := strings.Split(sdkPath, "/")
    return parts[len(parts)-1]
}
