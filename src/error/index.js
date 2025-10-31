// error/index.js
import { definitions } from './definitions.js';

// stringFromError()
//
// Accepts a sourceString and an index from an ohm error node,
// Then it walks back to the beginning of the line, and forwards to the end
// of the line to get the whole error line.
function stringFromError(sourceString, index) {
	let lineNumber = lineNumberFromError(sourceString, index);

	let str = sourceString;

	// get earliest index
	let idx = index;
	if (idx > 1) {
		let currentChar = "";
		while (idx > 0 && currentChar != "\n") {
			currentChar = str.slice(idx, idx+1);
			// console.log(`currentChar: [${idx}]:(${currentChar})`);
			if (currentChar == "\n") {
				idx = idx + 1;
				break;
			}
			idx = idx - 1;
		}
		// console.log(`earliest index: ${idx}`);
		// console.log(`(${str.slice(idx, idx+1)})`);
	}
	let earliest_index = idx;

	// get latest index
	idx = index;
	let limit = (sourceString.length - 1)
	if (idx > 1 && idx < limit) {
		let currentChar = "";
		while (idx < limit && currentChar != "\n") {
			currentChar = str.slice(idx, idx+1);
			// console.log(`currentChar: [${idx}]:(${currentChar})`);
			if (currentChar == "\n") {
				idx = idx - 1;
				break;
			}
			idx = idx + 1;
		}
		// console.log(`latest index: ${idx}`);
		// console.log(`(${str.slice(idx, idx-1)})`);
	}
	let latest_index = idx;

	let error_string = str.slice(earliest_index, latest_index);
	return `${lineNumber}| ${error_string}`;
}

// lineNumberFromError()
//
// Accepts a sourceString and an index from an ohm error node,
// Then get's the line number for that error, and drafts a string.
function lineNumberFromError(sourceString, index) {
	let str = sourceString;

	// get earliest index
	let idx = 0;
	let lineNumber = 1;
	const length = sourceString.length - 1;

	while (idx < index && idx < length) {
		if (str.slice(idx, idx+1) == "\n") {
			lineNumber = lineNumber + 1;
		}
		idx = idx + 1;
	}
	let lineStr = `${lineNumber}`
	while (lineStr.length < 4) {
		lineStr = `0${lineStr}`;
	}

	return `${lineStr}`;
}

// grabErrorString()
//
// grabs an error string from an errored node.
// Mostly wraps stringFromError
function grabErrorString(errorNode) {
	const errorString = stringFromError(errorNode.source.sourceString, errorNode.source.startIdx)
	return errorString;
}

// make()
//
// make an error array thing.
// An error Array, is a two member array where the first member shows an error message.
// The second member shows the line number and the line of code where the error is.
// [String, String]
function make(code, node) {
	return [definitions[code], grabErrorString(node)];
}

// syntaxError
//
// Alias for make() above.
function syntaxError(code, node) {
	return make(code, node);
}

// typeString()
//
// Makes a type string from a valid type declaration line.
function typeString(info) {
	let str = "(";
	let prefix = "";
	info.types.map((c) => {
		str = str.concat(`${prefix}${c}`);
		prefix = ", ";
	})
	str = str.concat(")");
	return str;
}

// compilerError()
function compilerError(key, info, node) {
	const str = typeString(info);
	return [`Type already exists:\n  ${key}: ${str};\n`, grabErrorString(node)];
}

// print()
//
// Prints an error. Expects an Error array. Which is an array
// with two strings. [string, string]
function print(errorArray) {
	console.log("");
	for (const errs of errorArray) {
		for (const err of errs) {
			console.log(err);
		}
		console.log("");
	}
}

// flagError()
//
// flags a type redeclaration error.
function flagError(key, info) {
	let str = "(";
	let prefix = "";
	info.types.map((c) => {
		str = str.concat(`${prefix}${c}`);
		prefix = ", ";
	})
	str = str.concat(")");
	// console.log(`Type already exists:\n  ${key}: ${str};\n`);
}

export {
	stringFromError,
	print,
	grabErrorString,
	lineNumberFromError,
	make,
	compilerError
}
