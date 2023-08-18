# Hello World Application

## Contributing
Please make sure you follow project code style rules. To prevent code issues please use pre-commit hooks.
Before you can run hooks, you need to have the pre-commit package manager installed.

```sh
make install-git-tools
```

You can extend pre-commit hooks by adding codegen scripts to `run_codegen()` function or adding checks scripts to `run_checks()` functions in `./scripts/pre-commit` file