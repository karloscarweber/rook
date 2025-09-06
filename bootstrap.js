// bootstrap.js
// This file is basically here to bootstrap the assembly of the primary rook
// files and binary and stuff.

import { makeTestFn } from './lib/test_helper.js';
const test = makeTestFn(import.meta.url);

// test('Bootstrap loads', async() => {
console.log("--- Rook ---");
// });

import * as grammar_tests from './test/grammar_test.js'
import * as small_compiler_tests from './test/small_compiler_test.js';
