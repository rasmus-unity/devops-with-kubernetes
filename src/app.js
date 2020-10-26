const { v4: uuidv4 } = require('uuid');

var id = uuidv4();

function outputId(arg) {
  console.log(`${new Date().toISOString()}: ${id}`);
}

outputId();
setInterval(outputId, 4995); // for some reason 5ms short interval, results in every 5 sec...
