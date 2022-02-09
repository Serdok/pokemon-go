import sys

import oyaml as yaml
import string


placeholder = 'firebase_config'
yaml_signal = 'py'
template = string.Template('{{$signal:$placeholder}}')


def process_app_template_file(firebase_config):
    with open('./app.template.yaml', 'r', encoding='utf-8') as stream:
        try:
            app = yaml.safe_load(stream)
            for key in app['env_variables']:
                value = app['env_variables'][key]
                if value == template.substitute(signal=yaml_signal, placeholder=placeholder):
                    app['env_variables'][key] = f'{firebase_config}'
        except yaml.YAMLError as err:
            print(err)
            return 1

    with open('./app.yaml', 'w', encoding='utf-8') as stream:
        try:
            yaml.safe_dump(app, stream, default_flow_style=False, allow_unicode=True)
        except yaml.YAMLError as err:
            print(err)
            return 1

    return 0


if __name__ == '__main__':
    process_app_template_file(sys.argv[1])
