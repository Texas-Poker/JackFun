using Google.Protobuf;
using UnityEngine;

namespace JackFun
{
    public class HttpManager
    {
        public static void Init()
        {
            Entry();
        }

        private static void Entry()
        {
            var req = new Pb.Http.ReqEntry {Secret = "天王盖地虎,宝塔镇河妖"};
            UnityHTTP.Request theRequest = new UnityHTTP.Request("post", "http://127.0.0.1:8088/entry", req.ToByteArray());
            theRequest.Send((request) =>
            {
                var resp = new Pb.Http.RespEntry();
                resp.MergeFrom(request.response.bytes);

                if (resp.ErrCode != Pb.Enum.ErrorCode.Ok)
                {
                    Debug.LogError(resp.ErrCode.ToString());
                    return;
                }

                JackFunUrl.LoginUrl = resp.LoginUrl;
                JackFunUrl.RegisterUrl = resp.RegisterUrl;
                JackFunUrl.TcpHost = resp.TcpUrl.Host;
                JackFunUrl.TcpPort = resp.TcpUrl.Port;
                Debug.Log(resp.ToString());

                Register();
            });
        }

        private static void Register()
        {
            var req = new Pb.Http.ReqRegister() {Account = "jack", Password = "123"};
            UnityHTTP.Request theRequest = new UnityHTTP.Request("post", JackFunUrl.RegisterUrl, req.ToByteArray());
            theRequest.Send((request) =>
            {
                var resp = new Pb.Http.RespRegister();
                resp.MergeFrom(request.response.bytes);

                if (resp.ErrCode != Pb.Enum.ErrorCode.Ok)
                {
                    Debug.LogError("register error," + resp);
                    Login();
                    return;
                }


                Debug.Log("register result=" + resp);

                Login();
            });
        }


        private static void Login()
        {
            var req = new Pb.Http.ReqLogin() {Account = "jack", Password = "123"};
            UnityHTTP.Request theRequest = new UnityHTTP.Request("post", JackFunUrl.LoginUrl, req.ToByteArray());
            theRequest.Send((request) =>
            {
                var resp = new Pb.Http.RespLogin();
                resp.MergeFrom(request.response.bytes);

                if (resp.ErrCode != Pb.Enum.ErrorCode.Ok)
                {
                    Debug.LogError(resp.ErrCode.ToString());
                    return;
                }


                Debug.Log("login result=" + resp);
                Session.Token = resp.Token;
                NetPitaya.Connect();
            });
        }
    }
}