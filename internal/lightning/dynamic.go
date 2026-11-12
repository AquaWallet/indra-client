//go:build dynamic || !static

package lightning

/*
#cgo LDFLAGS: -L${SRCDIR}/lib/bolt12/dist -lbolt12 -Wl,-rpath=${SRCDIR}/lib/bolt12/dist
*/
import "C"
