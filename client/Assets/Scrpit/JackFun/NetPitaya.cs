using System.Collections.Generic;
using JackFun.UI;
using Pitaya;
using UnityEngine;

namespace JackFun
{
    public static class NetPitaya
    {
        private static IPitayaClient _client;


        public static void Init()
        {
            _client = new PitayaClient();


            _client.NetWorkStateChangedEvent += (ev, error) =>
            {
                if (ev == PitayaNetWorkState.Connected)
                {
                    Debug.Log("Successfully connected!");
                    OnConnected();
                }
                else if (ev == PitayaNetWorkState.FailToConnect)
                {
                    Debug.Log("Failed to connect");
                }
            };
        }

        public static void Connect()
        {
            _client.Connect(JackFunUrl.TcpHost, JackFunUrl.TcpPort, new Dictionary<string, string>());
        }

        private static void OnConnected()
        {
            CallAuth();
        }

        private static void CallAuth()
        {
            var req = new Pb.Lobby.ReqAuth()
            {
                Token = Session.Token
            };
            Debug.Log("CallAuth ");
            _client.Request<Pb.Lobby.RespAuth>("ServerLobby.ComponentLobby.ReqAuth", req,
                (resp) =>
                {
                    if (resp.ErrCode != Pb.Enum.ErrorCode.Ok)
                    {
                        Debug.LogError("auth error, =>" + resp.ErrCode);
                        UITips.Open(ErrorCodeUtil.ToString(resp.ErrCode));
                        return;
                    }

                    Debug.Log("call auth success, resp=" + resp);
                },
                (resp) => { });
        }

        public static void Release()
        {
            if (_client != null)
            {
                _client.Dispose();
                _client = null;
            }
        }
    }
}