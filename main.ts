const ab = new ArrayBuffer(8);

const a8 = new Uint8Array(ab);

const a32 = new Uint32Array(ab);

a32[0] = 34;

console.log(ab);
