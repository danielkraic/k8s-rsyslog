import datetime
import socket
import time
import logging
import logging.handlers

# inspired by https://github.com/EasyPost/syslog-rfc5424-formatter/
class MyFormatter(logging.Formatter, object):
    def __init__(self, *args, **kwargs):
        return super(MyFormatter, self).__init__(*args, **kwargs)

    def format(self, record):
        record.__dict__['isotime'] = datetime.datetime.fromtimestamp(record.created).strftime(format='%b %d %H:%M:%S')

        header = '{isotime} {name}[{process}]: '.format(
            **record.__dict__
        )
        
        return header + super(MyFormatter, self).format(record)

def main():
    logger = logging.getLogger('rsyslog-app-py')
    logger.setLevel(logging.INFO)

    handler = logging.handlers.SysLogHandler('/dev/log')
    handler.setFormatter(MyFormatter())
    logger.addHandler(handler)

    while (True):
        logger.info('some py msg')
        time.sleep(3)

if __name__ == '__main__':
    main()
