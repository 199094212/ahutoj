#ifndef __RESULT_H__
#define __RESULT_H__
#define DEV_DEBUG 0
typedef enum
{
    OJ_AC = 1,
    OJ_WA,
    OJ_TLE,
    OJ_MLE,
    OJ_RE,
    OJ_PE,
    OJ_OLE,
    OJ_CE,
    OJ_JUDGE,
    OJ_REJUDGE,
    OJ_PENDING,
    OJ_FAILED,
} SubRes;
static const char *runningres[] = {"", "AC", "WA", "TLE", "MLE",
                                   "RE", "PE", "OLE", "CE", "JUDGING", "REJUDGING", "PENDING", "FAILED"};
typedef enum
{
    C = 1,
    CPP,
    CPP11,
    CPP17,
    JAVA,
    PYTHON3
} lanuage;
#define DATAPATH "./Data/" /*测试样例目录*/
#define DEC ".des"
#define LOGPATH "./log"
#define CONF "/etc/ahutoj/config.conf" /*配置文件目录*/
#define COMPDIR "./run%d"
#define IPC_PATH "./judge"

#endif