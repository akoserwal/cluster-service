// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package smtpdetails

import (
	"sync"
)

var (
	lockPasswordGeneratorMockGenerate sync.RWMutex
)

// Ensure, that PasswordGeneratorMock does implement PasswordGenerator.
// If this is not the case, regenerate this file with moq.
var _ PasswordGenerator = &PasswordGeneratorMock{}

// PasswordGeneratorMock is a mock implementation of PasswordGenerator.
//
//     func TestSomethingThatUsesPasswordGenerator(t *testing.T) {
//
//         // make and configure a mocked PasswordGenerator
//         mockedPasswordGenerator := &PasswordGeneratorMock{
//             GenerateFunc: func(length int, numDigits int, numSymbols int, noUpper bool, allowRepeat bool) (string, error) {
// 	               panic("mock out the Generate method")
//             },
//         }
//
//         // use mockedPasswordGenerator in code that requires PasswordGenerator
//         // and then make assertions.
//
//     }
type PasswordGeneratorMock struct {
	// GenerateFunc mocks the Generate method.
	GenerateFunc func(length int, numDigits int, numSymbols int, noUpper bool, allowRepeat bool) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// Generate holds details about calls to the Generate method.
		Generate []struct {
			// Length is the length argument value.
			Length int
			// NumDigits is the numDigits argument value.
			NumDigits int
			// NumSymbols is the numSymbols argument value.
			NumSymbols int
			// NoUpper is the noUpper argument value.
			NoUpper bool
			// AllowRepeat is the allowRepeat argument value.
			AllowRepeat bool
		}
	}
}

// Generate calls GenerateFunc.
func (mock *PasswordGeneratorMock) Generate(length int, numDigits int, numSymbols int, noUpper bool, allowRepeat bool) (string, error) {
	if mock.GenerateFunc == nil {
		panic("PasswordGeneratorMock.GenerateFunc: method is nil but PasswordGenerator.Generate was just called")
	}
	callInfo := struct {
		Length      int
		NumDigits   int
		NumSymbols  int
		NoUpper     bool
		AllowRepeat bool
	}{
		Length:      length,
		NumDigits:   numDigits,
		NumSymbols:  numSymbols,
		NoUpper:     noUpper,
		AllowRepeat: allowRepeat,
	}
	lockPasswordGeneratorMockGenerate.Lock()
	mock.calls.Generate = append(mock.calls.Generate, callInfo)
	lockPasswordGeneratorMockGenerate.Unlock()
	return mock.GenerateFunc(length, numDigits, numSymbols, noUpper, allowRepeat)
}

// GenerateCalls gets all the calls that were made to Generate.
// Check the length with:
//     len(mockedPasswordGenerator.GenerateCalls())
func (mock *PasswordGeneratorMock) GenerateCalls() []struct {
	Length      int
	NumDigits   int
	NumSymbols  int
	NoUpper     bool
	AllowRepeat bool
} {
	var calls []struct {
		Length      int
		NumDigits   int
		NumSymbols  int
		NoUpper     bool
		AllowRepeat bool
	}
	lockPasswordGeneratorMockGenerate.RLock()
	calls = mock.calls.Generate
	lockPasswordGeneratorMockGenerate.RUnlock()
	return calls
}
