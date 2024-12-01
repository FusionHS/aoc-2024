function loadData(fileName: string): string {
  return Deno.readTextFileSync(fileName);
}

function getRows(): { left: number[]; right: number[] } {
  const dataString = loadData("input-a.txt");

  const left: number[] = [];
  const right: number[] = [];

  dataString
    .split("\n")
    .forEach((item) => {
      let strings: string[] = item.split(/\s+/);
      left.push(Number(strings[0]));
      right.push(Number(strings[1]));
    });
  return { left, right };
}

function partA() {
  const { left, right } = getRows();

  left.sort();
  right.sort();

  return left
    .map((item, i) => Math.abs(item - right[i]))
    .reduce((a, b) => a + b, 0);
}
function partB() {
  const { left, right } = getRows();

  return left
    .map((item) => right.filter((it) => it == item).length * item)
    .reduce((a, b) => a + b, 0);
}

console.log("Answer part A:", partA());
console.log("Answer part B:", partB());
