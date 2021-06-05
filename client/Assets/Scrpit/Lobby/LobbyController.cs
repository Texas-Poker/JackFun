using System.Collections.Generic;
using JackFun.UI;
using UnityEngine;

namespace JackFun.Lobby
{
    public static class LobbyController
    {
        public static List<GameInfo> _listGameInfo = new List<GameInfo>();

        public static void CallLobbyInfo()
        {
            var req = new Pb.Lobby.ReqLobbyInfo();
            Debug.Log("CallLobbyInfo ");
            NetPitaya.Call<Pb.Lobby.RespLobbyInfo>("ServerLobby.ComponentLobby.CallLobbyInfo", req,
                (resp) =>
                {
                    _listGameInfo = new List<GameInfo>();
                    foreach (var info in resp.Infos)
                    {
                        _listGameInfo.Add(new GameInfo(info));
                    }

                    UIUtil.CloseUI<UILogin>();
                    UIUtil.OpenUI<UILobby>();
                });
        }
    }
}