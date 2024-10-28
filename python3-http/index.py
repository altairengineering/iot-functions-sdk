#!/usr/bin/env python
import sys
import os
import asyncio

from quart import Quart, request, utils
from info_handlers import info_bp
from responses import format_response
import inspect

app = Quart(__name__)

# Register blueprint health
app.register_blueprint(info_bp)


class Event:
    def __init__(self):
        self.body = None
        self.headers = None
        self.method = None
        self.query = None
        self.path = None

    async def init_async(self):
        self.body = await request.get_data()
        self.headers = request.headers
        self.method = request.method
        self.query = request.args
        self.path = request.path


class Context:
    def __init__(self):
        self.hostname = os.getenv('HOSTNAME', 'localhost')


@app.route('/', defaults={'path': ''}, methods=['GET', 'PUT', 'POST', 'PATCH', 'DELETE'])
@app.route('/<path:path>', methods=['GET', 'PUT', 'POST', 'PATCH', 'DELETE'])
async def call_handler(path):
    req = Event()

    await req.init_async()

    try:
        from function import handler

        # Check if handle function is async
        if inspect.iscoroutinefunction(handler.handle):
            response_data = await asyncio.gather(handler.handle(req))

            if isinstance(response_data, list):
                response_data = response_data[0]
        else:
            response_data = await utils.run_sync(handler.handle)(req)

    except Exception as e:
        response_data = error_handler(req, e)

    resp = format_response(response_data)

    return resp


def error_handler(req, exc):
    import traceback

    exc_type, exc_obj, exc_tb = sys.exc_info()
    lineno = traceback.extract_tb(exc_tb)[-1].lineno
    line_code = traceback.extract_tb(exc_tb)[-1].line

    return {
        "status_code": 500,
        "body": {
            "error": {
                "status": 500,
                "message": f"{exc_type.__name__} at line {lineno}: {exc}",
                "details": {
                    "exception_type": exc_type.__name__,
                    "exception_info": str(exc),
                    "exception_line_number": lineno,
                    "exception_line_code": line_code,
                }
            }
        },
    }


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8082)
