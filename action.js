const core = require("@actions/core");
// const fs = require("fs");

(() => {
  const nameToGreet = core.getInput("aws-key");

  console.log(process.cwd());

  // Read the contents of the 'function' folder
//   const functionDir = "functions";
//   const files = fs.readdirSync(functionDir);

//   // Print the list of files in the folder
//   console.log(files);

  console.log(`Hello ${nameToGreet}!`);
})();
