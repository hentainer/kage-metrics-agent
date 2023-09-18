package utils_test

import (
	"testing"

	"github.com/msales/kage/utils"
	"github.com/stretchr/testify/assert"
)

func TestSplitMap(t *testing.T) {
	tests := []struct {
		in     []string
		sep    string