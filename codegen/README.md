# Codegen

This directory contains the scripts and templates used to generate all the code & docs in this repository.

If you're running on windows with msys you should run the codegen script like so:

```sh
MSYS_NO_PATHCONV=1 ./codegen/codegen.sh
```

This prevents msys appending a directory path to docker arguments.

### TODO

Update to a newer version of `openapi-generator` so we get new go class generation.

There's a few errors in the `5.x` template generated code.
