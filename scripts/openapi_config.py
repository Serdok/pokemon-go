import sys

import oyaml as yaml
import string


placeholder = 'server_endpoint'
yaml_signal = 'py'
template = string.Template('{{$signal:$placeholder}}')


def process_openapi_template_file(app_engine_url):
    with open('./api/openapi.template.yaml', 'r', encoding='utf-8') as stream:
        try:
            spec = yaml.safe_load(stream)
            for server in spec['servers']:
                if server['url'] == template.substitute(signal=yaml_signal, placeholder=placeholder):
                    server['url'] = f'{app_engine_url}/api/v1'
        except yaml.YAMLError as err:
            print(err)
            return 1

    with open('./api/openapi.yaml', 'w', encoding='utf-8') as stream:
        try:
            yaml.safe_dump(spec, stream, default_flow_style=False, allow_unicode=True)
        except yaml.YAMLError as err:
            print(err)
            return 1

    return 0


if __name__ == '__main__':
    process_openapi_template_file(sys.argv[1])
