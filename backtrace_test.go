package ntest

/**
电话号码的字母组合
*/

var m = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	combinations := []string{}
	if len(digits) == 0 {
		return combinations
	}

	var backtrace func(digits string, index int, combination string)
	backtrace = func(digits string, index int, combination string) {
		// index 长度和传入数字一样长时，终止
		if index == len(digits) {
			combinations = append(combinations, combination)
			return
		}

		digit := string(digits[index])
		for _, char := range m[digit] {
			backtrace(digits, index+1, combination+string(char))
		}

	}
	backtrace(digits, 0, "")
	return combinations
}
