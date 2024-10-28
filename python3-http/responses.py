from quart import jsonify


def format_status_code(resp):
    if 'status_code' in resp:
        return resp['status_code']

    return 200


def format_body(resp):
    if 'body' not in resp:
        return ""
    elif type(resp['body']) == dict:
        return jsonify(resp['body'])
    else:
        return str(resp['body'])


def format_headers(resp):
    if 'headers' not in resp:
        return []
    elif type(resp['headers']) == dict:
        headers = []
        for key in resp['headers'].keys():
            header_tuple = (key, resp['headers'][key])
            headers.append(header_tuple)
        return headers

    return resp['headers']


def format_response(resp):
    if resp is None:
        headers = [('X-Invoked', 'true')]
        return '', 200, headers
    status_code = format_status_code(resp)
    body = format_body(resp)
    headers = format_headers(resp)

    headers.append(('X-Invoked', 'true'))

    return body, status_code, headers
