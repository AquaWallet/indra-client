//go:build static && !dynamic

package lightning

/*
#cgo LDFLAGS: ${SRCDIR}/lib/bolt12/dist/libbolt12.a -Wl,--no-as-needed -Wl,--allow-multiple-definition -ldl -lm
*/
import "C"
