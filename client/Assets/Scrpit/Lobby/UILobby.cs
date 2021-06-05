using FairyGUI;
using JackFun.UI;
using UnityEngine;

namespace JackFun.Lobby
{
    public class UILobby : UIPanel
    {
        private GLoader _headIcon;
        private GTextField _labelNickName;
        private GTextField _labelGold;
        private GTextField _labelDiamond;
        private GList _listGame;


        private void Awake()
        {
            packageName = "FP_Lobby";
            componentName = "FC_Lobby";
        }

        private void Start()
        {
            InitRoot();
            RefreshAll();
        }

        private void InitRoot()
        {
            _headIcon = ui.GetChild("Head_Icon").asCom.GetChild("Head").asLoader;
            _labelGold = ui.GetChild("Label_Gold").asTextField;
            _labelDiamond = ui.GetChild("Label_Diamond").asTextField;
            _labelNickName = ui.GetChild("Label_NickName").asTextField;
            _listGame = ui.GetChild("List_Game").asList;
        }


        private void RefreshAll()
        {
            RefreshNickName();
            RefreshGold();
            RefreshDiamond();
            RefreshListGame();
        }

        private void RefreshNickName()
        {
            _labelNickName.text = Session.User.NickName;
        }

        private void RefreshGold()
        {
            _labelGold.text = "金币：" + Session.User.Gold;
        }

        private void RefreshDiamond()
        {
            _labelDiamond.text = "钻石：" + Session.User.Diamond;
        }

        private void RefreshListGame()
        {
            _listGame.itemRenderer = (int index, GObject item) =>
            {
                var data = LobbyController._listGameInfo[index];
                var btn = item.asButton;
                btn.title = data.GameName;
                btn.onClick.Clear();
                btn.onClick.Add(() =>
                {
                    Debug.Log("加入游戏Id=" + data.GameId + ",游戏名=" + data.GameName);
                    if (!data.IsOpen)
                    {
                        UITips.Open("游戏暂未开放");
                    }
                });
            };

            _listGame.numItems = LobbyController._listGameInfo.Count;
        }


        private void OnDestroy()
        {
        }
    }
}