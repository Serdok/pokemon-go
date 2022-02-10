import os
import platform
import subprocess
import sys

from openapi_config import process_openapi_template_file

app_engine_url = 'http://localhost:8080'


def run_windows():
    try:
        environ = os.environ.copy()
        environ["FIREBASE_CONFIG"] = './firebase.key.json'
        environ["FIRESTORE_EMULATOR_HOST"] = 'localhost:9100'
        process = subprocess.Popen(
            [
                'python.exe',
                'C:\\Users\\yoyou\\AppData\\Local\\Google\\Cloud SDK\\google-cloud-sdk.staging\\bin\\dev_appserver.py',
                'app.yaml',
                '--port=8080'
            ],
            env=environ,
            stdout=sys.stdout
        )
        process.communicate()
        return 0
    except RuntimeError:
        return 1


if __name__ == '__main__':
    system = platform.system()
    process_openapi_template_file(app_engine_url)

    if system == "Windows":
        run_windows()
