/*
 * @Author: FunctionSir
 * @License: AGPLv3
 * @Date: 2025-04-05 23:31:47
 * @LastEditTime: 2025-06-20 09:09:16
 * @LastEditors: FunctionSir
 * @Description: Simple go library to read INI file.
 * @FilePath: /readini/readini.go
 */

package readini

import (
	"errors"
	"os"
	"strings"
)

// You should read from vars of this type only, do NOT write.
// Write operations might cause data-race.
type Sec map[string]string

// You should read from vars of this type only, do NOT write.
// Write operations might cause data-race.
type Conf map[string]Sec

// Wrong format in INI file.
var ErrWrongFormat = errors.New("wrong config file format")

func (s Sec) HasKey(key string) bool {
	_, hasKey := s[key]
	return hasKey
}

func (c Conf) HasSection(section string) bool {
	_, hasSection := c[section]
	return hasSection
}

func (c Conf) HasKey(section, key string) bool {
	if !c.HasSection(section) {
		return false
	}
	return c[section].HasKey(key)
}

// Split key and value.
func splitKeyAndValue(s *string) (string, string) {
	runeList := []rune(*s)
	key := make([]rune, 0)
	value := make([]rune, 0)
	processedValue := make([]rune, 0)
	var i int = 0
	for i = range runeList {
		if runeList[i] == '=' {
			break
		}
		key = append(key, runeList[i])
	}
	for i = i + 1; i < len(runeList); i++ {
		value = append(value, runeList[i])
	}
	trimedValue := strings.TrimSpace(string(value))
	tmp := []rune(trimedValue)
	l := 0
	r := len(tmp) - 1
	if tmp[l] == '"' {
		l++
	}
	if tmp[r] == '"' {
		r--
	}
	for i := l; i <= r; i++ {
		processedValue = append(processedValue, tmp[i])
	}
	return strings.TrimSpace(string(key)),
		string(processedValue)
}

// Load INI from file.
func LoadFromFile(path string) (Conf, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return make(Conf), err
	}
	return LoadFromBytes(file)
}

// Load INI from runes (rune slice)
func LoadFromRunes(r []rune) (Conf, error) {
	return LoadFromString(string(r))
}

// Load INI from bytes (byte slice)
func LoadFromBytes(b []byte) (Conf, error) {
	return LoadFromString(string(b))
}

// Load INI from string.
func LoadFromString(str string) (Conf, error) {
	lines := strings.Split(str, "\n")
	return LoadFromLines(lines)
}

// Load INI from lines (string slice).
func LoadFromLines(lines []string) (Conf, error) {
	conf := make(Conf)
	curSection := ""
	conf[curSection] = make(map[string]string)
	for _, line := range lines {
		tmp := strings.TrimSpace(line)
		// Empty line or comments.
		if len(tmp) <= 0 || tmp[0] == '#' || tmp[0] == ';' {
			continue
		}
		// K-V pairs.
		if tmp[0] != '[' {
			k, v := splitKeyAndValue(&tmp)
			conf[curSection][k] = v
			continue
		}
		// Sections.
		if tmp[len(tmp)-1] != ']' || len(tmp) == 2 {
			return conf, ErrWrongFormat
		}
		tmp = tmp[1 : len(tmp)-1]
		tmp = strings.TrimSpace(tmp)
		curSection = tmp
		conf[curSection] = make(map[string]string)
	}
	return conf, nil
}
