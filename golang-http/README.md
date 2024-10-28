# Go Function Template

This directory provides the template for running serverless functions in Altair®
IoT Studio™ using Go.

## Prerequisites
- Go 1.19+

## Installation

To set up the environment, run the following command (if you are using the
Makefile):

```sh
make install
```

Alternatively, you can run the following commands to install the dependencies
(make sure you have Go 1.19+ installed):

```sh
go mod tidy
```

### Additional libraries

If your function requires additional libraries, you can add them to the
`go.mod` file. After adding the libraries, run the installation command again.

> **Note**: Add only libraries (and exact versions) that are supported by
> IoT Studio or you could face issues when deploying the function. You can
> check the supported libraries in the [IoT Studio documentation](https://help.altair.com/altair-iot-studio/topics/functions/go_template.htm#reference_rsf_q3m_tpb__section_ivb_jjm_tpb).

## Editing the Code

Edit the code in the [`function/handler.go`](function/handler.go) file. This is
where the main logic for the Go function resides.

## Running the Function

To run the function locally, use the following command:

```sh
make run
```

If you are not using the Makefile, you can run the following command:

```sh
export _SECRETS_PATH=../variables   # (in Windows, use set _SECRETS_PATH=..\variables)
go run .
```

The function will start a local server at `http://localhost:8082` that you can
use to test the function.

## 📖 Documentation

Please refer to the [IoT Studio documentation](https://help.altair.com/altair-iot-studio/topics/functions/go_template.htm)
for more information on using this template.
