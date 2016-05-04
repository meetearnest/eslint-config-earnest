#!/usr/bin/env bash

function install {
  # install eslint-plugin-mocha which we use to prevent .onlys from passing lint
  npm install eslint-plugin-mocha --save-dev

  # create the eslintrc files
  echo '{ "extends": "@earnest/eslint-config/mocha" }' > test/.eslintrc
  echo '{ "extends": "@earnest/eslint-config" }' > .eslintrc

  # add scripts to package.json
  npm install -g npm-add-script
  npmAddScript -k lint -v "./node_modules/.bin/eslint ."
  npmAddScript -k lint-changed -v "git diff --name-only --cached --relative | grep '\\.js$' | xargs ./node_modules/.bin/eslint"
  npm uninstall -g npm-add-script
}

case "${1}" in
  install) install $@
  ;;
esac
