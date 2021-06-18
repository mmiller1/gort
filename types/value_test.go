/*
 * Copyright 2021 The Gort Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGuessTypedValue(t *testing.T) {
	tests := map[string]interface{}{
		`true`:      BoolValue{true},
		`false`:     BoolValue{false},
		`0.0`:       FloatValue{0.0},
		`.10`:       FloatValue{0.10},
		`-1.0`:      FloatValue{-1.0},
		`0`:         IntValue{0},
		`10`:        IntValue{10},
		`-1`:        IntValue{-1},
		`"testing"`: StringValue{"testing", '"'},
		`'testing'`: StringValue{"testing", '\''},
		`""`:        StringValue{"", '"'},
		`''`:        StringValue{"", '\''},
		`arbitrary`: StringValue{"arbitrary", '\u0000'},
	}

	for input, expected := range tests {
		actual, err := GuessTypedValue(input, false)
		if !assert.NoError(t, err, input) {
			continue
		}

		assert.Equal(t, expected, actual)
	}
}

func TestGuessTypesValueStrict(t *testing.T) {
	tests := map[string]interface{}{
		`"testing"`: StringValue{"testing", '"'},
		`'testing'`: StringValue{"testing", '\''},
		`""`:        StringValue{"", '"'},
		`''`:        StringValue{"", '\''},
	}

	for input, expected := range tests {
		actual, err := GuessTypedValue(input, true)
		if !assert.NoError(t, err, input) {
			continue
		}

		assert.Equal(t, expected, actual)
	}

	_, err := GuessTypedValue("arbitrary", true)
	assert.Error(t, err, "arbitrary")
}
