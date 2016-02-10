module.exports = {
  "extends": "@earnest/eslint-config-es7",
  "plugins": [
    "mocha"
  ],
  "env": {
    "mocha": true
  },
  "globals": {
    "should": true
  },
  "rules": {
    "mocha/no-exclusive-tests": 2
  }
};
