package merror

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestErrorStruct(t *testing.T) {
	err1 := errors.New("math: square root of negative number")
	statsErrJSONBody, _ := json.Marshal(ErrorMessage{Message: err1.Error(), Code: 500})

	errorJson := `{"message":"math: square root of negative number","code":500}`

	if string(statsErrJSONBody) != errorJson {
		t.Fatalf("Responses are not equal %s with %s", string(statsErrJSONBody), errorJson)
	}

}
