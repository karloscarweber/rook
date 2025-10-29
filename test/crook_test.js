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
		const types = cc.buildTypesList(crook, matchResult3);
		t.deepEqual(types.at(0).get("dude").get("types"), ['i32', 'u32']);
		t.deepEqual(types.at(0).get("nope").get("types"), ['i32']);

		// uncomment to print errors
		// errors.print(types.at(1));
	}

});
