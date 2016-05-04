# eslint-config-earnest
Shareable ESLint configuration for the Earnest JavaScript style guide.

If you are looking for an ES7 version, look [here](https://github.com/meetearnest/eslint-config-earnest-es7). 

## Usage

1. Install the following `"devDependencies"` in your repo:

    `npm i --save-dev @earnest/eslint-config@latest`

2. (Recommended) Setup your editor to support inline ESLint support. For Sublime Text, that means 
`npm install -g eslint` then installing `SublimeLinter` and `SublimeLinter-contrib-eslint` packages. 
For Vim, use [Syntastic](https://github.com/scrooloose/syntastic).

### Note: the install script does the following for you

1. Add a root level `.eslintrc` that references this package

    ```
    echo '{ "extends": "@earnest/eslint-config" }' > .eslintrc
    ```

2. Add another `.eslintrc` to your `test` folder that supports mocha
  
    ```
    npm install eslint-plugin-mocha --save-dev
    echo '{ "extends": "@earnest/eslint-config/mocha" }' > test/.eslintrc
    ```
    
3. (Recommended) Add the following entries to your `package.json` for simplified CLI access to linting:

    ```json
    "scripts": {
      "lint": "./node_modules/.bin/eslint .",
      "lint-changed": "git diff --name-only --cached --relative | grep '\\.js$' | xargs ./node_modules/.bin/eslint"
    }
    ```

