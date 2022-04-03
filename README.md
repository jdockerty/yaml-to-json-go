# YAML to JSON 

A small CLI application that lets you convert a YAML file into a JSON file or the other way around, depending on which you specify as the source.
You can also print out the result to the console, rather than specifying a target file in which to place the result.

Simple validation is also supported, based on whether the file was valid enough to be parsed into the program.

The supported extensions are:
* `.yaml`
* `.yml`
* `.json`

Various other useful methods are available in the `conversion` package that help when dealing with unstructured JSON or YAML. This may be useful when dealing with YAML or JSON in a programmatic sense.

    go get -u github.com/jdockerty/yaml-to-json-go/conversion

Can be used to retreive the packages if you have Go installed.

## Install

If you already have Go installed, then you can simply run

```sh
go install github.com/jdockerty/yaml-to-json-go@latest
mv $(go env GOPATH)/bin/yaml-to-json-go $(go env GOPATH)/bin/yamltojson
```

This installs the binary in your Go `bin` directory and renames it for easier access on the command line, feel free to alter this to whatever suits your needs, such as `y2j`.

An alternative method is to grab the binary file from the provided S3 bucket.

```
wget https://yaml-to-json-go.s3.eu-west-2.amazonaws.com/yamltojson.zip
unzip yamltojson.zip
sudo mv yamltojson /usr/local/bin
```

Or, if you have already cloned the repository and are currently within that directory.

    make install

## Usage

Converting YAML to JSON or JSON to YAML is simple in the CLI through the use of the `convert` sub-command.

    yamltojson convert <path/to/source_file> <output_filename>

This will read the source file and create/write to the target file specified. For example:

    yamltojson convert ~/project/config/data.yaml ~/another-project/output.json

If you do not want to write or create the target file, you can simple output the conversion to the console window.

    yamltojson convert --print a-nice-file.yml

A full directory can be converted with the `--directory` or `-d` flag too, although this works best when all of the containing files within the directory are of the same type, such as converting a folder full of JSON files into YAML.

    yamltojson convert --directory="<path/to/json-files/>,<output_directory>"

There is also simple validation on a list of files. This simply tells you whether a JSON or YAML file is valid or not.

    yamltojson validate path/to/config.json path/to/another/file.yml
