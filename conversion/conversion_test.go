package conversion

import (
	"encoding/json"
	"fmt"
	"github.com/matryer/is"
	"testing"
)

type entity struct {
	X int `gorm:"easting"`
	Y int `gorm:"northing"`
}

func (e entity) String() string {
	return fmt.Sprintf("[%v,%v]", e.X, e.Y)
}

type model struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (m model) String() string {
	str, _ := json.Marshal(m)
	return string(str)
}

func Test_Conversion(t *testing.T) {
	Is := is.New(t)

	t.Run("it should copy fields over if they match irrespective of tags", func(t *testing.T) {
		dbEntity := entity{3, 14}
		Is.Equal(dbEntity.X, 3)
		Is.Equal(dbEntity.Y, 14)
		Is.Equal(dbEntity.String(), "[3,14]")

		m := model(dbEntity)
		Is.Equal(m.X, 3)
		Is.Equal(m.Y, 14)
		Is.Equal(m.String(), `{"x":3,"y":14}`)
	})
}
