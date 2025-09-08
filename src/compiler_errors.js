// compiler_errors.js
//
// Compiler errors are errors that we find when we are compiling.
// The structure is:
// compiler_error {
//		type: "error",
//		filename: string,
//		line: int,
//		column: int,
//		message: string,
//		suggestion: string,
//		severity: enum { warning, error } ,
// }
//

function error() {
	const comp_error = undefined;

	return comp_error;
}

const types = {}

// errors for type compilation
types.types_list.undeclared_type;
types.types_list.redeclaration;
types.types_list.unterminated_declaration;
