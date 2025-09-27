// index.js


// makes an Error Object
// function errorObject()

function stringFromError(sourceString, index) {
	let lineNumber = lineNumberFromError(sourceString, index);

	let str = sourceString;

	// get earliest index
	let idx = index;
	if (idx > 1) {
		let currentChar = ""
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
	let earliest_index = idx

	// get latest index
	idx = 18
	let limit = 32
	if (idx > 1 && idx < limit) {
		let currentChar = ""
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

export {
	// error,
	stringFromError
}
