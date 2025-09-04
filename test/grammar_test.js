// grammar_test.js
import { makeTestFn, testExtractedExamples } from '../lib/test_helper.js';
import * as grammar from '../src/grammar.js';

const test = makeTestFn(import.meta.url);

test('Test the grammar', async() => {
	console.log("This is in grammar_test");
});


test('Extracted examples', () => testExtractedExamples(grammar.definition));
