// index.js





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
	print
}
