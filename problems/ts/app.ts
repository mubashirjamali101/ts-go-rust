import problem2 from "./problem2"
import problem3 from "./problem3"

let problemCount = 0;
function printProblem(problem: Function) {
  problemCount++;
  console.log(`\nProblem ${problemCount}:`);
  console.log("---------------------");
  problem()
  console.log("---------------------\n");
}

printProblem(problem2)
printProblem(problem3)
