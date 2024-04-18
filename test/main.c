#include <stdio.h>

#ifdef _WIN32 // 判断是否为Windows系统
#define LIB_NAME "./grpc_auth.dll"

#include <windows.h>
#define LOAD_LIBRARY(path) LoadLibrary(path)
#define GET_FUNCTION(handle, name) GetProcAddress(handle, name)
#define CLOSE_LIBRARY(handle) FreeLibrary(handle)

#else // Linux 或 Unix 系统
#define LIB_NAME "./grpc_auth.so"

#include <dlfcn.h>
#define LOAD_LIBRARY(path) dlopen(path, RTLD_LAZY)
#define GET_FUNCTION(handle, name) dlsym(handle, name)
#define CLOSE_LIBRARY(handle) dlclose(handle)

#endif

typedef int (*P_PluginInit)(char *addr);
typedef int (*P_PluginBasicAuth)(char *username, char *password, char *clientId, char *clientAddress);
typedef int (*P_PluginAclCheck)(char *username, char *clientId, char *topic, int access, int qos, int retain);

void load_execute_close_library()
{
    // 加载库文件
    void *handle = LOAD_LIBRARY(LIB_NAME);
    if (handle == NULL)
    {
        printf("Failed to load library: %s\n", LIB_NAME);
        return;
    }

    // 调用库函数

    P_PluginInit PluginInit = (P_PluginInit)GET_FUNCTION(handle, "PluginInit");
    if (PluginInit != NULL)
    {

        int code = PluginInit("127.0.0.1:10086"); // 调用库函数
        printf("Calling PluginInit code = %d\n", code);
    }
    else
    {
        printf("Failed to find PluginInit in library\n");
    }

    P_PluginBasicAuth PluginBasicAuth = (P_PluginBasicAuth)GET_FUNCTION(handle, "PluginBasicAuth");
    if (PluginBasicAuth != NULL)
    {
        int code = PluginBasicAuth("username", "password", "clientId", "clientAddress"); // 调用库函数
        printf("Calling PluginBasicAuth code = %d\n", code);
    }
    else
    {
        printf("Failed to find PluginBasicAuth in library\n");
    }

    P_PluginAclCheck PluginAclCheck = (P_PluginAclCheck)GET_FUNCTION(handle, "PluginAclCheck");
    if (PluginAclCheck != NULL)
    {
        int code = PluginAclCheck("username", "clientId", "topic", 0, 0, 0); // 调用库函数
        printf("Calling PluginAclCheck code = %d\n", code);
    }
    else
    {
        printf("Failed to find PluginAclCheck in library\n");
    }

    // 关闭库文件
    CLOSE_LIBRARY(handle);
}

int main()
{
    load_execute_close_library();
    return 0;
}
