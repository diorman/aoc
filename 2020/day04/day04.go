package day04

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passportMap map[string]string

type validatorFunc func(passportMap) bool

var passwordRegex = regexp.MustCompile(`\s([a-z]{3}):(\S+)`)

var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func Resolve(puzzlePath string) ([]interface{}, error) {
	passports, err := parsePuzzle(puzzlePath)
	if err != nil {
		return nil, err
	}

	return []interface{}{
		countValidPasswords(passports, hasRequiredFields),
		countValidPasswords(passports, matchesFieldRules),
	}, nil
}

func parseRawPassport(rawPassport string) passportMap {
	passport := make(passportMap)

	for _, match := range passwordRegex.FindAllStringSubmatch(rawPassport, -1) {
		passport[match[1]] = match[2]
	}

	return passport
}

func parsePuzzle(parsePuzzle string) ([]passportMap, error) {
	file, err := os.Open(parsePuzzle)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	passports := []passportMap{}
	rawPassport := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			rawPassport = fmt.Sprintf("%s %s", rawPassport, line)
			continue
		}

		passports = append(passports, parseRawPassport(rawPassport))
		rawPassport = ""
	}

	if rawPassport != "" {
		passports = append(passports, parseRawPassport(rawPassport))
	}

	return passports, scanner.Err()
}

func hasRequiredFields(p passportMap) bool {
	for _, field := range requiredFields {
		if _, ok := p[field]; !ok {
			return false
		}
	}

	return true
}

func isNumberInRange(value string, min, max int) bool {
	n, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return n >= min && n <= max
}

func matchesRegex(regex string, value string) bool {
	re := regexp.MustCompile(regex)
	return re.MatchString(value)
}

func matchesNDigits(value string, n int) bool {
	return matchesRegex(fmt.Sprintf(`^\d{%d}$`, n), value)
}

func isValidHeight(value string) bool {
	// hgt (Height) - a number followed by either cm or in:
	//   If cm, the number must be at least 150 and at most 193.
	//   If in, the number must be at least 59 and at most 76.
	if matchesRegex(`^\d+cm$`, value) {
		return isNumberInRange(strings.TrimSuffix(value, "cm"), 150, 193)
	}

	if matchesRegex(`^\d+in$`, value) {
		return isNumberInRange(strings.TrimSuffix(value, "in"), 59, 76)
	}

	return false
}

func matchesFieldRules(p passportMap) bool {
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	if !matchesNDigits(p["byr"], 4) || !isNumberInRange(p["byr"], 1920, 2002) {
		return false
	}

	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	if !matchesNDigits(p["iyr"], 4) || !isNumberInRange(p["iyr"], 2010, 2020) {
		return false
	}

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	if !matchesNDigits(p["eyr"], 4) || !isNumberInRange(p["eyr"], 2020, 2030) {
		return false
	}

	if !isValidHeight(p["hgt"]) {
		return false
	}

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	if !matchesRegex(`^#[0-9a-f]{6}$`, p["hcl"]) {
		return false
	}

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	if !matchesRegex(`^(amb|blu|brn|gry|grn|hzl|oth)$`, p["ecl"]) {
		return false
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	if !matchesNDigits(p["pid"], 9) {
		return false
	}

	return true
}

func countValidPasswords(passports []passportMap, validator validatorFunc) int {
	result := 0

	for _, p := range passports {
		if validator(p) {
			result++
		}
	}

	return result
}
