package errors

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/zeebo/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		code ErrorCode
	}
	type output struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want output
	}{
		{
			name: "ok",
			args: args{
				code: InvalidArgument,
			},
			want: output{
				err: &CustomError{
					code:   InvalidArgument,
					origin: fmt.Errorf("InvalidArgument"),
					stack:  "",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, New(tt.args.code).Error(), tt.want.err.Error())
		})
	}

}

func TestNewf(t *testing.T) {
	type args struct {
		code   ErrorCode
		format string
		args   []interface{}
	}
	type output struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want output
	}{
		{
			name: "ok",
			args: args{
				code:   InvalidArgument,
				format: "test",
				args:   []interface{}{"test"},
			},
			want: output{
				err: &CustomError{
					code:   InvalidArgument,
					origin: fmt.Errorf("test"),
					stack:  "",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, Newf(tt.args.code, tt.args.format, tt.args.args...).Error(), tt.want.err.Error())
		})
	}
}

func TestWrap(t *testing.T) {
	type args struct {
		err error
	}
	type output struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want output
	}{
		{
			name: "ok",
			args: args{
				err: fmt.Errorf("test"),
			},
			want: output{
				err: &CustomError{
					code:   Internal,
					origin: fmt.Errorf("test"),
					stack:  "",
				},
			},
		},
		{
			name: "ok: エラーを二重にラップ",
			args: args{
				err: Wrap(fmt.Errorf("test")),
			},
			want: output{
				err: &CustomError{
					code:   Internal,
					origin: fmt.Errorf("test"),
					stack:  "",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, Wrap(tt.args.err).Error(), tt.want.err.Error())
		})
	}
}

func TestWrapf(t *testing.T) {
	type args struct {
		err    error
		format string
		args   []interface{}
	}
	type output struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want output
	}{
		{
			name: "ok",
			args: args{
				err:    fmt.Errorf("test"),
				format: "test",
				args:   []interface{}{"test"},
			},
			want: output{
				err: &CustomError{
					code:   Internal,
					origin: fmt.Errorf("test"),
					stack:  "",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, Wrapf(tt.args.err, tt.args.format, tt.args.args...).Error(), tt.want.err.Error())
		})
	}
}

func TestEqualCode(t *testing.T) {
	type args struct {
		err  error
		code ErrorCode
	}
	type output struct {
		ok bool
	}
	tests := []struct {
		name string
		args args
		want output
	}{
		{
			name: "ok",
			args: args{
				err:  New(InvalidArgument),
				code: InvalidArgument,
			},
			want: output{
				ok: true,
			},
		},
		{
			name: "ng",
			args: args{
				err:  New(InvalidArgument),
				code: NotFound,
			},
			want: output{
				ok: false,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, EqualCode(tt.args.err, tt.args.code), tt.want.ok)
		})
	}
}

func TestIs(t *testing.T) {
	var ErrMyError = New(InvalidArgument)
	type args struct {
		err    error
		target error
	}
	type output struct {
		ok bool
	}
	tests := []struct {
		name string
		args args
		want output
	}{
		{
			name: "ok",
			args: args{
				err:    ErrMyError,
				target: Wrapf(ErrMyError, "test"),
			},
			want: output{
				ok: true,
			},
		},
		{
			name: "ng",
			args: args{
				err:    New(InvalidArgument),
				target: fmt.Errorf("test"),
			},
			want: output{
				ok: false,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, Is(tt.args.err, tt.args.target), tt.want.ok)
		})
	}
}

func TestCustomError_convert(t *testing.T) {
	t.Run("valid error", func(t *testing.T) {
		validate := validator.New()

		type input struct {
			Name string `validate:"required"`
		}
		err := validate.Struct(input{})

		ce := &CustomError{}
		converted := ce.convert(err, err, "")
		assert.Equal(t, converted.(*CustomError).code, InvalidArgument)
	})

}
