// crook_compiler.js
import * as grammar from './grammar.js';
import { stringToBytes, instr, magic, version } from './instructions.js';
import { u32 } from './numbers.js';
import { stringFromError } from './error/index.js';

// Crook Compiler doesn't actually spit out WebAssembly, it spits out errors.
// purpose is to help the user as much as possible so there are no errors,
// and perhaps to help them write better code.
//
// So it's built in sections that kind of mirror the other compiler, but just
// collects a lot of error information.

// buildPreludeTypes()
//
// builds a map of prelude types. A type without any children
// is a foundational type. It's built in, and can't be changed.
function buildPreludeTypes() {
	const types = new Map();
	types.set("i32", { types: [] });
	types.set("i64", { types: [] });
	types.set("f32", { types: [] });
	types.set("f64", { types: [] });
	return types
}

// builds a list of types, reports errors when we can.
// Accepts:
//		(grammar: Grammar, matchResult: MatchResult]
//
// 		import { crook } from './grammar.js';
function buildTypesList(grammar, matchResult) {
	const tempSemantics = grammar.createSemantics();
	const types = buildPreludeTypes();
	let lastType = ""
	let lineNumber = 1;
	const errors = [];

	// flagError:
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
	// grabs an error string from an errored node.
	function grabErrorString(errorNode) {
		const errorString = stringFromError(errorNode.source.sourceString, errorNode.source.startIdx)
		return errorString;
	}

	// const scopes = [new Map()];
	// Operate on types
	tempSemantics.addOperation('buildTypesList', {
		_default(...children) {
			return children.forEach((c) => c.buildTypesList());
		},
		TypeDecl(tdstart, tdbody, _nl) {
			const type_name = tdstart.buildTypesList();
			const internal_types = tdbody.buildTypesList();

			if ((typeof internal_types) != 'undefined') {
				types.set(type_name, { "types": internal_types });
			}
		},
		TDStart(name, _colon) {
			lastType = name.sourceString;
			return lastType;
		},
		TDBody_goodBody(_lparen, optArgs, _rparen) {
			const children = optArgs.child(0)?.buildTypesList();
			return children;
		},
		TDBody_missingLeftParen(_lparen, _iterAny, _rparen) {
			const errorString = grabErrorString(_lparen);
			errors.push([`missing left paren.`, errorString]);
			return undefined;
		},
		TDBody_missingRightParen(_lparen, _iterAny, _rparen) {
			const errorString = grabErrorString(_rparen);
			errors.push([`missing right paren.`, errorString]);
			return undefined;
		},
		// TDBody_badBody(_paren, _iterAny) {
		// 	errors.push([`expecting identifier on line: ${lineNumber} `]);
		// 	return undefined;
		// },

		TypeDeclNoColo(_ident, _lpar, _itparams, _rpar, _semi) {
			const errorString = grabErrorString(_lpar);
			errors.push([`missing colon.`, errorString]);
			return undefined;
		},
		TypeDeclNoSemi(tdstart, tbody, _any) {
			const errorString = grabErrorString(_any);
			errors.push([`missing semicolon.`, errorString]);
			return undefined;
		},
		Params(ident, _, iterIdent) {
			const childTypes = [];
			// grabs every valid identifier
			for (const id of [ident, ...iterIdent.children]) {
				const name = id.sourceString;
				childTypes.push(name);
			}
			const tt = new Map();
			tt.set("types", childTypes)
			types.set(lastType, tt);
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
	// compiler methods
	buildTypesList,
}
