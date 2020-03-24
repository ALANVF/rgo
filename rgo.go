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
	if int(C.Rf_isNull(this.Raw)) == 1 {
		return true
	} else {
		return false
	}
}

func (this *SEXP) IsSymbol() bool {
	if int(C.Rf_isSymbol(this.Raw)) == 1 {
		return true
	} else {
		return false
	}
}

func (this *SEXP) IsLogical() bool {
	if int(C.Rf_isLogical(this.Raw)) == 1 {
		return true
	} else {
		return false
	}
}

func (this *SEXP) IsReal() bool {
	if int(C.Rf_isReal(this.Raw)) == 1 {
		return true
	} else {
		return false
	}
}

func (this *SEXP) IsComplex() bool {
	if int(C.Rf_isComplex(this.Raw)) == 1 {
		return true
	} else {
		return false
	}
}

func (this *SEXP) IsExpression() bool {
	if int(C.Rf_isExpression(this.Raw)) == 1 {
		return true
	} else {
		return false
	}
}

func (this *SEXP) IsEnvironment() bool {
	if int(C.Rf_isEnvironment(this.Raw)) == 1 {
		return true
	} else {
		return false
	}
}

func (this *SEXP) IsString() bool {
	if int(C.Rf_isString(this.Raw)) == 1 {
		return true
	} else {
		return false
	}
}

func (this *SEXP) IsObject() bool {
	if int(C.Rf_isObject(this.Raw)) == 1 {
		return true
	} else {
		return false
	}
}
