# YAML to JSON 

A small CLI application that lets you convert a YAML file into a JSON file or the other way around, depending on which you specify as the source.
You can also print out the result to the console, rather than specifying a target file in which to place the result.

The supported extensions are:
* `.yaml`
* `.yml`
* `.json`

Various other useful methods are available in the `conversion` package that help when dealing with unstructured JSON or YAML. This may be useful when dealing with YAML or JSON in a programmatic sense.

## Usage

Converting YAML to JSON or JSON to YAML is simple in the CLI through the use of the `convert` sub-command.

    yamltojson convert <path/to/source_file> <path/to/target_file>

This will read the source file and create/write to the target file specified. For example:

    yamltojson convert ~/project/config/data.yaml ~/another-project/output.json

If you do not want to write or create the target file, you can simple output the conversion to the console window.

    yamltojson convert --print a-nice-file.yml

## TODO

* Add a `validate` sub-command which piggybacks from conversion, load yaml/json file to test its validity.