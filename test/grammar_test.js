// grammar_test.js
import {
	makeTestFn,
	testExtractedExamples,
	assert
} from '../lib/test_helper.js';
import * as grammar from '../src/grammar.js';

const test = makeTestFn(import.meta.url);

test('Test the grammar', async() => {
	console.log("This is in grammar_test");
});

test('Rook Extracted grammar examples', async() => testExtractedExamples(grammar.definition));

//
test('Rook Top Level declarations', () => {
	assert.ok(grammar.rook.match(`This is some bullshit`).failed());
	assert.ok(grammar.rook.match(`15`).failed());
	assert.ok(grammar.rook.match(`
		i32: (i32);
		String: (u32);
		Maginot: (i32);
		Nostromo: (i32);
		`).succeeded());
	assert.ok(grammar.rook.match(`"This sucks"`).failed());
	assert.ok(grammar.rook.match(`
		func zero() {
			0
		}
		func add(x, y) {
			x + y
		}
		`))
});

//+ "type tuple: (i32, i32);"
//- "type age (i32);"
//- "type slash ( )"
//+ "type morning: (i32, u32, f64);"
