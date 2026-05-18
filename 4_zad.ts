const b = 2.5;
const b3 = b * b * b;

const f = (x: number) => {
    const sum = b3 + Math.pow(x, 3);
    return (1 + Math.pow(Math.sin(sum), 2)) / Math.cbrt(sum);
};

console.log("Задача А");
console.log("   x   |     y");
console.log("-------|----------");

const xs = [];
for (let x = 1.28; x <= 3.29; x += 0.4) {
    xs.push(x);
}

for (let i = 0; i < xs.length; i++) {
    const y = f(xs[i]);
    console.log(`${xs[i].toFixed(2).padStart(6)} | ${y.toFixed(6).padStart(8)}`);
}

console.log("\nЗадача Б");
console.log("   x    |     y");
console.log("--------|----------");

const xs2 = [1.1, 2.4, 3.6, 1.7, 3.9];

for (let i = 0; i < xs2.length; i++) {
    const y = f(xs2[i]);
    console.log(`${xs2[i].toFixed(1).padStart(5)}   | ${y.toFixed(6).padStart(8)}`);
}