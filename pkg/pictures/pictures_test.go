package pictures

import (
	"testing"
	"time"
)

func TestPictureService(t *testing.T) {
	t.Run("MakeName", testMakeName)
}

func testFetch(t *testing.T) {

}

func testMakeName(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	cases := []struct {
		name string
		time time.Time
		want string
	}{
		{name: "hoge.jpg", time: time.Date(2020, time.January, 1, 0, 0, 0, 0, jst), want: "1577804400_hoge.jpg"},
		{name: "test.png", time: time.Date(2025, time.September, 24, 0, 0, 0, 0, jst), want: "1758639600_test.png"},
	}

	service := &PictureService{}
	for _, c := range cases {
		got := service.makeName(c.name, c.time)
		if got != c.want {
			t.Errorf("unexpected return value.\nmakeName(%v, %#v) =\ngot, %v", c.name, c.time, c.want)
		}
	}
}
