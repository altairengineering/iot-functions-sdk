# Altair IoT Studio Local Development Templates

This repository provides templates for running serverless functions in Altair®
IoT Studio™ using Python and Go.
These templates allow users to run functions locally just as they would on IoT
Studio, enabling the use of their preferred IDE, debugging, and rapid iteration.

## Python Function Template

### Prerequisites
- Python 3.10+

### Installation

To set up the environment, run the following command (if you are using the
Makefile):

```sh
make install-python
```

Alternatively, you can run the following commands to install the dependencies
(make sure you have Python 3.10+ installed):

```sh
cd python3-http
python3 -m venv env
source env/bin/activate    # (in Windows, use env\Scripts\activate)
pip install -r requirements.txt
```

This will create a Python environment in the `python3-http/env` directory and
install the required dependencies.

#### Additional libraries

If your function requires additional libraries, you can add them to the
`requirements.txt` file. After adding the libraries, run the installation
command again.

> **Note**: Add only libraries (and exact versions) that are supported by
> IoT Studio or you could face issues when deploying the function. You can
> check the supported libraries in the [IoT Studio documentation](https://help.altair.com/altair-iot-studio/topics/functions/python_template.htm#reference_h53_3jm_tpb__section_ydg_3nm_tpb).

### Editing the Code

Edit the code in the [`python3-http/function/handler.py`](python3-http/function/handler.py)
file. This is where the main logic for the Python function resides.

### Running the Function

To run the function locally, use the following command:

```sh
make run-python
```

If you are not using the Makefile, you can run the following command:

```sh
cd python3-http
export _SECRETS_PATH=../variables   # (in Windows, use set _SECRETS_PATH=..\variables)
python index.py
```

The function will start a local server at `http://localhost:8082` that you can
use to test the function.

## Go Function Template

### Prerequisites
- Go 1.19+

### Installation

To set up the environment, run the following command (if you are using the
Makefile):

```sh
make install-go
```

Alternatively, you can run the following commands to install the dependencies
(make sure you have Go 1.19+ installed):

```sh
cd golang-http
go mod tidy
```

#### Additional libraries

If your function requires additional libraries, you can add them to the
`go.mod` file. After adding the libraries, run the installation command again.

> **Note**: Add only libraries (and exact versions) that are supported by
> IoT Studio or you could face issues when deploying the function. You can
> check the supported libraries in the [IoT Studio documentation](https://help.altair.com/altair-iot-studio/topics/functions/go_template.htm#reference_rsf_q3m_tpb__section_ivb_jjm_tpb).

### Editing the Code

Edit the code in the [`golang-http/function/handler.go`](golang-http/function/handler.go)
file. This is where the main logic for the Go function resides.

### Running the Function

To run the function locally, use the following command:

```sh
make run-go
```

If you are not using the Makefile, you can run the following command:

```sh
cd golang-http
export _SECRETS_PATH=../variables   # (in Windows, use set _SECRETS_PATH=..\variables)
go run .
```

The function will start a local server at `http://localhost:8082` that you can
use to test the function.

## Variables

To use Variables locally as you would in IoT Studio, modify the
[`variables/function-variables`](variables/function-variables) file to include
the Variables your Function needs, with the same values and types as defined in
IoT Studio.

The `make` commands will automatically set up the functions so that they can
access Variables from this file. If you don't want to use the `make` commands,
you can set the `_SECRETS_PATH` environment variable to the path of the
`function-variables` file (e.g., `export _SECRETS_PATH=variables/function-variables`).

## About this repository

### Directory Structure

The repository contains the following directories:

- `python3-http/`: Contains the Python function template.
- `golang-http/`: Contains the Go function template.

### Makefile Targets

Use the Makefile targets to easily set up the environment and develop and test
your functions locally before deploying them to Altair IoT Studio. The Makefile
contains the following targets:

- `install-python`: Sets up the Python environment and installs dependencies.
- `install-go`: Sets up the Go environment and installs dependencies.
- `run-python`: Runs the Python function locally.
- `run-go`: Runs the Go function locally.

## 📖 Documentation

Please refer to the [IoT Studio documentation](https://help.altair.com/altair-iot-studio/topics/functions)
for more information on using these templates.
- [Python template documentation](https://help.altair.com/altair-iot-studio/topics/functions/python_template.htm)
- [Go template documentation](https://help.altair.com/altair-iot-studio/topics/functions/go_template.htm)
