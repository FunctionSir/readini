/*
 * @Author: FunctionSir
 * @License: AGPLv3
 * @Date: 2025-04-06 00:15:04
 * @LastEditTime: 2025-04-20 22:31:21
 * @LastEditors: FunctionSir
 * @Description: -
 * @FilePath: /readini/readini_test.go
 */

package readini

import (
	"testing"
)

var testStrOK string = "KeyA=ValA\n [     SomeSection   ]\n 第二个键🚩 = 初音 ミク／はつねミク Hatsune Miku 🚩 来自中文维基百科"
var testStrErr string = "KeyA = ValA\n [SomeSection \n KeyB=ValB"

var testFileOk string = "testFileOK.conf"
var testFileErr string = "testFileErr.conf"
var testFileNotExists string = "testFileNotExists.conf"

func TestLoadFromFile(t *testing.T) {
	res, err := LoadFromFile(testFileOk)
	if err != nil {
		t.Errorf("Should not get an error, but error \"%s\" got",
			err.Error())
	}
	if !res.HasSection("") {
		t.Errorf("Should have section \"\", but it said it doesn't have")
	}
	if res.HasSection("NoSuchSection") {
		t.Errorf("Should not have section \"NoSuchSection\", but it said it has")
	}
	if !res.HasSection("SomeSection") {
		t.Errorf("Should have section \"SomeSection\", but it said it doesn't have")
	}
	if res[""]["KeyA"] != "ValA" {
		t.Errorf("Val of \"KeyA\" be \"ValA\", but %s found",
			res[""]["KeyA"])
	}
	if !res.HasKey("SomeSection", "第二个键🚩") {
		t.Errorf("There has key \"第二个键🚩\" in section \"SomeSection\", but got false")
	}
	if res.HasKey("SomeSection", "第三个键🚩") {
		t.Errorf("There has no key \"第三个键🚩\" in section \"SomeSection\", but got true")
	}
	if res.HasKey("NoSuchSection", "第四个键🚩") {
		t.Errorf("There has no key \"第三个键🚩\" in section \"SomeSection\", but got true")
	}
	if res["SomeSection"]["第二个键🚩"] != "初音 ミク／はつねミク Hatsune Miku 🚩 来自中文维基百科" {
		t.Errorf("Val of \"第二个键🚩\" in section \"SomeSection\" should be \"初音 ミク／はつねミク Hatsune Miku 🚩 来自中文维基百科\", but %s found",
			res["SomeSection"]["第二个键🚩"])
	}
	sec := res["SomeSection"]
	if !sec.HasKey("第二个键🚩") {
		t.Errorf("There has key \"第二个键🚩\" in section \"SomeSection\", but got false")
	}
	if sec.HasKey("第三个键🚩") {
		t.Errorf("There has no key \"第三个键🚩\" in section \"SomeSection\", but got true")
	}
	_, err = LoadFromFile(testFileErr)
	if err == nil {
		t.Errorf("Should get an error, but \"nil\" got")
	}
	_, err = LoadFromFile(testFileNotExists)
	if err == nil {
		t.Errorf("Should get an error, but \"nil\" got")
	}
}

func TestLoadFromRunes(t *testing.T) {
	res, err := LoadFromRunes([]rune(testStrOK))
	if err != nil {
		t.Errorf("Should not get an error, but error \"%s\" got",
			err.Error())
	}
	if !res.HasSection("") {
		t.Errorf("Should have section \"\", but it said it doesn't have")
	}
	if !res.HasSection("SomeSection") {
		t.Errorf("Should have section \"SomeSection\", but it said it doesn't have")
	}
	if res[""]["KeyA"] != "ValA" {
		t.Errorf("Val of \"KeyA\" be \"ValA\", but %s found",
			res[""]["KeyA"])
	}
	if res["SomeSection"]["第二个键🚩"] != "初音 ミク／はつねミク Hatsune Miku 🚩 来自中文维基百科" {
		t.Errorf("Val of \"第二个键🚩\" in section \"SomeSection\" should be \"初音 ミク／はつねミク Hatsune Miku 🚩 来自中文维基百科\", but %s found",
			res["SomeSection"]["第二个键🚩"])
	}
	_, err = LoadFromRunes([]rune(testStrErr))
	if err == nil {
		t.Errorf("Should get an error, but \"nil\" got")
	}
}
