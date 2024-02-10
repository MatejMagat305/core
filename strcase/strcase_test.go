// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on https://github.com/ettle/strcase
// Copyright (c) 2020 Liyan David Chang under the MIT License

package strcase

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestOrignal(t *testing.T) {
	assertEqual(t, "nativeOrgURL", To("NativeOrgURL", CamelCase, 0))
	assertEqual(t, "nativeOrgUrl", To("NativeOrgUrl", CamelCase, 0))

	assertEqual(t, "native-org-url", To("NativeOrgURL", LowerCase, '-'))
	assertEqual(t, "json-string", To("JSONString", LowerCase, '-'))

	assertEqual(t, "jsonString", To("JSONString", CamelCase, 0))
	assertEqual(t, "JSONString", To("JSONString", TitleCase, 0))
	assertEqual(t, "JsonString", To("JsonString", TitleCase, 0))

	assertEqual(t, "nasa-rocket", To("NASARocket", LowerCase, '-'))
	assertEqual(t, "nasa-Rocket", To("NASARocket", CamelCase, '-'))
	assertEqual(t, "NASA-Rocket", To("NASARocket", TitleCase, '-'))

	assertEqual(t, "Ps4", To("ps4", TitleCase, '-'))
	assertEqual(t, "PS4", To("PS4", TitleCase, '-'))

	// Not great if you're coming from an all-caps case
	assertEqual(t, "SCREAMINGCASE", To("SCREAMING_CASE", TitleCase, 0))
}

func TestAll(t *testing.T) {
	// Instead of testing, we can generate the outputs to make it easier to
	// add more test cases or functions
	generate := false

	type data struct {
		input string
		snake string
		SNAKE string
		kebab string
		KEBAB string
		Camel string
		camel string
		title string
	}
	for _, test := range []data{
		{
			input: "Hello world!",
			snake: "hello_world!",
			SNAKE: "HELLO_WORLD!",
			kebab: "hello-world!",
			KEBAB: "HELLO-WORLD!",
			Camel: "HelloWorld!",
			camel: "helloWorld!",
			title: "Hello World!",
		},
		{
			input: "",
			snake: "",
			SNAKE: "",
			kebab: "",
			KEBAB: "",
			Camel: "",
			camel: "",
			title: "",
		},
		{
			input: ".",
			snake: "",
			SNAKE: "",
			kebab: "",
			KEBAB: "",
			Camel: "",
			camel: "",
			title: "",
		},
		{
			input: "A",
			snake: "a",
			SNAKE: "A",
			kebab: "a",
			KEBAB: "A",
			Camel: "A",
			camel: "a",
			title: "A",
		},
		{
			input: "a",
			snake: "a",
			SNAKE: "A",
			kebab: "a",
			KEBAB: "A",
			Camel: "A",
			camel: "a",
			title: "A",
		},
		{
			input: "foo",
			snake: "foo",
			SNAKE: "FOO",
			kebab: "foo",
			KEBAB: "FOO",
			Camel: "Foo",
			camel: "foo",
			title: "Foo",
		},
		{
			input: "snake_case",
			snake: "snake_case",
			SNAKE: "SNAKE_CASE",
			kebab: "snake-case",
			KEBAB: "SNAKE-CASE",
			Camel: "SnakeCase",
			camel: "snakeCase",
			title: "Snake Case",
		},
		{
			input: "SNAKE_CASE",
			snake: "snake_case",
			SNAKE: "SNAKE_CASE",
			kebab: "snake-case",
			KEBAB: "SNAKE-CASE",
			Camel: "SNAKECASE",
			camel: "snakeCASE",
			title: "SNAKE CASE",
		},
		{
			input: "kebab-case",
			snake: "kebab_case",
			SNAKE: "KEBAB_CASE",
			kebab: "kebab-case",
			KEBAB: "KEBAB-CASE",
			Camel: "KebabCase",
			camel: "kebabCase",
			title: "Kebab Case",
		},
		{
			input: "PascalCase",
			snake: "pascal_case",
			SNAKE: "PASCAL_CASE",
			kebab: "pascal-case",
			KEBAB: "PASCAL-CASE",
			Camel: "PascalCase",
			camel: "pascalCase",
			title: "Pascal Case",
		},
		{
			input: "camelCase",
			snake: "camel_case",
			SNAKE: "CAMEL_CASE",
			kebab: "camel-case",
			KEBAB: "CAMEL-CASE",
			Camel: "CamelCase",
			camel: "camelCase",
			title: "Camel Case",
		},
		{
			input: "Title Case",
			snake: "title_case",
			SNAKE: "TITLE_CASE",
			kebab: "title-case",
			KEBAB: "TITLE-CASE",
			Camel: "TitleCase",
			camel: "titleCase",
			title: "Title Case",
		},
		{
			input: "point.case",
			snake: "point_case",
			SNAKE: "POINT_CASE",
			kebab: "point-case",
			KEBAB: "POINT-CASE",
			Camel: "PointCase",
			camel: "pointCase",
			title: "Point Case",
		},
		{
			input: "snake_case_with_more_words",
			snake: "snake_case_with_more_words",
			SNAKE: "SNAKE_CASE_WITH_MORE_WORDS",
			kebab: "snake-case-with-more-words",
			KEBAB: "SNAKE-CASE-WITH-MORE-WORDS",
			Camel: "SnakeCaseWithMoreWords",
			camel: "snakeCaseWithMoreWords",
			title: "Snake Case With More Words",
		},
		{
			input: "SNAKE_CASE_WITH_MORE_WORDS",
			snake: "snake_case_with_more_words",
			SNAKE: "SNAKE_CASE_WITH_MORE_WORDS",
			kebab: "snake-case-with-more-words",
			KEBAB: "SNAKE-CASE-WITH-MORE-WORDS",
			Camel: "SNAKECASEWITHMOREWORDS",
			camel: "snakeCASEWITHMOREWORDS",
			title: "SNAKE CASE WITH MORE WORDS",
		},
		{
			input: "kebab-case-with-more-words",
			snake: "kebab_case_with_more_words",
			SNAKE: "KEBAB_CASE_WITH_MORE_WORDS",
			kebab: "kebab-case-with-more-words",
			KEBAB: "KEBAB-CASE-WITH-MORE-WORDS",
			Camel: "KebabCaseWithMoreWords",
			camel: "kebabCaseWithMoreWords",
			title: "Kebab Case With More Words",
		},
		{
			input: "PascalCaseWithMoreWords",
			snake: "pascal_case_with_more_words",
			SNAKE: "PASCAL_CASE_WITH_MORE_WORDS",
			kebab: "pascal-case-with-more-words",
			KEBAB: "PASCAL-CASE-WITH-MORE-WORDS",
			Camel: "PascalCaseWithMoreWords",
			camel: "pascalCaseWithMoreWords",
			title: "Pascal Case With More Words",
		},
		{
			input: "camelCaseWithMoreWords",
			snake: "camel_case_with_more_words",
			SNAKE: "CAMEL_CASE_WITH_MORE_WORDS",
			kebab: "camel-case-with-more-words",
			KEBAB: "CAMEL-CASE-WITH-MORE-WORDS",
			Camel: "CamelCaseWithMoreWords",
			camel: "camelCaseWithMoreWords",
			title: "Camel Case With More Words",
		},
		{
			input: "Title Case With More Words",
			snake: "title_case_with_more_words",
			SNAKE: "TITLE_CASE_WITH_MORE_WORDS",
			kebab: "title-case-with-more-words",
			KEBAB: "TITLE-CASE-WITH-MORE-WORDS",
			Camel: "TitleCaseWithMoreWords",
			camel: "titleCaseWithMoreWords",
			title: "Title Case With More Words",
		},
		{
			input: "point.case.with.more.words",
			snake: "point_case_with_more_words",
			SNAKE: "POINT_CASE_WITH_MORE_WORDS",
			kebab: "point-case-with-more-words",
			KEBAB: "POINT-CASE-WITH-MORE-WORDS",
			Camel: "PointCaseWithMoreWords",
			camel: "pointCaseWithMoreWords",
			title: "Point Case With More Words",
		},
		{
			input: "snake_case__with___multiple____delimiters",
			snake: "snake_case_with_multiple_delimiters",
			SNAKE: "SNAKE_CASE_WITH_MULTIPLE_DELIMITERS",
			kebab: "snake-case-with-multiple-delimiters",
			KEBAB: "SNAKE-CASE-WITH-MULTIPLE-DELIMITERS",
			Camel: "SnakeCaseWithMultipleDelimiters",
			camel: "snakeCaseWithMultipleDelimiters",
			title: "Snake Case With Multiple Delimiters",
		},
		{
			input: "SNAKE_CASE__WITH___multiple____DELIMITERS",
			snake: "snake_case_with_multiple_delimiters",
			SNAKE: "SNAKE_CASE_WITH_MULTIPLE_DELIMITERS",
			kebab: "snake-case-with-multiple-delimiters",
			KEBAB: "SNAKE-CASE-WITH-MULTIPLE-DELIMITERS",
			Camel: "SNAKECASEWITHMultipleDELIMITERS",
			camel: "snakeCASEWITHMultipleDELIMITERS",
			title: "SNAKE CASE WITH Multiple DELIMITERS",
		},
		{
			input: "kebab-case--with---multiple----delimiters",
			snake: "kebab_case_with_multiple_delimiters",
			SNAKE: "KEBAB_CASE_WITH_MULTIPLE_DELIMITERS",
			kebab: "kebab-case-with-multiple-delimiters",
			KEBAB: "KEBAB-CASE-WITH-MULTIPLE-DELIMITERS",
			Camel: "KebabCaseWithMultipleDelimiters",
			camel: "kebabCaseWithMultipleDelimiters",
			title: "Kebab Case With Multiple Delimiters",
		},
		{
			input: "Title Case  With   Multiple    Delimiters",
			snake: "title_case_with_multiple_delimiters",
			SNAKE: "TITLE_CASE_WITH_MULTIPLE_DELIMITERS",
			kebab: "title-case-with-multiple-delimiters",
			KEBAB: "TITLE-CASE-WITH-MULTIPLE-DELIMITERS",
			Camel: "TitleCaseWithMultipleDelimiters",
			camel: "titleCaseWithMultipleDelimiters",
			title: "Title Case With Multiple Delimiters",
		},
		{
			input: "point.case..with...multiple....delimiters",
			snake: "point_case_with_multiple_delimiters",
			SNAKE: "POINT_CASE_WITH_MULTIPLE_DELIMITERS",
			kebab: "point-case-with-multiple-delimiters",
			KEBAB: "POINT-CASE-WITH-MULTIPLE-DELIMITERS",
			Camel: "PointCaseWithMultipleDelimiters",
			camel: "pointCaseWithMultipleDelimiters",
			title: "Point Case With Multiple Delimiters",
		},
		{
			input: " leading space",
			snake: "leading_space",
			SNAKE: "LEADING_SPACE",
			kebab: "leading-space",
			KEBAB: "LEADING-SPACE",
			Camel: "LeadingSpace",
			camel: "leadingSpace",
			title: "Leading Space",
		},
		{
			input: "   leading spaces",
			snake: "leading_spaces",
			SNAKE: "LEADING_SPACES",
			kebab: "leading-spaces",
			KEBAB: "LEADING-SPACES",
			Camel: "LeadingSpaces",
			camel: "leadingSpaces",
			title: "Leading Spaces",
		},
		{
			input: "\t\t\r\n leading whitespaces",
			snake: "leading_whitespaces",
			SNAKE: "LEADING_WHITESPACES",
			kebab: "leading-whitespaces",
			KEBAB: "LEADING-WHITESPACES",
			Camel: "LeadingWhitespaces",
			camel: "leadingWhitespaces",
			title: "Leading Whitespaces",
		},
		{
			input: "trailing space ",
			snake: "trailing_space",
			SNAKE: "TRAILING_SPACE",
			kebab: "trailing-space",
			KEBAB: "TRAILING-SPACE",
			Camel: "TrailingSpace",
			camel: "trailingSpace",
			title: "Trailing Space",
		},
		{
			input: "trailing spaces   ",
			snake: "trailing_spaces",
			SNAKE: "TRAILING_SPACES",
			kebab: "trailing-spaces",
			KEBAB: "TRAILING-SPACES",
			Camel: "TrailingSpaces",
			camel: "trailingSpaces",
			title: "Trailing Spaces",
		},
		{
			input: "trailing whitespaces\t\t\r\n",
			snake: "trailing_whitespaces",
			SNAKE: "TRAILING_WHITESPACES",
			kebab: "trailing-whitespaces",
			KEBAB: "TRAILING-WHITESPACES",
			Camel: "TrailingWhitespaces",
			camel: "trailingWhitespaces",
			title: "Trailing Whitespaces",
		},
		{
			input: " on both sides ",
			snake: "on_both_sides",
			SNAKE: "ON_BOTH_SIDES",
			kebab: "on-both-sides",
			KEBAB: "ON-BOTH-SIDES",
			Camel: "OnBothSides",
			camel: "onBothSides",
			title: "On Both Sides",
		},
		{
			input: "    many on both sides  ",
			snake: "many_on_both_sides",
			SNAKE: "MANY_ON_BOTH_SIDES",
			kebab: "many-on-both-sides",
			KEBAB: "MANY-ON-BOTH-SIDES",
			Camel: "ManyOnBothSides",
			camel: "manyOnBothSides",
			title: "Many On Both Sides",
		},
		{
			input: "\rwhitespaces on both sides\t\t\r\n",
			snake: "whitespaces_on_both_sides",
			SNAKE: "WHITESPACES_ON_BOTH_SIDES",
			kebab: "whitespaces-on-both-sides",
			KEBAB: "WHITESPACES-ON-BOTH-SIDES",
			Camel: "WhitespacesOnBothSides",
			camel: "whitespacesOnBothSides",
			title: "Whitespaces On Both Sides",
		},
		{
			input: "  extraSpaces in_This TestCase Of MIXED_CASES\t",
			snake: "extra_spaces_in_this_test_case_of_mixed_cases",
			SNAKE: "EXTRA_SPACES_IN_THIS_TEST_CASE_OF_MIXED_CASES",
			kebab: "extra-spaces-in-this-test-case-of-mixed-cases",
			KEBAB: "EXTRA-SPACES-IN-THIS-TEST-CASE-OF-MIXED-CASES",
			Camel: "ExtraSpacesInThisTestCaseOfMIXEDCASES",
			camel: "extraSpacesInThisTestCaseOfMIXEDCASES",
			title: "Extra Spaces In This Test Case Of MIXED CASES",
		},
		{
			input: "CASEBreak",
			snake: "case_break",
			SNAKE: "CASE_BREAK",
			kebab: "case-break",
			KEBAB: "CASE-BREAK",
			Camel: "CASEBreak",
			camel: "caseBreak",
			title: "CASE Break",
		},
		{
			input: "ID",
			snake: "id",
			SNAKE: "ID",
			kebab: "id",
			KEBAB: "ID",
			Camel: "ID",
			camel: "id",
			title: "ID",
		},
		{
			input: "userID",
			snake: "user_id",
			SNAKE: "USER_ID",
			kebab: "user-id",
			KEBAB: "USER-ID",
			Camel: "UserID",
			camel: "userID",
			title: "User ID",
		},
		{
			input: "JSON_blob",
			snake: "json_blob",
			SNAKE: "JSON_BLOB",
			kebab: "json-blob",
			KEBAB: "JSON-BLOB",
			Camel: "JSONBlob",
			camel: "jsonBlob",
			title: "JSON Blob",
		},
		{
			input: "HTTPStatusCode",
			snake: "http_status_code",
			SNAKE: "HTTP_STATUS_CODE",
			kebab: "http-status-code",
			KEBAB: "HTTP-STATUS-CODE",
			Camel: "HTTPStatusCode",
			camel: "httpStatusCode",
			title: "HTTP Status Code",
		},
		{
			input: "FreeBSD and SSLError are not golang initialisms",
			snake: "free_bsd_and_ssl_error_are_not_golang_initialisms",
			SNAKE: "FREE_BSD_AND_SSL_ERROR_ARE_NOT_GOLANG_INITIALISMS",
			kebab: "free-bsd-and-ssl-error-are-not-golang-initialisms",
			KEBAB: "FREE-BSD-AND-SSL-ERROR-ARE-NOT-GOLANG-INITIALISMS",
			Camel: "FreeBSDAndSSLErrorAreNotGolangInitialisms",
			camel: "freeBSDAndSSLErrorAreNotGolangInitialisms",
			title: "Free BSD And SSL Error Are Not Golang Initialisms",
		},
		{
			input: "David's Computer",
			snake: "david's_computer",
			SNAKE: "DAVID'S_COMPUTER",
			kebab: "david's-computer",
			KEBAB: "DAVID'S-COMPUTER",
			Camel: "David'sComputer",
			camel: "david'sComputer",
			title: "David's Computer",
		},
		{
			input: "Ünicode support for Æthelred and Øyvind",
			snake: "ünicode_support_for_æthelred_and_øyvind",
			SNAKE: "ÜNICODE_SUPPORT_FOR_ÆTHELRED_AND_ØYVIND",
			kebab: "ünicode-support-for-æthelred-and-øyvind",
			KEBAB: "ÜNICODE-SUPPORT-FOR-ÆTHELRED-AND-ØYVIND",
			Camel: "ÜnicodeSupportForÆthelredAndØyvind",
			camel: "ünicodeSupportForÆthelredAndØyvind",
			title: "Ünicode Support For Æthelred And Øyvind",
		},
		{
			input: "http200",
			snake: "http200",
			SNAKE: "HTTP200",
			kebab: "http200",
			KEBAB: "HTTP200",
			Camel: "Http200",
			camel: "http200",
			title: "Http200",
		},
		{
			input: "NumberSplittingVersion1.0r3",
			snake: "number_splitting_version1.0r3",
			SNAKE: "NUMBER_SPLITTING_VERSION1.0R3",
			kebab: "number-splitting-version1.0r3",
			KEBAB: "NUMBER-SPLITTING-VERSION1.0R3",
			Camel: "NumberSplittingVersion1.0r3",
			camel: "numberSplittingVersion1.0r3",
			title: "Number Splitting Version1.0r3",
		},
		{
			input: "When you have a comma, odd results",
			snake: "when_you_have_a_comma,_odd_results",
			SNAKE: "WHEN_YOU_HAVE_A_COMMA,_ODD_RESULTS",
			kebab: "when-you-have-a-comma,-odd-results",
			KEBAB: "WHEN-YOU-HAVE-A-COMMA,-ODD-RESULTS",
			Camel: "WhenYouHaveAComma,OddResults",
			camel: "whenYouHaveAComma,OddResults",
			title: "When You Have A Comma, Odd Results",
		},
		{
			input: "Ordinal numbers work: 1st 2nd and 3rd place",
			snake: "ordinal_numbers_work:_1st_2nd_and_3rd_place",
			SNAKE: "ORDINAL_NUMBERS_WORK:_1ST_2ND_AND_3RD_PLACE",
			kebab: "ordinal-numbers-work:-1st-2nd-and-3rd-place",
			KEBAB: "ORDINAL-NUMBERS-WORK:-1ST-2ND-AND-3RD-PLACE",
			Camel: "OrdinalNumbersWork:1st2ndAnd3rdPlace",
			camel: "ordinalNumbersWork:1st2ndAnd3rdPlace",
			title: "Ordinal Numbers Work: 1st 2nd And 3rd Place",
		},
		{
			input: "BadUTF8\xe2\xe2\xa1",
			snake: "bad_utf8_���",
			SNAKE: "BAD_UTF8_���",
			kebab: "bad-utf8-���",
			KEBAB: "BAD-UTF8-���",
			Camel: "BadUTF8���",
			camel: "badUTF8���",
			title: "Bad UTF8 ���",
		},
		{
			input: "IDENT3",
			snake: "ident3",
			SNAKE: "IDENT3",
			kebab: "ident3",
			KEBAB: "IDENT3",
			Camel: "IDENT3",
			camel: "ident3",
			title: "IDENT3",
		},
		{
			input: "LogRouterS3BucketName",
			snake: "log_router_s3_bucket_name",
			SNAKE: "LOG_ROUTER_S3_BUCKET_NAME",
			kebab: "log-router-s3-bucket-name",
			KEBAB: "LOG-ROUTER-S3-BUCKET-NAME",
			Camel: "LogRouterS3BucketName",
			camel: "logRouterS3BucketName",
			title: "Log Router S3 Bucket Name",
		},
		{
			input: "PINEAPPLE",
			snake: "pineapple",
			SNAKE: "PINEAPPLE",
			kebab: "pineapple",
			KEBAB: "PINEAPPLE",
			Camel: "PINEAPPLE",
			camel: "pineapple",
			title: "PINEAPPLE",
		},
		{
			input: "Int8Value",
			snake: "int8_value",
			SNAKE: "INT8_VALUE",
			kebab: "int8-value",
			KEBAB: "INT8-VALUE",
			Camel: "Int8Value",
			camel: "int8Value",
			title: "Int8 Value",
		},
		{
			input: "first.last",
			snake: "first_last",
			SNAKE: "FIRST_LAST",
			kebab: "first-last",
			KEBAB: "FIRST-LAST",
			Camel: "FirstLast",
			camel: "firstLast",
			title: "First Last",
		},
	} {
		t.Run(test.input, func(t *testing.T) {
			output := data{
				input: test.input,
				snake: ToSnake(test.input),
				SNAKE: ToSNAKE(test.input),
				kebab: ToKebab(test.input),
				KEBAB: ToKEBAB(test.input),
				Camel: ToCamel(test.input),
				camel: ToLowerCamel(test.input),
				title: ToTitle(test.input),
			}
			if generate || test != output {
				line := fmt.Sprintf("%#v", output)
				line = strings.TrimPrefix(line, "strcase.data")
				line = strings.Replace(line, "\", ", "\",\n", -1)
				line = strings.Replace(line, "{", "{\n", -1)
				line = strings.Replace(line, "}", "\n},", -1)
				line = regexp.MustCompile("\"\n").ReplaceAllString(line, "\",\n")
				fmt.Println(line)
			}
			assertTrue(t, test == output)
		})
	}
}
