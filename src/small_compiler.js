// small_compiler.js
import * as grammar from './grammar.js';
import { stringToBytes, instr, magic, version } from './instructions.js';
import { u32 } from './numbers.js';
import { grabErrorString } from './error/index.js';


// section function
// compiles a section
function section(id, contents) {
	const sizeInBytes = contents.flat(Infinity).length;
	return [id, u32(sizeInBytes), contents]
}

// vector function
// vectors are arrays that are preceded by the length of the vector.
function vec(elements) {
	return [u32(elements.length), elements];
}

const SECTION_ID_TYPE = 1;

// returns a functype section
// it's defined as the 0x60 byte followed by a vector of parameters,
// and followed by a vector of result types.
function functype(paramTypes, resultTypes) {
	return [0x60, vec(paramTypes), vec(resultTypes)];
}

// returns a type section
function typesec(functypes) {
	return section(SECTION_ID_TYPE, vec(functypes));
}

const SECTION_ID_FUNCTION = 3;

// LEB128 encoded u32 value
const typeidx = (x) => u32(x);

// funcsec: a section with a section id3, that contains a vector
// of type indices. Which are like parameters and return types.
function funcsec(typeidxs) {
	return section(SECTION_ID_FUNCTION, vec(typeidxs));
}

const SECTION_ID_CODE = 10;

// encodes the code part, along with its length in bytes.
function code(func) {
	const sizeInBytes = func.flat(Infinity).length;
	return [u32(sizeInBytes), func];
}

// encodes the function part, so locals and the body in a vector.
// remember a vector has it's length.
function func(locals, body) {
	return [vec(locals), body];
}

// encodes a code section.
function codesec(codes) {
	return section(SECTION_ID_CODE, vec(codes));
}

const SECTION_ID_EXPORT = 7;

function name(s) {
	return vec(stringToBytes(s));
}

function export_(nm, exportdesc) {
	return [name(nm), exportdesc];
}

function exportsec(exports) {
	return section(SECTION_ID_EXPORT, vec(exports));
}

const funcidx = (x) => u32(x);

const exportdesc = {
	func (idx) {
		return [0x00, funcidx(idx)];
	},
};

// makes a web assembly module
function module(sections) {
	return [instr.magic, instr.version, sections];
}

// Small Semantics
// const semantics = grammar.rook.createSemantics();

function buildPreludeTypes() {
	const types = new Map();
	types.set("u32", { types: [] });
	types.set("u64", { types: [] });
	types.set("u128", { types: [] });
	types.set("i32", { types: [] });
	types.set("i64", { types: [] });
	types.set("f32", { types: [] });
	types.set("f64", { types: [] });
	return types
}

// builds a list of types, reports errors when we can. What is that?
function buildTypesList(grammar, matchResult) {
	const tempSemantics = grammar.createSemantics();
	const types = buildPreludeTypes();
	const errors = [];

	// flagError:
	// flags a type redeclaration error.
	function flagError(key, info, node) {
		let errorString = grabErrorString(node)

		let str = "(";
		let prefix = "";
		info.types.map((c) => {
			str = str.concat(`${prefix}${c}`);
			prefix = ", ";
		})
		str = str.concat(")");
		return `Type already exists:\n  ${key}: ${str};\n`;
	}

	// const scopes = [new Map()];


	// Operation: buildTypesList
	//
	// This operation iterates through top level declarations
	// and collects the data types used. Data types are aliases
	// for the builtin in number types: i32, i64, f32, f64, and the unsigned
	// numeric types: u32, u64, u128. (Rook assumes that i32, and i64
	// are signed, thus we have u32, and u64.)
	//
	// Type information is collected in this way to replace user defined types
	// later during compilation. Data types will expand over time.
	//
	// Reference: https://webassembly.github.io/spec/core/syntax/types.html
	tempSemantics.addOperation('buildTypesList', {
		_default(...children) {
			return children.forEach((c) => c.buildTypesList());
		},
		TypeDecl(name, _colon, _lparen, optParams, _rparen, _semicolon) {
			if (typeof optParams.child(0) !== "undefined" && typeof optParams.child(0) != undefined) {
				const childTypes = optParams.child(0).buildTypesList();
				const info = { types: childTypes };
				const key = name.sourceString
				if (types.has(key)) {
					flagError(key, info, name);
				} else {
					types.set(key, info);
				}
			}
		},
		Params(ident, _, iterIdent) {
			const childTypes = [];
			for (const id of [ident, ...iterIdent.children]) {
				const name = id.sourceString;
				childTypes.push(name);
			}
			return childTypes;
		}
	});
	tempSemantics(matchResult).buildTypesList();
	return [types, errors];
}

export {
	grammar,
	stringToBytes,
	instr,
	u32,
	section,
	vec,
	SECTION_ID_TYPE,
	functype,
	typesec,
	SECTION_ID_FUNCTION,
	typeidx,
	funcsec,
	SECTION_ID_CODE,
	code,
	func,
	codesec,
	SECTION_ID_EXPORT,
	name,
	export_,
	exportsec,
	funcidx,
	exportdesc,
	module,
	magic,
	version,
	buildTypesList,
}
