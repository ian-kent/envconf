package envconf

import (
	"flag"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFlagFromEnv(t *testing.T) {
	Convey("Replacing flags with env vars after parsing", t, func() {
		os.Clearenv()

		s := flag.String("test-flag", "default", "A test flag")
		i := flag.Int("test-flag-1", 10, "A test flag")
		flag.Parse()

		sv := FromEnvP("TEST_FLAG", *s)
		iv := FromEnvP("TEST_FLAG_1", *i)

		So(sv, ShouldEqual, "default")
		So(iv, ShouldEqual, 10)

		os.Setenv("TEST_FLAG", "foo")
		os.Setenv("TEST_FLAG_1", "15")

		sv = FromEnvP("TEST_FLAG", *s)
		iv = FromEnvP("TEST_FLAG_1", *i)

		So(sv, ShouldEqual, "foo")
		So(iv, ShouldEqual, 15)
	})

	Convey("Replacing flags with env vars after parsing", t, func() {
		os.Clearenv()

		s := flag.String("test-flag-2", FromEnvP("TEST_FLAG_2", "default").(string), "A test flag")
		i := flag.Int("test-flag-3", FromEnvP("TEST_FLAG_3", 10).(int), "A test flag")
		flag.Parse()

		So(*s, ShouldEqual, "default")
		So(*i, ShouldEqual, 10)

		os.Setenv("TEST_FLAG_4", "foo")
		os.Setenv("TEST_FLAG_5", "15")

		s = flag.String("test-flag-4", FromEnvP("TEST_FLAG_4", "default").(string), "A test flag")
		i = flag.Int("test-flag-5", FromEnvP("TEST_FLAG_5", 10).(int), "A test flag")
		flag.Parse()

		So(*s, ShouldEqual, "foo")
		So(*i, ShouldEqual, 15)
	})
}
