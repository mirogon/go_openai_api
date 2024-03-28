// Code generated by MockGen. DO NOT EDIT.
// Source: openai_communicator.go

// Package mock_openai_api is a generated GoMock package.
package mock_openai_api

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	error_system "github.com/mirogon/go_error_system"
	openai_data "github.com/mirogon/go_openai_api/data"
)

// MockOpenAiApiCommunicator is a mock of OpenAiApiCommunicator interface.
type MockOpenAiApiCommunicator struct {
	ctrl     *gomock.Controller
	recorder *MockOpenAiApiCommunicatorMockRecorder
}

// MockOpenAiApiCommunicatorMockRecorder is the mock recorder for MockOpenAiApiCommunicator.
type MockOpenAiApiCommunicatorMockRecorder struct {
	mock *MockOpenAiApiCommunicator
}

// NewMockOpenAiApiCommunicator creates a new mock instance.
func NewMockOpenAiApiCommunicator(ctrl *gomock.Controller) *MockOpenAiApiCommunicator {
	mock := &MockOpenAiApiCommunicator{ctrl: ctrl}
	mock.recorder = &MockOpenAiApiCommunicatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenAiApiCommunicator) EXPECT() *MockOpenAiApiCommunicatorMockRecorder {
	return m.recorder
}

// GenerateImage mocks base method.
func (m *MockOpenAiApiCommunicator) GenerateImage(input string, numImages int, resolution string) (string, error_system.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateImage", input, numImages, resolution)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error_system.Error)
	return ret0, ret1
}

// GenerateImage indicates an expected call of GenerateImage.
func (mr *MockOpenAiApiCommunicatorMockRecorder) GenerateImage(input, numImages, resolution interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateImage", reflect.TypeOf((*MockOpenAiApiCommunicator)(nil).GenerateImage), input, numImages, resolution)
}

// GptCompletion mocks base method.
func (m *MockOpenAiApiCommunicator) GptCompletion(messages []openai_data.GptMessage, maxToken int, gptModel string) (string, error_system.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GptCompletion", messages, maxToken, gptModel)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error_system.Error)
	return ret0, ret1
}

// GptCompletion indicates an expected call of GptCompletion.
func (mr *MockOpenAiApiCommunicatorMockRecorder) GptCompletion(messages, maxToken, gptModel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GptCompletion", reflect.TypeOf((*MockOpenAiApiCommunicator)(nil).GptCompletion), messages, maxToken, gptModel)
}

// TextToSpeech mocks base method.
func (m *MockOpenAiApiCommunicator) TextToSpeech(input, voice string) ([]byte, error_system.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TextToSpeech", input, voice)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error_system.Error)
	return ret0, ret1
}

// TextToSpeech indicates an expected call of TextToSpeech.
func (mr *MockOpenAiApiCommunicatorMockRecorder) TextToSpeech(input, voice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TextToSpeech", reflect.TypeOf((*MockOpenAiApiCommunicator)(nil).TextToSpeech), input, voice)
}
