from quart import Blueprint
from responses import format_response

info_bp = Blueprint('info', __name__)


@info_bp.route('/__/<path:path>', methods=['GET', 'PUT', 'POST', 'PATCH', 'DELETE'])
def undefined(path):
    return format_response({"status_code": 404})


@info_bp.route('/__/health', methods=['GET'])
def health():
    return format_response({
        "status_code": 200,
        "body": "OK"
    })
