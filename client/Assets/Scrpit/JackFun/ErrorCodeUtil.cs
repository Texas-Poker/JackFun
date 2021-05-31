using System.Diagnostics;
using Google.Protobuf.WellKnownTypes;
using Pb.Enum;

namespace JackFun
{
    public static class ErrorCodeUtil
    {
        public static string ToString(ErrorCode enumValue)
        {
            switch (enumValue)
            {
                case ErrorCode.Ok:
                    return "成功";
                case ErrorCode.EntryError:
                    return "您的游戏包非包，请到官方指定点重新下载";
                case ErrorCode.LoginPasswordError:
                    return "登录密码错误";
                case ErrorCode.LoginAccountUnExixtent:
                    return "登录账号不存在";
                case ErrorCode.RegisterAccountExit:
                    return "注册账号已存在";
                case ErrorCode.AuthFailed:
                    return "账号未登录或Token已过期";
                default:
                    return "未知错误";
            }

        }

    }
}