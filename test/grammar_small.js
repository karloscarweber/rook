// grammar.js
import * as ohm from 'ohm-js';
import assert from 'node:assert';

const parentDef = String.raw`
  Parent {
	start = "parent"
  }
`;
const parentGrammar = ohm.grammar(parentDef);

const childDef = String.raw`
  Child <: Parent {
	start := "child"
  }
`;

// Note that we use ohm.grammars, not ohm.grammar.
const childGrammar = ohm.grammars(childDef, {Parent: parentGrammar});

// childGrammar:
// {
//  Parent: object,
//  Child: object
// }

const combinedDef = parentDef.concat(childDef);
const grammars = ohm.grammars(combinedDef);
// grammars:
// {
//  Parent: object,
//  Child: object
// }

// console.log(parentGrammar)
assert.ok(parentGrammar.name == "Parent")
// console.log(childGrammar)
assert.ok(childGrammar.Child)
console.log(Object.keys(childGrammar))
assert.deepEqual([ 'Child' ], Object.keys(childGrammar))
// console.log(grammars)
assert.ok(grammars.Parent)
assert.ok(grammars.Child)
console.log(Object.keys(grammars))
assert.deepEqual([ 'Parent', 'Child' ], Object.keys(grammars))


