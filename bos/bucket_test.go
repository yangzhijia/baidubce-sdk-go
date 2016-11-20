package bos

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/guoyao/baidubce-sdk-go/bce"
	"github.com/guoyao/baidubce-sdk-go/util"
)

func TestNewObjectMetadataFromHeader(t *testing.T) {
	header := http.Header{
		"cache-control":       []string{"no-cache", "no-store"},
		"Content-Disposition": []string{"inline"},
		"Content-Length":      []string{"1024"},
		"Content-Range":       []string{"bytes=0-1024"},
		"Content-Type":        []string{"text/plain"},
		"Expires":             []string{"Tue, 03 Jan 2017 05:30:19 GMT"},
		"Etag":                []string{"abc123"},
		"x-bce-meta-name":     []string{"hello"},
	}
	metadata := NewObjectMetadataFromHeader(header)

	if metadata.CacheControl != "no-cache" {
		t.Error(util.FormatTest("NewObjectMetadataFromHeader", metadata.CacheControl, "no-cache"))
	}

	if metadata.ContentDisposition != "inline" {
		t.Error(util.FormatTest("NewObjectMetadataFromHeader", metadata.ContentDisposition, "inline"))
	}

	if metadata.ContentLength != 1024 {
		t.Error(util.FormatTest("NewObjectMetadataFromHeader", strconv.FormatInt(metadata.ContentLength, 10), strconv.Itoa(1024)))
	}

	if metadata.ContentRange != "bytes=0-1024" {
		t.Error(util.FormatTest("NewObjectMetadataFromHeader", metadata.ContentRange, "bytes=0-1024"))
	}

	if metadata.ContentType != "text/plain" {
		t.Error(util.FormatTest("NewObjectMetadataFromHeader", metadata.ContentType, "text/plain"))
	}

	if metadata.Expires != "Tue, 03 Jan 2017 05:30:19 GMT" {
		t.Error(util.FormatTest("NewObjectMetadataFromHeader", metadata.Expires, "Tue, 03 Jan 2017 05:30:19 GMT"))
	}

	if metadata.ETag != "abc123" {
		t.Error(util.FormatTest("NewObjectMetadataFromHeader", metadata.ETag, "abc123"))
	}

	if metadata.UserMetadata["x-bce-meta-name"] != "hello" {
		t.Error(util.FormatTest("NewObjectMetadataFromHeader", metadata.UserMetadata["x-bce-meta-name"], "hello"))
	}
}

func TestAddUserMetadata(t *testing.T) {
	metadata := &ObjectMetadata{}
	metadata.AddUserMetadata("x-bce-meta-name", "hello")

	if metadata.UserMetadata["x-bce-meta-name"] != "hello" {
		t.Error(util.FormatTest("AddUserMetadata", metadata.UserMetadata["x-bce-meta-name"], "hello"))
	}
}

func TestMergeToSignOption(t *testing.T) {
	option := &bce.SignOption{}
	header := http.Header{
		"cache-control":       []string{"no-cache", "no-store"},
		"Content-Disposition": []string{"inline"},
		"Content-Length":      []string{"1024"},
		"Content-Range":       []string{"bytes=0-1024"},
		"Content-Type":        []string{"text/plain"},
		"Expires":             []string{"Tue, 03 Jan 2017 05:30:19 GMT"},
		"Etag":                []string{"abc123"},
		"x-bce-meta-name":     []string{"hello"},
	}
	metadata := NewObjectMetadataFromHeader(header)
	metadata.AddUserMetadata("server", "nginx")
	metadata.mergeToSignOption(option)

	if option.Headers["Cache-Control"] != "no-cache" {
		t.Error(util.FormatTest("mergeToSignOption", option.Headers["Cache-Control"], "no-cache"))
	}

	if option.Headers["Content-Disposition"] != "inline" {
		t.Error(util.FormatTest("mergeToSignOption", option.Headers["Content-Disposition"], "inline"))
	}

	if option.Headers["x-bce-meta-name"] != "hello" {
		t.Error(util.FormatTest("mergeToSignOption", option.Headers["x-bce-meta-name"], "hello"))
	}

	if option.Headers["x-bce-meta-server"] != "nginx" {
		t.Error(util.FormatTest("mergeToSignOption", option.Headers["x-bce-meta-server"], "nginx"))
	}
}

func TestIsUserDefinedMetadata(t *testing.T) {
	expected := true
	result := IsUserDefinedMetadata("x-bce-meta-name")

	if result != expected {
		t.Error(util.FormatTest("IsUserDefinedMetadata", strconv.FormatBool(result), strconv.FormatBool(expected)))
	}

	expected = false
	result = IsUserDefinedMetadata("content-type")

	if result != expected {
		t.Error(util.FormatTest("IsUserDefinedMetadata", strconv.FormatBool(result), strconv.FormatBool(expected)))
	}
}

func TestToUserDefinedMetadata(t *testing.T) {
	expected := "x-bce-meta-name"
	result := ToUserDefinedMetadata("x-bce-meta-name")

	if result != expected {
		t.Error(util.FormatTest("ToUserDefinedMetadata", result, expected))
	}

	result = ToUserDefinedMetadata("name")

	if result != expected {
		t.Error(util.FormatTest("ToUserDefinedMetadata", result, expected))
	}

	expected = "x-bce-meta-content-type"
	result = ToUserDefinedMetadata("content-type")

	if result != expected {
		t.Error(util.FormatTest("ToUserDefinedMetadata", result, expected))
	}
}
