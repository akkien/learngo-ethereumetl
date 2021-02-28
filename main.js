const axios = require("axios");

axios({
  method: "post",
  url: "https://mainnet.infura.io/v3/2ee8969fa00742efb10051fc923552e1",
  data: {
    jsonrpc: "2.0",
    method: "eth_getBlockByNumber",
    params: ["0xB3AC19", true],
    id: 1
  }
})
  .then(function (response) {
    console.log(response);
  })
  .catch(function (error) {
    console.log(error);
  });
