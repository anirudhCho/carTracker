'use strict';

var utils = require('../utils/writer.js');
var Car = require('../service/CarService');

module.exports.createCar = function createCar (req, res, next, body) {
  console.log("controller")
  Car.createCar(body)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.getCarById = async function getCarById (req, res, next, carId) {
   await Car.getCarById(carId)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.updateCar = function updateCar (req, res, next, body) {
  Car.updateCar(body)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
