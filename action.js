const core = require("@actions/core");
const fs = require("fs");
const path = require("path");

(() => {
  const nameToGreet = core.getInput("aws-key");

  //joining path of directory
  const directoryPath = path.join(process.cwd());
  //passsing directoryPath and callback function
  fs.readdir(directoryPath, function (err, files) {
    //handling error
    if (err) {
      return console.log("Unable to scan directory: " + err);
    }
    //listing all files using forEach
    files.forEach(function (file) {
      // Do whatever you want to do with the file
      console.log(file);
    });
  });

  console.log(`Hello ${nameToGreet}!`);
})();
