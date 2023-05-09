const core = require('@actions/core');


(() => {
    const nameToGreet = core.getInput('aws-key');

    console.log(`Hello ${nameToGreet}!`);
})()