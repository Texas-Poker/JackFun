using JackFun.Lobby;
using Pb.Lobby;
using UnityEngine;

namespace JackFun.UI
{
    public static class LoginController
    {
        public static void CallAuth()
        {
            var req = new ReqAuth()
            {
                Token = Session.Token
            };
            Debug.Log("CallAuth ");
            var router = "ServerLobby.ComponentLobby.ReqAuth";
            NetPitaya.Call<RespAuth>(router, req,
                (resp) =>
                {
                    Session.User = new User(resp);
                    Debug.Log("[" + router + "] resp >>" + resp);
                    LobbyController.CallLobbyInfo();
                });
        }
    }
}