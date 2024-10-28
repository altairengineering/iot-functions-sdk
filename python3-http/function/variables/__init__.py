import json
import os

VARIABLES_SECRET = "function-variables"

_SECRETS_PATH = os.getenv('_SECRETS_PATH', '').rstrip('/')


def get(variable_name: str):
    """
    Returns the variable with the given name.

    :param variable_name: Name of the variable.
    :return: Value of the variable, or an empty string if the secret is not found.
    """
    variables = _load_variables()
    if not variables:
        return ""

    variable = variables.get(variable_name)
    if variable is None:
        return ""

    return variable.get("value")


def exists(variable_name: str) -> bool:
    """
    Returns whether the variable exists or not.

    :param variable_name: Name of the variable.
    :return: True if the variable exists, False otherwise.
    """
    variables = _load_variables()
    return variable_name in variables


def _load_variables() -> dict:
    try:
        with open(f"{_SECRETS_PATH}/{VARIABLES_SECRET}", "r") as file:
            variables = json.load(file)
    except (FileNotFoundError, json.JSONDecodeError):
        return {}

    return variables
