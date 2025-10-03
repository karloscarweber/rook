// numbers.js
// helpers for working with numbers.

const SEVEN_BIT_MASK_BIG_INT   = 0b01111111n;
const CONTINUATION_BIT = 0b10000000;

function leb128(v) {
	let val = BigInt(v);
	let more = true;
	const r = [];

	while (more) {
		const b = Number(val & SEVEN_BIT_MASK_BIG_INT);
		val = val >> 7n;
		more = val !== 0n;
		if (more) {
			r.push(b | CONTINUATION_BIT);
		} else {
			r.push(b);
		}
	}

	return r;
}

function sleb128(v) {
	let val = BigInt(v);
	let more = true;
	const r = [];

	while (more) {
		const b = Number(val & SEVEN_BIT_MASK_BIG_INT);
		const signBitSet = !!(b & 0x40);

		val = val >> 7n;

		if ((val === 0n && !signBitSet) || (val === -1n && signBitSet)) {
			more = false;
			r.push(b);
		} else {
			r.push(b | CONTINUATION_BIT)
		}
	}

	return r;
}

const MIN_U32 = 0;
const MAX_U32 = 2 ** 32 - 1;
function u32(v) {
	if (v < MIN_U32 || v > MAX_U32) {
		throw Error(`Value out of range for u32: ${v}`);
	}

	return leb128(v);
}

const MIN_I32 = -(2 ** 32 / 2);
const MAX_I32 = 2 ** 32 / 2 - 1;
const I32_NEG_OFFSET = 2 ** 32;
function i32(v) {
	if (v < MIN_I32 || v > MAX_U32) {
		throw Error(`Value out of range for i32: ${v}`);
	}

	if (v > MAX_I32) {
		return sleb128(v - I32_NEG_OFFSET);
	}

	return sleb128(v);
}

export {
	SEVEN_BIT_MASK_BIG_INT,
	CONTINUATION_BIT,
	leb128,
	sleb128,
	MIN_U32,
	MAX_U32,
	u32,
	MIN_I32,
	MAX_I32,
	I32_NEG_OFFSET,
	i32,
}

