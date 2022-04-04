const axios = require("axios");
const { Pool, Client } = require('pg')


const main = async () => {
  checkDb();
}

const checkDb = async () => {
  const client = new Client({
    user: 'postgres',
    host: '127.0.0.1',
    database: 'bikestore',
    password: 'mysecret',
    port: 5432,
  })  
  client.connect();

  client.query('SELECT NOW()', (err, res) => {
    console.log(err, res)
    client.end()
  })
}

const getBlock = () => {
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
}

main();