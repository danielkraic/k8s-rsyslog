import time
import logging
import logging.handlers

def main():
    logger = logging.getLogger('rsyslog-app-py')
    logger.setLevel(logging.INFO)

    #add handler to the logger
    handler = logging.handlers.SysLogHandler('/var/run/rsyslog/dev/log')

    #add formatter to the handler
    # formatter = logging.Formatter('Python: { "loggerName":"%(name)s", "timestamp":"%(asctime)s", "pathName":"%(pathname)s", "logRecordCreationTime":"%(created)f", "functionName":"%(funcName)s", "levelNo":"%(levelno)s", "lineNo":"%(lineno)d", "time":"%(msecs)d", "levelName":"%(levelname)s", "message":"%(message)s"}')
    formatter = logging.Formatter('%(filename)s: %(funcName)s: %(message)s')

    handler.formatter = formatter
    logger.addHandler(handler)

    logger.info("Test Message")

    while (True):
        logger.info('some py msg')
        time.sleep(5)

if __name__ == '__main__':
    main()
