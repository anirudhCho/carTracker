'use strict';

var request = require('request');
const BLOCKCHAIN_BASE_URL = 'BLOCKCHAIN_BASE_URL'


/**
 * Manufacture a new car
 * Add a new car by Manufacturer
 * returns Car
 **/
exports.createCar = function (body) {
  return new Promise(async function (resolve, reject) {
    
    var bodyAsJson = {
      channel: 'channelName',
      chaincode: 'CarChaincode',
      chaincodeVersion: '1.0',
      method: 'createCar',
      args: [{
        model: body.model,
        manufacturedOn: body.dateManufacture
      }]
    }

    var options = {
      uri: BLOCKCHAIN_BASE_URL,
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      json: true,
      body: bodyAsJson
    };

    await request(options, function (error, response, body) {
      if (!error && response.statusCode == 200) {
        resolve(body);
      }
      else {
        reject(error);
      }
    });
  });
}


/**
 * Find car by ID
 * returns Car
 **/
exports.getCarById = function (carId) {
  return new Promise(async function (resolve, reject) {
    
    var bodyAsJson = {
      channel: 'channelName',
      chaincode: 'CarChaincode',
      chaincodeVersion: '1.0',
      method: 'getCar',
      args: [{
        carId: body.carId
      }]
    }

    var options = {
      uri: BLOCKCHAIN_BASE_URL,
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      json: true,
      body: bodyAsJson
    };

    await request(options, function (error, response, body) {
      if (!error && response.statusCode == 200) {
        resolve(body);
      }
      else {
        reject(error);
      }
    });
  });
}


/**
 * Update ownership of an existing car
 **/
exports.updateCar = function (body) {
  return new Promise(async function (resolve, reject) {
    var bodyAsJson = {
      channel: 'channelName',
      chaincode: 'CarChaincode',
      chaincodeVersion: '1.0',
      method: 'transferCar',
      args: [{
        carId: body.carId,
        currentOwner: body.currentOwner,
        newOwner: body.newOwner
      }]
    }

    var options = {
      uri: BLOCKCHAIN_BASE_URL,
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      json: true,
      body: bodyAsJson
    };

    await request(options, function (error, response, body) {
      if (!error && response.statusCode == 200) {
        resolve(body);
      }
      else {
        reject(error);
      }
    });
  });
}

