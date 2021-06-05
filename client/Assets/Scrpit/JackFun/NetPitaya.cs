using System;
using System.Collections.Generic;
using Google.Protobuf;
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
            Debug.Log("Connect " + JackFunUrl.TcpHost + ":" + JackFunUrl.TcpPort + ",success");
        }

        public static void Call<TResp>(string route, IMessage req, Action<TResp> respAction, Action<PitayaError> errorAction = null)
        {
            _client.Request(route, req,
                (TResp resp) =>
                {
                    Debug.Log("call " + route + ",success, resp=" + resp);
                    respAction?.Invoke(resp);
                }, (err) =>
                {
                    Debug.LogError(err.Code);
                    errorAction?.Invoke(err);
                });
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