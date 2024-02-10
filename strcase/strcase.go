// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on https://github.com/ettle/strcase
// Copyright (c) 2020 Liyan David Chang under the MIT License

// Package strcase provides functions for manipulating the case of strings
// (CamelCase, kebab-case, snake_case, Sentence case, etc). It is based on
// https://github.com/ettle/strcase, which is Copyright (c) 2020 Liyan David
// Chang under the MIT License. Its principle difference from other strcase
// packages is that it preserves acronyms in input text for CamelCase and
// lowerCamelCase.
package strcase

// ToSnake returns words in snake_case (lower case words with underscores).
func ToSnake(s string) string {
	return To(s, LowerCase, '_')
}

// ToSNAKE returns words in SNAKE_CASE (upper case words with underscores).
// Also known as SCREAMING_SNAKE_CASE or UPPER_CASE.
func ToSNAKE(s string) string {
	return To(s, UpperCase, '_')
}

// ToKebab returns words in kebab-case (lower case words with dashes).
// Also known as dash-case.
func ToKebab(s string) string {
	return To(s, LowerCase, '-')
}

// ToKEBAB returns words in KEBAB-CASE (upper case words with dashes).
// Also known as SCREAMING-KEBAB-CASE or SCREAMING-DASH-CASE.
func ToKEBAB(s string) string {
	return To(s, UpperCase, '-')
}

// ToCamel returns words in CamelCase (capitalized words concatenated together).
// Also known as UpperCamelCase.
func ToCamel(s string) string {
	return To(s, TitleCase, 0)
}

// ToLowerCamel returns words in lowerCamelCase (capitalized words concatenated together,
// with first word lower case). Also known as camelCase or mixedCase.
func ToLowerCamel(s string) string {
	return To(s, CamelCase, 0)
}

// ToTitle returns words in Title Case (capitalized words with spaces).
func ToTitle(s string) string {
	return To(s, TitleCase, ' ')
}
