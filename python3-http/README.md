# Python Function Template

This directory provides the template for running serverless functions in Altair®
IoT Studio™ using Python.

## Prerequisites
- Python 3.10+

## Installation

To set up the environment, run the following command (if you are using the
Makefile):

```sh
make install
```

Alternatively, you can run the following commands to install the dependencies
(make sure you have Python 3.10+ installed):

```sh
python3 -m venv env
source env/bin/activate    # (in Windows, use env\Scripts\activate)
pip install -r requirements.txt
```

This will create a Python environment in the `env` directory and install the
required dependencies.

## Additional libraries

If your function requires additional libraries, you can add them to the
`requirements.txt` file. After adding the libraries, run the installation
command again.

> **Note**: Add only libraries (and exact versions) that are supported by
> IoT Studio or you could face issues when deploying the function. You can
> check the supported libraries in the [IoT Studio documentation](https://help.altair.com/altair-iot-studio/topics/functions/python_template.htm#reference_h53_3jm_tpb__section_ydg_3nm_tpb).

## Editing the Code

Edit the code in the [`function/handler.py`](function/handler.py) file. This is
where the main logic for the Python function resides.

## Running the Function

To run the function locally, use the following command:

```sh
make run
```

If you are not using the Makefile, you can run the following command:

```sh
export _SECRETS_PATH=../variables   # (in Windows, use set _SECRETS_PATH=..\variables)
python index.py
```

The function will start a local server at `http://localhost:8082` that you can
use to test the function.

## 📖 Documentation

Please refer to the [IoT Studio documentation](https://help.altair.com/altair-iot-studio/topics/functions/python_template.htm)
for more information on using this template.
