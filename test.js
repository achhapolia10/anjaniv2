const a = require("child_process");

const subprocess = a.spawn("./imanager64", undefined, {});

subprocess.stdout.on("data", (data) => {
  console.log(`Received chunk ${data}`);
});
