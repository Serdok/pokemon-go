import platform
import subprocess
import sys

from openapi_config import process_openapi_template_file
from app_config import process_app_template_file

app_engine_url = 'http://localhost:8080'
firebase_config = './firebase-adminsdk.json'


def run_windows():
    try:
        process = subprocess.Popen(
            [
                'python.exe',
                'C:\\Users\\yoyou\\AppData\\Local\\Google\\Cloud SDK\\google-cloud-sdk.staging\\bin\\dev_appserver.py',
                'app.yaml',
                '--port=8080'
            ],
            stdout=sys.stdout
        )
        process.communicate()
        return 0
    except RuntimeError:
        return 1


if __name__ == '__main__':
    system = platform.system()
    process_openapi_template_file(app_engine_url)
    process_app_template_file(firebase_config)

    if system == "Windows":
        run_windows()
