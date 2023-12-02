package main

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/codeLine"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/regex"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/times"
	"testing"
)

// go test -v -run="TestGetCountCodeLine"
func TestGetCountCodeLine(t *testing.T) {
	pathDirRel := "./app"
	countCodeLine := codeLine.GetCountCodeLine(pathDirRel)
	if countCodeLine < 100 {
		t.Errorf(
			"the countCodeline |%d| is invalid",
			countCodeLine)
	}
}

// go test -v -run="TestGetMsgCodeLine"
func TestGetMsgCodeLine(t *testing.T) {
	pathDirRel := "./app"
	msgCodeLine := codeLine.GetMsgCodeLine(pathDirRel)
	dd.IGNORE(msgCodeLine)
	fmt.Println(msgCodeLine)
}

// go test -v -run="TestIsEmail"
func TestIsEmail(t *testing.T) {
	email := "13683012872@163.com"
	isValid := regex.IsEmail(email)
	if !isValid {
		t.Errorf(
			"the email |%s| is invalid",
			email)
	}
}

// go test -v -run="TestIsMobileNo"
func TestIsMobileNo(t *testing.T) {
	mobileNo := "13683012872"
	isValid := regex.IsMobileNo(mobileNo)
	if !isValid {
		t.Errorf(
			"the mobileNo |%s| is invalid",
			mobileNo)
	}
}

// go test -bench="BenchmarkToValidType"
func BenchmarkToValidType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conv.ToValidType(0, "string")
	}
}

// go test -bench="BenchmarkIsTime"
func BenchmarkIsTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		now := times.GetNow()
		regex.IsTime(now)
	}
}
