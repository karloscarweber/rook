// grammar_test.js
import test from 'ava';
import { testExtractedExamples } from '../lib/test_helper.js';
import * as ohm from 'ohm-js';
import { extractExamples } from 'ohm-js/extras';
import * as grammar from '../src/grammar.js';
import * as sc from '../src/small_compiler.js';
import * as cc from '../src/crook_compiler.js';
import * as errors from '../src/error/index.js';

test('Crook Grammar', async t => {
	const grammarSource = grammar.strict.concat(grammar.loose);
	const crook = ohm.grammars(grammarSource).Crook
	for (const ex of extractExamples(grammarSource)) {
		const result = crook.match(ex.example, ex.rule);
		t.deepEqual(result.succeeded(), ex.shouldMatch, JSON.stringify(ex));
	}

	const sample = `morning: i32, u32, f64); morning: (i32, u32, f64);`;

	const matchResult = crook.match(sample);
	/// uncomment to trace
	// const trace = crook.trace(sample);
	// console.log(trace);

	t.assert(matchResult.succeeded());
});

// Test Crook in catching type declarations errors
test('Crook TypeDeclarations', t => {
	const matchResult = grammar.crook.match(`
		morning: i32, u32, f64);
		morning (i32, u32, f64);
		whatever: i32);
		morning: (i32, u32, f64);
	`)
	t.assert(matchResult.succeeded());

	// matchResult
});

const rook = sc.grammar.rook;
const crook = sc.grammar.crook;

test('Trigger a change to use the Crook Compiler when we encounter an error.', async t => {
	const program = `dude: (i32, u32);
whatever: i32);
nope: (i32;
nope: (i32);
nope: i32);
`;
	const matchResult2 = rook.match(program);
	t.assert(matchResult2.failed());

	if (matchResult2.failed()) {

		const matchResult3 = crook.match(program);
		t.assert(matchResult3.succeeded());

		// test the types
		const chunk = cc.buildTypesList(crook, matchResult3);
		t.deepEqual(chunk.types.get("dude").get("types"), ['i32', 'u32']);
		t.deepEqual(chunk.types.get("nope").get("types"), ['i32']);

		// uncomment to print errors
		// errors.print(types.at(1));
	}

});

test('collect the correct list of type errors when we encounter them.', async t => {
	const program = `dude: (i32, u32);
whatever: i32);
nope: (i32;
nope: (i32);
nope: i32);
nope (i32);
dude: (i32, u32)
`;

	const matchResult2 = rook.match(program);
	t.assert(matchResult2.failed());

	if (matchResult2.failed()) {

		const matchResult3 = crook.match(program);
		t.assert(matchResult3.succeeded());

		// test the types
		const chunk = cc.buildTypesList(crook, matchResult3);

		t.deepEqual(chunk.errors[0][0], "missing left paren: `(`.");
		t.deepEqual(chunk.errors[0][1], "0002| whatever: i32)");

		t.deepEqual(chunk.errors[1][0], "missing right paren: `)`.");
		t.deepEqual(chunk.errors[1][1], "0003| nope: (i32");

		t.deepEqual(chunk.errors[2][0], "missing left paren: `(`.");
		t.deepEqual(chunk.errors[2][1], "0005| nope: i32)");

		t.deepEqual(chunk.errors[3][0], "missing colon: `:`.");
		t.deepEqual(chunk.errors[3][1], "0006| nope (i32)");

		// TODO: semicolon is missing from error string.
		// t.deepEqual(chunk.errors[4][0], "missing semicolon: `;`.");
		// t.deepEqual(chunk.errors[4][1], "0007| nope (i32)");
	}

});
