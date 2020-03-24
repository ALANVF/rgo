package rgo

// USING_CGO

/*
#cgo CFLAGS: -I/usr/share/R/include
#cgo LDFLAGS: -lm -lR

#include <stdlib.h>

#include <R.h>
#include <Rembedded.h>
#include <Rinternals.h>
#include <R_ext/Parse.h>

const char* R_HOME = "/usr/lib/R";
const char* R_EXEC = "/usr/bin/R";

void Rext_open(void) {
	setenv("R_HOME", R_HOME, 1);
	const char *args[] = {R_EXEC, "--silent", "--gui=none"};
	Rf_initEmbeddedR(3, (char**) args);
}

SEXP Rext_eval(const char* code) {
	return R_ParseEvalString(code, R_GlobalEnv);
}
*/
import "C"

import "unsafe"

// Types:
type SEXP struct {
	Raw C.SEXP
}



// Methods:
func Open() {
	C.Rext_open()
}

func Eval(code string) *SEXP {
	cstr := C.CString(code)
	defer C.free(unsafe.Pointer(cstr))
	return &SEXP {
		Raw: C.Rext_eval(cstr),
	}
}

func Close() {
	C.Rf_endEmbeddedR(0)
}

func NewSEXP(value interface{}) *SEXP {
	switch val := value.(type) {
		case complex64: {
			return &SEXP {
				Raw: C.Rf_ScalarComplex(C.Rcomplex {r: C.double(real(val)), i: C.double(imag(val))}),
			}
		}
		
		case complex128: {
			return &SEXP {
				Raw: C.Rf_ScalarComplex(C.Rcomplex {r: C.double(real(val)), i: C.double(imag(val))}),
			}
		}

		case int: {
			return &SEXP {
				Raw: C.Rf_ScalarInteger(C.int(val)),
			}
		}

		case float32: {
			return &SEXP {
				Raw: C.Rf_ScalarReal(C.double(val)),
			}
		}

		case float64: {
			return &SEXP {
				Raw: C.Rf_ScalarReal(C.double(val)),
			}
		}

		case bool: {
			if val {
				return &SEXP {
					Raw: C.Rf_ScalarLogical(C.int(1)),
				}
			} else {
				return &SEXP {
					Raw: C.Rf_ScalarLogical(C.int(0)),
				}
			}
		}

		case byte: {
			return &SEXP {
				Raw: C.Rf_ScalarRaw(C.Rbyte(val)),
			}
		}

		case string: {
			str := C.CString(val)
			defer C.free(unsafe.Pointer(str))
			return &SEXP {
				Raw: C.Rf_mkString(str),
			}
		}

		default: {
			panic("Invalid R type!")
		}
	}
}

func (this *SEXP) IsNull() bool {
	return int(C.Rf_isNull(this.Raw)) == 1
}

func (this *SEXP) IsSymbol() bool {
	return int(C.Rf_isSymbol(this.Raw)) == 1
}

func (this *SEXP) IsLogical() bool {
	return int(C.Rf_isLogical(this.Raw)) == 1
}

func (this *SEXP) IsReal() bool {
	return int(C.Rf_isReal(this.Raw)) == 1
}

func (this *SEXP) IsComplex() bool {
	return int(C.Rf_isComplex(this.Raw)) == 1
}

func (this *SEXP) IsExpression() bool {
	return int(C.Rf_isExpression(this.Raw)) == 1
}

func (this *SEXP) IsEnvironment() bool {
	return int(C.Rf_isEnvironment(this.Raw)) == 1
}

func (this *SEXP) IsString() bool {
	return int(C.Rf_isString(this.Raw)) == 1
}

func (this *SEXP) IsObject() bool {
	return int(C.Rf_isObject(this.Raw)) == 1
}

func (this *SEXP) IsArray() bool {
    return int(C.Rf_isArray(this.Raw)) == 1
}

func (this *SEXP) IsFactor() bool {
    return int(C.Rf_isFactor(this.Raw)) == 1
}

func (this *SEXP) IsFrame() bool {
    return int(C.Rf_isFrame(this.Raw)) == 1
}

func (this *SEXP) IsFunction() bool {
    return int(C.Rf_isFunction(this.Raw)) == 1
}

func (this *SEXP) IsInteger() bool {
    return int(C.Rf_isInteger(this.Raw)) == 1
}

func (this *SEXP) IsLanguage() bool {
    return int(C.Rf_isLanguage(this.Raw)) == 1
}

func (this *SEXP) IsList() bool {
    return int(C.Rf_isList(this.Raw)) == 1
}

func (this *SEXP) IsMatrix() bool {
    return int(C.Rf_isMatrix(this.Raw)) == 1
}

func (this *SEXP) IsNewList() bool {
    return int(C.Rf_isNewList(this.Raw)) == 1
}

func (this *SEXP) IsNumber() bool {
    return int(C.Rf_isNumber(this.Raw)) == 1
}

func (this *SEXP) IsNumeric() bool {
    return int(C.Rf_isNumeric(this.Raw)) == 1
}

func (this *SEXP) IsPairList() bool {
    return int(C.Rf_isPairList(this.Raw)) == 1
}

func (this *SEXP) IsPrimitive() bool {
    return int(C.Rf_isPrimitive(this.Raw)) == 1
}

func (this *SEXP) IsTs() bool {
    return int(C.Rf_isTs(this.Raw)) == 1
}

func (this *SEXP) IsUserBinop() bool {
    return int(C.Rf_isUserBinop(this.Raw)) == 1
}

func (this *SEXP) IsValidString() bool {
    return int(C.Rf_isValidString(this.Raw)) == 1
}

func (this *SEXP) IsValidStringF() bool {
    return int(C.Rf_isValidStringF(this.Raw)) == 1
}

func (this *SEXP) IsVector() bool {
    return int(C.Rf_isVector(this.Raw)) == 1
}

func (this *SEXP) IsVectorAtomic() bool {
    return int(C.Rf_isVectorAtomic(this.Raw)) == 1
}

func (this *SEXP) IsVectorList() bool {
    return int(C.Rf_isVectorList(this.Raw)) == 1
}

func (this *SEXP) IsVectorizable() bool {
    return int(C.Rf_isVectorizable(this.Raw)) == 1
}



func (this *SEXP) Bool() bool {
	return int(C.Rf_asLogical(this.Raw)) != 0
}

func (this *SEXP) Int() int {
	return int(C.Rf_asInteger(this.Raw))
}

func (this *SEXP) Byte() byte {
	return byte(*C.RAW(this.Raw))
}

func (this *SEXP) Float32() float32 {
	return float32(C.Rf_asReal(this.Raw))
}

func (this *SEXP) Float64() float64 {
	return float64(C.Rf_asReal(this.Raw))
}

func (this *SEXP) Complex64() complex64 {
	cmplx := C.Rf_asComplex(this.Raw)
	return complex64(complex(float64(cmplx.r), float64(cmplx.i)))
}

func (this *SEXP) Complex128() complex128 {
	cmplx := C.Rf_asComplex(this.Raw)
	return complex(float64(cmplx.r), float64(cmplx.i))
}
