// instructions.js

// helpers
// Converts Javascript string to UTF-8
function stringToBytes(s) {
	const bytes = new TextEncoder().encode(s);
	return Array.from(bytes);
}

// returns the web assembly module prefix in UTF-8
function magic() {
	// [0x00, 0x61, 0x73, 0x6d]
	return stringToBytes('\0asm')
}

// returns the web assembly binary format version number, which is 1
function version() {
	return [0x01, 0x00, 0x00, 0x00]
}

const instr = {
	version: [0x01, 0x00, 0x00, 0x00],
	magic: [0x00, 0x61, 0x73, 0x6d],
	end: 0x0b, // the end instruction to end a function body
};

export {
	stringToBytes,
	instr,
	magic,
	version,
}
