package new

import "testing"

func TestCreate(t *testing.T) {
	err := Create("yaml", "./", "github.com/ggdream/home", "app", "1.0.0", "mocaraka@gmail.com", "test", "github.com/ggdream/home")
	if err != nil {
		panic(err.Error())
	}
}
