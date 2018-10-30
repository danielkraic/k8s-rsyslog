#include <syslog.h>
#include <unistd.h>
#include <stdio.h>

int main(void) {
    setlogmask (LOG_UPTO (LOG_NOTICE));
    openlog ("rsyslog-app-c", LOG_CONS | LOG_PID | LOG_NDELAY, LOG_LOCAL1);
 
    while (1) {
        syslog (LOG_ERR, "some c msg");
        sleep(3);
    }

    closelog ();
    return 0;
}