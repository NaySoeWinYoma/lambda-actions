const fs = require("fs");

fs.readdirSync("functions").forEach((file) => {
  console.log(file);
});
