package grafana

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestStackVal_UnmarshalJSON_GotTrue(t *testing.T) {
	var sampleOut struct {
		Val BoolString `json:"val"`
	}
	var sampleIn = []byte(`{"val":true}`)

	json.Unmarshal(sampleIn, &sampleOut)

	if !sampleOut.Val.Flag {
		t.Errorf("should be true but got false")
	}
	if sampleOut.Val.Value != "" {
		t.Error("string value should be empty")
	}
}

func TestStackVal_UnmarshalJSON_GotFalse(t *testing.T) {
	var sampleOut struct {
		Val BoolString `json:"val"`
	}
	var sampleIn = []byte(`{"val":false}`)

	json.Unmarshal(sampleIn, &sampleOut)

	if sampleOut.Val.Flag {
		t.Errorf("should be false but got true")
	}
	if sampleOut.Val.Value != "" {
		t.Error("string value should be empty")
	}
}

func TestStackVal_UnmarshalJSON_GotString(t *testing.T) {
	var sampleOut struct {
		Val BoolString `json:"val"`
	}
	var sampleIn = []byte(`{"val":"A"}`)

	json.Unmarshal(sampleIn, &sampleOut)

	if sampleOut.Val.Flag {
		t.Error("should be false but got true")
	}
	if sampleOut.Val.Value != "A" {
		t.Errorf("should be 'A' but got '%s'", sampleOut.Val.Value)
	}
}

func TestStackVal_MarshalJSON_GotTrue(t *testing.T) {
	var sampleInp struct {
		Val BoolString `json:"val"`
	}
	sampleInp.Val.Flag = true
	var sampleOut = []byte(`{"val":true}`)

	data, _ := json.Marshal(sampleInp)

	if bytes.Compare(data, sampleOut) != 0 {
		t.Errorf("should be %s but got %s", sampleOut, data)
	}
}

func TestStackVal_MarshalJSON_GotFalse(t *testing.T) {
	var sampleInp struct {
		Val BoolString `json:"val"`
	}
	sampleInp.Val.Flag = false
	var sampleOut = []byte(`{"val":false}`)

	data, _ := json.Marshal(sampleInp)

	if bytes.Compare(data, sampleOut) != 0 {
		t.Errorf("should be %s but got %s", sampleOut, data)
	}
}

func TestStackVal_MarshalJSON_GotString(t *testing.T) {
	var sampleInp struct {
		Val BoolString `json:"val"`
	}
	sampleInp.Val.Value = "A"
	var sampleOut = []byte(`{"val":"A"}`)

	data, _ := json.Marshal(sampleInp)

	if bytes.Compare(data, sampleOut) != 0 {
		t.Errorf("should be %s but got %s", sampleOut, data)
	}
}
