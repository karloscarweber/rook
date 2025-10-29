// small_compiler_test.js
import test from 'ava';
import * as sc from '../src/small_compiler.js';
import * as cc from '../src/crook_compiler.js';
import * as errors from '../src/error/index.js';

// nopLangBytes
// nopLag will only be here for a little bit to bootstrap
  // The bytes used to make a lang thing
  const nopLangBytes = [
	sc.instr.magic,
	sc.instr.version,

	// ------ type section ------

	1, // Section identifier
	4, // Section size in bytes
	1, // Number of entries that follow

	// type section - entry 0
	0x60, // Type `function`
	0, // Empty vector of parameters
	0, // Empty vector of return values

	// ------- function section -------

	3, // Section identifier
	2, // Section size in bytes
	1, // Number of entries that follow

	// function section - entry 0
	0, // Index of the type section entry

	// ------- code section -------

	10, // Section identifier
	4, // Section size in bytes
	1, // Number of entries that follow

	// code section - entry 0
	2, // Entry size in bytes
	0, // Empty vector of local variables
	11, // `end` instruction
  ].flat(Infinity);

function compileVoidLang(code) {
	if (code !== '') {
		throw new Error(`Expected empty code, got: "${code}"`);
	}

	// const bytes = [sc.instr.magic, sc.instr.version].flat(Infinity);
	const bytes = [sc.magic(), sc.version()].flat(Infinity);
	return Uint8Array.from(bytes);
}

function compileNopLang(source) {
	if (source !== '') {
		throw new Error(`Expected empty code, got: "${source}"`);
	}
	const mod = sc.module([
		sc.typesec([sc.functype([], [])] ),
		sc.funcsec([sc.typeidx(0)]),
		sc.exportsec([sc.export_('main', sc.exportdesc.func(0))]),
		sc.codesec([sc.code(sc.func([], [sc.instr.end]))]),
	]);
	return Uint8Array.from(mod.flat(Infinity));
}

test('Compiles an empty program', async t => {
	const bytes = compileVoidLang('');
	const {instance, module} = await WebAssembly.instantiate(
		compileVoidLang(''),
	);

	t.deepEqual(instance instanceof WebAssembly.Instance, true);
	t.deepEqual(module instanceof WebAssembly.Module, true);
});

test('hand crafted module with a function', async t => {
	const {instance, module} = await WebAssembly.instantiate(
		Uint8Array.from(nopLangBytes),
	);

	t.deepEqual(instance instanceof WebAssembly.Instance, true);
	t.deepEqual(module instanceof WebAssembly.Module, true);
});

test('compiledNopLang compiles to a wasm module', async t => {
  const {instance} = await WebAssembly.instantiate(compileNopLang(''));

  t.deepEqual(instance.exports.main(), undefined);
  t.throws(() => compileNopLang('42'));
});

const rook = sc.grammar.rook;

test('Build Types List, rejecting duplicates', async t => {
	const matchResult = rook.match(`i32: (i32, u32);
String: (u32);
Maginot: (i32);
Nostromo: (i32);
i32: ();
i32: ();
i32: ();
Maginot: (i32);
i32: ();
Nostromo: (i32);
`);
	t.assert(rook.match('i32: (i32, u32);').succeeded());
	t.assert(rook.match('i32: (i32, u32)').failed());
	const types = sc.buildTypesList(rook, matchResult);

	const mmap = new Map()
	mmap.set('u32', {types: []});
	mmap.set('u64', {types: []});
	mmap.set('u128', {types: []});
	mmap.set('i32', {types: []});
	mmap.set('i64', {types: []});
	mmap.set('f32', {types: []});
	mmap.set('f64', {types: []});
	mmap.set('String', {types: [ 'u32' ]});
	mmap.set('Maginot', {types: [ 'i32' ]});
	mmap.set('Nostromo', {types: [ 'i32' ]});
	t.deepEqual(types[0], mmap);
});

test('Protect against Type redeclarations, and begin a little error reporting', async t => {
	const matchResult2 = rook.match(`
		i32: (i32, u32)
		i32: (i32, u32)
		`);
	t.assert(matchResult2.failed());
});
