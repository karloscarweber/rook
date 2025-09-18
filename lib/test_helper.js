// lib/test_helper.js
// setup
import assert from 'node:assert';
import {basename} from 'node:path';
import process from 'node:process';
import {default as nodeTest} from 'node:test';
import {fileURLToPath} from 'node:url';
import * as ohm from 'ohm-js';
import {extractExamples} from 'ohm-js/extras';

function makeTestFn(url) {
  const filename = fileURLToPath(url);
  // Return a function with the same interface as Node's `test` function.
  return (name, ...args) => {
	  const filejs = basename(filename, '.js');
	  nodeTest(`[${filejs}] ${name}`, ...args);
  };
}

// helper to test the ohm grammar we define
function testExtractedExamples(grammarSource) {
	const grammar = ohm.grammar(grammarSource);
	console.log(grammar)
	for (const ex of extractExamples(grammarSource)) {
		const result = grammar.match(ex.example, ex.rule);
		assert.strictEqual(result.succeeded(), ex.shouldMatch, JSON.stringify(ex));
	}
}

function testExtractedExamplesFromConcatenatedGrammars(grammarSource) {
	const grammar = ohm.grammars(grammarSource);
	// console.log(grammar)
	// for (const ex of extractExamples(grammarSource)) {
//
	// 	const result = grammar.match(ex.example, ex.rule);
	// 	assert.strictEqual(result.succeeded(), ex.shouldMatch, JSON.stringify(ex));
	// }
}

export {
  makeTestFn,
  testExtractedExamples,
  testExtractedExamplesFromConcatenatedGrammars,
  assert,
};
