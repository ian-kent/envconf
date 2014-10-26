package envconf

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFromEnv(t *testing.T) {
	Convey("Value not replaced if environment variable not set", t, func() {
		os.Clearenv()
		v := interface{}("default")

		vl, err := FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, "default")
		So(err, ShouldBeNil)
	})

	Convey("Value replaced if environment variable is set", t, func() {
		os.Clearenv()
		v := interface{}("default")
		os.Setenv("FOO_VAR", "bar")

		vl, err := FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, "bar")
		So(err, ShouldBeNil)
	})

	Convey("Value replaced if environment variable is set to empty value", t, func() {
		os.Clearenv()
		v := interface{}("default")
		os.Setenv("FOO_VAR", "")

		vl, err := FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, "")
		So(err, ShouldBeNil)
	})

	Convey("FromEnvP panics on error", t, func() {
		os.Clearenv()
		v := interface{}(map[string]string{"foo": "bar"})
		os.Setenv("FOO_VAR", "true")

		So(func() {
			FromEnvP("FOO_VAR", v)
		}, ShouldPanic)
	})

	Convey("Unsupported type returns an error", t, func() {
		os.Clearenv()
		v := interface{}(map[string]string{"foo": "bar"})

		os.Setenv("FOO_VAR", "true")
		vl, err := FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, v)
		So(err, ShouldEqual, ErrUnsupportedType)
	})

	Convey("bool returns a bool", t, func() {
		os.Clearenv()
		v := interface{}(false)

		vl, err := FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, false)
		So(err, ShouldBeNil)

		os.Setenv("FOO_VAR", "true")
		vl, err = FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, true)
		So(vl, ShouldHaveSameTypeAs, false)
		So(vl, ShouldNotEqual, "test")
		So(err, ShouldBeNil)
	})

	Convey("string returns a string", t, func() {
		os.Clearenv()
		v := interface{}(string("test"))

		vl, err := FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, string("test"))
		So(err, ShouldBeNil)

		os.Setenv("FOO_VAR", "bar")
		vl, err = FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, string("bar"))
		So(vl, ShouldHaveSameTypeAs, string("bar"))
		So(vl, ShouldNotEqual, "test")
		So(err, ShouldBeNil)
	})

	Convey("int returns an int", t, func() {
		os.Clearenv()
		v := interface{}(int(10))

		vl, err := FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, int(10))
		So(err, ShouldBeNil)

		os.Setenv("FOO_VAR", "15")
		vl, err = FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, int(15))
		So(vl, ShouldHaveSameTypeAs, int(15))
		So(vl, ShouldNotEqual, "15")
		So(err, ShouldBeNil)
	})

	Convey("int8 returns an int8", t, func() {
		os.Clearenv()
		v := interface{}(int8(10))

		vl, err := FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, int8(10))
		So(err, ShouldBeNil)

		os.Setenv("FOO_VAR", "15")
		vl, err = FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, int8(15))
		So(vl, ShouldHaveSameTypeAs, int8(15))
		So(vl, ShouldNotEqual, "15")
		So(err, ShouldBeNil)
	})

	Convey("int16 returns an int16", t, func() {
		os.Clearenv()
		v := interface{}(int16(10))

		vl, err := FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, int16(10))
		So(err, ShouldBeNil)

		os.Setenv("FOO_VAR", "15")
		vl, err = FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, int16(15))
		So(vl, ShouldHaveSameTypeAs, int16(15))
		So(vl, ShouldNotEqual, "15")
		So(err, ShouldBeNil)
	})

	Convey("int32 returns an int32", t, func() {
		os.Clearenv()
		v := interface{}(int32(10))

		vl, err := FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, int32(10))
		So(err, ShouldBeNil)

		os.Setenv("FOO_VAR", "15")
		vl, err = FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, int32(15))
		So(vl, ShouldHaveSameTypeAs, int32(15))
		So(vl, ShouldNotEqual, "15")
		So(err, ShouldBeNil)
	})

	Convey("int64 returns an int64", t, func() {
		os.Clearenv()
		v := interface{}(int64(10))

		vl, err := FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, int64(10))
		So(err, ShouldBeNil)

		os.Setenv("FOO_VAR", "15")
		vl, err = FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, int64(15))
		So(vl, ShouldHaveSameTypeAs, int64(15))
		So(vl, ShouldNotEqual, "15")
		So(err, ShouldBeNil)
	})

	Convey("uint returns a uint", t, func() {
		os.Clearenv()
		v := interface{}(uint(10))

		vl, err := FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, uint(10))
		So(err, ShouldBeNil)

		os.Setenv("FOO_VAR", "15")
		vl, err = FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, uint(15))
		So(vl, ShouldHaveSameTypeAs, uint(15))
		So(vl, ShouldNotEqual, "15")
		So(err, ShouldBeNil)
	})

	Convey("uint8 returns a uint8", t, func() {
		os.Clearenv()
		v := interface{}(uint8(10))

		vl, err := FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, uint8(10))
		So(err, ShouldBeNil)

		os.Setenv("FOO_VAR", "15")
		vl, err = FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, uint8(15))
		So(vl, ShouldHaveSameTypeAs, uint8(15))
		So(vl, ShouldNotEqual, "15")
		So(err, ShouldBeNil)
	})

	Convey("uint16 returns a uint16", t, func() {
		os.Clearenv()
		v := interface{}(uint16(10))

		vl, err := FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, uint16(10))
		So(err, ShouldBeNil)

		os.Setenv("FOO_VAR", "15")
		vl, err = FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, uint16(15))
		So(vl, ShouldHaveSameTypeAs, uint16(15))
		So(vl, ShouldNotEqual, "15")
		So(err, ShouldBeNil)
	})

	Convey("uint32 returns a uint32", t, func() {
		os.Clearenv()
		v := interface{}(uint32(10))

		vl, err := FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, uint32(10))
		So(err, ShouldBeNil)

		os.Setenv("FOO_VAR", "15")
		vl, err = FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, uint32(15))
		So(vl, ShouldHaveSameTypeAs, uint32(15))
		So(vl, ShouldNotEqual, "15")
		So(err, ShouldBeNil)
	})

	Convey("uint64 returns a uint64", t, func() {
		os.Clearenv()
		v := interface{}(uint64(10))

		vl, err := FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, uint64(10))
		So(err, ShouldBeNil)

		os.Setenv("FOO_VAR", "15")
		vl, err = FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, uint64(15))
		So(vl, ShouldHaveSameTypeAs, uint64(15))
		So(vl, ShouldNotEqual, "15")
		So(err, ShouldBeNil)
	})

	Convey("float32 returns a float32", t, func() {
		os.Clearenv()
		v := interface{}(float32(10))

		vl, err := FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, float32(10))
		So(err, ShouldBeNil)

		os.Setenv("FOO_VAR", "15")
		vl, err = FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, float32(15))
		So(vl, ShouldHaveSameTypeAs, float32(15))
		So(vl, ShouldNotEqual, "15")
		So(err, ShouldBeNil)
	})

	Convey("float64 returns a float64", t, func() {
		os.Clearenv()
		v := interface{}(float64(10))

		vl, err := FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, float64(10))
		So(err, ShouldBeNil)

		os.Setenv("FOO_VAR", "15")
		vl, err = FromEnv("FOO_VAR", v)
		So(vl, ShouldEqual, float64(15))
		So(vl, ShouldHaveSameTypeAs, float64(15))
		So(vl, ShouldNotEqual, "15")
		So(err, ShouldBeNil)
	})
}
