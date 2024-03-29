package gohetznerdns

import (
	"testing"

	"gotest.tools/assert"
)

func TestNewClient(t *testing.T) {
	dns, _ := NewClient("asdadasdasd")
	assert.Assert(t, dns != nil)
	assert.Assert(t, dns.GetZoneService() != nil)
	assert.Assert(t, dns.GetRecordService() != nil)
}

func TestNewClientTokenError(t *testing.T) {
	_, err := NewClient("   ")
	assert.Error(t, err, "901 : token is empty")
}

func TestSetBaseUrl(t *testing.T) {
	dns, _ := NewClient("asdaasdada")
	err := dns.SetBaseURL("https://test.com")
	assert.NilError(t, err)
}
func TestSetBaseUrlInvalid(t *testing.T) {
	dns, _ := NewClient("asdasdsadas")
	err := dns.SetBaseURL("https://te|.^com")
	assert.Error(t, err, "parse \"https://te|.^com\": invalid character \"|\" in host name")
}
