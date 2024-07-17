package server

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"kratos-realworld-r/internal/errors"
	"testing"
)

func TestHTTPStruct(t *testing.T) {
	a := &errors.HTTPError{
		Errors: make(map[string][]string),
	}
	a.Errors["body"] = []string{"can't be empty"}
	b, err := json.Marshal(a)
	assert.NoError(t, err)
	fmt.Printf("%s", string(b))
	panic("face")
}
