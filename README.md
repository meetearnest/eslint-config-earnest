# eslint-config-earnest
Shareable ESLint configuration for the Earnest JavaScript style guide

## Usage

1. Add the dependency to your project
    `npm i --save-dev eslint-config-earnest`
2. Edit your `package.json` to change the dependency version to `latest`
    `"eslint-config-earnest": "latest"`
3. Add a root level `.eslintrc` that references this package
    `{"extends": "earnest"}`
4. Add another `.eslintrc` to your `test` folder that supports mocha
    ```
    {
      "extends": "earnest",
      "env": {
        "mocha": true
      }
    }
    ```
