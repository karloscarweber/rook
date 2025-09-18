// bootstrap.js
// This file is basically here to bootstrap the assembly of the primary rook
// files and binary and stuff.

import { makeTestFn } from './lib/test_helper.js';
const test = makeTestFn(import.meta.url);

test('does this work?', async() => {
	console.log("it better");
});

import * as grammar_tests from './test/grammar_test.js'
