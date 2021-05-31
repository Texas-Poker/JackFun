using FairyGUI;

namespace JackFun.Lobby
{
    public class UILobby : UIPanel
    {
        private GLoader _headIcon;
        private GLabel _labelNickName;
        private GLabel _labelGold;
        private GLabel _labelDiamond;


        private void Awake()
        {
            packageName = "FP_Lobby";
            componentName = "FC_Lobby";
        }

        private void Start()
        {
            InitRoot();
            AddEvent();
        }


        private void InitRoot()
        {
            
            _headIcon = ui.GetChild("Head_Icon").asCom.GetChild("Head").asLoader;
            _labelGold = ui.GetChild("Label_Gold").asLabel;
            _labelDiamond = ui.GetChild("Label_Diamond").asLabel;
            _labelNickName = ui.GetChild("Label_NickName").asLabel;
        }

        private void AddEvent()
        {
            
        }

        private void RefreshAll()
        {
            
        }

        private void RefreshNickName()
        {
            
        }

        private void OnDestroy()
        {
            
        }
    }
}