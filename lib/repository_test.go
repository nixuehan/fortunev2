package lib

import (
	"testing"
)

func TestRepository_Fetch(t *testing.T) {
	stockSource := &Repository{}
	content,_ := stockSource.Fetch("sz000333")
	t.Log(content.Name)
}
