package tests

import (
	"root_academy_go/internal/user/adapter/in/web/dto"
	"root_academy_go/internal/user/domain/rules"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_valid_email(t *testing.T) {
	input := dto.Register{Email: "wrong_email"}

	got := rules.ValidEmail(input.Email)

	assert.False(t, got)
}
