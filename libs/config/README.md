# Configuration

Shared library for project configuration based on YAML.

---------

To set local variables, use the _.env_ file in the root directory of the project.

The _config.yaml_ file must be available at the time the configuration is initialized.

The _.env_ and _config.yaml_ files must be in the working directory, otherwise the file
paths are determined from _ENV_FILE_ and _CONFIG_FILE_ environment variables.

