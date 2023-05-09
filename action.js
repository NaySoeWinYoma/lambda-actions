const core = require("@actions/core");
const fs = require("fs");
const path = require("path");

(() => {
  const nameToGreet = core.getInput("aws-key");

    const dd = process.cwd()

  // Construct the path to the 'function' folder
  const functionDir = path.join(dd, "..", "functions");

  // Read the contents of the 'function' folder
  const files = fs.readdirSync(functionDir);

  // Print the list of files in the folder
  console.log(files);

  console.log(`Hello ${nameToGreet}!`);
})();
