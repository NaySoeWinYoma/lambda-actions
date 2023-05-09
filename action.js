const core = require('@actions/core');
const fs = require("fs");


(() => {
    const nameToGreet = core.getInput('aws-key');

    fs.readdirSync("./functions").forEach(file => {
        console.log(file);
    });

    console.log(`Hello ${nameToGreet}!`);
})()