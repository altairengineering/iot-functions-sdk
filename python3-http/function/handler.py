import json
from iots import API
from function import variables


def handle(req):
    # Gets the client credentials used for API authentication
    # client_credentials = variables.get("my_client_credentials")
    # my_scopes = ["category", "thing"]

    # Creates a client instance to make requests to the IoT Studio API
    # api_url = "https://api.swx.altairone.com"
    # with API(api_url).set_credentials(client_credentials["client_id"],
    #                                   client_credentials["client_secret"],
    #                                   my_scopes) as api:
    #     ...
    #     The token will be revoked when the 'with' block ends
    #     or if the code returns or raises an exception

    return {
        "status_code": 200,
        "body": req.body.decode('utf-8')
    }
