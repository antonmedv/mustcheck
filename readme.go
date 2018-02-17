// ███╗   ███╗██╗   ██╗███████╗████████╗  ██████╗██╗  ██╗███████╗ ██████╗██╗  ██╗
// ████╗ ████║██║   ██║██╔════╝╚══██╔══╝ ██╔════╝██║  ██║██╔════╝██╔════╝██║ ██╔╝
// ██╔████╔██║██║   ██║███████╗   ██║    ██║     ███████║█████╗  ██║     █████╔╝
// ██║╚██╔╝██║██║   ██║╚════██║   ██║    ██║     ██╔══██║██╔══╝  ██║     ██╔═██╗
// ██║ ╚═╝ ██║╚██████╔╝███████║   ██║    ╚██████╗██║  ██║███████╗╚██████╗██║  ██╗
// ╚═╝     ╚═╝ ╚═════╝ ╚══════╝   ╚═╝     ╚═════╝╚═╝  ╚═╝╚══════╝ ╚═════╝╚═╝  ╚═╝

package mustcheck

// Installation:
//
//     go get github.com/antonmedv/mustcheck
//
// Import in code:
//
//     import . "github.com/antonmedv/mustcheck"
//

// Check compares err with nil and panics.
//
// In lots of cases, it's only necessary to log the error in fail fast.
// Copy-paste leads to errors and untested code (someone may forget to return,
// other may use log incorrectly). By using panic you can implement logging
// in one place (recovery) and also have a stack trace. If you code really
// has the second path or needs custom error handling just don't use it.
//
// Example:
//
//     f, err := os.Create(path)
//     check(err)
//
//
// ¯\_(ツ)_/¯
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// Must ensures what returned err is nil and returns the result, panics otherwise.
//
// This is useful for creating structures where you don't need explicit error handling
// and prefer to fail fast.
//
// Example:
//
//     return Must(strconv.Atoi(str)).(int)
//
// Note what you need to cast an interface to proper type.
//
// (╯°□°)╯︵ ┻━┻
func Must(val interface{}, err error) interface{} {
	Check(err)
	return val
}

// License under MIT.
