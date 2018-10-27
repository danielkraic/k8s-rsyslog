#include <syslog.h>
#include <unistd.h>
#include <stdio.h>

int main(void) {
    printf("hola\n");
    setlogmask (LOG_UPTO (LOG_NOTICE));

    printf("started\n");
    openlog ("rsyslog-app-c", LOG_CONS | LOG_PID | LOG_NDELAY, LOG_LOCAL1);
    printf("opened\n");

    while (1) {
        syslog (LOG_ERR, "some c msg");
        printf("msg send\n");
        sleep(5);
    }

    closelog ();
}