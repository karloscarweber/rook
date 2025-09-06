// grammar_test.js
import {
	makeTestFn,
	testExtractedExamples,
	assert
} from '../lib/test_helper.js';
import * as grammar from '../src/grammar.js';

const test = makeTestFn(import.meta.url);

test('Rook Extracted grammar examples', async() => testExtractedExamples(grammar.definition));

test('Rook Top Level declarations', () => {
	assert.ok(grammar.rook.match(`This is some bullshit`).failed());
	assert.ok(grammar.rook.match(`15`).failed());

	// Type declarations
	assert.ok(grammar.rook.match(`
		i32: (i32);
		String: (u32);
		Maginot: (i32);
		Nostromo: (i32);
		`).succeeded());
	assert.ok(grammar.rook.match(`"This sucks"`).failed());

	// function declarations
	assert.ok(grammar.rook.match(`
		func zero() {
			0
		}
		func add(x, y) {
			x + y
		}
		`));

	// external function declarations
	assert.ok(grammar.rook.match(`
		extern func console_log(x);
		extern func print(x);
		extern func do_something();
		extern func stop_thread(x,y);
		`));
});

test('Rook Statements', () => {
	assert.ok(grammar.rook.match(`
		let name = "Jim";
		let age = 19;
		`));

	assert.ok(grammar.rook.match(`
		if (5 == 5) {
			print(90);
		}
		`));

		//+ "while 0 {}", "while x < 10 { x := x + 1; }"
	assert.ok(grammar.rook.match(`
		while 0 {}
		while x < 10 {
			x := x + 1;
		}
		`));

	assert.ok(grammar.rook.match(`
		x := 3;
		y := 2 + 1;
		arr[x + 1] := 3;
		`));
});

// test('Rook Expressions', () => {
//
// })
