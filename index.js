const fs = require("fs");


fs.writeFileSync('test.tf', "Test");

const content = fs.readFileSync("test.tf", "utf8");

console.log(content)