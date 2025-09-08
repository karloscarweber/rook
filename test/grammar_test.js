// grammar_test.js
import {
	makeTestFn,
	testExtractedExamples,
	assert
} from '../lib/test_helper.js';
import * as grammar from '../src/grammar.js';

const test = makeTestFn(import.meta.url);

test('Rook Extracted grammar examples', async() => testExtractedExamples(grammar.strict));

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
	const matchR = grammar.rook.match(`
	func zero(): i32 {
		0
	}
	func add(x, y): i32 {
		x + y
	}
	`);
	if (matchR.failed() == true) {
		console.log(matchR.message);
	}
	assert.ok(matchR.succeeded());

	// external function declarations
	assert.ok(grammar.rook.match(`
		extern func console_log(x);
		extern func print(x);
		extern func do_something();
		extern func stop_thread(x,y);
		`).succeeded());
});

test('Rook Statements', () => {
	const matchR = grammar.rook.match(`
		func start(): void {
			let name = "Jim";
			let age = 19;
			19
		}
		`);
	if (matchR.failed() == true) {
		console.log(matchR.message);
	}
	assert.ok(matchR.succeeded());

	const matchR2 = grammar.rook.match(`
		func start(): void {
			if (5 == 5) {
				print(90);
			}
		}
		`);
	if (matchR2.failed() == true) {
		console.log(matchR2.message);
	}
	assert.ok(matchR2.succeeded());

		//+ "while 0 {}", "while x < 10 { x := x + 1; }"
	assert.ok(grammar.rook.match(`
		func start(): void {
			while 0 {}
			while x < 10 {
				x := x + 1;
			}
		}
		`).succeeded());

	assert.ok(grammar.rook.match(`
		func start(): void {
			x := 3;
			y := 2 + 1;
			arr[x + 1] := 3;
		}
		`).succeeded());
});

// test('Rook Expressions', () => {
//
// })

// Test looseGrammar
test('Loose Rook Extracted grammar examples', async() => testExtractedExamples(grammar.loose));
test('Test loose Rook Grammar', () => {
	assert.ok(grammar.rookLoose.match(`
		morning: i32, u32, f64);
		morning (i32, u32, f64)
	`).failed());
});
