const core = require("@actions/core");
const { exec } = require("child_process");

(() => {
  const nameToGreet = core.getInput("aws-key");

  exec("ls -la", (error, stdout, stderr) => {
      if (error) {
          console.log(`error: ${error.message}`);
          return;
      }
      if (stderr) {
          console.log(`stderr: ${stderr}`);
          return;
      }
      console.log(`stdout: ${stdout}`);
  });
  
  console.log(`Hello ${nameToGreet}!`);
})();
