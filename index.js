const fs = require("fs");

(() => {
  const files = fs.readdirSync("functions");

  files.forEach((f) => {
    console.log(f)
  })
})();
