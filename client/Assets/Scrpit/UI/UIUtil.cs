using FairyGUI;
using UnityEngine;

namespace JackFun.UI
{
    public static class UIUtil
    {
        public static void Init()
        {
            UIPackage.AddPackage("UI/FP_Common");
            UIPackage.AddPackage("UI/FP_Login_Register");
            OpenUI<UILoading>();
            OpenUI<UITips>();
        }


        public static void OpenUI<T>() where T : UIPanel
        {
            var go = new GameObject
            {
                layer = LayerMask.NameToLayer("UI"),
            };
            var uiPanel = go.AddComponent<T>();
            go.name = "Panel_" + uiPanel.componentName;
            UIPackage.AddPackage("UI/" + uiPanel.packageName);
            uiPanel.fitScreen = FitScreen.FitHeightAndSetCenter;
            uiPanel.container.renderMode = RenderMode.ScreenSpaceOverlay;
            uiPanel.ui.MakeFullScreen();
            uiPanel.CreateUI();
        }
    }
}