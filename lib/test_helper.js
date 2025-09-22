// lib/test_helper.js
// setup
import assert from 'node:assert';
import * as ohm from 'ohm-js';
import {extractExamples} from 'ohm-js/extras';

// helper to test the ohm grammar we define
function testExtractedExamples(test_context, grammarSource) {
	const t = test_context;
	const grammar = ohm.grammar(grammarSource);
	for (const ex of extractExamples(grammarSource)) {
		const result = grammar.match(ex.example, ex.rule);
		t.deepEqual(result.succeeded(), ex.shouldMatch, JSON.stringify(ex));
	}
}

export {
  testExtractedExamples,
};
