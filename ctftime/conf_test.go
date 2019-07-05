package ctftime

import (
	"io/ioutil"
	"os"
	"testing"
)

const testJson = `{
	"slack_config": {
		"api_token": "TEST_API_TOKEN",
		"channel_id": "TEST_CHANNEL_ID"
	}
}`

func TestUnmarshalConfig(t *testing.T) {
	type testCase struct {
		Test     string
		Result   string
		Expected string
	}

	tmpDir, _ := ioutil.TempDir("", "test")
	defer os.RemoveAll(tmpDir)
	ioutil.WriteFile(tmpDir+"/test_config.json", []byte(testJson), 0644)
	testConfigFilePath := tmpDir + "/test_config.json"
	config, err := NewConfig(testConfigFilePath)
	if err != nil {
		t.Error(err)
	}

	testCases := []testCase{
		{
			Test:     "test api_token",
			Result:   config.SlackConfig.APIToken,
			Expected: "TEST_API_TOKEN",
		},
	}

	for _, test := range testCases {
		if test.Expected != test.Result {
			t.Errorf("test:%v, expected:%v, result:%v", test.Test, test.Expected, test.Result)
		}
	}

}

func TestInvalidFilePath(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("", "test")
	defer os.RemoveAll(tmpDir)
	ioutil.WriteFile(tmpDir+"/test_config.json", []byte(testJson), 0644)
	testConfigFilePath := tmpDir + "/invalid_config.json"
	config, err := NewConfig(testConfigFilePath)
	if err == nil {
		t.Errorf("test: TestInvalidFilePath, expected: nil, result:%v", config)
	}
}

func TestInvalidFileContent(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("", "test")
	defer os.RemoveAll(tmpDir)
	ioutil.WriteFile(tmpDir+"/test_config.json", []byte(""), 0644)
	testConfigFilePath := tmpDir + "/test_config.json"
	config, err := NewConfig(testConfigFilePath)
	if err == nil {
		t.Errorf("test: TestInvalidFileContent, expected: nil, result:%v", config)
	}
}
