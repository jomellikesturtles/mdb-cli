const { spawn } = require("child_process");
const goProcess = spawn("./mdb-cli"); // After running 'go build'


console.log('starting node...')

goProcess.stdout.on("data", (data) => {
  console.log(`From Go: ${data.toString()}`);
});

goProcess.stdin.write("Hello from Node\n");
