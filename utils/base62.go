package utils

const base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func EncodeBase62(num int) string {

	if num == 0 {
		return "a"
	}

	result := ""

	for num > 0 {
		remainder := num % 62
		result = string(base62Chars[remainder]) + result
		num = num / 62
	}

	return result
}
